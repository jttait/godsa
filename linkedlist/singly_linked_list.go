package linkedlist

// SinglyLinkedListNode is a data structure that contains a value and a pointer to the next node in
// the list.
type SinglyLinkedListNode[T any] struct {
	Val  T
	Next *SinglyLinkedListNode[T]
}

// NewSinglyLinkedListNode instantiates a singly-linked list node and returns a pointer to it.
func NewSinglyLinkedListNode[T any](i T) *SinglyLinkedListNode[T] {
	n := SinglyLinkedListNode[T]{}
	n.Val = i
	return &n
}

type SinglyLinkedList[T any] struct {
	dummyHead *SinglyLinkedListNode[T]
}

// NewSinglyLinkedList instantiates a new singly-linked list and returns a pointer to it.
func NewSinglyLinkedList[T any](items ...T) *SinglyLinkedList[T] {
	q := SinglyLinkedList[T]{}
	var zeroValue T
	q.dummyHead = NewSinglyLinkedListNode[T](zeroValue)
	for _, i := range items {
		q.InsertLast(i)
	}
	return &q
}

// Size returns the number of items in the list.
func (q *SinglyLinkedList[T]) Size() int {
	if q.dummyHead.Next == nil {
		return 0
	}
	result := 0
	current := q.dummyHead
	for current.Next != nil {
		current = current.Next
		result += 1
	}
	return result
}

// InsertFront inserts a new item at the front of the list.
func (q *SinglyLinkedList[T]) InsertFront(i T) {
	if q.dummyHead.Next == nil {
		q.dummyHead.Next = NewSinglyLinkedListNode[T](i)
	} else {
		temp := q.dummyHead.Next
		q.dummyHead.Next = NewSinglyLinkedListNode[T](i)
		q.dummyHead.Next.Next = temp
	}
}

// InsertLast inserts a new item at the end of the list.
func (q *SinglyLinkedList[T]) InsertLast(i T) {
	current := q.dummyHead
	for current.Next != nil {
		current = current.Next
	}
	current.Next = NewSinglyLinkedListNode[T](i)
}

// RemoveFront removes and returns the item at the front of the singly-linked list.
func (q *SinglyLinkedList[T]) RemoveFront() (T, bool) {
	if q.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	result := q.dummyHead.Next.Val
	q.dummyHead.Next = q.dummyHead.Next.Next
	return result, true
}

// RemoveLast removes and returns the item at the end of the list.
func (q *SinglyLinkedList[T]) RemoveLast() (T, bool) {
	if q.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	current := q.dummyHead
	for current.Next.Next != nil {
		current = current.Next
	}
	result := current.Next.Val
	current.Next = nil
	return result, true
}

// PeekFront returns the item at the front of the list.
func (q *SinglyLinkedList[T]) PeekFront() (T, bool) {
	if q.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	return q.dummyHead.Next.Val, true
}

// PeekLast returns the item at the end of the list.
func (q *SinglyLinkedList[T]) PeekLast() (T, bool) {
	if q.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	current := q.dummyHead
	for current.Next != nil {
		current = current.Next
	}
	return current.Val, true
}

// Get returns the item at a given index of the list.
func (d *SinglyLinkedList[T]) Get(index int) (T, bool) {
	if d.dummyHead.Next == nil {
		var zeroValue T
		return zeroValue, false
	}
	current := d.dummyHead.Next
	currentIndex := 0
	for currentIndex < index {
		currentIndex += 1
		current = current.Next
	}
	if current == nil {
		var zeroValue T
		return zeroValue, false
	}
	return current.Val, true
}

// Map applies a given function to each item in the list.
func (d *SinglyLinkedList[T]) Map(f func(T) T) *SinglyLinkedList[T] {
	result := NewSinglyLinkedList[T]()
	current := d.dummyHead.Next
	for current != nil {
		result.InsertLast(f(current.Val))
		current = current.Next
	}
	return result
}

// Filter applies a given predicate function to each item in the list and returns a linked list of
// all items for which the predicate is true.
func (d *SinglyLinkedList[T]) Filter(f func(T) bool) *SinglyLinkedList[T] {
	result := NewSinglyLinkedList[T]()
	current := d.dummyHead.Next
	for current != nil {
		if f(current.Val) {
			result.InsertLast(current.Val)
		}
		current = current.Next
	}
	return result
}
