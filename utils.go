package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type ApiCast struct{}
type ApiUser struct{}
type MentionNotification struct{}
type ReplyNotification struct{}

type Streamable interface{}

type ExponentialCounter struct {
	maxCounter int
	counter    int
}

func (e *ExponentialCounter) reset() {
	e.counter = 0
}

func (e *ExponentialCounter) counterFunction() int {
	e.counter++
	if e.counter > e.maxCounter {
		e.counter = e.maxCounter
	}
	return e.counter
}

type BoundedSet struct {
	set map[string]bool
}

func NewBoundedSet(size int) *BoundedSet {
	return &BoundedSet{
		set: make(map[string]bool, size),
	}
}

func (b *BoundedSet) Add(attribute string) {
	if len(b.set) >= cap(b.set) {
		// Remove a random element to make space
		for k := range b.set {
			delete(b.set, k)
			break
		}
	}
	b.set[attribute] = true
}

func (b *BoundedSet) Contains(attribute string) bool {
	_, found := b.set[attribute]
	return found
}

func streamGenerator(
	function func(cursor *string, limit int) []Streamable,
	attributeName string,
	pauseAfter *int,
	skipExisting bool,
	maxCounter int,
	limit int,
	cursor *string,
) <-chan Streamable {

	output := make(chan Streamable)
	exponentialCounter := &ExponentialCounter{maxCounter: maxCounter}
	seenAttributes := NewBoundedSet(301)
	var beforeAttribute string
	withoutBeforeCounter := 0
	responsesWithoutNew := 0

	go func() {
		defer close(output)

		for {
			found := false
			var newestAttribute string
			dynamicLimit := limit

			if beforeAttribute == "" {
				dynamicLimit -= withoutBeforeCounter
				withoutBeforeCounter = (withoutBeforeCounter + 1) % (limit / 2)
			}

			log.Printf("Limit: %d", dynamicLimit)

			items := function(cursor, dynamicLimit)
			for i := len(items) - 1; i >= 0; i-- {
				item := items[i]
				attribute := fmt.Sprintf("%v", item) // Assuming attribute extraction logic here
				if seenAttributes.Contains(attribute) {
					continue
				}
				found = true
				seenAttributes.Add(attribute)
				newestAttribute = attribute
				if !skipExisting {
					output <- item
				}
			}

			beforeAttribute = newestAttribute
			skipExisting = false

			if pauseAfter != nil && *pauseAfter < 0 {
				output <- nil
			} else if found {
				exponentialCounter.reset()
				responsesWithoutNew = 0
			} else {
				responsesWithoutNew++
				if pauseAfter != nil && responsesWithoutNew > *pauseAfter {
					exponentialCounter.reset()
					responsesWithoutNew = 0
					output <- nil
				} else {
					time.Sleep(time.Duration(exponentialCounter.counterFunction()) * time.Second)
				}
			}
		}
	}()

	return output
}

func main() {
	// Example usage
	cursor := ""
	pauseAfter := 5
	limit := 50
	maxCounter := 16
	skipExisting := false
	attributeName := "hash"

	function := func(cursor *string, limit int) []Streamable {
		// Simulate function logic that returns a list of items
		return []Streamable{ApiUser{}, ApiCast{}, MentionNotification{}, ReplyNotification{}}
	}

	for item := range streamGenerator(function, attributeName, &pauseAfter, skipExisting, maxCounter, limit, &cursor) {
		fmt.Println(item)
	}
}
