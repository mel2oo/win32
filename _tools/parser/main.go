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
	headers map[string]*api.HeadersXml
	win32s  map[string]*api.ModuleXml

	headerMap map[string]*Header
	win32Map  map[string]*Win32
}

func new() *Result {
	return &Result{
		headers: make(map[string]*api.HeadersXml),
		win32s:  make(map[string]*api.ModuleXml),

		headerMap: make(map[string]*Header),
		win32Map:  make(map[string]*Win32),
	}
}

func (r *Result) parse() (err error) {

	err = filepath.Walk(api.InternalDir, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		v := &api.HeadersXml{File: path}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}

		r.headers[v.File] = v

		return nil
	})
	if err != nil {
		return
	}

	err = filepath.Walk(api.HeaderDir, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		v := &api.HeadersXml{File: path}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}

		r.headers[v.File] = v

		return nil
	})
	if err != nil {
		return
	}

	err = filepath.Walk(api.Win32Dir, func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		v := &api.ModuleXml{File: path}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}

		r.win32s[v.File] = v

		return nil
	})
	if err != nil {
		return
	}

	for _, v := range r.headers {
		r.headerMap[v.File] = r.buildHeader(v)
	}

	for _, v := range r.win32s {
		r.win32Map[v.File] = r.buildWin32(v)
	}

	return
}

type Header struct {
	Includes  map[string]*Header
	Variables []Variable
}

type Win32 struct {
	Includes  map[string]*Header
	Variables []Variable
	Apis      []*Api
}

func (r *Result) buildHeader(v *api.HeadersXml) *Header {
	h := &Header{
		Includes:  make(map[string]*Header),
		Variables: make([]Variable, 0),
	}

	for _, vv := range v.Include {
		key := filepath.Join("API", vv.Filename)
		header := r.buildHeader(r.headers[key])
		h.Includes[key] = header
		r.headerMap[key] = header
	}

	for _, vv := range v.Headers.Condition {
		if (vv.Architecture == api.X64 && *amd64) ||
			(vv.Architecture == api.X86 && !*amd64) {
			for _, vvv := range vv.Variable {
				h.Variables = append(h.Variables, SetVariable(v.Headers.Variable, vvv))
			}
		}
	}

	for _, vv := range v.Headers.Variable {
		h.Variables = append(h.Variables, SetVariable(v.Headers.Variable, vv))
	}

	return h
}

func (r *Result) buildWin32(v *api.ModuleXml) *Win32 {
	w := &Win32{
		Includes:  make(map[string]*Header),
		Variables: make([]Variable, 0),
		Apis:      make([]*Api, 0),
	}

	for _, vv := range v.Include {
		key := filepath.Join("API", vv.Filename)
		w.Includes[key] = r.headerMap[key]
	}

	for _, vv := range v.Module.Condition {
		if (vv.Architecture == api.X64 && *amd64) ||
			(vv.Architecture == api.X86 && !*amd64) {
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
