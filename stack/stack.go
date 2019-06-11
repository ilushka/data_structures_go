package stack

import (
    "github.com/ilushka/data_structures_go/linked_list"
)

type Stack struct {
    linked_list.LinkedList
}

func (s *Stack) Push(value int) {
    s.AppendAtIndex(value, 0)
}

func (s *Stack) Pop() (int, error) {
    val, err := s.Get(0)
    if err != nil {
        return 0, err
    }
    s.Remove(0)
    return val, nil
}

