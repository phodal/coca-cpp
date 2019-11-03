package main

import (
	. "./adapter"
)

func main()  {
	callApp := new(CPPCallApp)
	callApp.Analysis("examples/cpp-basic/function-call/")
}
