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

type Variable struct {
	Name    string  `xml:"Name,attr"`
	Type    string  `xml:"Type,attr"`
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

func (a *ApiXml) P() {}
