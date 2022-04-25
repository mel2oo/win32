package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mel2oo/win32/_tools/parser/api"
)

var (
	amd64   = flag.Bool("amd64", false, "build api in architeture amd64")
	xmlfile = flag.String("xmlfile", "Windows/Advapi32.xml", "parse xml file")
)

func main() {
	flag.Parse()

	res := new()
	if err := res.parse(); err != nil {
		panic(err)
	}

	res.generate()
}

type Result struct {
	headerMap map[string]*Header
	// interfaceMap map[string]int
	moduleMap map[string]*Module
}

type Header struct {
	Includes  map[string]*Header
	Variables []Variable
}

type Interface struct {
	Includes map[string]*Header
}

type Module struct {
	Includes  map[string]*Header
	Variables []Variable
	Apis      []*Api
}

func new() *Result {
	return &Result{
		headerMap: make(map[string]*Header),
		moduleMap: make(map[string]*Module),
	}
}

func (r *Result) parse() error {

	v, err := r.readModuleXml(*xmlfile)
	if err != nil {
		return err
	}

	r.moduleMap[v.File] = r.buildModule(v)

	return nil
}

func (r *Result) readHeaderXml(xmlpath string) (*api.HeadersXml, error) {
	v := &api.HeadersXml{File: xmlpath}

	data, err := os.ReadFile(v.File)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(data, v); err != nil {
		return nil, err
	}

	return v, nil
}

func (r *Result) buildHeader(xmlpath string) *Header {
	h := &Header{
		Includes:  make(map[string]*Header),
		Variables: make([]Variable, 0),
	}

	v, err := r.readHeaderXml(xmlpath)
	if err != nil {
		return nil
	}

	for _, vv := range v.Include {
		key := vv.Filename
		if _, ok := r.headerMap[key]; !ok {
			h.Includes[key] = r.buildHeader(vv.Filename)
		} else {
			h.Includes[key] = r.headerMap[key]
		}
	}

	for _, vv := range v.Headers.Condition {
		if (vv.Architecture == api.W64 && *amd64) ||
			(vv.Architecture == api.W32 && !*amd64) {
			for _, vvv := range vv.Variable {
				h.Variables = append(h.Variables, SetVariable(v.Headers.Variable, vvv))
			}
		}
	}

	for _, vv := range v.Headers.Variable {
		h.Variables = append(h.Variables, SetVariable(v.Headers.Variable, vv))
	}

	r.headerMap[xmlpath] = h

	return h
}

func (r *Result) readModuleXml(xmlpath string) (*api.ModuleXml, error) {
	v := &api.ModuleXml{File: xmlpath}

	data, err := os.ReadFile(v.File)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(data, v); err != nil {
		return nil, err
	}

	return v, nil
}

func (r *Result) buildModule(v *api.ModuleXml) *Module {
	w := &Module{
		Includes:  make(map[string]*Header),
		Variables: make([]Variable, 0),
		Apis:      make([]*Api, 0),
	}

	for _, vv := range v.Include {
		key := vv.Filename
		if _, ok := r.headerMap[key]; !ok {
			w.Includes[key] = r.buildHeader(vv.Filename)
		} else {
			w.Includes[key] = r.headerMap[key]
		}
	}

	for _, vv := range v.Module.Condition {
		if (vv.Architecture == api.W64 && *amd64) ||
			(vv.Architecture == api.W32 && !*amd64) {
			for _, vvv := range vv.Variable {
				w.Variables = append(w.Variables, SetVariable(v.Module.Variable, vvv))
			}
		}
	}

	for _, vv := range v.Module.Variable {
		w.Variables = append(w.Variables, SetVariable(v.Module.Variable, vv))
	}

	for _, vv := range v.Module.Api {
		w.Apis = append(w.Apis, GetApi(w.Includes, w.Variables, vv))
	}

	return w
}

func (r *Result) generate() error {
	for n, h := range r.headerMap {
		var (
			fdt = "package api\n\n"
			sp1 = strings.Split(n, "\\")
			sp2 = strings.Split(sp1[len(sp1)-1], ".")
			mdn = sp2[0] + ".go"
		)

		for _, v := range h.Variables {
			fdt += fmt.Sprintf("var %s %s\n\n", v.Name(), GetBase(v))

			if v.Set() != nil {
				var cst string
				for def, val := range v.Set() {
					cst += fmt.Sprintf("%s %s %s\n", def, v.Name(), val)
				}
				fdt += fmt.Sprintf("const (\n%s)\n\n", cst)
			}

			if v.Field() != nil {
				var cst string
				for _, field := range v.Field() {
					cst += fmt.Sprintf("%s %s\n", field.name, field.base.Name())
				}
				fdt += fmt.Sprintf("type %s struct{\n%s}\n\n", v.Name(), cst)
			}
		}

		fs, err := os.Create(mdn)
		if err != nil {
			return err
		}

		if _, err = fs.Write([]byte(fdt)); err != nil {
			return err
		}

		fs.Close()
	}

	for n, m := range r.moduleMap {
		fmt.Println(n, m)
	}

	return nil
}
