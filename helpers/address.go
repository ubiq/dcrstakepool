// Copyright (c) 2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package helpers

import (
	"github.com/decred/dcrd/chaincfg/v3"
	"github.com/decred/dcrd/dcrec"
	"github.com/decred/dcrd/dcrutil/v3"
	"github.com/decred/dcrd/hdkeychain/v3"
)

const (
	// ExternalBranch is a helper value that needs to
	// match dcrwallet's udb.ExternalBranch
	ExternalBranch uint32 = 0
)

// DCRUtilAddressFromExtendedKey parses the public address of a hd extended key
// using a secp256k1 elliptic curve into a ECDSA public key, compresses it using
// ripemd160, and wraps it in a dcrutil AddressPubKeyHash in order to easily
// obtain its human readable formats. Returns an error upon a parsing error or
// if key is for the wrong network.
func DCRUtilAddressFromExtendedKey(key *hdkeychain.ExtendedKey, params *chaincfg.Params) (*dcrutil.AddressPubKeyHash, error) {
	return dcrutil.NewAddressPubKeyHash(dcrutil.Hash160(key.SerializedPubKey()), params, dcrec.STEcdsaSecp256k1)
}
