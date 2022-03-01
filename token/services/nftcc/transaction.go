/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package nftcc

import (
	"encoding/json"
	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/services/state"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/flogging"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
	"github.com/hyperledger-labs/fabric-token-sdk/token"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/ttxcc"
	"github.com/pkg/errors"
)

var logger = flogging.MustGetLogger("token-sdk.nftcc")

type TxOption ttxcc.TxOption

func WithAuditor(auditor view.Identity) TxOption {
	return func(o *ttxcc.TxOptions) error {
		o.Auditor = auditor
		return nil
	}
}

type Transaction struct {
	*ttxcc.Transaction
}

func NewAnonymousTransaction(sp view.Context, opts ...TxOption) (*Transaction, error) {
	// convert opts to ttxcc.TxOption
	txOpts := make([]ttxcc.TxOption, len(opts))
	for i, opt := range opts {
		txOpts[i] = ttxcc.TxOption(opt)
	}
	tx, err := ttxcc.NewAnonymousTransaction(sp, txOpts...)
	if err != nil {
		return nil, err
	}

	return &Transaction{Transaction: tx}, nil
}

func Wrap(tx *ttxcc.Transaction) *Transaction {
	return &Transaction{Transaction: tx}
}

func (t Transaction) Issue(wallet *token.IssuerWallet, state interface{}, recipient view.Identity) error {
	// set state id first
	_, err := t.setStateID(state)
	if err != nil {
		return err
	}
	// marshal state to json
	stateJSON, err := json.Marshal(state)
	if err != nil {
		return errors.Wrap(err, "failed to marshal state")
	}
	stateJSONStr := string(stateJSON)

	// Issue
	return t.Transaction.Issue(wallet, recipient, stateJSONStr, 1)
}

func (t Transaction) Transfer(state interface{}, recipient view.Identity) error {
	panic("implement me")
}

func (t *Transaction) setStateID(s interface{}) (string, error) {
	logger.Debugf("setStateID %v...", s)
	defer logger.Debugf("setStateID...done")
	var key string
	var err error
	switch d := s.(type) {
	case AutoLinearState:
		logger.Debugf("AutoLinearState...")
		key, err = d.GetLinearID()
		if err != nil {
			return "", err
		}
	case LinearState:
		logger.Debugf("LinearState...")
		key = state.GenerateUUID()
		key = d.SetLinearID(key)
	default:
		return "", nil
	}
	return key, nil
}
