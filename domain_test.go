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
	"testing"

	"github.com/stretchr/testify/assert"
	e2types "github.com/wealdtech/go-eth2-types/v2"
)

func _hexStringToBytes(input string) []byte {
	res, _ := hex.DecodeString(input)
	return res
}

func TestDomain(t *testing.T) {
	genesisRoot := _hexStringToBytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	domain := e2types.Domain(e2types.DomainVoluntaryExit, []byte{0x01, 0x02, 0x03, 0x04}, genesisRoot)

	expectedDomain := _hexStringToBytes("0400000001020304ffffffffffffffffffffffffffffffffffffffffffffffff")
	assert.Equal(t, expectedDomain, domain)
}
