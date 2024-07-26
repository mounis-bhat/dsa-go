package queue

import "errors"

type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

var ErrQueueEmpty = errors.New("Queue is empty")

func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, ErrQueueEmpty
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, ErrQueueEmpty
	}
	return q.elements[0], nil
}

func (q *Queue) Display() []int {
	return q.elements
}
