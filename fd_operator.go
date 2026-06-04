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

// FDOperator is a collection of operations on file descriptors.
type FDOperator struct {
	// FD is file descriptor, poll will bind when register.
	FD int

	// The FDOperator provides three operations of reading, writing, and hanging.
	// The poll actively fire the FDOperator when fd changes, no check the return value of FDOperator.
	OnRead  func(p Poll) error
	OnWrite func(p Poll) error
	OnHup   func(p Poll) error

	// The following is the required fn, which must exist when used, or directly panic.
	// Fns are only called by the poll when handles connection events.
	Inputs   func(vs [][]byte) (rs [][]byte)
	InputAck func(n int) (err error)

	// Outputs will locked if len(rs) > 0, which need unlocked by OutputAck.
	// supportZeroCopy is not implemented, and it will be ignored
	Outputs   func(vs [][]byte) (rs [][]byte, supportZeroCopy bool)
	OutputAck func(n int) (err error)

	// poll is the registered location of the file descriptor.
	poll Poll

	// protect only detach once
	detached int32

	// private, used by operatorCache
	next  *FDOperator
	state int32 // CAS: 0(unused) 1(inuse) 2(do-done)
	index int32 // index in operatorCache
}

func (op *FDOperator) Control(event PollEvent) error { _ = "STUB: not implemented"; return nil }

func (op *FDOperator) Free() { _ = "STUB: not implemented"; return }

func (op *FDOperator) do() (can bool) { _ = "STUB: not implemented"; return false }

func (op *FDOperator) done() { _ = "STUB: not implemented"; return }

func (op *FDOperator) inuse() { _ = "STUB: not implemented"; return }

func (op *FDOperator) unused() { _ = "STUB: not implemented"; return }

func (op *FDOperator) isUnused() bool { _ = "STUB: not implemented"; return false }

func (op *FDOperator) reset() { _ = "STUB: not implemented"; return }
