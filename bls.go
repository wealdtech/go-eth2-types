// Copyright 2020 - 2023 Weald Technology Trading.
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

// Package types provides generic types for the Ethereum consensus system.
package types

import (
	bls "github.com/herumi/bls-eth-go-binary/bls"
)

// InitBLS initialises the BLS library with the appropriate curve and parameters for Ethereum 2.
func InitBLS() error {
	if err := bls.Init(bls.BLS12_381); err != nil {
		return err
	}
	return bls.SetETHmode(bls.EthModeDraft07)
}
