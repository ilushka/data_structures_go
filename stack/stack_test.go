package stack

import (
    "testing"
)

func TestStack(t *testing.T) {
    s := new(Stack)
    s.Push(555)
    s.Push(666)
    s.Push(777)
    s.Push(888)
    if s.Length != 4 {
        t.Errorf("Stack length should be 4")
    }
    if val, err := s.Pop(); err != nil || val != 888 {
        t.Errorf("Next element in stack should have been 888")
    }
    if val, err := s.Pop(); err != nil || val != 777 {
        t.Errorf("Next element in stack should have been 777")
    }
    if val, err := s.Pop(); err != nil || val != 666 {
        t.Errorf("Next element in stack should have been 666")
    }
    if s.Length != 1 {
        t.Errorf("Stack length should be 1")
    }

    // Try poping off of empty stack
    s.Pop()
    if _, err := s.Pop(); err == nil {
        t.Errorf("Should not be able to pop off element from empty stack")
    }
}


