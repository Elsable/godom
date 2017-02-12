// This file implements DOMRect interface
// https://developer.mozilla.org/en-US/docs/Web/API/DOMRect
package godom

import (
	"github.com/gopherjs/gopherjs/js"
)

type DOMRect struct {
	*js.Object
}

func (r *DOMRect) X() float64 {
	return r.Get("x").Float()
}

func (r *DOMRect) Y() float64 {
	return r.Get("y").Float()
}

func (r *DOMRect) Width() float64 {
	return r.Get("width").Float()
}

func (r *DOMRect) Height() float64 {
	return r.Get("height").Float()
}

func (r *DOMRect) Top() float64 {
	return r.Get("top").Float()
}

func (r *DOMRect) Right() float64 {
	return r.Get("right").Float()
}

func (r *DOMRect) Bottom() float64 {
	return r.Get("bottom").Float()
}

func (r *DOMRect) Left() float64 {
	return r.Get("left").Float()
}