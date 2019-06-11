package binary_tree

import (
    "testing"
    "github.com/ilushka/data_structures_go/linked_list"
)

/*
func TestInsert(t *testing.T) {
    expected := [...]int{10, 100, 12, 122, 2, 31, 333}
    bt := new(BinaryTree)
    for _, v := range expected {
        bt.Insert(v)
    }
}
*/

func TestPreorder(t *testing.T) {
    initial := [...]int{10, 100, 12, 122, 2, 31, 333, 9, 2}
    bt := new(BinaryTree)
    for _, v := range initial {
        bt.Insert(v)
    }
    ll := new(linked_list.LinkedList)
    bt.Preorder(func (v int) {
        ll.Append(v)
    })
    expected := [...]int{10, 2, 9, 2, 100, 12, 31, 122, 333}
    result := *ll.ToArray()
    for ii, vv := range result {
        if vv != expected[ii] {
            t.Errorf("result: %v vs expected: %v", result, expected) 
            break
        }
    }
}

