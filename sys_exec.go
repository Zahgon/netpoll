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
	"syscall"
)

// GetSysFdPairs creates and returns the fds of a pair of sockets.
func GetSysFdPairs() (r, w int) { _ = "STUB: not implemented"; return 0, 0 }

// setTCPNoDelay set the TCP_NODELAY flag on socket
func setTCPNoDelay(fd int, b bool) (err error) { _ = "STUB: not implemented"; return nil }

// Wrapper around the socket system call that marks the returned file
// descriptor as nonblocking and close-on-exec.
func sysSocket(family, sotype, proto int) (int, error) {
	_ = "STUB: not implemented"
	// See ../syscall/exec_unix.go for description of ForkLock.
	return 0, nil
}

const barriercap = 32

type barrier struct {
	bs  [][]byte
	ivs []syscall.Iovec
}

// writev wraps the writev system call.
func writev(fd int, bs [][]byte, ivs []syscall.Iovec) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// syscall

// readv wraps the readv system call.
// return 0, nil means EOF.
func readv(fd int, bs [][]byte, ivs []syscall.Iovec) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// syscall

// TODO: read from sysconf(_SC_IOV_MAX)? The Linux default is
//
//	1024 and this seems conservative enough for now. Darwin's
//	UIO_MAXIOV also seems to be 1024.
//
// iovecs limit length to 2GB(2^31)
func iovecs(bs [][]byte, ivs []syscall.Iovec) (iovLen int) { _ = "STUB: not implemented"; return 0 }

func resetIovecs(bs [][]byte, ivs []syscall.Iovec) { _ = "STUB: not implemented"; return }

// Boolean to int.
func boolint(b bool) int { _ = "STUB: not implemented"; return 0 }
