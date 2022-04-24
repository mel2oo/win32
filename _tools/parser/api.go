package main

import "github.com/mel2oo/win32/_tools/parser/api"

type Api struct {
	Name        string
	BothCharset bool
	Params      []Param
	Return      Variable
}

type Param struct {
	Name string
	Type Variable
}

func GetApi(ins map[string]*Header, vbs []Variable, v api.Api) *Api {
	ps := make([]Param, 0)
	for _, p := range v.Param {
		ps = append(ps, Param{
			Name: p.Name,
			Type: GetPriVariable(ins, vbs, p.Type),
		})
	}

	return &Api{
		Name:        v.Name,
		BothCharset: v.BothCharset,
		Params:      ps,
		Return:      GetPriVariable(ins, vbs, v.Return.Type),
	}
}

func GetPriVariable(ins map[string]*Header, vbs []Variable, name string) Variable {
	for _, v := range vbs {
		if v.Name() == name {
			return v
		}
	}

	return GetIncVariable(ins, name)
}

func GetIncVariable(ins map[string]*Header, name string) Variable {
	for _, v := range ins {
		if v == nil {
			return nil
		}

		for _, vv := range v.Variables {
			if vv.Name() == name {
				return vv
			}
		}

		vv := GetIncVariable(v.Includes, name)
		if vv != nil {
			return vv
		}
	}
	return nil
}
