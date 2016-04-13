package tester

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

type VMEnv struct {
	state *state.StateDB
	block *types.Block

	transactor *common.Address
	value      *big.Int

	depth int
	Gas   *big.Int
	time  *big.Int
	logs  []vm.StructLog

	evm *vm.EVM
}

func NewEnv(state *state.StateDB, transactor common.Address, value *big.Int, cfg vm.Config) *VMEnv {
	env := &VMEnv{
		state:      state,
		transactor: &transactor,
		value:      value,
		time:       big.NewInt(time.Now().Unix()),
	}
	cfg.Logger.Collector = env

	env.evm = vm.New(env, cfg)
	return env
}

// ruleSet implements vm.RuleSet and will always default to the homestead rule set.
type ruleSet struct{}

func (ruleSet) IsHomestead(*big.Int) bool { return true }

func (self *VMEnv) RuleSet() vm.RuleSet        { return ruleSet{} }
func (self *VMEnv) Vm() vm.Vm                  { return self.evm }
func (self *VMEnv) Db() vm.Database            { return self.state }
func (self *VMEnv) MakeSnapshot() vm.Database  { return self.state.Copy() }
func (self *VMEnv) SetSnapshot(db vm.Database) { self.state.Set(db.(*state.StateDB)) }
func (self *VMEnv) Origin() common.Address     { return *self.transactor }
func (self *VMEnv) BlockNumber() *big.Int      { return common.Big0 }
func (self *VMEnv) Coinbase() common.Address   { return *self.transactor }
func (self *VMEnv) Time() *big.Int             { return self.time }
func (self *VMEnv) Difficulty() *big.Int       { return common.Big1 }
func (self *VMEnv) BlockHash() []byte          { return make([]byte, 32) }
func (self *VMEnv) Value() *big.Int            { return self.value }
func (self *VMEnv) GasLimit() *big.Int         { return big.NewInt(1000000000) }
func (self *VMEnv) VmType() vm.Type            { return vm.StdVmTy }
func (self *VMEnv) Depth() int                 { return 0 }
func (self *VMEnv) SetDepth(i int)             { self.depth = i }
func (self *VMEnv) GetHash(n uint64) common.Hash {
	if self.block.Number().Cmp(big.NewInt(int64(n))) == 0 {
		return self.block.Hash()
	}
	return common.Hash{}
}
func (self *VMEnv) AddStructLog(log vm.StructLog) {
	self.logs = append(self.logs, log)
}
func (self *VMEnv) StructLogs() []vm.StructLog {
	return self.logs
}

func (self *VMEnv) AddLog(log *vm.Log) {
	self.state.AddLog(log)
}
func (self *VMEnv) CanTransfer(from common.Address, balance *big.Int) bool {
	return self.state.GetBalance(from).Cmp(balance) >= 0
}
func (self *VMEnv) Transfer(from, to vm.Account, amount *big.Int) {
	core.Transfer(from, to, amount)
}

func (self *VMEnv) Call(caller vm.ContractRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	self.Gas = gas
	return core.Call(self, caller, addr, data, gas, price, value)
}

func (self *VMEnv) CallCode(caller vm.ContractRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	return core.CallCode(self, caller, addr, data, gas, price, value)
}

func (self *VMEnv) DelegateCall(caller vm.ContractRef, addr common.Address, data []byte, gas, price *big.Int) ([]byte, error) {
	return core.DelegateCall(self, caller, addr, data, gas, price)
}

func (self *VMEnv) Create(caller vm.ContractRef, data []byte, gas, price, value *big.Int) ([]byte, common.Address, error) {
	return core.Create(self, caller, data, gas, price, value)
}
