// Code generated by "stringer -type=Element"; DO NOT EDIT

package main

import "fmt"

const _Element_name = "PromethiumCobaltCuriumRutheniumPlutoniumHydrogenLithium"

var _Element_index = [...]uint8{0, 10, 16, 22, 31, 40, 48, 55}

func (i Element) String() string {
	if i >= Element(len(_Element_index)-1) {
		return fmt.Sprintf("Element(%d)", i)
	}
	return _Element_name[_Element_index[i]:_Element_index[i+1]]
}