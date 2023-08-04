package agent_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid"
	"github.com/aviate-labs/agent-go/identity"
	"github.com/aviate-labs/agent-go/principal"
	"os"
	"testing"
)

func Example_anonymous_query() {
	ledgerID, _ := principal.Decode("ryjl3-tyaaa-aaaaa-aaaba-cai")
	a, _ := agent.New(agent.Config{})
	args, err := candid.EncodeValueString("record { account = \"9523dc824aa062dcd9c91b98f4594ff9c6af661ac96747daef2090b7fe87037d\" }")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.QueryString(ledgerID, "account_balance_dfx", args))
	// Output:
	// (record { 5035232 = 0 : nat64 }) <nil>
}

func Example_query() {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	id, _ := identity.NewEd25519Identity(publicKey, privateKey)
	ledgerID, _ := principal.Decode("ryjl3-tyaaa-aaaaa-aaaba-cai")
	a, _ := agent.New(agent.Config{
		Identity: id,
	})
	args, err := candid.EncodeValueString("record { account = \"9523dc824aa062dcd9c91b98f4594ff9c6af661ac96747daef2090b7fe87037d\" }")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.QueryString(ledgerID, "account_balance_dfx", args))
	// Output:
	// (record { 5035232 = 0 : nat64 }) <nil>
}

func TestLedgerBalance(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	id, _ := identity.NewEd25519Identity(publicKey, privateKey)
	ledgerID, _ := principal.Decode("ryjl3-tyaaa-aaaaa-aaaba-cai")
	a, _ := agent.New(agent.Config{
		Identity: id,
	})
	args, err := candid.EncodeValueString("record { account = \"b7d5655c950eb49315719f89d2f0cbbe392ab4db2ae2db0d15051f4d2f858486\" }")
	if err != nil {
		fmt.Println(err)
	}
	str, err := a.QueryString(ledgerID, "account_balance_dfx", args)

	fmt.Println(str, err)
}

func TestWalletBalance(t *testing.T) {

	data, _ := os.ReadFile("/Users/mohaijiang/.config/dfx/identity/126639987/identity.pem")
	id, err := identity.NewSecp256k1IdentityFromPEM(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(id.Sender().Encode())

	ledgerID, _ := principal.Decode("wjws2-vaaaa-aaaag-abtfa-cai")
	a, _ := agent.New(agent.Config{
		Identity: id,
	})

	str, err := a.QueryString(ledgerID, "wallet_balance128", []byte{})

	fmt.Println(str, err)

}
