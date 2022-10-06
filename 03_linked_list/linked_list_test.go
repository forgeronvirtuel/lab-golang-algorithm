package linkedlist

import (
	"bytes"
	"math/rand"
	"testing"
)

const dataSize = 2_000_000

func generateBytes() []byte {
	token := make([]byte, 4)
	rand.Read(token)
	return token
}

func TestBytesLinkedListItem_Add(t *testing.T) {
	var ll *BytesLinkedList

	ll = ll.Add([]byte("A"))
	if bytes.Compare(ll.last.msg, []byte("A")) != 0 {
		t.Errorf("Expected message to be A, got %s", string(ll.last.msg))
	}

	ll = ll.Add([]byte("B"))
	if bytes.Compare(ll.last.msg, []byte("B")) != 0 {
		t.Errorf("Expected message to be B, got %s", string(ll.last.msg))
	}
}

func TestBytesLinkedListItem_Get(t *testing.T) {
	var ll *BytesLinkedList

	msg1 := generateBytes()
	msg2 := generateBytes()

	lli := &bytesLinkedListItem{
		msg: msg2,
		ll: &bytesLinkedListItem{
			msg: msg1,
			ll:  nil,
		},
	}

	ll = &BytesLinkedList{
		last:   lli,
		length: 2,
	}

	if bytes.Compare(ll.Get(), msg2) != 0 {
		t.Fatalf("Not the same message")
	}
	if ll.Length() != 1 {
		t.Fatalf("not the same length: got %d, expected %d", ll.Length(), 1)
	}
}

func BenchmarkLinkedListItem_Add(b *testing.B) {
	var ll *BytesLinkedList

	data := make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = generateBytes()
	}

	ll = ll.Add(data[0])

	b.ResetTimer()
	b.StartTimer()
	for i := 1; i < dataSize; i++ {
		ll = ll.Add(data[i])
	}
	b.StopTimer()
}

func BenchmarkLinkedListItem_Get(b *testing.B) {
	var ll *BytesLinkedList

	for i := 0; i < dataSize; i++ {
		ll = ll.Add(generateBytes())
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 1; i < dataSize; i++ {
		ll.Get()
	}
	b.StopTimer()
}

func TestBytesLinkedListItem_Length(t *testing.T) {
	var ll *BytesLinkedList
	if ll.Length() != 0 {
		t.Errorf("length %d, expected 0", ll.Length())
	}

	ll = ll.Add(generateBytes())
	if ll.Length() != 1 {
		t.Errorf("Got length %d, expected 1", ll.Length())
	}

	ll = ll.Add(generateBytes())
	if ll.Length() != 2 {
		t.Errorf("Got length %d, expected 2", ll.Length())
	}
}
