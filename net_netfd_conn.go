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

//go:build darwin || netbsd || freebsd || openbsd || dragonfly || linux

package netpoll

import (
	"net"
	"time"
)

var _ Conn = &netFD{}

// Fd implements Conn.
func (c *netFD) Fd() (fd int) {
	_ = "STUB: not implemented"

	// Read implements Conn.
	return 0
}

func (c *netFD) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Write implements Conn.
func (c *netFD) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close will be executed only once.
func (c *netFD) Close() (err error) { _ = "STUB: not implemented"; return nil }

// LocalAddr implements Conn.
func (c *netFD) LocalAddr() (addr net.Addr) {
	_ = "STUB: not implemented"

	// RemoteAddr implements Conn.
	return *new(net.Addr)
}

func (c *netFD) RemoteAddr() (addr net.Addr) {
	_ = "STUB: not implemented"
	return *

	// SetKeepAlive implements Conn.
	// TODO: only tcp conn is ok.
	new(net.Addr)
}

func (c *netFD) SetKeepAlive(second int) error { _ = "STUB: not implemented"; return nil }

// SetDeadline implements Conn.
func (c *netFD) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline implements Conn.
func (c *netFD) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline implements Conn.
func (c *netFD) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
