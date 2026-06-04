// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”).
// All CloudWeGo Modifications are Copyright 2022 CloudWeGo authors.

package netpoll

func setDefaultSockopts(s, family, sotype int, ipv6only bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Allow both IP versions even if the OS default
// is otherwise. Note that some operating systems
// never admit this option.

// Allow broadcast.
