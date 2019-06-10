package linked_list

import (
    "testing"
)

func TestCreateFromSlice(t *testing.T) {
    testval := [...]int{100, 200, 300, 400, 500, 600, 700, 800, 900}
    ll := CreateFromSlice(testval[:])
    if ll.Length != len(testval) {
        t.Errorf("Linked list should have %d nodes", len(testval))
    }
    for ii, vv := range *ll.ToArray() {
        if vv != testval[ii] {
            t.Errorf("Linked list's node #%d should have value of %d", ii, testval[ii])
        }
    }
}

func TestAppends(t *testing.T) {
    ll := new(LinkedList)
    if ll.Length != 0 {
        t.Errorf("Linked list should be of length 0")
    }

    // Node at index 5 does not exist
    err := ll.AppendAtIndex(555, 5)
    if err == nil {
        t.Errorf("Should not be able to append node at index 5 in list of length %d", ll.Length)
    }

    // Add fist node
    ll.Append(666)
    if ll.Length != 1 {
        t.Errorf("Linked list should be of length 1")
    }
    val, err := ll.Get(0)
    if val != 666 {
        t.Errorf("Value at index 0 should be 666")
    }

    // Add second node specifically at index 1
    err = ll.AppendAtIndex(777, 1)
    if err != nil {
        t.Errorf("Failed to append at index 1: %v", err)
    }
    if ll.Length != 2 {
        t.Errorf("Linked list should be of length 2")
    }
    val, err = ll.Get(1)
    if val != 777 {
        t.Errorf("Value at index 1 should be 777")
    }

    // Append new head
    err = ll.AppendAtIndex(555, 0)
    if err != nil {
        t.Errorf("Failed to append at index 0: %v", err)
    }
    if ll.Length != 3 {
        t.Errorf("Linked list should be of length 3")
    }
    testval := [...]int{555, 666, 777}
    for ii, vv := range *ll.ToArray() {
        if vv != testval[ii] {
            t.Errorf("Linked list's node #%d should have value of %d", ii, testval[ii])
        }
    }
}

func TestQuickSort(t *testing.T) {
    expected := [...]int{-300, 200, 400, 500, 800, 880, 900, 950, 1000}
    ll := CreateFromSlice([]int{1000, 200, -300, 400, 500, 950, 800, 880, 900})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("Sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}

func TestQuickSortSingleItem(t *testing.T) {
    expected := [...]int{1000}
    ll := CreateFromSlice([]int{1000})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("Sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}

func TestQuickSortEqualValues(t *testing.T) {
    expected := [...]int{10, 10, 1000, 1200, 1200, 1220, 1500, 1500}
    ll := CreateFromSlice([]int{1200, 1220, 1200, 1000, 1500, 1500, 10, 10})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("Sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}

func TestQuickSortEqualValues2(t *testing.T) {
    expected := [...]int{1200, 1200, 1200}
    ll := CreateFromSlice([]int{1200, 1200, 1200})
    ll.QuickSort()
    for ii, vv := range *ll.ToArray() {
        if vv != expected[ii] {
            t.Errorf("Sorted linked list's node #%d should have value of %d", ii, expected[ii])
        }
    }
}
