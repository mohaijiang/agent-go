// Package assetstorage provides a client for the "assetstorage" canister.
// Do NOT edit this file. It was automatically generated by https://github.com/aviate-labs/agent-go.
package assetstorage

import (
	"github.com/mohaijiang/agent-go"
	"github.com/mohaijiang/agent-go/candid/idl"
	"github.com/mohaijiang/agent-go/principal"
)

// Agent is a client for the "assetstorage" canister.
type Agent struct {
	a          *agent.Agent
	canisterId principal.Principal
}

// NewAgent creates a new agent for the "assetstorage" canister.
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

// ApiVersion calls the "api_version" method on the "assetstorage" canister.
func (a Agent) ApiVersion() (*uint16, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 uint16
	if err := a.a.Query(
		a.canisterId,
		"api_version",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Authorize calls the "authorize" method on the "assetstorage" canister.
func (a Agent) Authorize(arg0 principal.Principal) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"authorize",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// CertifiedTree calls the "certified_tree" method on the "assetstorage" canister.
func (a Agent) CertifiedTree(arg0 struct {
}) (*struct {
	Certificate []byte `ic:"certificate"`
	Tree        []byte `ic:"tree"`
}, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		Certificate []byte `ic:"certificate"`
		Tree        []byte `ic:"tree"`
	}
	if err := a.a.Query(
		a.canisterId,
		"certified_tree",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Clear calls the "clear" method on the "assetstorage" canister.
func (a Agent) Clear(arg0 ClearArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"clear",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// CommitBatch calls the "commit_batch" method on the "assetstorage" canister.
func (a Agent) CommitBatch(arg0 CommitBatchArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"commit_batch",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// CommitProposedBatch calls the "commit_proposed_batch" method on the "assetstorage" canister.
func (a Agent) CommitProposedBatch(arg0 CommitProposedBatchArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"commit_proposed_batch",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// ComputeEvidence calls the "compute_evidence" method on the "assetstorage" canister.
func (a Agent) ComputeEvidence(arg0 ComputeEvidenceArguments) (**[]byte, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 *[]byte
	if err := a.a.Call(
		a.canisterId,
		"compute_evidence",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// CreateAsset calls the "create_asset" method on the "assetstorage" canister.
func (a Agent) CreateAsset(arg0 CreateAssetArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"create_asset",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// CreateBatch calls the "create_batch" method on the "assetstorage" canister.
func (a Agent) CreateBatch(arg0 struct {
}) (*struct {
	BatchId BatchId `ic:"batch_id"`
}, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		BatchId BatchId `ic:"batch_id"`
	}
	if err := a.a.Call(
		a.canisterId,
		"create_batch",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// CreateChunk calls the "create_chunk" method on the "assetstorage" canister.
func (a Agent) CreateChunk(arg0 struct {
	BatchId BatchId `ic:"batch_id"`
	Content []byte  `ic:"content"`
}) (*struct {
	ChunkId ChunkId `ic:"chunk_id"`
}, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		ChunkId ChunkId `ic:"chunk_id"`
	}
	if err := a.a.Call(
		a.canisterId,
		"create_chunk",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// Deauthorize calls the "deauthorize" method on the "assetstorage" canister.
func (a Agent) Deauthorize(arg0 principal.Principal) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"deauthorize",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// DeleteAsset calls the "delete_asset" method on the "assetstorage" canister.
func (a Agent) DeleteAsset(arg0 DeleteAssetArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"delete_asset",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// DeleteBatch calls the "delete_batch" method on the "assetstorage" canister.
func (a Agent) DeleteBatch(arg0 DeleteBatchArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"delete_batch",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// Get calls the "get" method on the "assetstorage" canister.
func (a Agent) Get(arg0 struct {
	Key             Key      `ic:"key"`
	AcceptEncodings []string `ic:"accept_encodings"`
}) (*struct {
	Content         []byte  `ic:"content"`
	ContentType     string  `ic:"content_type"`
	ContentEncoding string  `ic:"content_encoding"`
	Sha256          *[]byte `ic:"sha256,omitempty"`
	TotalLength     idl.Nat `ic:"total_length"`
}, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		Content         []byte  `ic:"content"`
		ContentType     string  `ic:"content_type"`
		ContentEncoding string  `ic:"content_encoding"`
		Sha256          *[]byte `ic:"sha256,omitempty"`
		TotalLength     idl.Nat `ic:"total_length"`
	}
	if err := a.a.Query(
		a.canisterId,
		"get",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetAssetProperties calls the "get_asset_properties" method on the "assetstorage" canister.
func (a Agent) GetAssetProperties(key Key) (*struct {
	MaxAge         *uint64        `ic:"max_age,omitempty"`
	Headers        *[]HeaderField `ic:"headers,omitempty"`
	AllowRawAccess *bool          `ic:"allow_raw_access,omitempty"`
	IsAliased      *bool          `ic:"is_aliased,omitempty"`
}, error) {
	args, err := idl.Marshal([]any{key})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		MaxAge         *uint64        `ic:"max_age,omitempty"`
		Headers        *[]HeaderField `ic:"headers,omitempty"`
		AllowRawAccess *bool          `ic:"allow_raw_access,omitempty"`
		IsAliased      *bool          `ic:"is_aliased,omitempty"`
	}
	if err := a.a.Query(
		a.canisterId,
		"get_asset_properties",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GetChunk calls the "get_chunk" method on the "assetstorage" canister.
func (a Agent) GetChunk(arg0 struct {
	Key             Key     `ic:"key"`
	ContentEncoding string  `ic:"content_encoding"`
	Index           idl.Nat `ic:"index"`
	Sha256          *[]byte `ic:"sha256,omitempty"`
}) (*struct {
	Content []byte `ic:"content"`
}, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 struct {
		Content []byte `ic:"content"`
	}
	if err := a.a.Query(
		a.canisterId,
		"get_chunk",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// GrantPermission calls the "grant_permission" method on the "assetstorage" canister.
func (a Agent) GrantPermission(arg0 GrantPermission) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"grant_permission",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// HttpRequest calls the "http_request" method on the "assetstorage" canister.
func (a Agent) HttpRequest(request HttpRequest) (*HttpResponse, error) {
	args, err := idl.Marshal([]any{request})
	if err != nil {
		return nil, err
	}
	var r0 HttpResponse
	if err := a.a.Query(
		a.canisterId,
		"http_request",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// HttpRequestStreamingCallback calls the "http_request_streaming_callback" method on the "assetstorage" canister.
func (a Agent) HttpRequestStreamingCallback(token StreamingCallbackToken) (**StreamingCallbackHttpResponse, error) {
	args, err := idl.Marshal([]any{token})
	if err != nil {
		return nil, err
	}
	var r0 *StreamingCallbackHttpResponse
	if err := a.a.Query(
		a.canisterId,
		"http_request_streaming_callback",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// List calls the "list" method on the "assetstorage" canister.
func (a Agent) List(arg0 struct {
}) (*[]struct {
	Key         Key    `ic:"key"`
	ContentType string `ic:"content_type"`
	Encodings   []struct {
		ContentEncoding string  `ic:"content_encoding"`
		Sha256          *[]byte `ic:"sha256,omitempty"`
		Length          idl.Nat `ic:"length"`
		Modified        Time    `ic:"modified"`
	} `ic:"encodings"`
}, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 []struct {
		Key         Key    `ic:"key"`
		ContentType string `ic:"content_type"`
		Encodings   []struct {
			ContentEncoding string  `ic:"content_encoding"`
			Sha256          *[]byte `ic:"sha256,omitempty"`
			Length          idl.Nat `ic:"length"`
			Modified        Time    `ic:"modified"`
		} `ic:"encodings"`
	}
	if err := a.a.Query(
		a.canisterId,
		"list",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListAuthorized calls the "list_authorized" method on the "assetstorage" canister.
func (a Agent) ListAuthorized() (*[]principal.Principal, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 []principal.Principal
	if err := a.a.Query(
		a.canisterId,
		"list_authorized",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ListPermitted calls the "list_permitted" method on the "assetstorage" canister.
func (a Agent) ListPermitted(arg0 ListPermitted) (*[]principal.Principal, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 []principal.Principal
	if err := a.a.Query(
		a.canisterId,
		"list_permitted",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ProposeCommitBatch calls the "propose_commit_batch" method on the "assetstorage" canister.
func (a Agent) ProposeCommitBatch(arg0 CommitBatchArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"propose_commit_batch",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// RevokePermission calls the "revoke_permission" method on the "assetstorage" canister.
func (a Agent) RevokePermission(arg0 RevokePermission) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"revoke_permission",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// SetAssetContent calls the "set_asset_content" method on the "assetstorage" canister.
func (a Agent) SetAssetContent(arg0 SetAssetContentArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"set_asset_content",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// SetAssetProperties calls the "set_asset_properties" method on the "assetstorage" canister.
func (a Agent) SetAssetProperties(arg0 SetAssetPropertiesArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"set_asset_properties",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// Store calls the "store" method on the "assetstorage" canister.
func (a Agent) Store(arg0 struct {
	Key             Key     `ic:"key"`
	ContentType     string  `ic:"content_type"`
	ContentEncoding string  `ic:"content_encoding"`
	Content         []byte  `ic:"content"`
	Sha256          *[]byte `ic:"sha256,omitempty"`
}) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"store",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// TakeOwnership calls the "take_ownership" method on the "assetstorage" canister.
func (a Agent) TakeOwnership() error {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"take_ownership",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// UnsetAssetContent calls the "unset_asset_content" method on the "assetstorage" canister.
func (a Agent) UnsetAssetContent(arg0 UnsetAssetContentArguments) error {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return err
	}
	if err := a.a.Call(
		a.canisterId,
		"unset_asset_content",
		args,
		[]any{},
	); err != nil {
		return err
	}
	return nil
}

// ValidateCommitProposedBatch calls the "validate_commit_proposed_batch" method on the "assetstorage" canister.
func (a Agent) ValidateCommitProposedBatch(arg0 CommitProposedBatchArguments) (*ValidationResult, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 ValidationResult
	if err := a.a.Call(
		a.canisterId,
		"validate_commit_proposed_batch",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ValidateGrantPermission calls the "validate_grant_permission" method on the "assetstorage" canister.
func (a Agent) ValidateGrantPermission(arg0 GrantPermission) (*ValidationResult, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 ValidationResult
	if err := a.a.Call(
		a.canisterId,
		"validate_grant_permission",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ValidateRevokePermission calls the "validate_revoke_permission" method on the "assetstorage" canister.
func (a Agent) ValidateRevokePermission(arg0 RevokePermission) (*ValidationResult, error) {
	args, err := idl.Marshal([]any{arg0})
	if err != nil {
		return nil, err
	}
	var r0 ValidationResult
	if err := a.a.Call(
		a.canisterId,
		"validate_revoke_permission",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

// ValidateTakeOwnership calls the "validate_take_ownership" method on the "assetstorage" canister.
func (a Agent) ValidateTakeOwnership() (*ValidationResult, error) {
	args, err := idl.Marshal([]any{})
	if err != nil {
		return nil, err
	}
	var r0 ValidationResult
	if err := a.a.Call(
		a.canisterId,
		"validate_take_ownership",
		args,
		[]any{&r0},
	); err != nil {
		return nil, err
	}
	return &r0, nil
}

type BatchId = idl.Nat

type BatchOperationKind = struct {
	CreateAsset        *CreateAssetArguments        `ic:"CreateAsset,variant"`
	SetAssetContent    *SetAssetContentArguments    `ic:"SetAssetContent,variant"`
	SetAssetProperties *SetAssetPropertiesArguments `ic:"SetAssetProperties,variant"`
	UnsetAssetContent  *UnsetAssetContentArguments  `ic:"UnsetAssetContent,variant"`
	DeleteAsset        *DeleteAssetArguments        `ic:"DeleteAsset,variant"`
	Clear              *ClearArguments              `ic:"Clear,variant"`
}

type ChunkId = idl.Nat

type ClearArguments = struct {
}

type CommitBatchArguments = struct {
	BatchId    BatchId              `ic:"batch_id"`
	Operations []BatchOperationKind `ic:"operations"`
}

type CommitProposedBatchArguments = struct {
	BatchId  BatchId `ic:"batch_id"`
	Evidence []byte  `ic:"evidence"`
}

type ComputeEvidenceArguments = struct {
	BatchId       BatchId `ic:"batch_id"`
	MaxIterations *uint16 `ic:"max_iterations,omitempty"`
}

type CreateAssetArguments = struct {
	Key            Key            `ic:"key"`
	ContentType    string         `ic:"content_type"`
	MaxAge         *uint64        `ic:"max_age,omitempty"`
	Headers        *[]HeaderField `ic:"headers,omitempty"`
	EnableAliasing *bool          `ic:"enable_aliasing,omitempty"`
	AllowRawAccess *bool          `ic:"allow_raw_access,omitempty"`
}

type DeleteAssetArguments = struct {
	Key Key `ic:"key"`
}

type DeleteBatchArguments = struct {
	BatchId BatchId `ic:"batch_id"`
}

type GrantPermission = struct {
	ToPrincipal principal.Principal `ic:"to_principal"`
	Permission  Permission          `ic:"permission"`
}

type HeaderField = struct {
	field0 string `ic:"field0"`
	field1 string `ic:"field1"`
}

type HttpRequest = struct {
	Method             string        `ic:"method"`
	Url                string        `ic:"url"`
	Headers            []HeaderField `ic:"headers"`
	Body               []byte        `ic:"body"`
	CertificateVersion *uint16       `ic:"certificate_version,omitempty"`
}

type HttpResponse = struct {
	StatusCode        uint16             `ic:"status_code"`
	Headers           []HeaderField      `ic:"headers"`
	Body              []byte             `ic:"body"`
	StreamingStrategy *StreamingStrategy `ic:"streaming_strategy,omitempty"`
}

type Key = string

type ListPermitted = struct {
	Permission Permission `ic:"permission"`
}

type Permission = struct {
	Commit            *struct{} `ic:"Commit,variant"`
	ManagePermissions *struct{} `ic:"ManagePermissions,variant"`
	Prepare           *struct{} `ic:"Prepare,variant"`
}

type RevokePermission = struct {
	OfPrincipal principal.Principal `ic:"of_principal"`
	Permission  Permission          `ic:"permission"`
}

type SetAssetContentArguments = struct {
	Key             Key       `ic:"key"`
	ContentEncoding string    `ic:"content_encoding"`
	ChunkIds        []ChunkId `ic:"chunk_ids"`
	Sha256          *[]byte   `ic:"sha256,omitempty"`
}

type SetAssetPropertiesArguments = struct {
	Key            Key             `ic:"key"`
	MaxAge         **uint64        `ic:"max_age,omitempty"`
	Headers        **[]HeaderField `ic:"headers,omitempty"`
	AllowRawAccess **bool          `ic:"allow_raw_access,omitempty"`
	IsAliased      **bool          `ic:"is_aliased,omitempty"`
}

type StreamingCallbackHttpResponse = struct {
	Body  []byte                  `ic:"body"`
	Token *StreamingCallbackToken `ic:"token,omitempty"`
}

type StreamingCallbackToken = struct {
	Key             Key     `ic:"key"`
	ContentEncoding string  `ic:"content_encoding"`
	Index           idl.Nat `ic:"index"`
	Sha256          *[]byte `ic:"sha256,omitempty"`
}

type StreamingStrategy = struct {
	Callback *struct {
		Callback struct { /* NOT SUPPORTED */
		} `ic:"callback"`
		Token StreamingCallbackToken `ic:"token"`
	} `ic:"Callback,variant"`
}

type Time = idl.Int

type UnsetAssetContentArguments = struct {
	Key             Key    `ic:"key"`
	ContentEncoding string `ic:"content_encoding"`
}

type ValidationResult = struct {
	Ok  *string `ic:"Ok,variant"`
	Err *string `ic:"Err,variant"`
}
