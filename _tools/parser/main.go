package main

import (
	"encoding/xml"
	"flag"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mel2oo/win32/_tools/parser/api"
)

var amd64 = flag.Bool("amd64", true, "build api in architeture amd64")

func main() {
	flag.Parse()

	res := new()
	if err := res.parse(); err != nil {
		panic(err)
	}
}

type Result struct {
	headers map[string]*api.HeaderXml
	win32s  map[string]*api.Win32Xml

	headerMap map[string]*Header
	win32Map  map[string]*Win32
}

func new() *Result {
	return &Result{
		headers:   make(map[string]*api.HeaderXml),
		win32s:    make(map[string]*api.Win32Xml),
		headerMap: make(map[string]*Header),
		win32Map:  make(map[string]*Win32),
	}
}

func (r *Result) parse() (err error) {

	err = filepath.Walk(api.HeaderDir, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		v := &api.HeaderXml{File: path}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}

		r.headers[filepath.Base(v.File)] = v

		return nil
	})
	if err != nil {
		return
	}

	err = filepath.Walk(api.Win32Dir, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		v := &api.Win32Xml{File: path}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}

		r.win32s[filepath.Base(v.File)] = v

		return nil
	})
	if err != nil {
		return
	}

	for _, v := range r.headers {
		r.headerMap[v.File] = r.buildHeader(v.File, v)
	}

	return
}

type Header struct {
	Includes  map[string]*Header
	Variables []Variable
}

type Win32 struct {
	Includes  []string
	Variables []Variable
	Apis      []Api
}

type Api struct{}

func (r *Result) buildHeader(key string, v *api.HeaderXml) *Header {
	if v == nil {
		return nil
	}

	h := &Header{
		Includes:  make(map[string]*Header),
		Variables: make([]Variable, 0),
	}

	for _, vv := range v.Include {
		key := filepath.Base(vv.Filename)
		header := r.buildHeader(key, r.headers[key])
		h.Includes[key] = header
		r.headerMap[key] = header
	}

	for _, vv := range v.Headers.Condition {
		if (vv.Architecture == api.X64 && *amd64) ||
			(vv.Architecture == api.X86 && !*amd64) {
			for _, vvv := range vv.Variable {
				h.Variables = append(h.Variables, GetVariable(vvv))
			}
		}
	}

	for _, vv := range v.Headers.Variable {
		GetVariable(vv)
	}

	return h
}

func (r *Result) buildWin32() {

}
