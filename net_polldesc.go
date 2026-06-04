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
)

func newPollDesc(fd int) *pollDesc { _ = "STUB: not implemented"; return nil }

type pollDesc struct {
	operator *FDOperator
	// The write event is OneShot, then mark the writable to skip duplicate calling.
	writeTrigger chan struct{}
	closeTrigger chan struct{}
}

// WaitWrite .
func (pd *pollDesc) WaitWrite(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil

	// add ET|Write|Hup
}

// triggered by poller
// triggered by poller
// no need to detach, since poller has done it in OnHup.

// triggered by ctx
// deregister from poller, upper caller function will close fd

// double check close trigger

func (pd *pollDesc) onwrite(p Poll) error { _ = "STUB: not implemented"; return nil }

func (pd *pollDesc) onhup(p Poll) error { _ = "STUB: not implemented"; return nil }

func (pd *pollDesc) detach() { _ = "STUB: not implemented"; return }
