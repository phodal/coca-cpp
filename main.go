package main

import (
	. "./adapter"
)

func main()  {
	callApp := new(CPPCallApp)
	callApp.Analysis("examples/hello")
}
