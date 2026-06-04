// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”).
// All CloudWeGo Modifications are Copyright 2022 CloudWeGo authors.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package netpoll

func setDefaultSockopts(s, family, sotype int, ipv6only bool) error {
	_ = "STUB: not implemented"
	return nil
}

// On DragonFly BSD, we adjust the ephemeral port
// range because unlike other BSD systems its default
// port range doesn't conform to IANA recommendation
// as described in RFC 6056 and is pretty narrow.

// Allow broadcast.
