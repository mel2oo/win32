package main

import "encoding/xml"

type Parser interface {
	P()
}

// Api xml
type ApiXml struct {
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
	Integer          Type = "Integer"
	ModuleHandle     Type = "ModuleHandle"
	Pointer          Type = "Pointer"
	Alias            Type = "Alias"
	Void             Type = "Void"
	Struct           Type = "Struct"
	Union            Type = "Union"
	Array            Type = "Array"
	Floating         Type = "Floating"
	Guid             Type = "Guid"
	Interface        Type = "Interface"
	Character        Type = "Character"
	UnicodeCharacter Type = "UnicodeCharacter"
	TCharacter       Type = "TCharacter"
)

type Variable struct {
	Name    string  `xml:"Name,attr"`
	Type    Type    `xml:"Type,attr"`
	Base    string  `xml:"Base,attr"`
	Count   string  `xml:"Count,attr"`
	Display Display `xml:"Display"`
	Field   []Field `xml:"Field"`
	Enum    Enum    `xml:"Enum"`
	Flag    Flag    `xml:"Flag"`
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
	Variable1    []Variable1  `xml:"Variable"`
}

type Size int

const (
	Size4 Size = 4
	Size8 Size = 8
)

type Variable1 struct {
	Name     string `xml:"Name,attr"`
	Type     Type   `xml:"Type,attr"`
	Size     Size   `xml:"Size,attr"`
	Unsigned bool   `xml:"Unsigned,attr"`
}
