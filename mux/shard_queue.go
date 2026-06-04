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

package mux

import (
	"runtime"
	"sync"

	"github.com/cloudwego/netpoll"
)

/* DOC:
 * ShardQueue uses the netpoll's nocopy API to merge and send data.
 * The Data Flush is passively triggered by ShardQueue.Add and does not require user operations.
 * If there is an error in the data transmission, the connection will be closed.
 *
 * ShardQueue.Add: add the data to be sent.
 * NewShardQueue: create a queue with netpoll.Connection.
 * ShardSize: the recommended number of shards is 32.
 */
var ShardSize int

func init() {
	ShardSize = runtime.GOMAXPROCS(0)
}

// NewShardQueue .
func NewShardQueue(size int, conn netpoll.Connection) (queue *ShardQueue) {
	_ = "STUB: not implemented"
	return nil
}

// WriterGetter is used to get a netpoll.Writer.
type WriterGetter func() (buf netpoll.Writer, isNil bool)

// ShardQueue uses the netpoll's nocopy API to merge and send data.
// The Data Flush is passively triggered by ShardQueue.Add and does not require user operations.
// If there is an error in the data transmission, the connection will be closed.
// ShardQueue.Add: add the data to be sent.
type ShardQueue struct {
	conn      netpoll.Connection
	idx, size int32
	getters   [][]WriterGetter // len(getters) = size
	swap      []WriterGetter   // use for swap
	locks     []int32          // len(locks) = size
	queueTrigger
}

const (
	// queueTrigger state
	active  = 0
	closing = 1
	closed  = 2
)

// here for trigger
type queueTrigger struct {
	trigger  int32
	state    int32 // 0: active, 1: closing, 2: closed
	runNum   int32
	w, r     int32      // ptr of list
	list     []int32    // record the triggered shard
	listLock sync.Mutex // list total lock
}

// Add adds to q.getters[shard]
func (q *ShardQueue) Add(gts ...WriterGetter) { _ = "STUB: not implemented"; return }

func (q *ShardQueue) Close() error { _ = "STUB: not implemented"; return nil }

// wait for all tasks finished

// triggering shard.
func (q *ShardQueue) triggering(shard int32) { _ = "STUB: not implemented"; return }

// foreach swap r & w. It's not concurrency safe.
func (q *ShardQueue) foreach() { _ = "STUB: not implemented"; return }

// is negative number of triggerNum

// lock & swap

// deal

// quit & check again

// if state is closing, change it to closed

// deal is used to get deal of netpoll.Writer.
func (q *ShardQueue) deal(gts []WriterGetter) { _ = "STUB: not implemented"; return }

// flush is used to flush netpoll.Writer.
func (q *ShardQueue) flush() { _ = "STUB: not implemented"; return }

// lock shard.
func (q *ShardQueue) lock(shard int32) { _ = "STUB: not implemented"; return }

// unlock shard.
func (q *ShardQueue) unlock(shard int32) { _ = "STUB: not implemented"; return }
