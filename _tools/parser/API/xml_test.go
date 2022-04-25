package api

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestReadHeader(t *testing.T) {
	v := HeadersXml{
		File: "Headers/windows.h.xml",
	}

	data, err := os.ReadFile(v.File)
	if err != nil {
		t.Fail()
		return
	}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		t.Fail()
		return
	}
}

func TestReadInterface(t *testing.T) {
	v := InterfaceXml{
		File: "Interfaces/IClassFactory.xml",
	}

	data, err := os.ReadFile(v.File)
	if err != nil {
		t.Fail()
		return
	}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		t.Fail()
		return
	}
}

func TestReadApiSetSchema(t *testing.T) {

}

func TestReadUnsupportedModules(t *testing.T) {

}

func TestReadModule(t *testing.T) {

}
