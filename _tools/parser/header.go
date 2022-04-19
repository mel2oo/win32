package main

import "encoding/xml"

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

type Condition struct {
	Architecture Architecture `xml:"Architecture,attr"`
	Variable1    []Variable1  `xml:"Variable"`
}

type Architecture int

const (
	X64 Architecture = 64
	X86 Architecture = 86
)

type Variable1 struct {
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Type,attr"`
	Size     int    `xml:"Size,attr"`
	Unsigned bool   `xml:"Unsigned,attr"`
}

func (h *HeaderXml) P() {}
