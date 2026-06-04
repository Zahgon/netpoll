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

//go:build darwin || netbsd || freebsd || openbsd || dragonfly

package netpoll

import (
	"sync"
)

func openPoll() (Poll, error) { _ = "STUB: not implemented"; return *new(Poll), nil }

func openDefaultPoll() (*defaultPoll, error) { _ = "STUB: not implemented"; return nil, nil }

type defaultPoll struct {
	fd      int
	trigger uint32
	m       sync.Map       //nolint:unused // only used in go:race
	opcache *operatorCache // operator cache
	hups    []func(p Poll) error
}

// Wait implements Poll.
func (p *defaultPoll) Wait() error {
	_ = "STUB: not implemented"
	// init
	return nil
}

// wait

// exit gracefully

// trigger

// clean trigger

// for non-connection

// only for connection

// read all left data if peer send and close

// only close connection if no further read bytes

// for non-connection

// only for connection

// TODO: Let the upper layer pass in whether to use ZeroCopy.

// hup conns together to avoid blocking the poll.

// TODO: Close will bad file descriptor here
func (p *defaultPoll) Close() error { _ = "STUB: not implemented"; return nil }

// Trigger implements Poll.
func (p *defaultPoll) Trigger() error { _ = "STUB: not implemented"; return nil }

// Control implements Poll.
func (p *defaultPoll) Control(operator *FDOperator, event PollEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// means WaitWrite finished
