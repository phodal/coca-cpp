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

func (s *CPPCallListener) EnterFunctionbody(ctx *FunctionbodyContext) {

}

// swap(a,b)
func (s *CPPCallListener) EnterAssignmentexpression(ctx *AssignmentexpressionContext) {
	//fmt.Println(ctx.GetText())
}

func (s *CPPCallListener) EnterLogicalorexpression(ctx *LogicalorexpressionContext) {
	//fmt.Println(ctx.GetText())
}

func (s *CPPCallListener) EnterUnaryexpression(ctx *UnaryexpressionContext) {
	//fmt.Println(ctx.GetText())
}

func (s *CPPCallListener) EnterPostfixexpression(ctx *PostfixexpressionContext) {
	if ctx.GetText() == "swap(a,b)" {
		fmt.Println(".......................")
		fmt.Println(ctx.GetChildren())
		fmt.Println(ctx.Postfixexpression())
		fmt.Println(ctx.Simpletypespecifier())
		fmt.Println(ctx.Typenamespecifier())
		fmt.Println(ctx.Expressionlist().GetText())
		fmt.Println("////////////////////////")
	}
}

