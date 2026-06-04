// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”).
// All CloudWeGo Modifications are Copyright 2022 CloudWeGo authors.

//go:build aix || darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris

package netpoll

import (
	"context"
	"errors"
	"net"
	"syscall"
	"time"
)

// nonDeadline and noCancel are just zero values for
// readability with functions taking too many parameters.
var noDeadline = time.Time{}

type netFD struct {
	// file descriptor
	fd int
	// When calling netFD.dial(), fd will be registered into poll in some scenarios, such as dialing tcp socket,
	// but not in other scenarios, such as dialing unix socket.
	// This leads to a different behavior in register poller at after, so use this field to mark it.
	pd *pollDesc
	// closed marks whether fd has expired
	closed uint32
	// Whether this is a streaming descriptor. Immutable.
	isStream bool
	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
	zeroReadIsEOF bool
	family        int    // AF_INET, AF_INET6, syscall.AF_UNIX
	sotype        int    // syscall.SOCK_STREAM, syscall.SOCK_DGRAM, syscall.SOCK_RAW
	isConnected   bool   // handshake completed or use of association with peer
	network       string // tcp, tcp4, tcp6, unix, unixgram, unixpacket
	localAddr     net.Addr
	remoteAddr    net.Addr
	// for detaching conn from poller
	detaching bool
}

func newNetFD(fd, family, sotype int, net string) *netFD { _ = "STUB: not implemented"; return nil }

// if dial connection error, you need exec netFD.Close actively
func (c *netFD) dial(ctx context.Context, laddr, raddr sockaddr) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// bind local address

// remote address from the user
// remote address we actually connected to

// remote address we actually connected to

// Record the local and remote addresses from the actual socket.
// Get the local address by calling Getsockname.
// For the remote address, use
// 1) the one returned by the connect method, if any; or
// 2) the one from Getpeername, if it succeeds; or
// 3) the one passed to us as the raddr parameter.

func (c *netFD) connect(ctx context.Context, la, ra syscall.Sockaddr) (rsa syscall.Sockaddr, retErr error) {
	_ = "STUB: not implemented"
	// Do not need to call c.writing here,
	// because c is not yet accessible to user,
	// so no concurrent operations are possible.
	return *new(syscall.Sockaddr), nil
}

// On Solaris we can see EINVAL if the socket has
// already been accepted and closed by the server.
// Treat this as a successful connection--writes to
// the socket will see EOF.  For details and a test
// case in C see https://golang.org/issue/6828.

// free operator to avoid leak

// Performing multiple connect system calls on a
// non-blocking socket under Unix variants does not
// necessarily result in earlier errors being
// returned. Instead, once runtime-integrated network
// poller tells us that the socket is ready, get the
// SO_ERROR socket option to see if the connection
// succeeded or failed. See issue 7474 for further
// details.

// The runtime poller can wake us up spuriously;
// see issues 14548 and 19289. Check that we are
// really connected; if not, wait again.

// Various errors contained in OpError.
var (
	errMissingAddress = errors.New("missing address")
	errCanceled       = errors.New("operation was canceled")
	errIOTimeout      = errors.New("i/o timeout")
)

// mapErr maps from the context errors to the historical internal net
// error values.
//
// TODO(bradfitz): get rid of this after adjusting tests and making
// context.DeadlineExceeded implement net.Error?
func mapErr(err error) error { _ = "STUB: not implemented"; return nil }
