/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package exchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger-labs/fabric-token-sdk/token/driver"
	"github.com/pkg/errors"
)

// ClaimSignature is the claim signature of an exchange script
type ClaimSignature struct {
	RecipientSignature []byte
	Preimage           []byte
}

// ClaimSigner is the signer for the claim of an exchange script
type ClaimSigner struct {
	Recipient driver.Signer
	Preimage  []byte
}

// Sign returns a signature of the recipient over the token request and preimage
func (cs *ClaimSigner) Sign(tokenRequestAndTxID []byte) ([]byte, error) {
	msg := concatTokenRequestTxIDPreimage(tokenRequestAndTxID, cs.Preimage)
	sigma, err := cs.Recipient.Sign(msg)
	if err != nil {
		return nil, err
	}

	claimSignature := ClaimSignature{
		Preimage:           cs.Preimage,
		RecipientSignature: sigma,
	}
	return json.Marshal(claimSignature)
}

func concatTokenRequestTxIDPreimage(tokenRequestAndTxID []byte, preImage []byte) []byte {
	var msg []byte
	msg = append(msg, tokenRequestAndTxID...)
	msg = append(msg, preImage...)
	return msg
}

// ClaimVerifier is the verifier of a ClaimSignature
type ClaimVerifier struct {
	Recipient driver.Verifier
	HashInfo  HashInfo
}

// Verify verifies that the passed signature is valid and that the contained preimage matches the hash info
func (cv *ClaimVerifier) Verify(tokenRequestAndTxID, claimSignature []byte) error {
	sig := &ClaimSignature{}
	err := json.Unmarshal(claimSignature, sig)
	if err != nil {
		return err
	}

	msg := concatTokenRequestTxIDPreimage(tokenRequestAndTxID, sig.Preimage)
	if err := cv.Recipient.Verify(msg, sig.RecipientSignature); err != nil {
		return err
	}

	hash := cv.HashInfo.HashFunc.New()
	if _, err = hash.Write(sig.Preimage); err != nil {
		return err
	}
	image := hash.Sum(nil)
	image = []byte(cv.HashInfo.HashEncoding.New().EncodeToString(image))

	if !bytes.Equal(cv.HashInfo.Hash, image) {
		return fmt.Errorf("hash mismatch: SHA(%x) = %x != %x", sig.Preimage, image, cv.HashInfo.Hash)
	}

	return nil
}

// ExchangeVerifier checks if an exchange script can be claimed or reclaimed
type ExchangeVerifier struct {
	Recipient driver.Verifier
	Sender    driver.Verifier
	Deadline  time.Time
	HashInfo  HashInfo
}

// Verify verifies the claim or reclaim signature
func (v *ExchangeVerifier) Verify(msg []byte, sigma []byte) error {
	// if timeout has not elapsed, only claim is allowed
	if time.Now().Before(v.Deadline) {
		cv := &ClaimVerifier{Recipient: v.Recipient, HashInfo: HashInfo{
			Hash:         v.HashInfo.Hash,
			HashFunc:     v.HashInfo.HashFunc,
			HashEncoding: v.HashInfo.HashEncoding,
		}}
		if err := cv.Verify(msg, sigma); err != nil {
			return errors.WithMessagef(err, "failed verifying exchange claim signature")
		}
		return nil
	}
	// if timeout has elapsed, only a reclaim is possible
	if err := v.Sender.Verify(msg, sigma); err != nil {
		return errors.WithMessagef(err, "failed verifying exchange reclaim signature")
	}
	return nil
}
