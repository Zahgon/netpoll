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
	"sync"
)

// newServer wrap listener into server, quit will be invoked when server exit.
func newServer(ln Listener, opts *options, onQuit func(err error)) *server {
	_ = "STUB: not implemented"
	return nil
}

type server struct {
	operator    FDOperator
	ln          Listener
	opts        *options
	onQuit      func(err error)
	connections sync.Map // key=fd, value=connection
}

// Run this server.
func (s *server) Run() (err error) { _ = "STUB: not implemented"; return nil }

// Close this server with deadline.
func (s *server) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// all connections have been closed

// smart control graceful shutdown check internal
// we should wait for more time if there are more active connections

// max wait time is 1000 ms

// min wait time is 50 ms

// OnRead implements FDOperator.
func (s *server) OnRead(p Poll) error {
	_ = "STUB: not implemented"
	// accept socket
	return nil
}

// EAGAIN | EWOULDBLOCK if conn and err both nil

// delay accept when too many open files

// since we use Epoll LT, we have to detach listener fd from epoll first
// and re-register it when accept successfully or there is no available connection

// ms

// recovery accept poll loop

// shut down

// OnHup implements FDOperator.
func (s *server) OnHup(p Poll) error { _ = "STUB: not implemented"; return nil }

func (s *server) onAccept(conn Conn) {
	_ = "STUB: not implemented"
	// store & register connection
	return
}

// trigger onConnect asynchronously

func isOutOfFdErr(err error) bool { _ = "STUB: not implemented"; return false }
