// Copyright 2014 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "errors"

func (u UUID) MarshalJSON() ([]byte, error) {
	if u == emptyUUID {
		return []byte(`""`), nil
	}
	return []byte(`"` + u.String() + `"`), nil
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
