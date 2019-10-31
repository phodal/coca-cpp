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

func (s *CPPCallListener) EnterClassname(ctx *ClassnameContext) {
	fmt.Println(ctx.Identifier().GetText())
}

func (s *CPPCallListener) EnterFunctiondefinition(ctx *FunctiondefinitionContext) {
	if ctx.Declspecifierseq() != nil {
		fmt.Println(ctx.Declspecifierseq().GetText())
	}
	fmt.Println(ctx.Declarator().GetText())
}