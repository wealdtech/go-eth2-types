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
	"crypto/rand"
	"encoding/binary"

	g1 "github.com/phoreproject/bls/g1pubs"
	"github.com/pkg/errors"
	bytesutil "github.com/wealdtech/go-bytesutil"
)

// BLSPrivateKey is a private key in Ethereum 2.
// It is a point on the BLS12-381 curve.
type BLSPrivateKey struct {
	key *g1.SecretKey
}

// BLSPrivateKeyFromBytes creates a BLS private key from a byte slice.
func BLSPrivateKeyFromBytes(priv []byte) (*BLSPrivateKey, error) {
	if len(priv) != 32 {
		return nil, errors.New("private key must be 32 bytes")
	}
	val := g1.DeserializeSecretKey(bytesutil.ToBytes32(priv))
	if val.GetFRElement() == nil {
		return nil, errors.New("invalid private key")
	}
	return &BLSPrivateKey{key: val}, nil
}

// GenerateBLSPrivateKey generates a random BLS private key.
func GenerateBLSPrivateKey() (*BLSPrivateKey, error) {
	key, err := g1.RandKey(rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate private key")
	}
	return &BLSPrivateKey{key: key}, nil
}

// Marshal a secret key into a byte slice.
func (p *BLSPrivateKey) Marshal() []byte {
	k := p.key.Serialize()
	return k[:]
}

// PublicKey obtains the public key corresponding to the BLS secret key.
func (p *BLSPrivateKey) PublicKey() PublicKey {
	return &BLSPublicKey{key: g1.PrivToPub(p.key)}
}

// Sign a message using a secret key.
func (p *BLSPrivateKey) Sign(msg []byte, domain uint64) Signature {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, domain)
	sig := g1.SignWithDomain(bytesutil.ToBytes32(msg), p.key, bytesutil.ToBytes8(b))
	return &BLSSignature{sig}
}
