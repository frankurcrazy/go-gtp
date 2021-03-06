// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ies

// NewCause creates a new Cause IE.
func NewCause(cause uint8, pce, bce, cs uint8, offendingIE *IE) *IE {
	i := New(Cause, 0x00, make([]byte, 2))
	i.Payload[0] = cause
	i.Payload[1] = ((pce << 2) & 0x04) | ((bce << 1) & 0x02) | cs&0x01

	if offendingIE != nil {
		i.Payload = append(i.Payload, offendingIE.Type)
		i.SetLength()
	}
	return i
}

// Cause returns Cause in uint8 if the type of IE matches.
func (i *IE) Cause() uint8 {
	if i.Type != Cause {
		return 0
	}

	return i.Payload[0]
}

// IsRemoteCause returns IsRemoteCause in bool if the type of IE matches.
func (i *IE) IsRemoteCause() bool {
	if i.Type != Cause {
		return false
	}

	if i.Payload[1]>>2&0x01 == 1 {
		return true
	}
	return false
}

// IsBearerContextIEError returns IsBearerContextIEError in bool if the type of IE matches.
func (i *IE) IsBearerContextIEError() bool {
	if i.Type != Cause {
		return false
	}

	if i.Payload[1]>>1&0x01 == 1 {
		return true
	}
	return false
}

// IsPDNConnectionIEError returns IsPDNConnectionIEError in bool if the type of IE matches.
func (i *IE) IsPDNConnectionIEError() bool {
	if i.Type != Cause {
		return false
	}

	if i.Payload[1]&0x01 == 1 {
		return true
	}
	return false
}
