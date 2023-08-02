// Copyright © 2019 - 2023 Weald Technology Trading.
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

import "github.com/pkg/errors"

// DomainType defines the type of the domain, as per https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/beacon-chain.md#custom-types
type DomainType [4]byte

// ZeroForkVersion is an empty fork version.
var ZeroForkVersion = []byte{0, 0, 0, 0}

// ZeroGenesisValidatorsRoot is an empty genesis validators root.
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
	DomainSelectionProof = DomainType{5, 0, 0, 0}
	// DomainAggregateAndProof is a domain constant.
	DomainAggregateAndProof = DomainType{6, 0, 0, 0}
	// DomainSyncCommittee is a domain constant.
	DomainSyncCommittee = DomainType{7, 0, 0, 0}
	// DomainSyncCommitteeSelectionProof is a domain constant.
	DomainSyncCommitteeSelectionProof = DomainType{8, 0, 0, 0}
	// DomainContributionAndProof is a domain constant.
	DomainContributionAndProof = DomainType{9, 0, 0, 0}
	// DomainBlsToExecutionChange is a domain constant.
	DomainBlsToExecutionChange = DomainType{0x0A, 0, 0, 0}
	// DomainBlobSidecar is a domain constant.
	DomainBlobSidecar = DomainType{0x0B, 0, 0, 0}
)

// ComputeDomain computes a domain.
func ComputeDomain(domainType DomainType, forkVersion []byte, genesisValidatorsRoot []byte) ([]byte, error) {
	if len(forkVersion) != 4 {
		return nil, errors.New("fork version must be 4 bytes in length")
	}
	if len(genesisValidatorsRoot) != 32 {
		return nil, errors.New("genesis validators root must be 32 bytes in length")
	}

	// Generate fork data root from fork version and genesis validators root.
	forkData := &ForkData{
		CurrentVersion:        forkVersion,
		GenesisValidatorsRoot: genesisValidatorsRoot,
	}
	forkDataRoot, err := forkData.HashTreeRoot()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate fork data hash tree root")
	}

	res := make([]byte, 32)
	copy(res[0:4], domainType[:])
	copy(res[4:32], forkDataRoot[:])

	return res, nil
}

// Domain returns a complete domain.
// Deprecated: due to panicking on error.  Use ComputeDomain() instead.
func Domain(domainType DomainType, forkVersion []byte, genesisValidatorsRoot []byte) []byte {
	// Generate fork data root from fork version and genesis validators root.
	forkData := &ForkData{
		CurrentVersion:        forkVersion,
		GenesisValidatorsRoot: genesisValidatorsRoot,
	}
	forkDataRoot, err := forkData.HashTreeRoot()
	if err != nil {
		panic(err)
	}

	res := make([]byte, 32)
	copy(res[0:4], domainType[:])
	copy(res[4:32], forkDataRoot[:])

	return res
}
