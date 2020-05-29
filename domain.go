// Copyright © 2019, 2020 Weald Technology Trading
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

import "github.com/prysmaticlabs/go-ssz"

// DomainType defines the type of the domain, as per https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
type DomainType [4]byte

// ZeroForkVersion is used where there is no requirement for a fork version, e.g. deposits.
var ZeroForkVersion = []byte{0, 0, 0, 0}

// ZeroGenesisValidatorsRoot is used where there is no requirement for a genesis validators root, e.g. deposits.
var ZeroGenesisValidatorsRoot = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var (
	// DomainBeaconProposer is a domain constant.
	DomainBeaconProposer = DomainType{0, 0, 0, 0}
	// DomainBeaconAttester is a domain constant.
	DomainBeaconAttester = DomainType{1, 0, 0, 0}
	// DomainRANDAO is a domain constant.
	DomainRANDAO = DomainType{2, 0, 0, 0}
	// DomainDeposit is a domain constant.
	DomainDeposit = DomainType{3, 0, 0, 0}
	// DomainVoluntaryExit is a domain constant.
	DomainVoluntaryExit = DomainType{4, 0, 0, 0}
	// DomainSelectionProof is a domain constant.
	DomainSelectionProof = []byte{5, 0, 0, 0}
	// DomainAggregateAndProof is a domain constant.
	DomainAggregateAndProof = []byte{6, 0, 0, 0}
)

// Domain returns a complete domain.
func Domain(domainType DomainType, forkVersion []byte, genesisValidatorsRoot []byte) []byte {
	// Generate fork data root from fork version and genesis validators root.
	forkDataRoot, err := ssz.HashTreeRoot(struct {
		CurrentVersion        []byte `ssz-size:"4"`
		GenesisValidatorsRoot []byte `ssz-size:"32"`
	}{
		CurrentVersion:        forkVersion,
		GenesisValidatorsRoot: genesisValidatorsRoot,
	})
	if err != nil {
		panic(err)
	}

	res := make([]byte, 32)
	copy(res[0:4], domainType[:])
	copy(res[4:32], forkDataRoot[:])

	return res
}
