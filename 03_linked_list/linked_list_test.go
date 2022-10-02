package linkedlist

import "testing"

func TestStringLinkedListItem_Add(t *testing.T) {
	ll := &StringLinkedListItem{
		msg: "A",
		ll:  nil,
	}

	nll, err := ll.Add("B")
	if err != nil {
		t.Fatal(err)
	}

	if nll.msg != "B" {
		t.Errorf("Expected message to be B, got %s", ll.msg)
	}
}
