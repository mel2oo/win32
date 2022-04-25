package api

import (
	"encoding/xml"
)

// Headers xml
type HeadersXml struct {
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

type Include struct {
	Filename string `xml:"Filename,attr"`
}

type Headers struct {
	Condition []Condition `xml:"Condition"`
	Variable  []Variable  `xml:"Variable"`
}

type Architecture int

const (
	W64 Architecture = 64
	W32 Architecture = 32
)

type Condition struct {
	Architecture Architecture `xml:"Architecture,attr"`
	Variable     []Variable   `xml:"Variable"`
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
	Count      int     `xml:"Count,attr"`
	Display    Display `xml:"Display"`
	Field      []Field `xml:"Field"`
	Enum       Enum    `xml:"Enum"`
	Flag       Flag    `xml:"Flag"`
}

type Display struct {
	Name string `xml:"Name,attr"`
}

type Field struct {
	Type       string `xml:"Type,attr"`
	Name       string `xml:"Name,attr"`
	Count      string `xml:"Count,attr"`
	PostCount  string `xml:"PostCount,attr"`
	Length     string `xml:"Length,attr"`
	PostLength string `xml:"PostLength,attr"`
	Display    string `xml:"Display,attr"`
	OutputOnly bool   `xml:"OutputOnly,attr"`
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

// Interface xml
type InterfaceXml struct {
	File      string
	Root      xml.Name  `xml:"ApiMonitor"`
	Include   []Include `xml:"Include"`
	Interface Interface `xml:"Interface"`
}

type Interface struct {
	Name          string `xml:"Name,attr"`
	Id            string `xml:"Id,attr"`
	BaseInterface string `xml:"BaseInterface,attr"`
	OnlineHelp    string `xml:"OnlineHelp,attr"`
	ErrorFunc     string `xml:"ErrorFunc,attr"`
	Category      string `xml:"Category,attr"`
	Api           []Api  `xml:"Api"`
}

// ApiSetSchema xml
type ApiSetSchemaXml struct {
	Module      string      `xml:"Module,attr"`
	Redirection Redirection `xml:"Redirection"`
}

type Redirection struct {
	Module string `xml:"Module,attr"`
}

// UnsupportedModules xml
type UnsupportedModulesXml struct {
	UnsupportedModules UnsupportedModules `xml:"UnsupportedModules"`
}

type UnsupportedModules struct {
	Module Module `xml:"Module"`
}

// Module xml
type ModuleXml struct {
	File    string
	Root    xml.Name  `xml:"ApiMonitor"`
	Include []Include `xml:"Include"`
	Module  Module    `xml:"Module"`
}

type Module struct {
	Name              string      `xml:"Name,attr"`
	CallingConvention string      `xml:"CallingConvention,attr"`
	ErrorFunc         string      `xml:"ErrorFunc,attr"`
	OnlineHelp        string      `xml:"OnlineHelp,attr"`
	Condition         []Condition `xml:"Condition"`
	Variable          []Variable  `xml:"Variable"`
	Api               []Api       `xml:"Api"`
}

type Api struct {
	Name        string  `xml:"Name,attr"`
	BothCharset bool    `xml:"BothCharset,attr"`
	Param       []Param `xml:"Param"`
	Return      Return  `xml:"Return"`
	Success     Success `xml:"Success"`
}

type Param struct {
	Type        string `xml:"Type,attr"`
	Name        string `xml:"Name,attr"`
	InterfaceId string `xml:"InterfaceId,attr"`
}

type Return struct {
	Type string `xml:"Type,attr"`
}

type Success struct {
	Return string `xml:"Return,attr"`
	Value  string `xml:"Value,attr"`
}
