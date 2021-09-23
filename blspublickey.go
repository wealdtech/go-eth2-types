// Copyright 2019 - 2021 Weald Technology Trading.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"fmt"
	"sync"

	bls "github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
)

// Size of an Ethereum BLS public key, in bytes.
var blsPubKeySize = 48

// BLSPublicKey used in the BLS signature scheme.
type BLSPublicKey struct {
	key        *bls.PublicKey
	serialized []byte
	accessMu   sync.RWMutex
}

// BLSPublicKeyFromBytes creates a BLS public key from a byte slice.
func BLSPublicKeyFromBytes(pub []byte) (*BLSPublicKey, error) {
	if len(pub) != blsPubKeySize {
		return nil, fmt.Errorf("public key must be %d bytes", blsPubKeySize)
	}
	var key bls.PublicKey
	if err := key.Deserialize(pub); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize public key")
	}
	return &BLSPublicKey{
		key: &key,
	}, nil
}

// Aggregate two public keys.  This updates the value of the existing key.
func (k *BLSPublicKey) Aggregate(other PublicKey) {
	k.accessMu.Lock()
	k.key.Add(other.(*BLSPublicKey).key)
	k.serialized = nil
	k.accessMu.Unlock()
}

// Marshal a BLS public key into a byte slice.
func (k *BLSPublicKey) Marshal() []byte {
	k.accessMu.Lock()
	if k.serialized == nil {
		k.serialized = k.key.Serialize()
	}
	res := make([]byte, blsPubKeySize)
	copy(res, k.serialized)
	k.accessMu.Unlock()

	return res
}

// Copy creates a copy of the public key.
func (k *BLSPublicKey) Copy() PublicKey {
	k.accessMu.Lock()

	if k.serialized == nil {
		k.serialized = k.key.Serialize()
	}

	var newKey bls.PublicKey
	//#nosec G104
	_ = newKey.Deserialize(k.serialized)

	key := &BLSPublicKey{
		key: &newKey,
	}

	if k.serialized != nil {
		key.serialized = make([]byte, blsPubKeySize)
		copy(key.serialized, k.serialized)
	}

	k.accessMu.Unlock()

	return key
}
