package queue

import (
    "testing"
)

func TestQueue(t *testing.T) {
    q := new(Queue)
    q.Push(555)
    q.Push(666)
    q.Push(777)
    q.Push(888)
    if q.Length != 4 {
        t.Errorf("Queue length should be 4")
    }
    if val, err := q.Pop(); err != nil || val != 555 {
        t.Errorf("Next element in queue should have been 555")
    }
    if val, err := q.Pop(); err != nil || val != 666 {
        t.Errorf("Next element in queue should have been 666")
    }
    if val, err := q.Pop(); err != nil || val != 777 {
        t.Errorf("Next element in queue should have been 777")
    }
    if q.Length != 1 {
        t.Errorf("Queue length should be 1")
    }

    // Try poping of empty queue
    q.Pop()
    if _, err := q.Pop(); err == nil {
        t.Errorf("Should not be able to pop off element from empty queue")
    }
}

