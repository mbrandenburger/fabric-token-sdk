/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package nftcc

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hyperledger-labs/fabric-token-sdk/token"
	token2 "github.com/hyperledger-labs/fabric-token-sdk/token/token"
	"github.com/pkg/errors"
	"github.com/thedevsaddam/gojsonq"
)

type vault interface {
	GetTokens(inputs ...*token2.ID) ([]*token2.Token, error)
}

type selector interface {
	Filter(filter token.SelectorFilterByUnspentToken, q string) ([]*token2.ID, error)
}

type QueryExecutor struct {
	selector
	vault
}

func (s *QueryExecutor) QueryByKey(house interface{}, key string, value string) error {
	ids, err := s.selector.Filter(&jsonFilter{
		q:     gojsonq.New(),
		key:   key,
		value: value,
	}, "1")
	if err != nil {
		return errors.Wrap(err, "failed to filter")
	}
	tokens, err := s.vault.GetTokens(ids...)
	if err != nil {
		return errors.Wrap(err, "failed to get tokens")
	}
	for _, t := range tokens {
		q, err := token2.ToQuantity(t.Quantity, 64)
		if err != nil {
			return errors.Wrap(err, "failed to convert quantity")
		}
		if q.Cmp(token2.NewQuantityFromUInt64(1)) == 0 {
			// this is the token
			decoded, err := base64.StdEncoding.DecodeString(t.Type)
			if err != nil {
				return errors.Wrap(err, "failed to decode type")
			}
			if err := json.Unmarshal(decoded, house); err == nil {
				return errors.Wrap(err, "failed to unmarshal state")
			}
		}
	}
	return errors.Wrap(err, "no suitable token found")
}

type jsonFilter struct {
	q          *gojsonq.JSONQ
	key, value string
}

func (j jsonFilter) ContainsToken(token *token2.UnspentToken) bool {
	decoded, err := base64.StdEncoding.DecodeString(token.Type)
	if err != nil {
		logger.Debugf("failed to decode token type [%s]", token.Type)
		return false
	}
	logger.Debugf("decoded token type [%s]", string(decoded))
	jq := j.q.FromString(string(decoded))
	res := jq.From(j.key).Where(j.key, "==", j.value).Get()
	logger.Debugf("res [%s] for [%s,%s]", res, j.key, j.value)
	return res != nil
}