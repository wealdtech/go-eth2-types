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

package types_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e2types "github.com/wealdtech/go-eth2-types/v2"
)

func TestBLSPublicKeyFromBytes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   string
	}{
		{
			name:  "empty",
			input: "",
			err:   "public key must be 48 bytes",
		},
		{
			name:  "short",
			input: "a99a76ed7796f7be22d5b7e85deeb7c5677e88e511e0b337618f8c4eb61349b4bf2d153f649f7b53359fe8b94a38e4",
			err:   "public key must be 48 bytes",
		},
		{
			name:  "long",
			input: "a99a76ed7796f7be22d5b7e85deeb7c5677e88e511e0b337618f8c4eb61349b4bf2d153f649f7b53359fe8b94a38e44c4c",
			err:   "public key must be 48 bytes",
		},
		{
			name:  "invalid",
			input: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			err:   "failed to deserialize public key: err blsPublicKeyDeserialize ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		},
		{
			name:  "good",
			input: "a99a76ed7796f7be22d5b7e85deeb7c5677e88e511e0b337618f8c4eb61349b4bf2d153f649f7b53359fe8b94a38e44c",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes, err := hex.DecodeString(test.input)
			require.NoError(t, err)
			_, err = e2types.BLSPublicKeyFromBytes(bytes)
			if test.err == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, test.err)
			}
		})
	}
}

func TestBLSPublicKeyOperations(t *testing.T) {
	privKey1, err := e2types.GenerateBLSPrivateKey()
	require.Nil(t, err)

	pubKey1 := privKey1.PublicKey()
	bytes := pubKey1.Marshal()

	pubKey1Copy, err := e2types.BLSPublicKeyFromBytes(bytes)
	require.Nil(t, err)

	assert.Equal(t, pubKey1.Marshal(), pubKey1Copy.Marshal())

	privKey2, err := e2types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pubKey2 := privKey2.PublicKey()

	aggPubKey1 := pubKey1.Copy()
	aggPubKey1.Aggregate(pubKey2)
	aggPubKey2 := pubKey2.Copy()
	aggPubKey2.Aggregate(pubKey1)
	assert.Equal(t, aggPubKey1.Marshal(), aggPubKey2.Marshal())
}

func TestPublicKeyMarshal(t *testing.T) {
	privKey, err := e2types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pubKey := privKey.PublicKey()

	// Obtain the public key.
	val1 := pubKey.Marshal()
	// Obtain it again.
	val2 := pubKey.Marshal()
	// Mutate it.
	val2[0] = 0x00
	val2[1] = 0x00
	val2[2] = 0x00
	val2[3] = 0x00
	// Ensure that the mutation has not changed the marshalled data.
	val3 := pubKey.Marshal()
	assert.Equal(t, val1, val3)
}

func TestPublicKeyAggregate(t *testing.T) {
	privKey1, err := e2types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pubKey1 := privKey1.PublicKey()
	privKey2, err := e2types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pubKey2 := privKey2.PublicKey()

	// Obtain the public key bytes.
	bytes1 := pubKey1.Marshal()
	bytes2 := pubKey2.Marshal()

	// Aggregate the keys.
	pubKey1.Aggregate(pubKey2)

	// Ensure that the first key's marshalled data has changed.
	require.NotEqual(t, bytes1, pubKey1.Marshal())

	// Ensure the second key's marshalled data has not changed.
	require.Equal(t, bytes2, pubKey2.Marshal())
}

func BenchmarkMarshal(b *testing.B) {
	privKey, err := e2types.GenerateBLSPrivateKey()
	require.Nil(b, err)
	pubKey := privKey.PublicKey()

	for i := 0; i < b.N; i++ {
		pubKey.Marshal()
	}
}
