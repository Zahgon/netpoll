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
	"os"
)

// CreateListener return a new Listener.
func CreateListener(network, addr string) (l Listener, err error) {
	_ = "STUB: not implemented"
	return *new(Listener), nil
}

// tcp, tcp4, tcp6, unix

// ConvertListener converts net.Listener to Listener
func ConvertListener(l net.Listener) (nl Listener, err error) {
	_ = "STUB: not implemented"
	return *new(Listener), nil
}

var _ net.Listener = &listener{}

type listener struct {
	fd   int
	addr net.Addr     // listener's local addr
	ln   net.Listener // tcp|unix listener
	file *os.File
}

// Accept implements Listener.
func (ln *listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

/* https://man7.org/linux/man-pages/man2/accept.2.html
EAGAIN or EWOULDBLOCK
  The socket is marked nonblocking and no connections are
  present to be accepted.  POSIX.1-2001 and POSIX.1-2008
  allow either error to be returned for this case, and do
  not require these constants to have the same value, so a
  portable application should check for both possibilities.
*/

// Close implements Listener.
func (ln *listener) Close() error { _ = "STUB: not implemented"; return nil }

// Addr implements Listener.
func (ln *listener) Addr() net.Addr {
	_ = "STUB: not implemented"

	// Fd implements Listener.
	return *new(net.Addr)
}

func (ln *listener) Fd() (fd int) { _ = "STUB: not implemented"; return 0 }

func (ln *listener) parseFD() (err error) { _ = "STUB: not implemented"; return nil }
