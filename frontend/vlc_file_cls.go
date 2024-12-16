package frontend

import (
	list "github.com/duke-git/lancet/v2/datastructure/list"
	stack "github.com/duke-git/lancet/v2/datastructure/stack"
	"os"
)

type VlcFile struct {
	Name       string
	Data       string
	File       *os.File
	Info       os.FileInfo
	Line       int
	Column     int
	Index      int
	Tokens     []*VlcToken
	ScopeStack stack.LinkedStack[*VlcScope]
	Scopes     list.List[*VlcScope]
}
