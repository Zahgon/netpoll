// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”).
// All CloudWeGo Modifications are Copyright 2022 CloudWeGo authors.

//go:build !windows

package netpoll

import (
	"context"
	"net"
	"syscall"
)

// BUG(mikio): On JS, NaCl and Plan 9, methods and functions related
// to UnixConn and UnixListener are not implemented.

// BUG(mikio): On Windows, methods and functions related to UnixConn
// and UnixListener don't work for "unixgram" and "unixpacket".

// UnixAddr represents the address of a Unix domain socket end point.
type UnixAddr struct {
	net.UnixAddr
}

func (a *UnixAddr) isWildcard() bool { _ = "STUB: not implemented"; return false }

func (a *UnixAddr) opAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (a *UnixAddr) family() int { _ = "STUB: not implemented"; return 0 }

func (a *UnixAddr) sockaddr(family int) (syscall.Sockaddr, error) {
	_ = "STUB: not implemented"
	return *new(syscall.Sockaddr), nil
}

func (a *UnixAddr) toLocal(net string) sockaddr {
	_ = "STUB: not implemented"

	// ResolveUnixAddr returns an address of Unix domain socket end point.
	//
	// The network must be a Unix network name.
	//
	// See func Dial for a description of the network and address
	// parameters.
	return *new(sockaddr)
}

func ResolveUnixAddr(network, address string) (*UnixAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnixConnection implements Connection.
type UnixConnection struct {
	connection
}

// newUnixConnection wraps UnixConnection.
func newUnixConnection(conn Conn) (connection *UnixConnection, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DialUnix acts like Dial for Unix networks.
//
// The network must be a Unix network name; see func Dial for details.
//
// If laddr is non-nil, it is used as the local address for the
// connection.
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sd *sysDialer) dialUnix(ctx context.Context, laddr, raddr *UnixAddr) (*UnixConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unixSocket(ctx context.Context, network string, laddr, raddr sockaddr, mode string) (conn *netFD, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
