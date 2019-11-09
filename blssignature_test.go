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

package types_test

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	types "github.com/wealdtech/go-eth2-types"
)

func TestInvalidSignatureFromBytes(t *testing.T) {
	_, err := types.BLSSignatureFromBytes([]byte{0x00})
	require.NotNil(t, err)
}

func TestAggregate(t *testing.T) {
	pk1, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pk2, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pk3, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)

	msgs := make([][32]byte, 3)
	msg0 := [32]byte{}
	_, err = rand.Read(msg0[:])
	require.Nil(t, err)
	msgs[0] = msg0
	msg1 := [32]byte{}
	_, err = rand.Read(msg1[:])
	require.Nil(t, err)
	msgs[1] = msg1
	msg2 := [32]byte{}
	_, err = rand.Read(msg2[:])
	require.Nil(t, err)
	msgs[2] = msg2

	pubKeys := make([]types.PublicKey, 3)
	pubKeys[0] = pk1.PublicKey()
	pubKeys[1] = pk2.PublicKey()
	pubKeys[2] = pk3.PublicKey()

	domain := uint64(15432456)
	sigs := make([]types.Signature, 3)
	sigs[0] = pk1.Sign(msgs[0][:], domain)
	sigs[1] = pk2.Sign(msgs[1][:], domain)
	sigs[2] = pk3.Sign(msgs[2][:], domain)

	sig := types.AggregateSignatures(sigs)

	verified := sig.VerifyAggregate(msgs, pubKeys, domain)
	assert.True(t, verified)
}

func TestAggregateNoSigs(t *testing.T) {
	pk1, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)

	pubKeys := make([]types.PublicKey, 0)

	msg := []byte("A test message")
	domain := uint64(15432456)
	sig := pk1.Sign(msg, domain)

	verified := sig.VerifyAggregateCommon(msg, pubKeys, domain)
	assert.False(t, verified)

	msgs := make([][32]byte, 0)
	verified = sig.VerifyAggregate(msgs, pubKeys, domain)
	assert.False(t, verified)
}

func TestAggregateCommon(t *testing.T) {
	pk1, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pk2, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pk3, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)

	pubKeys := make([]types.PublicKey, 3)
	pubKeys[0] = pk1.PublicKey()
	pubKeys[1] = pk2.PublicKey()
	pubKeys[2] = pk3.PublicKey()

	msg := []byte("A test message")
	domain := uint64(15432456)
	sigs := make([]types.Signature, 3)
	sigs[0] = pk1.Sign(msg, domain)
	sigs[1] = pk2.Sign(msg, domain)
	sigs[2] = pk3.Sign(msg, domain)

	sig := types.AggregateSignatures(sigs)

	verified := sig.VerifyAggregateCommon(msg, pubKeys, domain)
	assert.True(t, verified)
}
