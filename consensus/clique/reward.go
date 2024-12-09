package clique

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

const (
	decimals = 1e18
)

// reward 为 Clique 机制中的奖励分配函数
func (c *Clique) reward(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB) {
	coinbases := chain.Config().Clique.Coinbases
	coinbaseLen := len(coinbases)
	if coinbaseLen != 0 {
		index := header.Number.Uint64() % uint64(coinbaseLen)
		state.AddBalance(common.HexToAddress(coinbases[index]), uint256.NewInt(chain.Config().Clique.Reward*decimals))
	}
}
