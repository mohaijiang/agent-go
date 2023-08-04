package agent_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/aviate-labs/agent-go"
)

// var ic0, _ = url.Parse("https://ic0.app/")
var ic0, _ = url.Parse("http://127.0.0.1:4939/")

func ExampleClient_Status() {
	c := agent.NewClient(agent.ClientConfig{Host: ic0})
	status, _ := c.Status()
	fmt.Println(status.Version)
	// Output:
	// 0.18.0
}

func TestClient(t *testing.T) {
	//c := agent.NewClient(agent.ClientConfig{Host: ic0})
}
