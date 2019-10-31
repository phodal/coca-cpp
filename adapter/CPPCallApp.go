package adapter

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "../language/cplus"
)

type CPPCallApp struct {

}

func (j *CPPCallApp) Analysis(codeDir string) {
	files := (*CPPCallApp)(nil).getFiles(codeDir)
	for _, file := range files {
		fmt.Println(file)
		parser := (*CPPCallApp)(nil).processFile(file)
		context := parser.Translationunit()
		listener := NewCPPCallListener()

		antlr.NewParseTreeWalker().Walk(listener, context)
	}
}

func (j *CPPCallApp) getFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".cpp") || strings.Contains(path, ".h") {
			files = append(files, path)
		}
		return nil
	})
	return files
}


func (j *CPPCallApp) processFile(path string) *CPP14Parser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewCPP14Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewCPP14Parser(stream)
	return parser
}