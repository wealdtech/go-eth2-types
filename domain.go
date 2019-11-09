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
	"encoding/binary"
)

var (
	// DomainBeaconProposer is a domain constant.
	DomainBeaconProposer = []byte{0, 0, 0, 0}
	// DomainBeaconAttester is a domain constant.
	DomainBeaconAttester = []byte{1, 0, 0, 0}
	// DomainRANDAO is a domain constant.
	DomainRANDAO = []byte{2, 0, 0, 0}
	// DomainDeposit is a domain constant.
	DomainDeposit = []byte{3, 0, 0, 0}
	// DomainVoluntaryExit is a domain constant.
	DomainVoluntaryExit = []byte{4, 0, 0, 0}
)

// Domain returns a uint64 domain
func Domain(domainType []byte, forkVersion []byte) uint64 {
	res := make([]byte, 8)
	copy(res[0:4], domainType)
	copy(res[4:8], forkVersion)
	return binary.LittleEndian.Uint64(res)
}
