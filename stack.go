package stack

import (
	"container/list"
)

type Stack struct {
	Data *list.List
}

func (stack *Stack) New() {
	stack.Data = list.New()
}

func (stack *Stack) Push(v interface{}) {
	stack.Data.PushBack(v)
}

func (stack *Stack) Pop() interface{} {
	tmp := stack.Data.Back()
	if tmp != nil {
		stack.Data.Remove(tmp)
		return tmp.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.Data.Len()
}
