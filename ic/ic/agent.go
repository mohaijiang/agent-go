// Package ic provides a client for the "ic" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package ic

import (
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/candid/marshal"
	"github.com/aviate-labs/agent-go/principal"
)

// Agent is a client for the "ic" canister.
type Agent struct {
	a          agent.Agent
	canisterId principal.Principal
}

// NewAgent creates a new agent for the "ic" canister.
func NewAgent(canisterId principal.Principal, config agent.Config) Agent {
	return Agent{
		a:          agent.New(config),
		canisterId: canisterId,
	}
}

// BitcoinGetBalance calls the "bitcoin_get_balance" method on the "ic" canister.
func (a Agent) BitcoinGetBalance(arg0 GetBalanceRequest) (*Satoshi, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 Satoshi
	if err := a.a.Call(
		a.canisterId,
		"bitcoin_get_balance",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// BitcoinGetCurrentFeePercentiles calls the "bitcoin_get_current_fee_percentiles" method on the "ic" canister.
func (a Agent) BitcoinGetCurrentFeePercentiles(arg0 GetCurrentFeePercentilesRequest) (*[]MillisatoshiPerByte, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 []MillisatoshiPerByte
	if err := a.a.Call(
		a.canisterId,
		"bitcoin_get_current_fee_percentiles",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// BitcoinGetUtxos calls the "bitcoin_get_utxos" method on the "ic" canister.
func (a Agent) BitcoinGetUtxos(arg0 GetUtxosRequest) (*GetUtxosResponse, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 GetUtxosResponse
	if err := a.a.Call(
		a.canisterId,
		"bitcoin_get_utxos",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// BitcoinSendTransaction calls the "bitcoin_send_transaction" method on the "ic" canister.
func (a Agent) BitcoinSendTransaction(arg0 SendTransactionRequest) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"bitcoin_send_transaction",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// CanisterStatus calls the "canister_status" method on the "ic" canister.
func (a Agent) CanisterStatus(arg0 struct {
	CanisterId CanisterId `ic:"canister_id"`
}) (*struct {
	Status struct {
		running  *struct{} `ic:"running,variant"`
		stopping *struct{} `ic:"stopping,variant"`
		stopped  *struct{} `ic:"stopped,variant"`
	} `ic:"status"`
	Settings               DefiniteCanisterSettings `ic:"settings"`
	ModuleHash             *[]byte                  `ic:"module_hash,omitempty"`
	MemorySize             idl.Nat                  `ic:"memory_size"`
	Cycles                 idl.Nat                  `ic:"cycles"`
	IdleCyclesBurnedPerDay idl.Nat                  `ic:"idle_cycles_burned_per_day"`
}, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		Status struct {
			running  *struct{} `ic:"running,variant"`
			stopping *struct{} `ic:"stopping,variant"`
			stopped  *struct{} `ic:"stopped,variant"`
		} `ic:"status"`
		Settings               DefiniteCanisterSettings `ic:"settings"`
		ModuleHash             *[]byte                  `ic:"module_hash,omitempty"`
		MemorySize             idl.Nat                  `ic:"memory_size"`
		Cycles                 idl.Nat                  `ic:"cycles"`
		IdleCyclesBurnedPerDay idl.Nat                  `ic:"idle_cycles_burned_per_day"`
	}
	if err := a.a.Call(
		a.canisterId,
		"canister_status",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// CreateCanister calls the "create_canister" method on the "ic" canister.
func (a Agent) CreateCanister(arg0 struct {
	Settings *CanisterSettings `ic:"settings,omitempty"`
}) (*struct {
	CanisterId CanisterId `ic:"canister_id"`
}, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		CanisterId CanisterId `ic:"canister_id"`
	}
	if err := a.a.Call(
		a.canisterId,
		"create_canister",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// DeleteCanister calls the "delete_canister" method on the "ic" canister.
func (a Agent) DeleteCanister(arg0 struct {
	CanisterId CanisterId `ic:"canister_id"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"delete_canister",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// DepositCycles calls the "deposit_cycles" method on the "ic" canister.
func (a Agent) DepositCycles(arg0 struct {
	CanisterId CanisterId `ic:"canister_id"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"deposit_cycles",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// EcdsaPublicKey calls the "ecdsa_public_key" method on the "ic" canister.
func (a Agent) EcdsaPublicKey(arg0 struct {
	CanisterId     *CanisterId `ic:"canister_id,omitempty"`
	DerivationPath [][]byte    `ic:"derivation_path"`
	KeyId          struct {
		Curve EcdsaCurve `ic:"curve"`
		Name  string     `ic:"name"`
	} `ic:"key_id"`
}) (*struct {
	PublicKey []byte `ic:"public_key"`
	ChainCode []byte `ic:"chain_code"`
}, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		PublicKey []byte `ic:"public_key"`
		ChainCode []byte `ic:"chain_code"`
	}
	if err := a.a.Call(
		a.canisterId,
		"ecdsa_public_key",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// HttpRequest calls the "http_request" method on the "ic" canister.
func (a Agent) HttpRequest(arg0 struct {
	Url              string  `ic:"url"`
	MaxResponseBytes *uint64 `ic:"max_response_bytes,omitempty"`
	Method           struct {
		get  *struct{} `ic:"get,variant"`
		head *struct{} `ic:"head,variant"`
		post *struct{} `ic:"post,variant"`
	} `ic:"method"`
	Headers   []HttpHeader `ic:"headers"`
	Body      *[]byte      `ic:"body,omitempty"`
	Transform *struct {
		Function struct { /* NOT SUPPORTED */
		} `ic:"function"`
		Context []byte `ic:"context"`
	} `ic:"transform,omitempty"`
}) (*HttpResponse, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 HttpResponse
	if err := a.a.Call(
		a.canisterId,
		"http_request",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// InstallCode calls the "install_code" method on the "ic" canister.
func (a Agent) InstallCode(arg0 struct {
	Mode struct {
		install   *struct{} `ic:"install,variant"`
		reinstall *struct{} `ic:"reinstall,variant"`
		upgrade   *struct{} `ic:"upgrade,variant"`
	} `ic:"mode"`
	CanisterId CanisterId `ic:"canister_id"`
	WasmModule WasmModule `ic:"wasm_module"`
	Arg        []byte     `ic:"arg"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"install_code",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// ProvisionalCreateCanisterWithCycles calls the "provisional_create_canister_with_cycles" method on the "ic" canister.
func (a Agent) ProvisionalCreateCanisterWithCycles(arg0 struct {
	Amount      *idl.Nat          `ic:"amount,omitempty"`
	Settings    *CanisterSettings `ic:"settings,omitempty"`
	SpecifiedId *CanisterId       `ic:"specified_id,omitempty"`
}) (*struct {
	CanisterId CanisterId `ic:"canister_id"`
}, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		CanisterId CanisterId `ic:"canister_id"`
	}
	if err := a.a.Call(
		a.canisterId,
		"provisional_create_canister_with_cycles",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ProvisionalTopUpCanister calls the "provisional_top_up_canister" method on the "ic" canister.
func (a Agent) ProvisionalTopUpCanister(arg0 struct {
	CanisterId CanisterId `ic:"canister_id"`
	Amount     idl.Nat    `ic:"amount"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"provisional_top_up_canister",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// RawRand calls the "raw_rand" method on the "ic" canister.
func (a Agent) RawRand() (*[]byte, error) {
	args, err := marshal.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 []byte
	if err := a.a.Call(
		a.canisterId,
		"raw_rand",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// SignWithEcdsa calls the "sign_with_ecdsa" method on the "ic" canister.
func (a Agent) SignWithEcdsa(arg0 struct {
	MessageHash    []byte   `ic:"message_hash"`
	DerivationPath [][]byte `ic:"derivation_path"`
	KeyId          struct {
		Curve EcdsaCurve `ic:"curve"`
		Name  string     `ic:"name"`
	} `ic:"key_id"`
}) (*struct {
	Signature []byte `ic:"signature"`
}, error) {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		Signature []byte `ic:"signature"`
	}
	if err := a.a.Call(
		a.canisterId,
		"sign_with_ecdsa",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// StartCanister calls the "start_canister" method on the "ic" canister.
func (a Agent) StartCanister(arg0 struct {
	CanisterId CanisterId `ic:"canister_id"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"start_canister",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// StopCanister calls the "stop_canister" method on the "ic" canister.
func (a Agent) StopCanister(arg0 struct {
	CanisterId CanisterId `ic:"canister_id"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"stop_canister",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// UninstallCode calls the "uninstall_code" method on the "ic" canister.
func (a Agent) UninstallCode(arg0 struct {
	CanisterId CanisterId `ic:"canister_id"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"uninstall_code",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// UpdateSettings calls the "update_settings" method on the "ic" canister.
func (a Agent) UpdateSettings(arg0 struct {
	CanisterId principal.Principal `ic:"canister_id"`
	Settings   CanisterSettings    `ic:"settings"`
}) error {
	args, err := marshal.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"update_settings",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

type BitcoinAddress = string

type BitcoinNetwork = struct {
	mainnet *struct{} `ic:"mainnet,variant"`
	testnet *struct{} `ic:"testnet,variant"`
}

type BlockHash = []byte

type CanisterId = principal.Principal

type CanisterSettings = struct {
	Controllers       *[]principal.Principal `ic:"controllers,omitempty"`
	ComputeAllocation *idl.Nat               `ic:"compute_allocation,omitempty"`
	MemoryAllocation  *idl.Nat               `ic:"memory_allocation,omitempty"`
	FreezingThreshold *idl.Nat               `ic:"freezing_threshold,omitempty"`
}

type DefiniteCanisterSettings = struct {
	Controllers       []principal.Principal `ic:"controllers"`
	ComputeAllocation idl.Nat               `ic:"compute_allocation"`
	MemoryAllocation  idl.Nat               `ic:"memory_allocation"`
	FreezingThreshold idl.Nat               `ic:"freezing_threshold"`
}

type EcdsaCurve = struct {
	secp256k1 *struct{} `ic:"secp256k1,variant"`
}

type GetBalanceRequest = struct {
	Address          BitcoinAddress `ic:"address"`
	Network          BitcoinNetwork `ic:"network"`
	MinConfirmations *uint32        `ic:"min_confirmations,omitempty"`
}

type GetCurrentFeePercentilesRequest = struct {
	Network BitcoinNetwork `ic:"network"`
}

type GetUtxosRequest = struct {
	Address BitcoinAddress `ic:"address"`
	Network BitcoinNetwork `ic:"network"`
	Filter  *struct {
		MinConfirmations *uint32 `ic:"min_confirmations,variant"`
		Page             *[]byte `ic:"page,variant"`
	} `ic:"filter,omitempty"`
}

type GetUtxosResponse = struct {
	Utxos        []Utxo    `ic:"utxos"`
	TipBlockHash BlockHash `ic:"tip_block_hash"`
	TipHeight    uint32    `ic:"tip_height"`
	NextPage     *[]byte   `ic:"next_page,omitempty"`
}

type HttpHeader = struct {
	Name  string `ic:"name"`
	Value string `ic:"value"`
}

type HttpResponse = struct {
	Status  idl.Nat      `ic:"status"`
	Headers []HttpHeader `ic:"headers"`
	Body    []byte       `ic:"body"`
}

type MillisatoshiPerByte = uint64

type Outpoint = struct {
	Txid []byte `ic:"txid"`
	Vout uint32 `ic:"vout"`
}

type Satoshi = uint64

type SendTransactionRequest = struct {
	Transaction []byte         `ic:"transaction"`
	Network     BitcoinNetwork `ic:"network"`
}

type UserId = principal.Principal

type Utxo = struct {
	Outpoint Outpoint `ic:"outpoint"`
	Value    Satoshi  `ic:"value"`
	Height   uint32   `ic:"height"`
}

type WasmModule = []byte
