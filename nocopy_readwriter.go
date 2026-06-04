// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netpoll

import (
	"io"
)

const maxReadCycle = 16

func newZCReader(r io.Reader) *zcReader { _ = "STUB: not implemented"; return nil }

var _ Reader = &zcReader{}

// zcReader implements Reader.
type zcReader struct {
	r   io.Reader
	buf *LinkBuffer
}

// Next implements Reader.
func (r *zcReader) Next(n int) (p []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// Peek implements Reader.
func (r *zcReader) Peek(n int) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// Skip implements Reader.
func (r *zcReader) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

// Release implements Reader.
func (r *zcReader) Release() (err error) { _ = "STUB: not implemented"; return nil }

// Slice implements Reader.
func (r *zcReader) Slice(n int) (reader Reader, err error) {
	_ = "STUB: not implemented"
	return *new(Reader), nil
}

// Len implements Reader.
func (r *zcReader) Len() (length int) {
	_ = "STUB: not implemented"

	// ReadString implements Reader.
	return 0
}

func (r *zcReader) ReadString(n int) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ReadBinary implements Reader.
func (r *zcReader) ReadBinary(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadByte implements Reader.
func (r *zcReader) ReadByte() (b byte, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *zcReader) Until(delim byte) (line []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *zcReader) waitRead(n int) (err error) { _ = "STUB: not implemented"; return nil }

// fill buffer to greater than n, range no more than 16 times.
func (r *zcReader) fill(n int) (err error) { _ = "STUB: not implemented"; return nil }

func newZCWriter(w io.Writer) *zcWriter { _ = "STUB: not implemented"; return nil }

var _ Writer = &zcWriter{}

// zcWriter implements Writer.
type zcWriter struct {
	w   io.Writer
	buf *LinkBuffer
}

// Malloc implements Writer.
func (w *zcWriter) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil,

		// MallocLen implements Writer.
		nil
}

func (w *zcWriter) MallocLen() (length int) { _ = "STUB: not implemented"; return 0 }

// Flush implements Writer.
func (w *zcWriter) Flush() (err error) { _ = "STUB: not implemented"; return nil }

// MallocAck implements Writer.
func (w *zcWriter) MallocAck(n int) (err error) { _ = "STUB: not implemented"; return nil }

// Append implements Writer.
func (w *zcWriter) Append(w2 Writer) (err error) { _ = "STUB: not implemented"; return nil }

// WriteString implements Writer.
func (w *zcWriter) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil

	// WriteBinary implements Writer.
}

func (w *zcWriter) WriteBinary(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil

	// WriteDirect implements Writer.
}

func (w *zcWriter) WriteDirect(p []byte, remainCap int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteByte implements Writer.
func (w *zcWriter) WriteByte(b byte) (err error) { _ = "STUB: not implemented"; return nil }

// zcWriter implements ReadWriter.
type zcReadWriter struct {
	*zcReader
	*zcWriter
}

func newIOReader(r Reader) *ioReader { _ = "STUB: not implemented"; return nil }

var _ io.Reader = &ioReader{}

// ioReader implements io.Reader.
//
// Deprecated: connection already implements Read directly with optimized buffer access.
// This wrapper exists only for external Reader implementations.
type ioReader struct {
	r Reader
}

// Read implements io.Reader.
//
// BUG: Read calls Release which invalidates any slices previously returned by Next or Peek
// on the same Reader. Do not mix Next/Peek and Read on the same Reader without first
// calling Release.
func (r *ioReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// read min(len(p), buffer.Len)

func newIOWriter(w Writer) *ioWriter { _ = "STUB: not implemented"; return nil }

var _ io.Writer = &ioWriter{}

// ioWriter implements io.Writer.
type ioWriter struct {
	w Writer
}

// Write implements io.Writer.
func (w *ioWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// ioReadWriter implements io.ReadWriter.
type ioReadWriter struct {
	io.Reader
	io.Writer
}
