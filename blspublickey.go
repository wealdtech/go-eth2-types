// Copyright © 2019 Weald Technology Trading
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
	g1 "github.com/phoreproject/bls/g1pubs"
	"github.com/pkg/errors"
	bytesutil "github.com/wealdtech/go-bytesutil"
)

// BLSPublicKey used in the BLS signature scheme.
type BLSPublicKey struct {
	key *g1.PublicKey
}

// BLSPublicKeyFromBytes creates a BLS public key from a byte slice.
func BLSPublicKeyFromBytes(pub []byte) (*BLSPublicKey, error) {
	if len(pub) != 48 {
		return nil, errors.New("public key must be 48 bytes")
	}
	key, err := g1.DeserializePublicKey(bytesutil.ToBytes48(pub))
	if err != nil {
		return nil, errors.Wrap(err, "failed to deserialize public key")
	}
	return &BLSPublicKey{key: key}, nil
}

// Aggregate two public keys.
func (p *BLSPublicKey) Aggregate(other PublicKey) PublicKey {
	pubKey := p.key.Copy()
	pubKey.Aggregate(other.(*BLSPublicKey).key)
	return &BLSPublicKey{key: pubKey}
}

// Marshal a BLS public key into a byte slice.
func (p *BLSPublicKey) Marshal() []byte {
	data := p.key.Serialize()
	return data[:]
}
