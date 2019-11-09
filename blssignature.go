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

// BLSSignature is a BLS signature.
type BLSSignature struct {
	sig *g1.Signature
}

// BLSSignatureFromBytes creates a BLS signature from a byte slice.
func BLSSignatureFromBytes(data []byte) (Signature, error) {
	sig, err := g1.DeserializeSignature(bytesutil.ToBytes96(data))
	if err != nil {
		return nil, errors.Wrap(err, "failed to deserialize signature")
	}
	return &BLSSignature{sig: sig}, nil
}

// Verify a bls signature given a public key, a message, and a domain.
func (s *BLSSignature) Verify(msg []byte, pubKey PublicKey, domain uint64) bool {
	domainBytes := bytesutil.ToBytes8(bytesutil.Bytes8(domain))
	return g1.VerifyWithDomain(bytesutil.ToBytes32(msg), pubKey.(*BLSPublicKey).key, s.sig, domainBytes)
}

// VerifyAggregate verifies each public key against its respective message.
// Note: this is vulnerable to a rogue public-key attack.
func (s *BLSSignature) VerifyAggregate(msgs [][32]byte, pubKeys []PublicKey, domain uint64) bool {
	if len(pubKeys) == 0 {
		return false
	}
	var keys []*g1.PublicKey
	for _, v := range pubKeys {
		keys = append(keys, v.(*BLSPublicKey).key)
	}
	domainBytes := bytesutil.ToBytes8(bytesutil.Bytes8(domain))
	return s.sig.VerifyAggregateWithDomain(keys, msgs, domainBytes)
}

// VerifyAggregateCommon verifies each public key against a single message.
// Note: this is vulnerable to a rogue public-key attack.
func (s *BLSSignature) VerifyAggregateCommon(msg []byte, pubKeys []PublicKey, domain uint64) bool {
	if len(pubKeys) == 0 {
		return false
	}
	var keys []*g1.PublicKey
	for _, v := range pubKeys {
		keys = append(keys, v.(*BLSPublicKey).key)
	}
	domainBytes := bytesutil.ToBytes8(bytesutil.Bytes8(domain))
	return s.sig.VerifyAggregateCommonWithDomain(keys, bytesutil.ToBytes32(msg), domainBytes)
}

// Marshal a signature into a byte slice.
func (s *BLSSignature) Marshal() []byte {
	k := s.sig.Serialize()
	return k[:]
}

// AggregateSignatures aggregates a slice of signatures.
func AggregateSignatures(sigs []Signature) Signature {
	var ss []*g1.Signature
	for _, v := range sigs {
		if v == nil {
			continue
		}
		ss = append(ss, v.(*BLSSignature).sig)
	}
	return &BLSSignature{sig: g1.AggregateSignatures(ss)}
}
