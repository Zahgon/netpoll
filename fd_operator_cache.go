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

func newOperatorCache() *operatorCache { _ = "STUB: not implemented"; return nil }

type operatorCache struct {
	first  *FDOperator
	cache  []*FDOperator
	locked int32
	// freelist store the freeable operator
	// to reduce GC pressure, we only store op index here
	freelocked int32
	freelist   []int32
}

func (c *operatorCache) alloc() *FDOperator { _ = "STUB: not implemented"; return nil }

// freeable mark the operator that could be freed
// only poller could do the real free action
func (c *operatorCache) freeable(op *FDOperator) {
	_ = "STUB: not implemented"
	// reset all state
	return
}

func (c *operatorCache) free() { _ = "STUB: not implemented"; return }

func lock(locked *int32) { _ = "STUB: not implemented"; return }

func unlock(locked *int32) { _ = "STUB: not implemented"; return }
