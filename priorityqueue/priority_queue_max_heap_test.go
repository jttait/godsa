package priorityqueue

import (
	"testing"

	"github.com/jttait/godsa/assert"
)

func TestShouldBeSizeZeroForNewlyInstantiatedPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	result := q.Size()
	assert.AssertEqual(result, 0, t)
}

func TestShouldBeSizeOneWhenItemAddedToNewlyInstantiatedPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	result := q.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeSizeTwoWhenAddingTwoItems(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	q.Insert(6)
	result := q.Size()
	assert.AssertEqual(result, 2, t)
}

func TestShouldBeSizeTwoWhenAddingTwoIdenticalItems(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	q.Insert(5)
	result := q.Size()
	assert.AssertEqual(result, 2, t)
}

func TestShouldBeHighestPriorityItemWhenPollingPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	q.Insert(3)
	result, _ := q.Extract()
	assert.AssertEqual(result, 5, t)
}

func TestShouldBeTrueWhenPollingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	_, ok := q.Extract()
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenPollingEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	_, ok := q.Extract()
	assert.AssertFalse(ok, t)
}

func TestShouldBeHighestPriorityItemWhenPeekingPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	q.Insert(3)
	result, _ := q.Peek()
	assert.AssertEqual(result, 5, t)
}

func TestShouldBeTrueWhenPeekingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	_, ok := q.Peek()
	assert.AssertTrue(ok, t)
}

func TestShouldBeFalseWhenPeekingEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	_, ok := q.Peek()
	assert.AssertFalse(ok, t)
}

func TestShouldBeSameSizeAfterPeekingNonEmptyPriorityQueue(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	_, _ = q.Peek()
	result := q.Size()
	assert.AssertEqual(result, 1, t)
}

func TestShouldBeSizeZeroAfterPollingPriorityQueueOfSizeOne(t *testing.T) {
	q := NewPriorityQueueMaxHeap[int]()
	q.Insert(5)
	_, _ = q.Extract()
	result := q.Size()
	assert.AssertEqual(result, 0, t)
}
