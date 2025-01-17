/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package exchange

import (
	"crypto"
	"encoding/json"
	"time"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/assert"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
	"github.com/hyperledger-labs/fabric-token-sdk/token"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/interop/encoding"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/interop/exchange"
)

// Scan contains the input information for a scan of a matching preimage
type Scan struct {
	// TMSID identifies the TMS to use to perform the token operation
	TMSID token.TMSID
	// Timeout of the scan
	Timeout time.Duration
	// Hash is the hash to use in the scan
	Hash []byte
	// HashFunc is the hash function to use in the scan
	HashFunc crypto.Hash
}

type ScanView struct {
	*Scan
}

func (s *ScanView) Call(context view.Context) (interface{}, error) {
	preImage, err := exchange.ScanForPreImage(context, s.Hash, s.HashFunc, encoding.None, s.Timeout, token.WithTMSID(s.TMSID))
	if err != nil {
		return nil, err
	}
	return preImage, nil
}

type ScanViewFactory struct {
}

func (s *ScanViewFactory) NewView(in []byte) (view.View, error) {
	f := &ScanView{Scan: &Scan{}}
	err := json.Unmarshal(in, f.Scan)
	assert.NoError(err, "failed unmarshalling input")

	return f, nil
}
