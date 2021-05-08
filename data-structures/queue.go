package ds

import "errors"

type Queue struct {
	data     []int
	size     int
	capacity int
	front    int
	rear     int
}

func MakeQueue() Queue { 
	return Queue{
		data : make([]int, 10),
		size: 0,
		capacity: 10,
		front: 0,
		rear: 0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) AddFront(value int) {
	if q.capacity == q.size {
		q.resize()
	}

	q.front = q.front - 1
	if q.front == -1 {
		q.front = q.capacity - 1
	}

	q.data[q.front] = value
	q.size++
}

func (q *Queue) AddBack(value int) {
	if q.capacity == q.size {
		q.resize()
	}

	q.data[q.rear] = value
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

func (q *Queue) PeekFront() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("tried to peek from empty queue")
	}

	return q.data[q.front], nil
}

func (q *Queue) PeekBack() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("tried to peek from empty queue")
	}

	last := q.rear - 1
	if last == -1 {
		last = q.capacity - 1
	}

	return q.data[last], nil
}

func (q *Queue) PopBack() (int, error) {
	value, err := q.PeekBack()

	if err != nil {
		return 0, errors.New("tried to pop from empty queue")
	}

	q.rear--
	if q.rear == -1 {
		q.rear = q.capacity - 1
	}
	q.size--
	return value, nil
}

func (q *Queue) PopFront() (int, error) {
	value, err := q.PeekFront()

	if err != nil {
		return 0, errors.New("tried to peek from empty queue")
	}

	q.front = (q.front + 1) % q.capacity
	q.size--
	return value, nil
}

func (q *Queue) resize() {
	current := q.front
	data := make([]int, q.capacity*2)
	for i := 0; i < q.size; i++ {
		data[i] = q.data[current]
		current = (current + 1) % q.capacity
	}

	q.capacity *= 2
	q.front = 0
	q.rear = q.size
	q.data = data
}
