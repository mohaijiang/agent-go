// Package cmc provides a client for the "cmc" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package cmc

import (
	"github.com/mohaijiang/agent-go"
	"github.com/mohaijiang/agent-go/candid/idl"
	"github.com/mohaijiang/agent-go/principal"
)

type AccountIdentifier = struct {
	Bytes []byte `ic:"bytes"`
}

// Agent is a client for the "cmc" canister.
type Agent struct {
	a          *agent.Agent
	canisterId principal.Principal
}

// NewAgent creates a new agent for the "cmc" canister.
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

// GetIcpXdrConversionRate calls the "get_icp_xdr_conversion_rate" method on the "cmc" canister.
func (a Agent) GetIcpXdrConversionRate() (*IcpXdrConversionRateResponse, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 IcpXdrConversionRateResponse
	if err := a.a.Query(
		a.canisterId,
		"get_icp_xdr_conversion_rate",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetSubnetTypesToSubnets calls the "get_subnet_types_to_subnets" method on the "cmc" canister.
func (a Agent) GetSubnetTypesToSubnets() (*SubnetTypesToSubnetsResponse, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 SubnetTypesToSubnetsResponse
	if err := a.a.Query(
		a.canisterId,
		"get_subnet_types_to_subnets",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// NotifyCreateCanister calls the "notify_create_canister" method on the "cmc" canister.
func (a Agent) NotifyCreateCanister(arg0 NotifyCreateCanisterArg) (*NotifyCreateCanisterResult, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 NotifyCreateCanisterResult
	if err := a.a.Call(
		a.canisterId,
		"notify_create_canister",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// NotifyTopUp calls the "notify_top_up" method on the "cmc" canister.
func (a Agent) NotifyTopUp(arg0 NotifyTopUpArg) (*NotifyTopUpResult, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 NotifyTopUpResult
	if err := a.a.Call(
		a.canisterId,
		"notify_top_up",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type BlockIndex = uint64

type Cycles = idl.Nat

type CyclesCanisterInitPayload = struct {
	LedgerCanisterId       *principal.Principal  `ic:"ledger_canister_id,omitempty"`
	GovernanceCanisterId   *principal.Principal  `ic:"governance_canister_id,omitempty"`
	MintingAccountId       *AccountIdentifier    `ic:"minting_account_id,omitempty"`
	LastPurgedNotification *uint64               `ic:"last_purged_notification,omitempty"`
	ExchangeRateCanister   *ExchangeRateCanister `ic:"exchange_rate_canister,omitempty"`
}

type ExchangeRateCanister = struct {
	Set   *principal.Principal `ic:"Set,variant"`
	Unset *struct{}            `ic:"Unset,variant"`
}

type IcpXdrConversionRate = struct {
	TimestampSeconds   uint64 `ic:"timestamp_seconds"`
	XdrPermyriadPerIcp uint64 `ic:"xdr_permyriad_per_icp"`
}

type IcpXdrConversionRateResponse = struct {
	Data        IcpXdrConversionRate `ic:"data"`
	HashTree    []byte               `ic:"hash_tree"`
	Certificate []byte               `ic:"certificate"`
}

type NotifyCreateCanisterArg = struct {
	BlockIndex BlockIndex          `ic:"block_index"`
	Controller principal.Principal `ic:"controller"`
	SubnetType *string             `ic:"subnet_type,omitempty"`
}

type NotifyCreateCanisterResult = struct {
	Ok  *principal.Principal `ic:"Ok,variant"`
	Err *NotifyError         `ic:"Err,variant"`
}

type NotifyError = struct {
	Refunded *struct {
		Reason     string      `ic:"reason"`
		BlockIndex *BlockIndex `ic:"block_index,omitempty"`
	} `ic:"Refunded,variant"`
	Processing         *struct{}   `ic:"Processing,variant"`
	TransactionTooOld  *BlockIndex `ic:"TransactionTooOld,variant"`
	InvalidTransaction *string     `ic:"InvalidTransaction,variant"`
	Other              *struct {
		ErrorCode    uint64 `ic:"error_code"`
		ErrorMessage string `ic:"error_message"`
	} `ic:"Other,variant"`
}

type NotifyTopUpArg = struct {
	BlockIndex BlockIndex          `ic:"block_index"`
	CanisterId principal.Principal `ic:"canister_id"`
}

type NotifyTopUpResult = struct {
	Ok  *Cycles      `ic:"Ok,variant"`
	Err *NotifyError `ic:"Err,variant"`
}

type SubnetTypesToSubnetsResponse = struct {
	Data []struct {
		field0 string                `ic:"field0"`
		field1 []principal.Principal `ic:"field1"`
	} `ic:"data"`
}
