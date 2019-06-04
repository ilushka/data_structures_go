package data_structures_go

import (
    "testing"
)

func TestAppendTail(t *testing.T) {
    testval := [...]int{100, 200, 300, 400, 500, 600, 700, 800, 900}
    ll := new(LinkedList)
    for _, v := range testval {
        ll.AppendTail(v)
    }
    if ll.size != len(testval) {
        t.Errorf("linked list should have %d nodes", len(testval))
    }
    for ii, vv := range *ll.ToArray() {
        if vv != testval[ii] {
            t.Errorf("linked list's node #%d should have value of %d", ii, testval[ii])
        }
    }
}

