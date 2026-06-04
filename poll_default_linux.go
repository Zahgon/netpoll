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

package netpoll

import (
	"sync"
)

func openPoll() (Poll, error) { _ = "STUB: not implemented"; return *new(Poll), nil }

func openDefaultPoll() (*defaultPoll, error) { _ = "STUB: not implemented"; return nil, nil }

type defaultPoll struct {
	pollArgs
	fd      int            // epoll fd
	wop     *FDOperator    // eventfd, wake epoll_wait
	buf     []byte         // read wfd trigger msg
	trigger uint32         // trigger flag
	m       sync.Map       //nolint:unused // only used in go:race
	opcache *operatorCache // operator cache
	// fns for handle events
	Reset   func(size, caps int)
	Handler func(events []epollevent) (closed bool)
}

type pollArgs struct {
	size     int
	caps     int
	events   []epollevent
	barriers []barrier
	hups     []func(p Poll) error
}

func (a *pollArgs) reset(size, caps int) { _ = "STUB: not implemented"; return }

// Wait implements Poll.
func (p *defaultPoll) Wait() (err error) {
	_ = "STUB: not implemented"
	// init
	return nil
}

// wait

// we can make sure that there is no op remaining if Handler finished

func (p *defaultPoll) handler(events []epollevent) (closed bool) {
	_ = "STUB: not implemented"
	return false
}

// trigger or exit gracefully

// must clean trigger first

// if closed & exit

// for non-connection

// for connection

// read all left data if peer send and close

// read all left data if peer send and close

// only close connection if no further read bytes

// Under block-zerocopy, the kernel may give an error callback, which is not a real error, just an EAGAIN.
// So here we need to check this error, if it is EAGAIN then do nothing, otherwise still mark as hup.

// for non-connection

// for connection

// hup conns together to avoid blocking the poll.

// Close will write 10000000
func (p *defaultPoll) Close() error { _ = "STUB: not implemented"; return nil }

// Trigger implements Poll.
func (p *defaultPoll) Trigger() error { _ = "STUB: not implemented"; return nil }

// MAX(eventfd) = 0xfffffffffffffffe

// Control implements Poll.
func (p *defaultPoll) Control(operator *FDOperator, event PollEvent) error {
	_ = "STUB: not implemented"
	// DON'T move `fd=operator.FD` behind inuse() call, we can only access operator before op.inuse() for avoid race
	// G1:              G2:
	// op.inuse()       op.unused()
	// op.FD  -- T1     op.FD = 0  -- T2
	// T1 and T2 may happen together
	return nil
}

// server accept a new connection and wait read

// client create a new connection and wait connect finished

// deregister

// connection wait read/write

// connection wait read
