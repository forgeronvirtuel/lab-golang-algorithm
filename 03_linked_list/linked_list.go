package linkedlist

import "fmt"

type LinkedListError struct {
	msg       string
	operation string
}

func (err *LinkedListError) Error() string {
	return fmt.Sprintf("%s: %s", err.operation, err.msg)
}

type bytesLinkedListItem struct {
	msg []byte
	ll  *bytesLinkedListItem
}

func (ll *bytesLinkedListItem) add(msg []byte) *bytesLinkedListItem {
	if ll == nil {
		return &bytesLinkedListItem{
			msg: msg,
			ll:  nil,
		}
	}

	return &bytesLinkedListItem{msg: msg, ll: ll}
}

func (ll *bytesLinkedListItem) length() int {
	if ll == nil {
		return 0
	}
	var i int
	for i = 1; ll.ll != nil; i++ {
		ll = ll.ll
	}
	return i
}

type BytesLinkedList struct {
	last   *bytesLinkedListItem
	length int
}

func (ll *BytesLinkedList) Add(msg []byte) *BytesLinkedList {
	if ll == nil {
		return &BytesLinkedList{
			last: &bytesLinkedListItem{
				msg: msg,
				ll:  nil,
			},
			length: 1,
		}
	}

	ll.last = ll.last.add(msg)
	ll.length += 1
	return ll
}

func (ll *BytesLinkedList) Get() []byte {
	msg := ll.last.msg
	ll.last = ll.last.ll
	ll.length -= 1
	return msg
}

func (ll *BytesLinkedList) Length() int {
	if ll == nil {
		return 0
	}
	return ll.length
}
