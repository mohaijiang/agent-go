// Package icparchive provides a client for the "icparchive" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package icparchive

import (
	"github.com/mohaijiang/agent-go"
	"github.com/mohaijiang/agent-go/candid/idl"
	"github.com/mohaijiang/agent-go/principal"
)

type AccountIdentifier = []byte

// Agent is a client for the "icparchive" canister.
type Agent struct {
	a          *agent.Agent
	canisterId principal.Principal
}

// NewAgent creates a new agent for the "icparchive" canister.
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

// GetBlocks calls the "get_blocks" method on the "icparchive" canister.
func (a Agent) GetBlocks(arg0 GetBlocksArgs) (*GetBlocksResult, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 GetBlocksResult
	if err := a.a.Query(
		a.canisterId,
		"get_blocks",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type Block = struct {
	ParentHash  *[]byte     `ic:"parent_hash,omitempty"`
	Transaction Transaction `ic:"transaction"`
	Timestamp   Timestamp   `ic:"timestamp"`
}

type BlockIndex = uint64

type BlockRange = struct {
	Blocks []Block `ic:"blocks"`
}

type GetBlocksArgs = struct {
	Start  BlockIndex `ic:"start"`
	Length uint64     `ic:"length"`
}

type GetBlocksError = struct {
	BadFirstBlockIndex *struct {
		RequestedIndex  BlockIndex `ic:"requested_index"`
		FirstValidIndex BlockIndex `ic:"first_valid_index"`
	} `ic:"BadFirstBlockIndex,variant"`
	Other *struct {
		ErrorCode    uint64 `ic:"error_code"`
		ErrorMessage string `ic:"error_message"`
	} `ic:"Other,variant"`
}

type GetBlocksResult = struct {
	Ok  *BlockRange     `ic:"Ok,variant"`
	Err *GetBlocksError `ic:"Err,variant"`
}

type Memo = uint64

type Operation = struct {
	Mint *struct {
		To     AccountIdentifier `ic:"to"`
		Amount Tokens            `ic:"amount"`
	} `ic:"Mint,variant"`
	Burn *struct {
		From   AccountIdentifier `ic:"from"`
		Amount Tokens            `ic:"amount"`
	} `ic:"Burn,variant"`
	Transfer *struct {
		From   AccountIdentifier `ic:"from"`
		To     AccountIdentifier `ic:"to"`
		Amount Tokens            `ic:"amount"`
		Fee    Tokens            `ic:"fee"`
	} `ic:"Transfer,variant"`
	Approve *struct {
		From         AccountIdentifier `ic:"from"`
		Spender      AccountIdentifier `ic:"spender"`
		AllowanceE8s idl.Int           `ic:"allowance_e8s"`
		Fee          Tokens            `ic:"fee"`
		ExpiresAt    *Timestamp        `ic:"expires_at,omitempty"`
	} `ic:"Approve,variant"`
	TransferFrom *struct {
		From    AccountIdentifier `ic:"from"`
		To      AccountIdentifier `ic:"to"`
		Spender AccountIdentifier `ic:"spender"`
		Amount  Tokens            `ic:"amount"`
		Fee     Tokens            `ic:"fee"`
	} `ic:"TransferFrom,variant"`
}

type Timestamp = struct {
	TimestampNanos uint64 `ic:"timestamp_nanos"`
}

type Tokens = struct {
	E8s uint64 `ic:"e8s"`
}

type Transaction = struct {
	Memo          Memo       `ic:"memo"`
	Icrc1Memo     *[]byte    `ic:"icrc1_memo,omitempty"`
	Operation     *Operation `ic:"operation,omitempty"`
	CreatedAtTime Timestamp  `ic:"created_at_time"`
}
