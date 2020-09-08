// +build !debug
// +build !2k
// +build !testground

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/miner"
)

const UpgradeBreezeHeight = 41280
const BreezeGasTampingDuration = 120

func init() {
	builtin.SealProofPolicies[abi.RegisteredSealProof_StackedDrg32GiBV1].ConsensusMinerMinPower = big.NewInt(10 << 40)
	builtin.SealProofPolicies[abi.RegisteredSealProof_StackedDrg64GiBV1].ConsensusMinerMinPower = big.NewInt(10 << 40)
	miner.SupportedProofTypes = map[abi.RegisteredSealProof]struct{}{
		abi.RegisteredSealProof_StackedDrg32GiBV1: {},
		abi.RegisteredSealProof_StackedDrg64GiBV1: {},
	}
}

const BlockDelaySecs = uint64(builtin.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
