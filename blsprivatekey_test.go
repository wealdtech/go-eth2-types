// Copyright 2019, 2020 Weald Technology Trading
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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	types "github.com/wealdtech/go-eth2-types"
)

func _byteArray(input string) []byte {
	res, _ := hex.DecodeString(input)
	return res
}

func _blsPrivateKey(input string) *types.BLSPrivateKey {
	data, _ := hex.DecodeString(input)
	res, _ := types.BLSPrivateKeyFromBytes(data)
	return res
}

func TestBLSPrivateKeyFromBytes(t *testing.T) {
	goodBytes, err := hex.DecodeString("25295f0d1d592a90b333e26e85149708208e9f8e8bc18f6c77bd62f8ad7a6866")
	require.Nil(t, err)
	_, err = types.BLSPrivateKeyFromBytes(goodBytes)
	assert.Nil(t, err)

	badBytes, err := hex.DecodeString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	require.Nil(t, err)
	_, err = types.BLSPrivateKeyFromBytes(badBytes)
	assert.NotNil(t, err)
}

func TestBLSSignature(t *testing.T) {
	tests := []struct {
		name string
		key  *types.BLSPrivateKey
		msg  []byte
		err  error
	}{
		{
			name: "Nil",
			key:  _blsPrivateKey("25295f0d1d592a90b333e26e85149708208e9f8e8bc18f6c77bd62f8ad7a6866"),
			msg:  _byteArray("0102030405060708090a0b0c0d0e0f"),
			err:  errors.New("no path"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sig := test.key.Sign(test.msg)
			verified := sig.Verify(test.msg, test.key.PublicKey())
			assert.Equal(t, verified, true)

			sig2, err := types.BLSSignatureFromBytes(sig.Marshal())
			require.Nil(t, err)
			assert.Equal(t, sig.Marshal(), sig2.Marshal())
		})
	}
}

func TestGenerateBLSPrivateKey(t *testing.T) {
	key, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	assert.NotNil(t, key)
	assert.NotNil(t, key.Marshal())

	_, err = types.BLSPrivateKeyFromBytes(key.Marshal()[:31])
	assert.NotNil(t, err)
}
