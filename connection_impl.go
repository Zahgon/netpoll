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

//go:build !windows

package netpoll

import (
	"sync"
	"syscall"
	"time"
)

type connState = int32

const (
	connStateNone         = 0
	connStateConnected    = 1
	connStateDisconnected = 2
)

// connection is the implementation of Connection
type connection struct {
	netFD
	onEvent
	locker
	operator      *FDOperator
	readTimeout   time.Duration
	readDeadline  int64 // UnixNano(). it overwrites readTimeout. 0 if not set.
	readTimer     *time.Timer
	readTrigger   chan error
	waitReadSize  int64
	writeTimeout  time.Duration
	writeDeadline int64 // UnixNano(). it overwrites writeTimeout. 0 if not set.
	writeTimer    *time.Timer
	writeTrigger  chan error
	inputBuffer   *LinkBuffer
	outputBuffer  *LinkBuffer
	outputBarrier *barrier
	maxSize       int       // The maximum size of data between two Release().
	bookSize      int       // The size of data that can be read at once.
	state         connState // Connection state should be changed sequentially.
}

var (
	_ Connection = &connection{}
	_ Reader     = &connection{}
	_ Writer     = &connection{}
)

// Reader implements Connection.
func (c *connection) Reader() Reader {
	_ = "STUB: not implemented"

	// Writer implements Connection.
	return *new(Reader)
}

func (c *connection) Writer() Writer {
	_ = "STUB: not implemented"

	// IsActive implements Connection.
	return *new(Writer)
}

func (c *connection) IsActive() bool { _ = "STUB: not implemented"; return false }

// SetIdleTimeout implements Connection.
func (c *connection) SetIdleTimeout(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// SetReadTimeout implements Connection.
func (c *connection) SetReadTimeout(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// SetWriteTimeout implements Connection.
func (c *connection) SetWriteTimeout(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// SetDeadline implements net.Conn.SetDeadline
func (c *connection) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline implements net.Conn.SetReadDeadline
func (c *connection) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline implements net.Conn.SetWriteDeadline
func (c *connection) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// ------------------------------------------ implement zero-copy reader ------------------------------------------

// Next implements Connection.
func (c *connection) Next(n int) (p []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// Peek implements Connection.
func (c *connection) Peek(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip implements Connection.
func (c *connection) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

// Release implements Connection.
func (c *connection) Release() (err error) {
	_ = "STUB: not implemented"
	// Check inputBuffer length first to reduce contention in mux situation.
	// c.operator.do competes with c.inputs/c.inputAck
	return nil
}

// Set the maximum value of maxsize equal to mallocMax to prevent GC pressure.

// Double check length to reset tail node

// Slice implements Connection.
func (c *connection) Slice(n int) (r Reader, err error) {
	_ = "STUB: not implemented"
	return *new(Reader), nil
}

// Len implements Connection.
func (c *connection) Len() (length int) { _ = "STUB: not implemented"; return 0 }

// Until implements Connection.
func (c *connection) Until(delim byte) (line []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return all the data in the buffer

// skip all exists bytes

// ReadString implements Connection.
func (c *connection) ReadString(n int) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ReadBinary implements Connection.
func (c *connection) ReadBinary(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadByte implements Connection.
func (c *connection) ReadByte() (b byte, err error) { _ = "STUB: not implemented"; return 0, nil }

// ------------------------------------------ implement zero-copy writer ------------------------------------------

// Malloc implements Connection.
func (c *connection) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MallocLen implements Connection.
func (c *connection) MallocLen() (length int) { _ = "STUB: not implemented"; return 0 }

// Flush will send all malloc data to the peer,
// so must confirm that the allocated bytes have been correctly assigned.
//
// Flush first checks whether the out buffer is empty.
// If empty, it will call syscall.Write to send data directly,
// otherwise the buffer will be sent asynchronously by the epoll trigger.
func (c *connection) Flush() error { _ = "STUB: not implemented"; return nil }

// MallocAck implements Connection.
func (c *connection) MallocAck(n int) (err error) { _ = "STUB: not implemented"; return nil }

// Append implements Connection.
func (c *connection) Append(w Writer) (err error) { _ = "STUB: not implemented"; return nil }

// WriteString implements Connection.
func (c *connection) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteBinary implements Connection.
func (c *connection) WriteBinary(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// WriteDirect implements Connection.
func (c *connection) WriteDirect(p []byte, remainCap int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// WriteByte implements Connection.
func (c *connection) WriteByte(b byte) (err error) { _ = "STUB: not implemented"; return nil }

// ------------------------------------------ implement net.Conn ------------------------------------------

// Read behavior is the same as net.Conn, it will return io.EOF if buffer is empty.
func (c *connection) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Write will Flush soon.
func (c *connection) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close implements Connection.
func (c *connection) Close() error {
	_ = "STUB: not implemented"

	// Detach detaches the connection from poller but doesn't close it.
	return nil
}

func (c *connection) Detach() error { _ = "STUB: not implemented"; return nil }

// ------------------------------------------ private ------------------------------------------

var barrierPool = sync.Pool{
	New: func() interface{} {
		return &barrier{
			bs:  make([][]byte, barriercap),
			ivs: make([]syscall.Iovec, barriercap),
		}
	},
}

// init initializes the connection with options
func (c *connection) init(conn Conn, opts *options) (err error) {
	// init buffer, barrier, finalizer
	c.readTrigger = make(chan error, 1)
	c.writeTrigger = make(chan error, 1)
	c.bookSize, c.maxSize = defaultLinkBufferSize, defaultLinkBufferSize
	c.inputBuffer, c.outputBuffer = NewLinkBuffer(defaultLinkBufferSize), NewLinkBuffer()
	c.outputBarrier = barrierPool.Get().(*barrier)
	c.state = connStateNone

	c.initNetFD(conn) // conn must be *netFD{}
	c.initFDOperator()
	c.initFinalizer()

	syscall.SetNonblock(c.fd, true)
	// enable TCP_NODELAY by default
	switch c.network {
	case "tcp", "tcp4", "tcp6":
		setTCPNoDelay(c.fd, true)
	}

	// connection initialized and prepare options
	return c.onPrepare(opts)
}

func (c *connection) initNetFD(conn Conn) { _ = "STUB: not implemented"; return }

func (c *connection) initFDOperator() { _ = "STUB: not implemented"; return }

func (c *connection) initFinalizer() { _ = "STUB: not implemented"; return }

func (c *connection) triggerRead(err error) { _ = "STUB: not implemented"; return }

func (c *connection) triggerWrite(err error) { _ = "STUB: not implemented"; return }

// waitRead will wait full n bytes.
func (c *connection) waitRead(n int) (err error) { _ = "STUB: not implemented"; return nil }

// wait full n

// waitReadWithTimeout will wait full n bytes or until timeout.
func (c *connection) waitReadWithTimeout(n int, timeout time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// cannot return directly, stop timer first!

// cannot return directly, stop timer first!

// double check if there is enough data to be read

// clean timer.C

// flush writes data directly.
func (c *connection) flush() error { _ = "STUB: not implemented"; return nil }

// return if write all buffer.

func (c *connection) waitFlush() (err error) { _ = "STUB: not implemented"; return nil }

// set write timeout

// clean timer

// try fetch writeTrigger if both cases fires

// if timeout, remove write event from poller
// we cannot flush it again, since we don't if the poller is still process outputBuffer

func (c *connection) getState() connState { _ = "STUB: not implemented"; return *new(connState) }

func (c *connection) setState(newState connState) { _ = "STUB: not implemented"; return }

func (c *connection) changeState(from, to connState) bool { _ = "STUB: not implemented"; return false }
