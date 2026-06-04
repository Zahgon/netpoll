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
	"sync/atomic"
)

// ------------------------------------ implement OnPrepare, OnRequest, CloseCallback ------------------------------------

type gracefulExit interface {
	isIdle() (yes bool)
	Close() (err error)
}

// onEvent is the collection of event processing.
// OnPrepare, OnRequest, CloseCallback share the lock processing,
// which is a CAS lock and can only be cleared by OnRequest.
type onEvent struct {
	ctx                  context.Context
	onConnectCallback    atomic.Value
	onDisconnectCallback atomic.Value
	onRequestCallback    atomic.Value
	closeCallbacks       atomic.Value // value is latest *callbackNode
}

type callbackNode struct {
	fn  CloseCallback
	pre *callbackNode
}

// SetOnConnect set the OnConnect callback.
func (c *connection) SetOnConnect(onConnect OnConnect) error { _ = "STUB: not implemented"; return nil }

// SetOnDisconnect set the OnDisconnect callback.
func (c *connection) SetOnDisconnect(onDisconnect OnDisconnect) error {
	_ = "STUB: not implemented"
	return nil
}

// SetOnRequest initialize ctx when setting OnRequest.
func (c *connection) SetOnRequest(onRequest OnRequest) error { _ = "STUB: not implemented"; return nil }

// fix: trigger OnRequest if there is already input data.

// AddCloseCallback adds a CloseCallback to this connection.
func (c *connection) AddCloseCallback(callback CloseCallback) error {
	_ = "STUB: not implemented"
	return nil
}

// onPrepare supports close connection, but not read/write data.
// connection will be registered by this call after preparing.
func (c *connection) onPrepare(opts *options) (err error) { _ = "STUB: not implemented"; return nil }

// calling prepare first and then register.

// prepare may close the connection.

// onConnect is responsible for executing onRequest if there is new data coming after onConnect callback finished.
func (c *connection) onConnect() { _ = "STUB: not implemented"; return }

// it never happens because onDisconnect will not lock connecting if c.connected == 0

// when onDisconnect called, c.IsActive() must return false
func (c *connection) onDisconnect() { _ = "STUB: not implemented"; return }

// no need lock if onConnect is nil
// it's ok to force set state to disconnected since onConnect is nil

// check if OnConnect finished when onConnect != nil && onDisconnect != nil
// means OnConnect already finished
// protect onDisconnect run once
// if CAS return false, means OnConnect already helps to run onDisconnect

// OnConnect is not finished yet, return and let onConnect helps to call onDisconnect

// onRequest is responsible for executing the closeCallbacks after the connection has been closed.
func (c *connection) onRequest() (needTrigger bool) { _ = "STUB: not implemented"; return false }

// wait onConnect finished first

// let onConnect to call onRequest

// if not processed, should trigger read

// onProcess is responsible for executing the onConnect/onRequest function serially,
// and make sure the connection has been closed correctly if user call c.Close() in onConnect/onRequest function.
func (c *connection) onProcess(onConnect OnConnect, onRequest OnRequest) (processed bool) {
	_ = "STUB: not implemented"
	// task already exists
	return false
}

// cannot use recover() here, since we don't want to break the panic stack

// trigger onConnect first

// since we hold connecting lock, so we should help to call onDisconnect here

// The `onRequest` must be executed at least once if conn have any readable data,
// which is in order to cover the `send & close by peer` case.

// The processing loop must ensure that the connection meets `IsActive`.
// `onRequest` must either eventually read all the input data or actively Close the connection,
// otherwise the goroutine will fall into a dead loop.

// close by user or not processable

// handling callback if connection has been closed.

//  if closed by user when processing, it "may" needs detach

// Here is a corner case that operator will be detached twice:
//   If server closed the connection(client OnHup will detach op first and closeBy=poller),
//   and then client's OnRequest function also closed the connection(closeBy=user).
// But operator already prevent that detach twice will not cause any problem

// Note: Poller's closeCallback call will try to get processing lock failed but here already near to unlock processing.
//       So here we need to check connection state again, to avoid connection leak
// double check close state

// poller will get the processing lock failed, here help poller do closeCallback
// fd must already detach by poller

// double check is processable

// task exits

// end of task closure func

// add new task

// closeCallback .
// It can be confirmed that closeCallback and onRequest will not be executed concurrently.
// If onRequest is still running, it will trigger closeCallback on exit.
func (c *connection) closeCallback(needLock, needDetach bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If Close is called during OnPrepare, poll is not registered.
// PollDetach only happen when user call conn.Close() or poller detect error

// register only use for connection register into poll.
func (c *connection) register() (err error) { _ = "STUB: not implemented"; return nil }

// isIdle implements gracefulExit.
func (c *connection) isIdle() (yes bool) { _ = "STUB: not implemented"; return false }
