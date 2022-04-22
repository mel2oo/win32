package main

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mel2oo/win32/_tools/parser/api"
)

func main() {
	if err := parse(); err != nil {
		panic(err)
	}
}

func parse() (err error) {
	headers := make([]*api.HeaderXml, 0)
	win32s := make([]*api.Win32Xml, 0)

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

		headers = append(headers, v)

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

		win32s = append(win32s, v)

		return nil
	})
	if err != nil {
		return
	}

	hmap := make(map[string]Header)

	for _, v := range headers {
		h := Header{
			Includes: make([]string, 0),
		}

		for _, vv := range v.Include {
			h.Includes = append(h.Includes, vv.Filename)
		}

		for _, vv := range v.Headers.Variable {
			fmt.Println(vv)
			// todo

			
		}

		hmap[v.File] = h
	}

	return
}

var (
	headerMap = map[string]Header
	win32Map = map[string]Win32
)

func init(){
	headerMap = make(map[string]Header)
	win32Map = make(map[string]Win32)
}

type Header struct {
	Includes  []string
	Variables []Variable
}

type Win32 struct {
	Includes  []string
	Variables []Variable
	Apis      []Api
}

type Api struct{}
