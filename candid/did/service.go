package did

import (
	"fmt"
	"strings"

	"github.com/di-wu/parser/ast"
	"github.com/mohaijiang/agent-go/candid/internal/candid"
)

// Method is a public method of a service.
type Method struct {
	// Name describes the method.
	Name string
	// Func is a function type describing its signature.
	Func *Func
	// ID is a reference to a type definition naming a function reference type.
	// It is NOT possible to have both a function type and a reference.
	ID *string
	// Description
	Description string
}

func (m Method) String() string {
	s := fmt.Sprintf("%s : ", m.Name)
	if id := m.ID; id != nil {
		return s + *id
	}
	return s + m.Func.String()
}

// Service can be used to declare the complete interface of a service. A service is a standalone actor on the platform
// that can communicate with other services via sending and receiving messages. Messages are sent to a service by
// invoking one of its methods, i.e., functions that the service provides.
//
// Example:
//
//	service : {
//		addUser : (name : text, age : nat8) -> (id : nat64);
//		userName : (id : nat64) -> (text) query;
//		userAge : (id : nat64) -> (nat8) query;
//		deleteUser : (id : nat64) -> () oneway;
//	}
type Service struct {
	// ID represents the optional name given to the service. This only serves as documentation.
	ID *string

	// Methods is the list of methods that the service provides.
	Methods []Method
	// MethodId is the reference to the name of a type definition for an actor reference type.
	// It is NOT possible to have both a list of methods and a reference.
	MethodId *string
}

func convertService(n *ast.Node) Service {
	var actor Service
	for _, n := range n.Children() {
		switch n.Type {
		case candid.IdT:
			id := n.Value
			if actor.ID == nil {
				actor.ID = &id
				continue
			}
			actor.MethodId = &id
		case candid.TupTypeT:
		case candid.ActorTypeT:
			nextFunctionDescription := ""
			for _, n := range n.Children() {
				if n.Type == candid.CommentTextT {
					nextFunctionDescription += strings.TrimSpace(n.Value)
					continue
				}

				name := removeOuterQuotes(n.FirstChild.Value)
				switch n := n.LastChild; n.Type {
				case candid.FuncTypeT:
					f := convertFunc(n)
					actor.Methods = append(
						actor.Methods,
						Method{
							Name:        name,
							Func:        &f,
							Description: nextFunctionDescription,
						},
					)
				case candid.IdT, candid.TextT:
					id := n.Value
					actor.Methods = append(
						actor.Methods,
						Method{
							Name: name,
							ID:   &id,
						},
					)
				default:
					panic(n)
				}
				nextFunctionDescription = ""
			}
		case candid.MethTypeT:

			fmt.Println(n)
			var id string
			var f Func
			for _, c := range n.Children() {
				fmt.Println(c)
				switch c.Type {
				case candid.IdT:
					id = c.Value
				case candid.FuncTypeT:

					fmt.Println(c)
					f = convertFunc(c)
				}

			}
			actor.Methods = append(
				actor.Methods,
				Method{
					ID:   &id,
					Name: id,
					Func: &f,
				},
			)

		default:
			panic(n)
		}
	}
	return actor
}

func (a Service) String() string {
	s := "service "
	if id := a.ID; id != nil {
		s += fmt.Sprintf("%s ", *id)
	}
	s += ": "
	if id := a.MethodId; id != nil {
		return s + *id
	}
	s += "{\n"
	for _, m := range a.Methods {
		s += fmt.Sprintf("  %s;\n", m.String())
	}
	return s + "}"
}

func removeOuterQuotes(input string) string {
	// 去除首尾的空白字符
	input = strings.TrimSpace(input)

	// 检查字符串长度是否足够包含至少两个引号
	if len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"' {
		// 去除外围的引号
		return input[1 : len(input)-1]
	}

	return input
}
