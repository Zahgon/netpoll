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

// A sockaddr represents a TCP, IP or Unix network endpoint
// address that can be converted into a syscall.Sockaddr.
type sockaddr interface {
	net.Addr

	// family returns the platform-dependent address family
	// identifier.
	family() int

	// isWildcard reports whether the address is a wildcard
	// address.
	isWildcard() bool

	// sockaddr returns the address converted into a syscall
	// sockaddr type that implements syscall.Sockaddr
	// interface. It returns a nil interface when the address is nil.
	sockaddr(family int) (syscall.Sockaddr, error)

	// toLocal maps the zero address to a local system address (127.0.0.1 or ::1)
	toLocal(net string) sockaddr
}

func internetSocket(ctx context.Context, net string, laddr, raddr sockaddr, sotype, proto int, mode string) (conn *netFD, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// favoriteAddrFamily returns the appropriate address family for the
// given network, laddr, raddr and mode.
//
// If mode indicates "listen" and laddr is a wildcard, we assume that
// the user wants to make a passive-open connection with a wildcard
// address family, both AF_INET and AF_INET6, and a wildcard address
// like the following:
//
//   - A listen for a wildcard communication domain, "tcp",
//     with a wildcard address: If the platform supports
//     both IPv6 and IPv4-mapped IPv6 communication capabilities,
//     or does not support IPv4, we use a dual stack, AF_INET6 and
//     IPV6_V6ONLY=0, wildcard address listen. The dual stack
//     wildcard address listen may fall back to an IPv6-only,
//     AF_INET6 and IPV6_V6ONLY=1, wildcard address listen.
//     Otherwise we prefer an IPv4-only, AF_INET, wildcard address
//     listen.
//
//   - A listen for a wildcard communication domain, "tcp",
//     with an IPv4 wildcard address: same as above.
//
//   - A listen for a wildcard communication domain, "tcp",
//     with an IPv6 wildcard address: same as above.
//
//   - A listen for an IPv4 communication domain, "tcp4",
//     with an IPv4 wildcard address: We use an IPv4-only, AF_INET,
//     wildcard address listen.
//
//   - A listen for an IPv6 communication domain, "tcp6",
//     with an IPv6 wildcard address: We use an IPv6-only, AF_INET6
//     and IPV6_V6ONLY=1, wildcard address listen.
//
// Otherwise guess: If the addresses are IPv4 then returns AF_INET,
// or else returns AF_INET6. It also returns a boolean value what
// designates IPV6_V6ONLY option.
//
// Note that the latest DragonFly BSD and OpenBSD kernels allow
// neither "net.inet6.ip6.v6only=1" change nor IPPROTO_IPV6 level
// IPV6_V6ONLY socket option setting.
func favoriteAddrFamily(network string, laddr, raddr sockaddr) (family int, ipv6only bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// socket returns a network file descriptor that is ready for
// asynchronous I/O using the network poller.
func socket(ctx context.Context, net string, family, sotype, proto int, ipv6only bool, laddr, raddr sockaddr) (netfd *netFD, err error) {
	_ = "STUB: not implemented"
	// syscall.Socket & set socket options
	return nil, nil
}

// sockaddrToAddr returns a go/net friendly address
func sockaddrToAddr(sa syscall.Sockaddr) net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// if zone == "" && sa.ZoneId != 0 {
// }
