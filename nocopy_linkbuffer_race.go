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

//go:build race

package netpoll

import (
	"sync"
)

type LinkBuffer = SafeLinkBuffer

// SafeLinkBuffer only used to in go tests with -race
type SafeLinkBuffer struct {
	sync.Mutex
	UnsafeLinkBuffer
}

// ------------------------------------------ implement copy reader ------------------------------------------

func (b *SafeLinkBuffer) readCopy(p []byte) int { _ = "STUB: not implemented"; return 0 }

// ------------------------------------------ implement zero-copy reader ------------------------------------------

// Next implements Reader.
func (b *SafeLinkBuffer) Next(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Peek implements Reader.
func (b *SafeLinkBuffer) Peek(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip implements Reader.
func (b *SafeLinkBuffer) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

// Until implements Reader.
func (b *SafeLinkBuffer) Until(delim byte) (line []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Release implements Reader.
func (b *SafeLinkBuffer) Release() (err error) { _ = "STUB: not implemented"; return nil }

// ReadString implements Reader.
func (b *SafeLinkBuffer) ReadString(n int) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ReadBinary implements Reader.
func (b *SafeLinkBuffer) ReadBinary(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadByte implements Reader.
func (b *SafeLinkBuffer) ReadByte() (p byte, err error) { _ = "STUB: not implemented"; return 0, nil }

// Slice implements Reader.
func (b *SafeLinkBuffer) Slice(n int) (r Reader, err error) {
	_ = "STUB: not implemented"
	return *new(Reader), nil
}

// ------------------------------------------ implement zero-copy writer ------------------------------------------

// Malloc implements Writer.
func (b *SafeLinkBuffer) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MallocLen implements Writer.
func (b *SafeLinkBuffer) MallocLen() (length int) { _ = "STUB: not implemented"; return 0 }

// MallocAck implements Writer.
func (b *SafeLinkBuffer) MallocAck(n int) (err error) { _ = "STUB: not implemented"; return nil }

// Flush implements Writer.
func (b *SafeLinkBuffer) Flush() (err error) { _ = "STUB: not implemented"; return nil }

// Append implements Writer.
func (b *SafeLinkBuffer) Append(w Writer) (err error) { _ = "STUB: not implemented"; return nil }

// WriteBuffer implements Writer.
func (b *SafeLinkBuffer) WriteBuffer(buf *LinkBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// WriteString implements Writer.
func (b *SafeLinkBuffer) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteBinary implements Writer.
func (b *SafeLinkBuffer) WriteBinary(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteDirect cannot be mixed with WriteString or WriteBinary functions.
func (b *SafeLinkBuffer) WriteDirect(p []byte, remainLen int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteByte implements Writer.
func (b *SafeLinkBuffer) WriteByte(p byte) (err error) { _ = "STUB: not implemented"; return nil }

// Close will recycle all buffer.
func (b *SafeLinkBuffer) Close() (err error) { _ = "STUB: not implemented"; return nil }

// ------------------------------------------ implement connection interface ------------------------------------------

// Bytes returns all the readable bytes of this SafeLinkBuffer.
func (b *SafeLinkBuffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// GetBytes will read and fill the slice p as much as possible.
func (b *SafeLinkBuffer) GetBytes(p [][]byte) (vs [][]byte) { _ = "STUB: not implemented"; return nil }

// book will grow and malloc buffer to hold data.
//
// bookSize: The size of data that can be read at once.
// maxSize: The maximum size of data between two Release(). In some cases, this can
//
//	guarantee all data allocated in one node to reduce copy.
func (b *SafeLinkBuffer) book(bookSize, maxSize int) (p []byte) {
	_ = "STUB: not implemented"
	return nil
}

// bookAck will ack the first n malloc bytes and discard the rest.
//
// length: The size of data in inputBuffer. It is used to calculate the maxSize
func (b *SafeLinkBuffer) bookAck(n int) (length int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// calcMaxSize will calculate the data size between two Release()
func (b *SafeLinkBuffer) calcMaxSize() (sum int) { _ = "STUB: not implemented"; return 0 }

func (b *SafeLinkBuffer) resetTail(maxSize int) { _ = "STUB: not implemented"; return }

func (b *SafeLinkBuffer) indexByte(c byte, skip int) int { _ = "STUB: not implemented"; return 0 }
