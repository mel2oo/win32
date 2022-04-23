package main

import (
	"fmt"

	"github.com/mel2oo/win32/_tools/parser/api"
)

type Variable interface {
	Name() string
	Base() Variable
	Size() int
}

func GetVariable(v api.Variable) Variable {
	switch v.Type {
	case api.TypeInteger:
		return &VarTypeInteger{
			name: v.Name,
			// base:
			size:       v.Size,
			displayhex: v.DisplayHex,
		}

	case api.TypeModuleHandle:
		return &VarTypeModuleHandle{
			name: v.Name,
		}

	case api.TypePointer:
		return &VarTypePointer{
			name: v.Name,
			// base:
		}

	case api.TypeAlias:
		return &VarTypeAlias{
			name: v.Name,
			// base:
			display: v.Display.Name,
		}

	case api.TypeVoid:
		return &VarTypeVoid{
			name: v.Name,
		}

	case api.TypeCharacter:
		return &VarTypeCharacter{
			name: v.Name,
		}

	case api.TypeUnicodeCharacter:
		return &VarTypeUnicodeCharacter{
			name: v.Name,
		}

	case api.TypeTCharacter:
		return &VarTypeTCharacter{
			name: v.Name,
		}

	case api.TypeFloating:
		return &VarTypeFloating{
			name: v.Name,
			size: v.Size,
		}

	case api.TypeArray:
		return &VarTypeArray{
			name: v.Name,
			// base:
			count: v.Count,
		}

	case api.TypeInterface:
		return &VarTypeInterface{
			name: v.Name,
		}

	case api.TypeStruct:
		return &VarTypeStruct{
			name: v.Name,
		}

	case api.TypeUnion:
		return &VarTypeUnion{
			name: v.Name,
		}

	case api.TypeGuid:
		return &VarTypeGuid{
			name: v.Name,
		}

	default:
		fmt.Println("undefine type", v.Type)
		return nil
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

type VarTypeModuleHandle struct {
	name string
}

func (v *VarTypeModuleHandle) Name() string {
	return v.name
}

func (v *VarTypeModuleHandle) Base() Variable {
	return nil
}

func (v *VarTypeModuleHandle) Size() int {
	if *amd64 {
		return 8
	}
	return 4
}

type VarTypePointer struct {
	name string
	base Variable
}

func (v *VarTypePointer) Name() string {
	return v.name
}

func (v *VarTypePointer) Base() Variable {
	return v.base
}

func (v *VarTypePointer) Size() int {
	if *amd64 {
		return 8
	}
	return 4
}

type VarTypeAlias struct {
	name    string
	base    Variable
	display string
}

func (v *VarTypeAlias) Name() string {
	return v.name
}

func (v *VarTypeAlias) Base() Variable {
	return v.base
}

func (v *VarTypeAlias) Size() int {
	return v.base.Size()
}

type VarTypeVoid struct {
	name string
}

func (v *VarTypeVoid) Name() string {
	return v.name
}

func (v *VarTypeVoid) Base() Variable {
	return nil
}

func (v *VarTypeVoid) Size() int {
	return 0
}

type VarTypeCharacter struct {
	name string
}

func (v *VarTypeCharacter) Name() string {
	return v.name
}

func (v *VarTypeCharacter) Base() Variable {
	return nil
}

func (v *VarTypeCharacter) Size() int {
	return 1
}

type VarTypeUnicodeCharacter struct {
	name string
}

func (v *VarTypeUnicodeCharacter) Name() string {
	return v.name
}

func (v *VarTypeUnicodeCharacter) Base() Variable {
	return nil
}

func (v *VarTypeUnicodeCharacter) Size() int {
	return 2
}

type VarTypeTCharacter struct {
	name string
}

func (v *VarTypeTCharacter) Name() string {
	return v.name
}

func (v *VarTypeTCharacter) Base() Variable {
	return nil
}

func (v *VarTypeTCharacter) Size() int {
	return 0
}

type VarTypeFloating struct {
	name string
	size int
}

func (v *VarTypeFloating) Name() string {
	return v.name
}

func (v *VarTypeFloating) Base() Variable {
	return nil
}

func (v *VarTypeFloating) Size() int {
	return v.size
}

type VarTypeArray struct {
	name  string
	base  Variable
	count int
}

func (v *VarTypeArray) Name() string {
	return v.name
}

func (v *VarTypeArray) Base() Variable {
	return nil
}

func (v *VarTypeArray) Size() int {
	return v.base.Size() * v.count
}

type VarTypeInterface struct {
	name string
}

func (v *VarTypeInterface) Name() string {
	return v.name
}

func (v *VarTypeInterface) Base() Variable {
	return nil
}

func (v *VarTypeInterface) Size() int {
	if *amd64 {
		return 8
	}
	return 4
}

type VarTypeStruct struct {
	name   string
	fields []api.Field
}

func (v *VarTypeStruct) Name() string {
	return v.name
}

func (v *VarTypeStruct) Base() Variable {
	return nil
}

func (v *VarTypeStruct) Size() int {
	// todo
	return 0
}

type VarTypeUnion struct {
	name   string
	fields []api.Field
}

func (v *VarTypeUnion) Name() string {
	return v.name
}

func (v *VarTypeUnion) Base() Variable {
	return nil
}

func (v *VarTypeUnion) Size() int {
	// todo
	return 0
}

type VarTypeGuid struct {
	name string
}

func (v *VarTypeGuid) Name() string {
	return v.name
}

func (v *VarTypeGuid) Base() Variable {
	return nil
}

func (v *VarTypeGuid) Size() int {
	return 16
}
