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

const (
	managerUninitialized = iota
	managerInitializing
	managerInitialized
)

func newManager(numLoops int) *manager { _ = "STUB: not implemented"; return nil }

// LoadBalance is used to do load balancing among multiple pollers.
// a single poller may not be optimal if the number of cores is large (40C+).
type manager struct {
	numLoops int32
	status   int32       // 0: uninitialized, 1: initializing, 2: initialized
	balance  loadbalance // load balancing method
	polls    []Poll      // all the polls
}

// SetNumLoops will return error when set numLoops < 1
func (m *manager) SetNumLoops(numLoops int) (err error) { _ = "STUB: not implemented"; return nil }

// note: set new numLoops first and then change the status

// SetLoadBalance set load balance.
func (m *manager) SetLoadBalance(lb LoadBalance) error { _ = "STUB: not implemented"; return nil }

// Close release all resources.
func (m *manager) Close() (err error) { _ = "STUB: not implemented"; return nil }

// Run all pollers.
func (m *manager) Run() (err error) { _ = "STUB: not implemented"; return nil }

// shrink polls

// close redundant polls

// growth polls

// LoadBalance must be set before calling Run, otherwise it will panic.

// Reset pollers, this operation is very dangerous, please make sure to do this when calling !
func (m *manager) Reset() error { _ = "STUB: not implemented"; return nil }

// Pick will select the poller for use each time based on the LoadBalance.
func (m *manager) Pick() Poll {
	_ = "STUB: not implemented"

	// fast path
	return *new(Poll)
}

// slow path
// try to get initializing lock failed, wait others finished the init work, and try again

// adjust polls
// m.Run() will finish very quickly, so will not many goroutines block on Pick.

//nolint:staticcheck // SA9003: empty branch

// SetNumLoops called during m.Run() which cause CAS failed
// The polls will be adjusted next Pick
