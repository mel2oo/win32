package api

import (
	"encoding/xml"
)

const (
	Win32Dir  = "API/Windows"
	HeaderDir = "API/Headers"
)

// Api xml
type Win32Xml struct {
	File    string
	Root    xml.Name  `xml:"ApiMonitor"`
	Include []Include `xml:"Include"`
	Module  Module    `xml:"Module"`
}

type Include struct {
	Filename string `xml:"Filename,attr"`
}

type Module struct {
	Name              string     `xml:"Name,attr"`
	CallingConvention string     `xml:"CallingConvention,attr"`
	ErrorFunc         string     `xml:"ErrorFunc,attr"`
	OnlineHelp        string     `xml:"OnlineHelp,attr"`
	Category          Category   `xml:"Category"`
	Api               []Api      `xml:"Api"`
	Variable          []Variable `xml:"Variable"`
}

type Category struct {
	Name string `xml:"Name,attr"`
}

type Api struct {
	Name        string  `xml:"Name,attr"`
	BothCharset bool    `xml:"BothCharset,attr"`
	Param       []Param `xml:"Param"`
	Return      Return  `xml:"Return"`
	Success     Success `xml:"Success"`
}

type Param struct {
	Type string `xml:"Type,attr"`
	Name string `xml:"Name,attr"`
}

type Return struct {
	Type string `xml:"Type,attr"`
}

type Success struct {
	Return string `xml:"Return,attr"`
	Value  string `xml:"Value,attr"`
}

type Type string

const (
	TypeInteger          Type = "Integer"
	TypeModuleHandle     Type = "ModuleHandle"
	TypePointer          Type = "Pointer"
	TypeAlias            Type = "Alias"
	TypeVoid             Type = "Void"
	TypeStruct           Type = "Struct"
	TypeEnum             Type = "Enum"
	TypeFLAG             Type = "flag"
	TypeUnion            Type = "Union"
	TypeArray            Type = "Array"
	TypeFloating         Type = "Floating"
	TypeGuid             Type = "Guid"
	TypeInterface        Type = "Interface"
	TypeCharacter        Type = "Character"
	TypeUnicodeCharacter Type = "UnicodeCharacter"
	TypeTCharacter       Type = "TCharacter"
)

type Variable struct {
	Name       string  `xml:"Name,attr"`
	Type       Type    `xml:"Type,attr"`
	Base       string  `xml:"Base,attr"`
	Size       int     `xml:"Size,attr"`
	Unsigned   bool    `xml:"Unsigned,attr"`
	DisplayHex bool    `xml:"DisplayHex,attr"`
	Count      string  `xml:"Count,attr"`
	Display    Display `xml:"Display"`
	Field      []Field `xml:"Field"`
	Enum       Enum    `xml:"Enum"`
	Flag       Flag    `xml:"Flag"`
}

type Display struct {
	Name string `xml:"Name,attr"`
}

type Field struct {
	Type string `xml:"Type,attr"`
	Name string `xml:"Name,attr"`
}

type Enum struct {
	Set []Set `xml:"Set"`
}

type Flag struct {
	Set []Set `xml:"Set"`
}

type Set struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

// Header xml
type HeaderXml struct {
	File    string
	Root    xml.Name  `xml:"ApiMonitor"`
	HelpUrl []HelpUrl `xml:"HelpUrl"`
	Include []Include `xml:"Include"`
	Headers Headers   `xml:"Headers"`
}

type HelpUrl struct {
	Name string `xml:"Name,attr"`
	Url  string `xml:"Url,attr"`
}

type Headers struct {
	Condition []Condition `xml:"Condition"`
	Variable  []Variable  `xml:"Variable"`
}

type Architecture int

const (
	X64 Architecture = 64
	X86 Architecture = 86
)

type Condition struct {
	Architecture Architecture `xml:"Architecture,attr"`
	Variable1    []Variable   `xml:"Variable"`
}

// const (
// 	_pkg = "package main\n"
// )

// func (a *ApiXml) deeparse() error {
// 	for _, inc := range a.Include {
// 		file := fmt.Sprintf("API\\%s", inc.Filename)

// 		data, err := os.ReadFile(file)
// 		if err != nil {
// 			return err
// 		}

// 		v := new(HeaderXml)
// 		if err = xml.Unmarshal(data, v); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (a *ApiXml) makefile() {
// 	data := _pkg

// 	for _, v := range a.Module.Variable {
// 		switch v.Type {
// 		case _Struct:
// 			h := fmt.Sprintf("\ntype %s struct {", v.Name)
// 			for _, vv := range v.Field {
// 				h += fmt.Sprintf("\n	%s %s", vv.Name, vv.Type)
// 			}
// 			h += "\n}\n"
// 			data += h
// 		case _Enum:
// 			h := fmt.Sprintf("type %s %s", v.Type)
// 		case _FLAG:
// 		}
// 	}
// }
