package candidtest_test

import (
	"os"
	"testing"

	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/mohaijiang/agent-go/candid/internal/candidtest"
)

func TestData(t *testing.T) {
	rawDid, _ := os.ReadFile("../../idl/testdata/prim.test.did")
	p, _ := ast.New(rawDid)
	if _, err := candidtest.TestData(p); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		t.Error(err)
	}
}
