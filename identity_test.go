package agent

import (
	"fmt"
	"github.com/mohaijiang/agent-go/identity"
	"os"
	"testing"
)

func TestIdentity(t *testing.T) {

	id, _ := identity.NewRandomSecp256k1Identity()
	data, _ := id.ToPEM()
	fmt.Println(string(data))
}

func TestLoadIdentity(t *testing.T) {
	data, _ := os.ReadFile("/Users/mohaijiang/.config/dfx/identity/126639987/identity.pem")
	id, err := identity.NewSecp256k1IdentityFromPEM(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}
