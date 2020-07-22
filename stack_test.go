package stack

import "testing"

func TestLen(t *testing.T) {
	c := Stack{}
	c.New()

	num := 5

	for i := 0; i < num; i++ {
		c.Push(i)
	}

	got := c.Len()
	want := num
	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}

func TestPopPush(t *testing.T) {
	c := Stack{}
	c.New()

	num := 5

	for i := 0; i < num; i++ {
		c.Push(i)
	}

	for i := num - 1; i >= 0; i-- {
		tmp := c.Pop()
		got := tmp.(int)
		want := i

		if want != got {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	}

	got := c.Len()
	want := 0
	if want != got {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}
