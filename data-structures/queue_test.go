package ds

import "testing"

func TestAddFront(t *testing.T) {
	q := MakeQueue()
	for i := 0; i < 50; i++ {
		q.AddFront(i)

		j, err := q.PeekFront()
		if j != i || err != nil {
			t.Errorf("Peek() failed: Expected %v, was %v", i, j)
		}
	}
}

func TestAddBack(t *testing.T) {
	q := MakeQueue()
	for i := 0; i < 50; i++ {
		q.AddBack(i)

		j, err := q.PeekBack()
		if j != i || err != nil {
			t.Errorf("Peek() failed: Expected %v, was %v", i, j)
		}
	}
}

func TestPopFront(t *testing.T) {
	q := MakeQueue()
	for i := 0; i < 50; i++ {
		q.AddFront(i)
		j, err := q.PopFront()
		if j != i || err != nil {
			t.Errorf("PopFront() failed: Expected %v, was %v", i, j)
		}
	}

	if (!q.IsEmpty()) {
		t.Error("IsEmpty() failed: Expected true.  Got false.")
	}

	_, err := q.PopFront()
	if (err == nil) {
		t.Error("PopFront() failed. Expected error. Got nil.")
	}

	_, err = q.PeekFront()
	if (err == nil) {
		t.Error("PeekFront() failed. Expected error. Got nil.")
	}
}

func TestPopBack(t *testing.T) {
	q := MakeQueue()
	for i := 0; i < 50; i++ {
		q.AddBack(i)
		j, err := q.PopBack()
		if j != i || err != nil {
			t.Errorf("PopBack() failed: Expected %v, was %v", i, j)
		}
	}

	if (!q.IsEmpty()) {
		t.Error("IsEmpty() failed: Expected true.  Got false.")
	}

	_, err := q.PopBack()
	if (err == nil) {
		t.Error("PopBack() failed. Expected error. Got nil.")
	}

	_, err = q.PeekBack()
	if (err == nil) {
		t.Error("PeekBack() failed. Expected error. Got nil.")
	}
}