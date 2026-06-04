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

// ------------------------------------------ implement FDOperator ------------------------------------------

// onHup means close by poller.
func (c *connection) onHup(p Poll) error { _ = "STUB: not implemented"; return nil }

// call Disconnect callback first

// It depends on closing by user if OnConnect and OnRequest is nil, otherwise it needs to be released actively.
// It can be confirmed that the OnRequest goroutine has been exited before closeCallback executing,
// and it is safe to close the buffer at this time.

// already PollDetach when call OnHup

// onClose means close by user.
func (c *connection) onClose() error {
	_ = "STUB: not implemented"
	// user code close the connection
	return nil
}

// Detach from poller when processing finished, otherwise it will cause race

// closed by poller
// still need to change closing status to `user` since OnProcess should not be processed again

// user code should actively close the connection to recycle resources.
// poller already detached operator

// closeBuffer recycle input & output LinkBuffer.
func (c *connection) closeBuffer() { _ = "STUB: not implemented"; return }

// if client close the connection, we cannot ensure that the poller is not process the buffer,
// so we need to check the buffer length, and if it's an "unclean" close operation, let's give up to reuse the buffer

// inputs implements FDOperator.
func (c *connection) inputs(vs [][]byte) (rs [][]byte) { _ = "STUB: not implemented"; return nil }

// inputAck implements FDOperator.
func (c *connection) inputAck(n int) (err error) { _ = "STUB: not implemented"; return nil }

// Auto size bookSize.

// first start onRequest

// outputs implements FDOperator.
func (c *connection) outputs(vs [][]byte) (rs [][]byte, _ bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// outputAck implements FDOperator.
func (c *connection) outputAck(n int) (err error) { _ = "STUB: not implemented"; return nil }

// rw2r removed the monitoring of write events.
func (c *connection) rw2r() { _ = "STUB: not implemented"; return }
