package adapter

import (
	. "../language/cplus"
	"fmt"
)

func NewCPPCallListener() *CPPCallListener {
	return &CPPCallListener{}
}

type CPPCallListener struct {
	BaseCPP14Listener
}

func (s *CPPCallListener) EnterNamespacename(ctx *NamespacenameContext) {
	fmt.Println(ctx.GetText())
}

func (s *CPPCallListener) EnterFunctiondefinition(ctx *FunctiondefinitionContext) {
	fmt.Println(ctx.Declspecifierseq().GetText())
	fmt.Println(ctx.Declarator().GetText())
}