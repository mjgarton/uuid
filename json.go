// Copyright 2014 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "errors"

func (u UUID) MarshalJSON() ([]byte, error) {
	if u == emptyUUID {
		return []byte(`""`), nil
	}

	b := [38]byte{}

	for i, n := range []int{
		1, 3, 5, 7,
		10, 12,
		15, 17,
		20, 22,
		25, 27, 29, 31, 33, 35,
	} {
		b[n] = halfbyte2hexchar[(u[i]>>4)&0x0f]
		b[n+1] = halfbyte2hexchar[u[i]&0x0f]
	}

	b[0] = '"'
	b[9] = '-'
	b[14] = '-'
	b[19] = '-'
	b[24] = '-'
	b[37] = '"'

	return b[:], nil
}

var halfbyte2hexchar = []byte{
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 101, 102,
}

func (u *UUID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == `""` {
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("invalid UUID format")
	}
	data = data[1 : len(data)-1]
	uu, err := Parse(string(data))
	if err != nil {
		return errors.New("invalid UUID format")
	}
	*u = uu
	return nil
}
