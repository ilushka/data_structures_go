package queue

import (
    "github.com/ilushka/data_structures_go/linked_list"
)

type Queue struct {
    linked_list.LinkedList
}

func (q *Queue) Push(value int) {
    q.Append(value)
}

func (q *Queue) Pop() (int, error) {
    val, err := q.Get(0)
    if err != nil {
        return 0, err
    }
    q.Remove(0)
    return val, nil
}

