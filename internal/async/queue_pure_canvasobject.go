// Code generated by go run gen.go; DO NOT EDIT.
//go:build js
// +build js

package async

import (
	"github.com/Anwert/fyne/v2"
)

// CanvasObjectQueue implements lock-free FIFO freelist based queue.
//
// Reference: https://dl.acm.org/citation.cfm?doid=248052.248106
type CanvasObjectQueue struct {
	head *itemCanvasObject
	tail *itemCanvasObject
	len  uint64
}

// NewCanvasObjectQueue returns a queue for caching values.
func NewCanvasObjectQueue() *CanvasObjectQueue {
	head := &itemCanvasObject{next: nil, v: nil}
	return &CanvasObjectQueue{
		tail: head,
		head: head,
	}
}

type itemCanvasObject struct {
	next *itemCanvasObject
	v    fyne.CanvasObject
}

func loadCanvasObjectItem(p **itemCanvasObject) *itemCanvasObject {
	return *p
}

func casCanvasObjectItem(p **itemCanvasObject, _, new *itemCanvasObject) bool {
	*p = new
	return true
}
