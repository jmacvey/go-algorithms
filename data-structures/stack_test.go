package ds

import "testing"

func TestPush(t *testing.T) {
	s := MakeStack()
	for i := 0; i < 50; i++ {
		s.Push(i)

		j, err := s.Peek()
		if j != i || err != nil {
			t.Errorf("Peek() failed: Expected %v, was %v", i, j)
		}
	}
}

func TestPop(t *testing.T) {
	s := MakeStack()
	for i := 0; i < 50; i++ {
		s.Push(i)
		j, err := s.Pop()
		if j != i || err != nil {
			t.Errorf("Pop() failed: Expected %v, was %v", i, j)
		}
	}

	if (!s.IsEmpty()) {
		t.Error("IsEmpty() failed: Expected true.  Got false.")
	}

	_, err := s.Pop()
	if (err == nil) {
		t.Error("Pop() failed. Expected error. Got nil.")
	}

	_, err = s.Peek()
	if (err == nil) {
		t.Error("Peek() failed. Expected error. Got nil.")
	}
}