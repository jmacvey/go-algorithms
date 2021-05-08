package ds

import "testing"

func TestLinkedList(t *testing.T) {
	ll := MakeLinkedList()

	for i := 0; i < 10; i++ {
		ll.Add(i)
		value, err := ll.Get(i)
		if value != i || err != nil {
			t.Errorf("Get(%v) failed: Expected %v, Got %v", value, value, i)
		}
	}

	for i := 0; i < 9; i++ {
		// insert 0, 1 -> 0, 0, 1, 2, 3
		ll.Insert(i * 10, i * 2 + 1)
	}

	for i := 0; i < 18; i += 2 {
		value, err := ll.Get(i)
		value2, err2 := ll.Get(i + 1)

		if value * 10 != value2 || err != nil || err2 != nil {
			t.Errorf("Get(%v) * 10 ==  Get(%v) failed, Get(%v) = %v, Get(%v) = %v", i, i + 1, i, value, i + 1, value2)
		} 
	}
}

func TestRemove(t *testing.T) {
	ll := MakeLinkedList()

	for i := 0; i < 10; i++ {
		ll.Add(i)
	}

	for i := 9; i >= 0; i-- {
		v, err := ll.Remove(i)

		if (v != i || err != nil) {
			t.Errorf("Remove(%v) failed: Expected %v, Got %v", i, i, v)
		}
	}

	for i := 0; i < 10; i++ {
		ll.Add(i)
	}

	for i := 0; i < 5; i++ {
		v, err := ll.Remove(5)
		if (v != i + 5 || err != nil) {
			t.Errorf("Remove(%v) failed: Expected %v, Got %v", i + 5, i + 5, v)
		}
	}

}

func TestReverse(t *testing.T) {
	ll := MakeLinkedList()

	for i := 0; i < 10; i++ {
		ll.Add(i)
		v, err := ll.Get(i)

		if v != i || err != nil {
			t.Errorf("Get(%v) failed. Expected %v, Got %v", i, i, v)
		}
	}

	ll.Reverse()
	for i := 0; i < 10; i++ {
		ll.Add(i)
		v, err := ll.Get(i)

		if v != 9 - i || err != nil {
			t.Errorf("Get(%v) failed. Expected %v, Got %v", i, 9 - i, v)
		}
	}

	ll.Clear()

	for i := 0; i < 2; i++ {
		ll.Add(i)
		v, err := ll.Get(i)

		if v != i || err != nil {
			t.Errorf("Get(%v) failed. Expected %v, Got %v", i, i, v)
		}
	}

	ll.Reverse()

	for i := 0; i < 2; i++ {
		ll.Add(i)
		v, err := ll.Get(i)

		if v != 1 - i || err != nil {
			t.Errorf("Get(%v) failed. Expected %v, Got %v", i, 1 - i, v)
		}
	}

}