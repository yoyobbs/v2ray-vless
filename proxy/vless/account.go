// +build !confonly

package vless

import (
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/uuid"
)

// AsAccount implements protocol.Account.AsAccount().
func (a *Account) AsAccount() (protocol.Account, error) {
	id, err := uuid.ParseString(a.Id)
	if err != nil {
		return nil, newError("failed to parse ID").Base(err).AtError()
	}
	return &MemoryAccount{
		ID:         protocol.NewID(id),
		Mess:       a.Mess,       // need parser?
		Encryption: a.Encryption, // need parser?
	}, nil
}

// MemoryAccount is an in-memory form of VLess account.
type MemoryAccount struct {
	// ID of the account.
	ID *protocol.ID
	// Mess of the account. Used for client connections for now.
	Mess string
	// Encryption of the account. Used for client connections, and only accepts "none" for now.
	Encryption string
}

// Equals implements protocol.Account.Equals().
func (a *MemoryAccount) Equals(account protocol.Account) bool {
	vlessAccount, ok := account.(*MemoryAccount)
	if !ok {
		return false
	}
	return a.ID.Equals(vlessAccount.ID)
}
