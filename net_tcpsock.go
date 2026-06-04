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

// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	net.TCPAddr
}

func (a *TCPAddr) isWildcard() bool { _ = "STUB: not implemented"; return false }

func (a *TCPAddr) opAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (a *TCPAddr) family() int { _ = "STUB: not implemented"; return 0 }

func (a *TCPAddr) sockaddr(family int) (syscall.Sockaddr, error) {
	_ = "STUB: not implemented"
	return *new(syscall.Sockaddr), nil
}

func (a *TCPAddr) toLocal(network string) sockaddr {
	_ = "STUB: not implemented"
	return *new(sockaddr)
}

func loopbackIP(network string) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func ipToSockaddr(family int, ip net.IP, port int, zone string) (syscall.Sockaddr, error) {
	_ = "STUB: not implemented"
	return *new(syscall.Sockaddr), nil
}

// In general, an IP wildcard address, which is either
// "0.0.0.0" or "::", means the entire IP addressing
// space. For some historical reason, it is used to
// specify "any available address" on some operations
// of IP node.
//
// When the IP node supports IPv4-mapped IPv6 address,
// we allow an listener to listen to the wildcard
// address of both IP addressing spaces by specifying
// IPv6 wildcard address.

// We accept any IPv6 address including IPv4-mapped
// IPv6 address.

// TODO: sa := &syscall.SockaddrInet6{Port: port, ZoneId: uint32(zoneCache.index(zone))}

// ResolveTCPAddr returns an address of TCP end point.
//
// The network must be a TCP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveTCPAddr resolves the
// address to an address of TCP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveTCPAddr(network, address string) (*TCPAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TCPConnection implements Connection.
type TCPConnection struct {
	connection
}

// newTCPConnection wraps *TCPConnection.
func newTCPConnection(conn Conn) (connection *TCPConnection, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DialTCP acts like Dial for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialTCP(ctx context.Context, network string, laddr, raddr *TCPAddr) (*TCPConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sd *sysDialer) dialTCP(ctx context.Context, laddr, raddr *TCPAddr) (*TCPConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TCP has a rarely used mechanism called a 'simultaneous connection' in
// which Dial("tcp", addr1, addr2) run on the machine at addr1 can
// connect to a simultaneous Dial("tcp", addr2, addr1) run on the machine
// at addr2, without either machine executing Listen. If laddr == nil,
// it means we want the kernel to pick an appropriate originating local
// address. Some Linux kernels cycle blindly through a fixed range of
// local ports, regardless of destination port. If a kernel happens to
// pick local port 50001 as the source for a Dial("tcp", "", "localhost:50001"),
// then the Dial will succeed, having simultaneously connected to itself.
// This can only happen when we are letting the kernel pick a port (laddr == nil)
// and when there is no listener for the destination address.
// It's hard to argue this is anything other than a kernel bug. If we
// see this happen, rather than expose the buggy effect to users, we
// close the conn and try again. If it happens twice more, we relent and
// use the result. See also:
// 	https://golang.org/issue/2690
// 	https://stackoverflow.com/questions/4949858/
//
// The opposite can also happen: if we ask the kernel to pick an appropriate
// originating local address, sometimes it picks one that is already in use.
// So if the error is EADDRNOTAVAIL, we have to try again too, just for
// a different reason.
//
// The kernel socket code is no doubt enjoying watching us squirm.

func selfConnect(conn *netFD, err error) bool {
	_ = "STUB: not implemented"
	// If the connect failed, we clearly didn't connect to ourselves.
	return false
}

// The socket constructor can return an conn with raddr nil under certain
// unknown conditions. The errors in the calls there to Getpeername
// are discarded, but we can't catch the problem there because those
// calls are sometimes legally erroneous with a "socket not connected".
// Since this code (selfConnect) is already trying to work around
// a problem, we make sure if this happens we recognize trouble and
// ask the DialTCP routine to try again.
// TODO: try to understand what's really going on.

func spuriousENOTAVAIL(err error) bool { _ = "STUB: not implemented"; return false }
