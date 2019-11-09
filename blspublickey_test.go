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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	types "github.com/wealdtech/go-eth2-types"
)

func TestBLSPublicKey(t *testing.T) {
	privKey1, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)

	pubKey1 := privKey1.PublicKey()
	bytes := pubKey1.Marshal()

	pubKey1Copy, err := types.BLSPublicKeyFromBytes(bytes)
	require.Nil(t, err)

	assert.Equal(t, pubKey1.Marshal(), pubKey1Copy.Marshal())

	_, err = types.BLSPublicKeyFromBytes(bytes[:46])
	require.NotNil(t, err)

	privKey2, err := types.GenerateBLSPrivateKey()
	require.Nil(t, err)
	pubKey2 := privKey2.PublicKey()

	aggPubKey1 := pubKey1.Aggregate(pubKey2)
	aggPubKey2 := pubKey2.Aggregate(pubKey1)
	assert.Equal(t, aggPubKey1.Marshal(), aggPubKey2.Marshal())
}
