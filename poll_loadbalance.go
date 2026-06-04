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

// LoadBalance sets the load balancing method.
type LoadBalance int

const (
	// RoundRobin requests that connections are distributed to a Poll
	// in a round-robin fashion.
	RoundRobin LoadBalance = iota
	// Random requests that connections are randomly distributed.
	Random
)

// loadbalance sets the load balancing method for []*polls
type loadbalance interface {
	LoadBalance() LoadBalance
	// Pick choose the most qualified Poll
	Pick() (poll Poll)

	Rebalance(polls []Poll)
}

func newLoadbalance(lb LoadBalance, polls []Poll) loadbalance {
	_ = "STUB: not implemented"
	return *new(loadbalance)
}

func newRandomLB(polls []Poll) loadbalance { _ = "STUB: not implemented"; return *new(loadbalance) }

type randomLB struct {
	polls    []Poll
	pollSize int
}

func (b *randomLB) LoadBalance() LoadBalance { _ = "STUB: not implemented"; return *new(LoadBalance) }

func (b *randomLB) Pick() (poll Poll) { _ = "STUB: not implemented"; return *new(Poll) }

func (b *randomLB) Rebalance(polls []Poll) { _ = "STUB: not implemented"; return }

func newRoundRobinLB(polls []Poll) loadbalance { _ = "STUB: not implemented"; return *new(loadbalance) }

type roundRobinLB struct {
	polls    []Poll
	accepted uintptr // accept counter
	pollSize int
}

func (b *roundRobinLB) LoadBalance() LoadBalance {
	_ = "STUB: not implemented"
	return *new(LoadBalance)
}

func (b *roundRobinLB) Pick() (poll Poll) { _ = "STUB: not implemented"; return *new(Poll) }

func (b *roundRobinLB) Rebalance(polls []Poll) { _ = "STUB: not implemented"; return }
