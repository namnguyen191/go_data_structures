package main

import "errors"

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) enqueue(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) dequeue(value T) (T, error) {
	var defaultValue T
	if q.isEmpty() {
		return defaultValue, errors.New("Queue is empty")
	}

	returnVal := q.data[0]
	q.data = q.data[1:]

	return returnVal, nil
}

func (q *Queue[T]) peek() (T, error) {
	var defaultValue T
	if q.isEmpty() {
		return defaultValue, errors.New("Queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue[T]) isEmpty() bool {
	return len(q.data) == 0
}
