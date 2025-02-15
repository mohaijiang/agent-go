// Automatically generated by https://github.com/aviate-labs/agent-go.
package wallet_test

import (
	"github.com/mohaijiang/agent-go"
	"github.com/mohaijiang/agent-go/candid/idl"
	"github.com/mohaijiang/agent-go/mock"
	"github.com/mohaijiang/agent-go/principal"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mohaijiang/agent-go/ic/wallet"
)

// Test_ApiVersion tests the "api_version" method on the "wallet" canister.
func Test_ApiVersion(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "api_version",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(uint16)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.ApiVersion(); err != nil {
		t.Fatal(err)
	}

}

// Test_Authorize tests the "authorize" method on the "wallet" canister.
func Test_Authorize(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "authorize",
			Arguments: []any{new(principal.Principal)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 principal.Principal
	if err := a.Authorize(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_CertifiedTree tests the "certified_tree" method on the "wallet" canister.
func Test_CertifiedTree(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name: "certified_tree",
			Arguments: []any{new(struct {
			})},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(struct {
					Certificate []byte `ic:"certificate"`
					Tree        []byte `ic:"tree"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 struct {
	}
	if _, err := a.CertifiedTree(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_Clear tests the "clear" method on the "wallet" canister.
func Test_Clear(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "clear",
			Arguments: []any{new(wallet.ClearArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.ClearArguments
	if err := a.Clear(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_CommitBatch tests the "commit_batch" method on the "wallet" canister.
func Test_CommitBatch(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "commit_batch",
			Arguments: []any{new(wallet.CommitBatchArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.CommitBatchArguments
	if err := a.CommitBatch(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_CommitProposedBatch tests the "commit_proposed_batch" method on the "wallet" canister.
func Test_CommitProposedBatch(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "commit_proposed_batch",
			Arguments: []any{new(wallet.CommitProposedBatchArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.CommitProposedBatchArguments
	if err := a.CommitProposedBatch(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_ComputeEvidence tests the "compute_evidence" method on the "wallet" canister.
func Test_ComputeEvidence(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "compute_evidence",
			Arguments: []any{new(wallet.ComputeEvidenceArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(*[]byte)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.ComputeEvidenceArguments
	if _, err := a.ComputeEvidence(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_CreateAsset tests the "create_asset" method on the "wallet" canister.
func Test_CreateAsset(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "create_asset",
			Arguments: []any{new(wallet.CreateAssetArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.CreateAssetArguments
	if err := a.CreateAsset(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_CreateBatch tests the "create_batch" method on the "wallet" canister.
func Test_CreateBatch(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name: "create_batch",
			Arguments: []any{new(struct {
			})},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(struct {
					BatchId wallet.BatchId `ic:"batch_id"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 struct {
	}
	if _, err := a.CreateBatch(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_CreateChunk tests the "create_chunk" method on the "wallet" canister.
func Test_CreateChunk(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name: "create_chunk",
			Arguments: []any{new(struct {
				BatchId wallet.BatchId `ic:"batch_id"`
				Content []byte         `ic:"content"`
			})},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(struct {
					ChunkId wallet.ChunkId `ic:"chunk_id"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 struct {
		BatchId wallet.BatchId `ic:"batch_id"`
		Content []byte         `ic:"content"`
	}
	if _, err := a.CreateChunk(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_Deauthorize tests the "deauthorize" method on the "wallet" canister.
func Test_Deauthorize(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "deauthorize",
			Arguments: []any{new(principal.Principal)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 principal.Principal
	if err := a.Deauthorize(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_DeleteAsset tests the "delete_asset" method on the "wallet" canister.
func Test_DeleteAsset(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "delete_asset",
			Arguments: []any{new(wallet.DeleteAssetArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.DeleteAssetArguments
	if err := a.DeleteAsset(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_DeleteBatch tests the "delete_batch" method on the "wallet" canister.
func Test_DeleteBatch(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "delete_batch",
			Arguments: []any{new(wallet.DeleteBatchArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.DeleteBatchArguments
	if err := a.DeleteBatch(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_Get tests the "get" method on the "wallet" canister.
func Test_Get(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name: "get",
			Arguments: []any{new(struct {
				Key             wallet.Key `ic:"key"`
				AcceptEncodings []string   `ic:"accept_encodings"`
			})},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(struct {
					Content         []byte  `ic:"content"`
					ContentType     string  `ic:"content_type"`
					ContentEncoding string  `ic:"content_encoding"`
					Sha256          *[]byte `ic:"sha256,omitempty"`
					TotalLength     idl.Nat `ic:"total_length"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 struct {
		Key             wallet.Key `ic:"key"`
		AcceptEncodings []string   `ic:"accept_encodings"`
	}
	if _, err := a.Get(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_GetAssetProperties tests the "get_asset_properties" method on the "wallet" canister.
func Test_GetAssetProperties(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "get_asset_properties",
			Arguments: []any{new(wallet.Key)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(struct {
					MaxAge         *uint64               `ic:"max_age,omitempty"`
					Headers        *[]wallet.HeaderField `ic:"headers,omitempty"`
					AllowRawAccess *bool                 `ic:"allow_raw_access,omitempty"`
					IsAliased      *bool                 `ic:"is_aliased,omitempty"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.Key
	if _, err := a.GetAssetProperties(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_GetChunk tests the "get_chunk" method on the "wallet" canister.
func Test_GetChunk(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name: "get_chunk",
			Arguments: []any{new(struct {
				Key             wallet.Key `ic:"key"`
				ContentEncoding string     `ic:"content_encoding"`
				Index           idl.Nat    `ic:"index"`
				Sha256          *[]byte    `ic:"sha256,omitempty"`
			})},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(struct {
					Content []byte `ic:"content"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 struct {
		Key             wallet.Key `ic:"key"`
		ContentEncoding string     `ic:"content_encoding"`
		Index           idl.Nat    `ic:"index"`
		Sha256          *[]byte    `ic:"sha256,omitempty"`
	}
	if _, err := a.GetChunk(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_GrantPermission tests the "grant_permission" method on the "wallet" canister.
func Test_GrantPermission(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "grant_permission",
			Arguments: []any{new(wallet.GrantPermission)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.GrantPermission
	if err := a.GrantPermission(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_HttpRequest tests the "http_request" method on the "wallet" canister.
func Test_HttpRequest(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "http_request",
			Arguments: []any{new(wallet.HttpRequest)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(wallet.HttpResponse)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.HttpRequest
	if _, err := a.HttpRequest(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_HttpRequestStreamingCallback tests the "http_request_streaming_callback" method on the "wallet" canister.
func Test_HttpRequestStreamingCallback(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "http_request_streaming_callback",
			Arguments: []any{new(wallet.StreamingCallbackToken)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(*wallet.StreamingCallbackHttpResponse)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.StreamingCallbackToken
	if _, err := a.HttpRequestStreamingCallback(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_List tests the "list" method on the "wallet" canister.
func Test_List(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name: "list",
			Arguments: []any{new(struct {
			})},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new([]struct {
					Key         wallet.Key `ic:"key"`
					ContentType string     `ic:"content_type"`
					Encodings   []struct {
						ContentEncoding string      `ic:"content_encoding"`
						Sha256          *[]byte     `ic:"sha256,omitempty"`
						Length          idl.Nat     `ic:"length"`
						Modified        wallet.Time `ic:"modified"`
					} `ic:"encodings"`
				})}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 struct {
	}
	if _, err := a.List(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_ListAuthorized tests the "list_authorized" method on the "wallet" canister.
func Test_ListAuthorized(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "list_authorized",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new([]principal.Principal)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.ListAuthorized(); err != nil {
		t.Fatal(err)
	}

}

// Test_ListPermitted tests the "list_permitted" method on the "wallet" canister.
func Test_ListPermitted(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "list_permitted",
			Arguments: []any{new(wallet.ListPermitted)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new([]principal.Principal)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.ListPermitted
	if _, err := a.ListPermitted(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_ProposeCommitBatch tests the "propose_commit_batch" method on the "wallet" canister.
func Test_ProposeCommitBatch(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "propose_commit_batch",
			Arguments: []any{new(wallet.CommitBatchArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.CommitBatchArguments
	if err := a.ProposeCommitBatch(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_RevokePermission tests the "revoke_permission" method on the "wallet" canister.
func Test_RevokePermission(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "revoke_permission",
			Arguments: []any{new(wallet.RevokePermission)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.RevokePermission
	if err := a.RevokePermission(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_SetAssetContent tests the "set_asset_content" method on the "wallet" canister.
func Test_SetAssetContent(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "set_asset_content",
			Arguments: []any{new(wallet.SetAssetContentArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.SetAssetContentArguments
	if err := a.SetAssetContent(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_SetAssetProperties tests the "set_asset_properties" method on the "wallet" canister.
func Test_SetAssetProperties(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "set_asset_properties",
			Arguments: []any{new(wallet.SetAssetPropertiesArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.SetAssetPropertiesArguments
	if err := a.SetAssetProperties(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_Store tests the "store" method on the "wallet" canister.
func Test_Store(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name: "store",
			Arguments: []any{new(struct {
				Key             wallet.Key `ic:"key"`
				ContentType     string     `ic:"content_type"`
				ContentEncoding string     `ic:"content_encoding"`
				Content         []byte     `ic:"content"`
				Sha256          *[]byte    `ic:"sha256,omitempty"`
			})},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 struct {
		Key             wallet.Key `ic:"key"`
		ContentType     string     `ic:"content_type"`
		ContentEncoding string     `ic:"content_encoding"`
		Content         []byte     `ic:"content"`
		Sha256          *[]byte    `ic:"sha256,omitempty"`
	}
	if err := a.Store(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_TakeOwnership tests the "take_ownership" method on the "wallet" canister.
func Test_TakeOwnership(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "take_ownership",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := a.TakeOwnership(); err != nil {
		t.Fatal(err)
	}

}

// Test_UnsetAssetContent tests the "unset_asset_content" method on the "wallet" canister.
func Test_UnsetAssetContent(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "unset_asset_content",
			Arguments: []any{new(wallet.UnsetAssetContentArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.UnsetAssetContentArguments
	if err := a.UnsetAssetContent(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_ValidateCommitProposedBatch tests the "validate_commit_proposed_batch" method on the "wallet" canister.
func Test_ValidateCommitProposedBatch(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "validate_commit_proposed_batch",
			Arguments: []any{new(wallet.CommitProposedBatchArguments)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(wallet.ValidationResult)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.CommitProposedBatchArguments
	if _, err := a.ValidateCommitProposedBatch(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_ValidateGrantPermission tests the "validate_grant_permission" method on the "wallet" canister.
func Test_ValidateGrantPermission(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "validate_grant_permission",
			Arguments: []any{new(wallet.GrantPermission)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(wallet.ValidationResult)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.GrantPermission
	if _, err := a.ValidateGrantPermission(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_ValidateRevokePermission tests the "validate_revoke_permission" method on the "wallet" canister.
func Test_ValidateRevokePermission(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "validate_revoke_permission",
			Arguments: []any{new(wallet.RevokePermission)},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(wallet.ValidationResult)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	var a0 wallet.RevokePermission
	if _, err := a.ValidateRevokePermission(a0); err != nil {
		t.Fatal(err)
	}

}

// Test_ValidateTakeOwnership tests the "validate_take_ownership" method on the "wallet" canister.
func Test_ValidateTakeOwnership(t *testing.T) {
	a, err := newAgent([]mock.Method{
		{
			Name:      "validate_take_ownership",
			Arguments: []any{},
			Handler: func(request mock.Request) ([]any, error) {
				return []any{*new(wallet.ValidationResult)}, nil
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.ValidateTakeOwnership(); err != nil {
		t.Fatal(err)
	}

}

// newAgent creates a new agent with the given (mock) methods.
// Runs a mock replica in the background.
func newAgent(methods []mock.Method) (*wallet.Agent, error) {
	replica := mock.NewReplica()
	canisterId := principal.Principal{Raw: []byte("wallet")}
	replica.AddCanister(canisterId, methods)
	s := httptest.NewServer(replica)
	u, _ := url.Parse(s.URL)
	a, err := wallet.NewAgent(canisterId, agent.Config{
		ClientConfig: &agent.ClientConfig{Host: u},
		FetchRootKey: true,
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}
