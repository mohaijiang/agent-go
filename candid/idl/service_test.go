package idl_test

import (
	"github.com/mohaijiang/agent-go/candid/idl"
	"github.com/mohaijiang/agent-go/principal"
)

func ExampleService() {
	p, _ := principal.Decode("w7x7r-cok77-xa")
	test(
		[]idl.Type{idl.NewServiceType(
			map[string]*idl.FunctionType{
				"foo": idl.NewFunctionType(
					[]idl.Type{new(idl.TextType)},
					[]idl.Type{new(idl.NatType)},
					nil,
				),
			},
		)},
		[]any{
			p,
		},
	)
	// Output:
	// 4449444c026a0171017d00690103666f6f0001010103caffee
}
