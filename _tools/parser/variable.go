package main

import (
	"fmt"
	"strings"

	"github.com/mel2oo/win32/_tools/parser/api"
)

var vbs map[string]Variable

func init() {
	vbs = make(map[string]Variable)
}

type Variable interface {
	Name() string
	Base() Variable
	Size() int
	Pack() int
	Set() []Set
	Field() []Field
}

func SetVariable(vs []api.Variable, v api.Variable) Variable {
	var vb Variable

	switch v.Type {
	case api.TypeInteger:
		vv := &VarTypeInteger{
			name:       v.Name,
			size:       v.Size,
			unsigned:   v.Unsigned,
			displayhex: v.DisplayHex,
		}

		if v.Enum.Set != nil {
			set := make([]Set, 0)
			for _, s := range v.Enum.Set {
				set = append(set, Set{
					key:   s.Name,
					value: s.Value,
				})
			}
			vv.set = set

		} else if v.Flag.Set != nil {
			set := make([]Set, 0)

			for _, s := range v.Flag.Set {
				set = append(set, Set{
					key:   s.Name,
					value: s.Value,
				})
			}
			vv.set = set
		}

		vb = vv

	case api.TypeModuleHandle:
		vb = &VarTypeModuleHandle{
			name: v.Name,
		}

	case api.TypePointer:
		base := vbs[v.Base]
		if base == nil {
		loop1:
			for _, vv := range vs {
				if vv.Name == v.Base {
					base = SetVariable(vs, vv)
					break loop1
				}
			}
		}
		vb = &VarTypePointer{
			name: v.Name,
			base: base,
		}

	case api.TypeAlias:

		base := vbs[v.Base]
		if base == nil {
		loop2:
			for _, vv := range vs {
				if vv.Name == v.Base {
					base = SetVariable(vs, vv)
					break loop2
				}
			}
		}
		if base == nil && strings.HasSuffix(v.Base, "*") {
			bname := strings.TrimSuffix(v.Base, "*")
		loop3:
			for _, vv := range vs {
				if vv.Name == bname {
					base = SetVariable(vs, vv)
					break loop3
				}
			}
		}

		vv := &VarTypeAlias{
			name:    v.Name,
			base:    base,
			display: v.Display.Name,
		}

		if v.Enum.Set != nil {
			set := make([]Set, 0)
			for _, s := range v.Enum.Set {
				set = append(set, Set{
					key:   s.Name,
					value: s.Value,
				})
			}
			vv.set = set

		} else if v.Flag.Set != nil {
			set := make([]Set, 0)
			for _, s := range v.Flag.Set {
				set = append(set, Set{
					key:   s.Name,
					value: s.Value,
				})
			}
			vv.set = set
		}

		vb = vv

	case api.TypeVoid:
		vb = &VarTypeVoid{
			name: v.Name,
		}

	case api.TypeCharacter:
		vb = &VarTypeCharacter{
			name: v.Name,
		}

	case api.TypeUnicodeCharacter:
		vb = &VarTypeUnicodeCharacter{
			name: v.Name,
		}

	case api.TypeTCharacter:
		vb = &VarTypeTCharacter{
			name: v.Name,
		}

	case api.TypeFloating:
		vb = &VarTypeFloating{
			name: v.Name,
			size: v.Size,
		}

	case api.TypeArray:
		base := vbs[v.Base]
		if base == nil {
		loop4:
			for _, vv := range vs {
				if vv.Name == v.Base {
					base = SetVariable(vs, vv)
					break loop4
				}
			}
		}
		vb = &VarTypeArray{
			name:  v.Name,
			base:  base,
			count: v.Count,
		}

	case api.TypeInterface:

		vb = &VarTypeInterface{
			name: v.Name,
		}

	case api.TypeStruct:
		fields := make([]Field, 0)
		for _, f := range v.Field {
			base := vbs[f.Type]
			if base == nil {
			loop5:
				for _, vv := range vs {
					if vv.Name == f.Type {
						base = SetVariable(vs, vv)
						break loop5
					}
				}
			}

			fields = append(fields, Field{
				name: f.Name,
				base: base,
			})
		}

		vv := &VarTypeStruct{
			name:   v.Name,
			fields: fields,
		}

		vv.Init()
		vb = vv

	case api.TypeUnion:
		fields := make([]Field, 0)
		for _, f := range v.Field {
			base := vbs[f.Type]
			if base == nil {
			loop6:
				for _, vv := range vs {
					if vv.Name == f.Type {
						base = SetVariable(vs, vv)
						break loop6
					}
				}
			}
			fields = append(fields, Field{
				name: f.Name,
				base: base,
			})
		}

		vv := &VarTypeUnion{
			name:   v.Name,
			fields: fields,
		}

		vv.Init()
		vb = vv

	case api.TypeGuid:
		vb = &VarTypeGuid{
			name: v.Name,
		}

	default:
		fmt.Println("undefine type", v.Type)
		return nil
	}

	vbs[v.Name] = vb
	return vb
}

type VarTypeInteger struct {
	name       string
	size       int
	unsigned   bool
	displayhex bool
	set        []Set
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

func (v *VarTypeInteger) Pack() int {
	return v.size
}

type Set struct {
	key   string
	value string
}

func (v *VarTypeInteger) Set() []Set {
	return v.set
}

func (v *VarTypeInteger) Field() []Field {
	return nil
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

func (v *VarTypeModuleHandle) Pack() int {
	if *amd64 {
		return 8
	}
	return 4
}

func (v *VarTypeModuleHandle) Set() []Set {
	return nil
}

func (v *VarTypeModuleHandle) Field() []Field {
	return nil
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

func (v *VarTypePointer) Pack() int {
	if *amd64 {
		return 8
	}
	return 4
}

func (v *VarTypePointer) Set() []Set {
	return nil
}

func (v *VarTypePointer) Field() []Field {
	return nil
}

type VarTypeAlias struct {
	name    string
	base    Variable
	display string
	set     []Set
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

func (v *VarTypeAlias) Pack() int {
	return v.base.Pack()
}

func (v *VarTypeAlias) Set() []Set {
	return v.set
}

func (v *VarTypeAlias) Field() []Field {
	return nil
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

func (v *VarTypeVoid) Pack() int {
	return 0
}

func (v *VarTypeVoid) Set() []Set {
	return nil
}

func (v *VarTypeVoid) Field() []Field {
	return nil
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

func (v *VarTypeCharacter) Pack() int {
	return 1
}

func (v *VarTypeCharacter) Set() []Set {
	return nil
}

func (v *VarTypeCharacter) Field() []Field {
	return nil
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

func (v *VarTypeUnicodeCharacter) Pack() int {
	return 2
}

func (v *VarTypeUnicodeCharacter) Set() []Set {
	return nil
}

func (v *VarTypeUnicodeCharacter) Field() []Field {
	return nil
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

func (v *VarTypeTCharacter) Pack() int {
	return 0
}

func (v *VarTypeTCharacter) Set() []Set {
	return nil
}

func (v *VarTypeTCharacter) Field() []Field {
	return nil
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

func (v *VarTypeFloating) Pack() int {
	return v.size
}

func (v *VarTypeFloating) Set() []Set {
	return nil
}

func (v *VarTypeFloating) Field() []Field {
	return nil
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

func (v *VarTypeArray) Pack() int {
	return v.base.Pack()
}

func (v *VarTypeArray) Set() []Set {
	return nil
}

func (v *VarTypeArray) Field() []Field {
	return nil
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

func (v *VarTypeInterface) Pack() int {
	if *amd64 {
		return 8
	}
	return 4
}

func (v *VarTypeInterface) Set() []Set {
	return nil
}

func (v *VarTypeInterface) Field() []Field {
	return nil
}

type Field struct {
	name   string
	base   Variable
	offset int
}

func (f *Field) Offset() int {
	return f.offset
}

type VarTypeStruct struct {
	name   string
	fields []Field
	size   int
	pack   int
}

func (v *VarTypeStruct) Init() {
	var (
		size   int
		pack   int
		offset int
	)

	if v.pack == 0 {
		v.pack = 8
	}

	for i := range v.fields {
		base := v.fields[i].base
		if base == nil {
			continue
		}

		m := base.Pack()
		if m > v.pack {
			m = v.pack
		}
		if pack < m {
			pack = m
		}

		offset = (offset + m - 1) &^ (m - 1)
		v.fields[i].offset = offset
		offset += base.Size()
	}

	size = (offset + pack - 1) &^ (pack - 1)

	v.size = size
	v.pack = pack
}

func (v *VarTypeStruct) Name() string {
	return v.name
}

func (v *VarTypeStruct) Base() Variable {
	return nil
}

func (v *VarTypeStruct) Size() int {
	return v.size
}

func (v *VarTypeStruct) Pack() int {
	return v.pack
}

func (v *VarTypeStruct) Set() []Set {
	return nil
}

func (v *VarTypeStruct) Field() []Field {
	return v.fields
}

type VarTypeUnion struct {
	name   string
	fields []Field
	size   int
	pack   int
}

func (v *VarTypeUnion) Init() {
	var (
		size int
		pack int
	)

	if v.pack == 0 {
		v.pack = 8
	}

	for i := range v.fields {
		base := v.fields[i].base
		if base == nil {
			continue
		}

		m := base.Pack()
		if m > v.pack {
			m = v.pack
		}
		if pack < m {
			pack = m
		}

		n := base.Size()
		if size < n {
			size = n
		}
	}

	v.size = size
	v.pack = pack
}

func (v *VarTypeUnion) Name() string {
	return v.name
}

func (v *VarTypeUnion) Base() Variable {
	return nil
}

func (v *VarTypeUnion) Size() int {
	return v.size
}

func (v *VarTypeUnion) Pack() int {
	return v.pack
}

func (v *VarTypeUnion) Set() []Set {
	return nil
}

func (v *VarTypeUnion) Field() []Field {
	return v.fields
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

func (v *VarTypeGuid) Pack() int {
	return 4
}

func (v *VarTypeGuid) Set() []Set {
	return nil
}

func (v *VarTypeGuid) Field() []Field {
	return nil
}

// private
func GetBase(v Variable) string {
	for {
		if v.Base() == nil {
			return v.Name()
		}
		v = v.Base()
	}
}
