/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package interop

import (
	"encoding/json"

	view2 "github.com/hyperledger-labs/fabric-smart-client/platform/view"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
	"github.com/hyperledger-labs/fabric-token-sdk/token/core/identity"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/interop/exchange"
	"github.com/pkg/errors"
)

// ScriptInfo includes info about the sender and the recipient
type ScriptInfo struct {
	Sender    []byte
	Recipient []byte
}

// GetOwnerAuditInfo returns the audit info of the owner
func GetOwnerAuditInfo(raw []byte, s view2.ServiceProvider) ([]byte, error) {
	if len(raw) == 0 {
		// this is a redeem
		return nil, nil
	}

	owner, err := identity.UnmarshallRawOwner(raw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal owner of input token")
	}
	if owner.Type == identity.SerializedIdentityType {
		auditInfo, err := view2.GetSigService(s).GetAuditInfo(raw)
		if err != nil {
			return nil, errors.Wrapf(err, "failed getting audit info for recipient identity [%s]", view.Identity(raw).String())
		}
		return auditInfo, nil
	}

	sender, recipient, err := GetScriptSenderAndRecipient(owner)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting script sender and recipient")
	}

	auditInfo := &ScriptInfo{}
	auditInfo.Sender, err = view2.GetSigService(s).GetAuditInfo(sender)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting audit info for exchange script [%s]", view.Identity(raw).String())
	}

	auditInfo.Recipient, err = view2.GetSigService(s).GetAuditInfo(recipient)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting audit info for script [%s]", view.Identity(raw).String())
	}
	raw, err = json.Marshal(auditInfo)
	if err != nil {
		return nil, errors.Wrapf(err, "failed marshaling audit info for script")
	}
	return raw, nil
}

// GetScriptSenderAndRecipient returns the script's sender and recipient according to the type of the given owner
func GetScriptSenderAndRecipient(ro *identity.RawOwner) (sender, recipient view.Identity, err error) {
	if ro.Type == exchange.ScriptTypeExchange {
		script := &exchange.Script{}
		err = json.Unmarshal(ro.Identity, script)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to unmarshal exchange script")
		}
		return script.Sender, script.Recipient, nil
	}
	return nil, nil, errors.New("unknown identity type")
}
