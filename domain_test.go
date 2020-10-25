// Copyright © 2020 Weald Technology Trading
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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e2types "github.com/wealdtech/go-eth2-types/v2"
)

func _hexStringToBytes(input string) []byte {
	res, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic(err)
	}
	return res
}

func TestDomain(t *testing.T) {
	genesisValidatorsRoot := _hexStringToBytes("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	domain := e2types.Domain(e2types.DomainDeposit, []byte{0x01, 0x02, 0x03, 0x04}, genesisValidatorsRoot)

	expectedDomain := _hexStringToBytes("0x03000000d1b9515995b783401c69f4b529f86de082d38f078019c37e6262ecb5")
	assert.Equal(t, expectedDomain, domain)
}

func TestComputeDomain(t *testing.T) {
	tests := []struct {
		name                  string
		domainType            e2types.DomainType
		forkVersion           []byte
		genesisValidatorsRoot []byte
		err                   string
		res                   []byte
	}{
		{
			name:                  "ForkVersionMissing",
			domainType:            e2types.DomainDeposit,
			genesisValidatorsRoot: _hexStringToBytes("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			err:                   "fork version must be 4 bytes in length",
		},
		{
			name:                  "ForkVersionShort",
			domainType:            e2types.DomainDeposit,
			forkVersion:           _hexStringToBytes("0x010203"),
			genesisValidatorsRoot: _hexStringToBytes("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			err:                   "fork version must be 4 bytes in length",
		},
		{
			name:                  "ForkVersionLong",
			domainType:            e2types.DomainDeposit,
			forkVersion:           _hexStringToBytes("0x0102030405"),
			genesisValidatorsRoot: _hexStringToBytes("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			err:                   "fork version must be 4 bytes in length",
		},
		{
			name:        "GenesisValidatorsRootMissing",
			domainType:  e2types.DomainDeposit,
			forkVersion: _hexStringToBytes("0x01020304"),
			err:         "genesis validators root must be 32 bytes in length",
		},
		{
			name:                  "GenesisValidatorsRootShort",
			domainType:            e2types.DomainDeposit,
			forkVersion:           _hexStringToBytes("0x01020304"),
			genesisValidatorsRoot: _hexStringToBytes("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"),
			err:                   "genesis validators root must be 32 bytes in length",
		},
		{
			name:                  "GenesisValidatorsRootLong",
			domainType:            e2types.DomainDeposit,
			forkVersion:           _hexStringToBytes("0x01020304"),
			genesisValidatorsRoot: _hexStringToBytes("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f2021"),
			err:                   "genesis validators root must be 32 bytes in length",
		},
		{
			name:                  "Good",
			domainType:            e2types.DomainDeposit,
			forkVersion:           _hexStringToBytes("0x01020304"),
			genesisValidatorsRoot: _hexStringToBytes("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			res:                   _hexStringToBytes("0x03000000d1b9515995b783401c69f4b529f86de082d38f078019c37e6262ecb5"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := e2types.ComputeDomain(test.domainType, test.forkVersion, test.genesisValidatorsRoot)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.res, res)
			}
		})
	}
}
