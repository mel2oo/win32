package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var (
		xmlfile = flag.String("xml", "API/Headers/common.h.xml", "xml file with api/header description")
		xmltype = flag.String("type", "header", "api/header")
	)
	flag.Parse()

	file, err := os.Open(*xmlfile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	switch *xmltype {
	case "api":
		v := new(ApiXml)
		if err = xml.Unmarshal(data, v); err != nil {
			panic(err)
		}

	case "header":
		v := new(HeaderXml)
		if err = xml.Unmarshal(data, v); err != nil {
			panic(err)
		}

		fmt.Println()
	}

}
