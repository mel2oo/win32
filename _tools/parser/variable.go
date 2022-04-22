package main

import (
	"strconv"

	"github.com/mel2oo/win32/_tools/parser/api"
)

type Variable interface {
	Name() string
	Base() Variable
	Size() int
	Text(int) string
}

func GetVariable(v api.Variable) Variable {
	if v.Type == api.TypeInteger {
		return &VarTypeInteger{
			name:       v.Name,
			base:       v.Base,
			size:       v.Size,
			displayhex: v.DisplayHex,
		}
	}
}

type VarTypeInteger struct {
	name       string
	size       int
	unsigned   bool
	displayhex bool
}

func (v *VarTypeInteger) Name() string {
	return v.name
}

func (v *VarTypeInteger) Base() Variable {
	return nil
}

func (v *VarTypeInteger) Size() int {
	return v.size
}

func (v *VarTypeInteger) Text(int) string {
	base := 10
	if vt.displayHex {
		base = 16
	}
	if vt.unsigned {
		return strconv.FormatUint(uint64(n), base)
	}
	return strconv.FormatInt(int64(n), base)
}
