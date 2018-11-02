// Copyright 2017 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rbytes // import "go-hep.org/x/hep/groot/rbytes"

import (
	"go-hep.org/x/hep/groot/root"
)

// StreamerInfo describes a ROOT Streamer.
type StreamerInfo interface {
	root.Named
	CheckSum() int
	ClassVersion() int
	Elements() []StreamerElement
}

// StreamerElement describes a ROOT StreamerElement
type StreamerElement interface {
	root.Named
	ArrayDim() int
	ArrayLen() int
	Type() int
	Offset() uintptr
	Size() uintptr
	TypeName() string
}

type StreamerInfoContext interface {
	StreamerInfo(name string) (StreamerInfo, error)
}

// Unmarshaler is the interface implemented by an object that can
// unmarshal itself from a ROOT buffer
type Unmarshaler interface {
	UnmarshalROOT(r *RBuffer) error
}

// Marshaler is the interface implemented by an object that can
// marshal itself into a ROOT buffer
type Marshaler interface {
	MarshalROOT(w *WBuffer) (int, error)
}
