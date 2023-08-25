package agent

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/fxamacker/cbor/v2"
	"github.com/mohaijiang/agent-go/principal"
)

// Client is a client for the IC agent.
type Client struct {
	client http.Client
	config ClientConfig
}

// NewClient creates a new client based on the given configuration.
func NewClient(cfg ClientConfig) Client {
	return Client{
		client: http.Client{},
		config: cfg,
	}
}

// Status returns the status of the IC.
func (c Client) Status() (*Status, error) {
	raw, err := c.get("/api/v2/status")
	if err != nil {
		return nil, err
	}
	var status Status
	return &status, cbor.Unmarshal(raw, &status)
}

func (c Client) call(canisterID principal.Principal, data []byte) ([]byte, error) {
	u := c.url(fmt.Sprintf("/api/v2/canister/%s/call", canisterID.Encode()))
	resp, err := c.client.Post(u, "application/cbor", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusAccepted:
		return io.ReadAll(resp.Body)
	case http.StatusOK:
		body, _ := io.ReadAll(resp.Body)
		var err preprocessingError
		if err := cbor.Unmarshal(body, &err); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("(%d) %s: %s", err.RejectCode, err.Message, err.ErrorCode)
	default:
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("(%d) %s: %s", resp.StatusCode, resp.Status, body)
	}
}

func (c Client) get(path string) ([]byte, error) {
	resp, err := c.client.Get(c.url(path))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}

func (c Client) post(path string, canisterID principal.Principal, data []byte) ([]byte, error) {
	u := c.url(fmt.Sprintf("/api/v2/canister/%s/%s", canisterID.Encode(), path))
	resp, err := c.client.Post(u, "application/cbor", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		return io.ReadAll(resp.Body)
	default:
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("(%d) %s: %s", resp.StatusCode, resp.Status, body)
	}
}

func (c Client) query(canisterID principal.Principal, data []byte) ([]byte, error) {
	return c.post("query", canisterID, data)
}

func (c Client) readState(canisterID principal.Principal, data []byte) ([]byte, error) {
	return c.post("read_state", canisterID, data)
}

func (c Client) url(p string) string {
	u := *c.config.Host
	u.Path = path.Join(u.Path, p)
	return u.String()
}

// ClientConfig is the configuration for a client.
type ClientConfig struct {
	Host *url.URL
}

type preprocessingError struct {
	// The reject code.
	RejectCode uint64 `cbor:"reject_code"`
	// A textual diagnostic message.
	Message string `cbor:"reject_message"`
	// An optional implementation-specific textual error code.
	ErrorCode string `cbor:"error_code"`
}
