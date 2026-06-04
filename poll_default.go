// Copyright 2023 CloudWeGo Authors
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

//go:build darwin || netbsd || freebsd || openbsd || dragonfly || linux

package netpoll

func (p *defaultPoll) Alloc() (operator *FDOperator) { _ = "STUB: not implemented"; return nil }

func (p *defaultPoll) Free(operator *FDOperator) { _ = "STUB: not implemented"; return }

func (p *defaultPoll) appendHup(operator *FDOperator) { _ = "STUB: not implemented"; return }

func (p *defaultPoll) detach(operator *FDOperator) { _ = "STUB: not implemented"; return }

func (p *defaultPoll) onhups() { _ = "STUB: not implemented"; return }

// readall read all left data before close connection
func readall(op *FDOperator, br barrier) (total int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
