package linkedlist

import "fmt"

type LinkedListError struct {
	msg       string
	operation string
}

func (err *LinkedListError) Error() string {
	return fmt.Sprintf("%s: %s", err.operation, err.msg)
}

type StringLinkedListItem struct {
	msg string
	ll  *StringLinkedListItem
}

func (ll *StringLinkedListItem) Add(msg string) (*StringLinkedListItem, error) {
	// defensive add
	if ll.ll != nil {
		return nil, &LinkedListError{
			msg:       "Already an item present",
			operation: "Add",
		}
	}
	ll.ll = &StringLinkedListItem{msg: msg}
	return ll.ll, nil
}
