package frontend

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strings"
	"sync"
	"vilac/consts"
)

type VlcCompiler struct {
	Files       []VlcFile
	CurrentFile *VlcFile
}

var _vc *VlcCompiler
var _vcOnce sync.Once

func NewVlcCompiler(files []VlcFile) *VlcCompiler {
	_vcOnce.Do(func() {
		_vc = new(VlcCompiler)
		_vc.Files = files
	})
	return _vc
}

func (vc *VlcCompiler) Sources() []*VlcFile {
	var sources = make([]*VlcFile, 0)
	for _, file := range vc.Files {
		sources = append(sources, &file)
	}
	return sources
}

func (vc *VlcCompiler) Compile() error {

	for idx, file := range vc.Files {
		err := vc.readFile(idx, file.Name)
		if err != nil {
			return err
		}
		// Token解析
		err = vc.parseTokens()
		if err != nil {
			return err
		}
		// Scope分析
		err = vc.parseScopes()
		if err != nil {
			return err
		}
		// 语法分析
	}
	return nil
}

func (vc *VlcCompiler) readFile(idx int, name string) error {
	var wd, err = os.Getwd()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/%s", wd, name), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	vc.Files[idx].File = file
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	vc.Files[idx].Data = string(data)
	vc.CurrentFile = &vc.Files[idx]
	return vc.Files[idx].File.Close()
}

func (vc *VlcCompiler) modifyToken(value string, modifier uint8) {
	if len(vc.CurrentFile.Tokens) == 0 {
		return
	}
	vc.CurrentFile.Tokens[len(vc.CurrentFile.Tokens)-1].Modifier |= modifier
}

func (vc *VlcCompiler) pushToken(value string, tokenType int, isArray bool) {
	if len(value) == 0 {
		return
	}
	if slices.Contains(keywords, value) {
		tokenType = consts.TokenTypeKeyword
	} else if slices.Contains(dataTypes, value) {
		tokenType = consts.TokenTypeDataType
	} else if modifiers[value] > 0 {
		vc.modifyToken(value, modifiers[value])
	}
	vc.CurrentFile.Tokens = append(vc.CurrentFile.Tokens, &VlcToken{
		Token:   value,
		Type:    tokenType,
		IsArray: isArray,
	})
}

func (vc *VlcCompiler) parseTokens() error {
	var vcf = vc.CurrentFile
	var tokenValue strings.Builder
	var isString = false
	for i := 0; i < len(vcf.Data); i++ {
		var token = rune(vcf.Data[i])
		if token == '\n' {
			vcf.Line++
			vcf.Column = 0
		} else {
			vcf.Column++
		}
		switch token {
		case '"':
			isString = !isString
			if !isString {
				vc.pushToken(tokenValue.String(), consts.TokenTypeString, false)
				tokenValue.Reset()
			}
			break
		case ' ', '\t', '\n', '\r':
			if isString {
				tokenValue.WriteRune(token)
				break
			}
			vc.pushToken(tokenValue.String(), consts.TokenTypeIdentifier, false)
			tokenValue.Reset()
			break
		case '[':
			if vcf.Data[i+1] == ']' {
				vc.pushToken(tokenValue.String(), consts.TokenTypeIdentifier, true)
				tokenValue.Reset()
				i++
				break
			}
			break
		case ',', '(', ')', ';', '{', '}':
			if isString {
				tokenValue.WriteRune(token)
				break
			}
			vc.pushToken(tokenValue.String(), consts.TokenTypeIdentifier, false)
			tokenValue.Reset()
			vc.pushToken(string(token), int(token), false)
			break
		default:
			tokenValue.WriteRune(token)
		}
	}
	return nil
}

func (vc *VlcCompiler) pushScope(scope *VlcScope) {
	vc.CurrentFile.Scopes.Push(scope)
	vc.CurrentFile.ScopeStack.Push(scope)
}

func (vc *VlcCompiler) popScope() *VlcScope {
	v, err := vc.CurrentFile.ScopeStack.Pop()
	if err != nil {
		return nil
	}
	return *v
}

func (vc *VlcCompiler) currentScope() *VlcScope {
	v, err := vc.CurrentFile.ScopeStack.Peak()
	if err != nil {
		return nil
	}
	return *v
}

func (vc *VlcCompiler) parseScopes() error {
	// 函数
	for i := vc.CurrentFile.Index; i < len(vc.CurrentFile.Tokens); i = vc.CurrentFile.Index {
		var token = vc.CurrentFile.Tokens[i]
		if token.Type == consts.TokenTypeDataType {
			if is, scope := vc.parseScopeFunc(); is {
				vc.pushScope(scope)
			}
		} else if token.Type == consts.TokenTypeIdentifier {
			if v := vc.currentScope(); v != nil {
				// 放入Scope中
			}
			fmt.Println(vc.parseCall())
		} else if token.Token == "}" {
			vc.popScope()
			vc.CurrentFile.Index++
		}
	}
	return nil
}

func (vc *VlcCompiler) parseScope(sm [][]uint8, ei int, scopeType int) (bool, *VlcScope) {
	var si = 0
	var start = vc.CurrentFile.Index
	for si < ei {
		si = int(sm[si][vc.CurrentFile.Tokens[vc.CurrentFile.Index].Type])
		if si == math.MaxUint8 {
			return false, nil
		} else {
			vc.CurrentFile.Index++
		}
	}
	return si == ei, NewVlcScope(scopeType, vc.CurrentFile.Tokens[start:vc.CurrentFile.Index])
}

func (vc *VlcCompiler) parseCall() (bool, *VlcScope) {
	sm, ei := CallSM()
	return vc.parseScope(sm, ei, consts.ScopeTypeCall)
}

func (vc *VlcCompiler) parseScopeFunc() (bool, *VlcScope) {
	sm, ei := FuncSM()
	return vc.parseScope(sm, ei, consts.ScopeTypeFunc)
}
