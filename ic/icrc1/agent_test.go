// Automatically generated by https://github.com/aviate-labs/agent-go.
package icrc1_test

import (
	"github.com/mohaijiang/agent-go"
	"github.com/mohaijiang/agent-go/candid/idl"
	"github.com/mohaijiang/agent-go/mock"
	"github.com/mohaijiang/agent-go/principal"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mohaijiang/agent-go/ic/icrc1"
)

// Test_Icrc1BalanceOf tests the "icrc1_balance_of" method on the "icrc1" canister.
func Test_Icrc1BalanceOf(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_balance_of",
			Arguments: []any{new(icrc1.Account)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(idl.Nat)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 icrc1.Account
	if _, err := a.Icrc1BalanceOf(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1Decimals tests the "icrc1_decimals" method on the "icrc1" canister.
func Test_Icrc1Decimals(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_decimals",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(uint8)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1Decimals(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1Fee tests the "icrc1_fee" method on the "icrc1" canister.
func Test_Icrc1Fee(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_fee",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(idl.Nat)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1Fee(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1Metadata tests the "icrc1_metadata" method on the "icrc1" canister.
func Test_Icrc1Metadata(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_metadata",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new([]struct {
					field0 string      `ic:"field0"`
					field1 icrc1.Value `ic:"field1"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1Metadata(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1MintingAccount tests the "icrc1_minting_account" method on the "icrc1" canister.
func Test_Icrc1MintingAccount(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_minting_account",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(*icrc1.Account)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1MintingAccount(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1Name tests the "icrc1_name" method on the "icrc1" canister.
func Test_Icrc1Name(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_name",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(string)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1Name(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1SupportedStandards tests the "icrc1_supported_standards" method on the "icrc1" canister.
func Test_Icrc1SupportedStandards(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_supported_standards",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new([]struct {
					Name string `ic:"name"`
					Url  string `ic:"url"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1SupportedStandards(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1Symbol tests the "icrc1_symbol" method on the "icrc1" canister.
func Test_Icrc1Symbol(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_symbol",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(string)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1Symbol(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1TotalSupply tests the "icrc1_total_supply" method on the "icrc1" canister.
func Test_Icrc1TotalSupply(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_total_supply",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(idl.Nat)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Icrc1TotalSupply(); err != nil {
		t.Fatal(err)
	}

}

// Test_Icrc1Transfer tests the "icrc1_transfer" method on the "icrc1" canister.
func Test_Icrc1Transfer(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "icrc1_transfer",
			Arguments: []any{new(icrc1.TransferArgs)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(struct {
					Ok  *idl.Nat             `ic:"Ok,variant"`
					Err *icrc1.TransferError `ic:"Err,variant"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 icrc1.TransferArgs
	if _, err := a.Icrc1Transfer(a0); err != nil {
		t.Fatal(err)
	}

}

// newAgent creates a new agent with the given (mock) methods.
// Runs a mock replica in the background.
func newAgent(methods []mock.Method) (*icrc1.Agent, error) {
	replica := mock.NewReplica()
	canisterId := principal.Principal{Raw: []byte("icrc1")}
	replica.AddCanister(canisterId, methods)
	s := httptest.NewServer(replica)
	u, _ := url.Parse(s.URL)
	a, err := icrc1.NewAgent(canisterId, agent.Config{
		ClientConfig: &agent.ClientConfig{Host: u},
		FetchRootKey: true,
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}
