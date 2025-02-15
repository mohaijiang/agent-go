// Package icrc1 provides a client for the "icrc1" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package icrc1

import (
	"github.com/mohaijiang/agent-go"
	"github.com/mohaijiang/agent-go/candid/idl"
	"github.com/mohaijiang/agent-go/principal"
)

type Account = struct {
	Owner      principal.Principal `ic:"owner"`
	Subaccount *Subaccount         `ic:"subaccount,omitempty"`
}

// Agent is a client for the "icrc1" canister.
type Agent struct {
	a          *agent.Agent
	canisterId principal.Principal
}

// NewAgent creates a new agent for the "icrc1" canister.
func NewAgent(canisterId principal.Principal, config agent.Config) (*Agent, error) {
	a, err := agent.New(config)
	if err != nil {
		return nil, err
	}
	return &Agent{
		a:          a,
		canisterId: canisterId,
	}, nil
}

// Icrc1BalanceOf calls the "icrc1_balance_of" method on the "icrc1" canister.
func (a Agent) Icrc1BalanceOf(arg0 Account) (*idl.Nat, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 idl.Nat
	if err := a.a.Query(
		a.canisterId,
		"icrc1_balance_of",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Decimals calls the "icrc1_decimals" method on the "icrc1" canister.
func (a Agent) Icrc1Decimals() (*uint8, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 uint8
	if err := a.a.Query(
		a.canisterId,
		"icrc1_decimals",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Fee calls the "icrc1_fee" method on the "icrc1" canister.
func (a Agent) Icrc1Fee() (*idl.Nat, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 idl.Nat
	if err := a.a.Query(
		a.canisterId,
		"icrc1_fee",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Metadata calls the "icrc1_metadata" method on the "icrc1" canister.
func (a Agent) Icrc1Metadata() (*[]struct {
	field0 string `ic:"field0"`
	field1 Value  `ic:"field1"`
}, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 []struct {
		field0 string `ic:"field0"`
		field1 Value  `ic:"field1"`
	}
	if err := a.a.Query(
		a.canisterId,
		"icrc1_metadata",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1MintingAccount calls the "icrc1_minting_account" method on the "icrc1" canister.
func (a Agent) Icrc1MintingAccount() (**Account, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 *Account
	if err := a.a.Query(
		a.canisterId,
		"icrc1_minting_account",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Name calls the "icrc1_name" method on the "icrc1" canister.
func (a Agent) Icrc1Name() (*string, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 string
	if err := a.a.Query(
		a.canisterId,
		"icrc1_name",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1SupportedStandards calls the "icrc1_supported_standards" method on the "icrc1" canister.
func (a Agent) Icrc1SupportedStandards() (*[]struct {
	Name string `ic:"name"`
	Url  string `ic:"url"`
}, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 []struct {
		Name string `ic:"name"`
		Url  string `ic:"url"`
	}
	if err := a.a.Query(
		a.canisterId,
		"icrc1_supported_standards",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Symbol calls the "icrc1_symbol" method on the "icrc1" canister.
func (a Agent) Icrc1Symbol() (*string, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 string
	if err := a.a.Query(
		a.canisterId,
		"icrc1_symbol",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1TotalSupply calls the "icrc1_total_supply" method on the "icrc1" canister.
func (a Agent) Icrc1TotalSupply() (*idl.Nat, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 idl.Nat
	if err := a.a.Query(
		a.canisterId,
		"icrc1_total_supply",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Icrc1Transfer calls the "icrc1_transfer" method on the "icrc1" canister.
func (a Agent) Icrc1Transfer(arg0 TransferArgs) (*struct {
	Ok  *idl.Nat       `ic:"Ok,variant"`
	Err *TransferError `ic:"Err,variant"`
}, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		Ok  *idl.Nat       `ic:"Ok,variant"`
		Err *TransferError `ic:"Err,variant"`
	}
	if err := a.a.Call(
		a.canisterId,
		"icrc1_transfer",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type Duration = uint64

type Subaccount = []byte

type Timestamp = uint64

type TransferArgs = struct {
	FromSubaccount *Subaccount `ic:"from_subaccount,omitempty"`
	To             Account     `ic:"to"`
	Amount         idl.Nat     `ic:"amount"`
	Fee            *idl.Nat    `ic:"fee,omitempty"`
	Memo           *[]byte     `ic:"memo,omitempty"`
	CreatedAtTime  *Timestamp  `ic:"created_at_time,omitempty"`
}

type TransferError = struct {
	BadFee *struct {
		ExpectedFee idl.Nat `ic:"expected_fee"`
	} `ic:"BadFee,variant"`
	BadBurn *struct {
		MinBurnAmount idl.Nat `ic:"min_burn_amount"`
	} `ic:"BadBurn,variant"`
	InsufficientFunds *struct {
		Balance idl.Nat `ic:"balance"`
	} `ic:"InsufficientFunds,variant"`
	TooOld          *struct{} `ic:"TooOld,variant"`
	CreatedInFuture *struct {
		LedgerTime Timestamp `ic:"ledger_time"`
	} `ic:"CreatedInFuture,variant"`
	Duplicate *struct {
		DuplicateOf idl.Nat `ic:"duplicate_of"`
	} `ic:"Duplicate,variant"`
	TemporarilyUnavailable *struct{} `ic:"TemporarilyUnavailable,variant"`
	GenericError           *struct {
		ErrorCode idl.Nat `ic:"error_code"`
		Message   string  `ic:"message"`
	} `ic:"GenericError,variant"`
}

type Value = struct {
	Nat  *idl.Nat `ic:"Nat,variant"`
	Int  *idl.Int `ic:"Int,variant"`
	Text *string  `ic:"Text,variant"`
	Blob *[]byte  `ic:"Blob,variant"`
}
