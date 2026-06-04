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
	"context"
	"net"
	"time"
)

// DialConnection is a default implementation of Dialer.
func DialConnection(network, address string, timeout time.Duration) (connection Connection, err error) {
	_ = "STUB: not implemented"
	return *new(Connection), nil
}

// NewFDConnection create a Connection initialized by any fd
// It's useful for writing unit tests for functions that have args with the type of netpoll.Connection
// The typical usage is like:
//
//	rfd, wfd := netpoll.GetSysFdPairs()
//	rconn, _ = netpoll.NewFDConnection(rfd)
//	wconn, _ = netpoll.NewFDConnection(wfd)
func NewFDConnection(fd int) (Connection, error) {
	_ = "STUB: not implemented"
	return *new(Connection), nil
}

// NewDialer only support TCP and unix socket now.
func NewDialer() Dialer { _ = "STUB: not implemented"; return *new(Dialer) }

var defaultDialer = NewDialer()

type dialer struct{}

// DialTimeout implements Dialer.
func (d *dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// DialConnection implements Dialer.
func (d *dialer) DialConnection(network, address string, timeout time.Duration) (connection Connection, err error) {
	_ = "STUB: not implemented"
	return *new(Connection), nil
}

func (d *dialer) dialTCP(ctx context.Context, network, address string) (connection *TCPConnection, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// host maybe empty if address is :12345

// The error from the first address is most relevant.

// check timeout error

// sysDialer contains a Dial's parameters and configuration.
type sysDialer struct {
	net.Dialer
	network, address string
}
