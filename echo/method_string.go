// generated by stringer -type=Method method.go; DO NOT EDIT

package main

import "fmt"

const _Method_name = "GETPOST"

var _Method_index = [...]uint8{0, 3, 7}

func (i Method) String() string {
	if i < 0 || i >= Method(len(_Method_index)-1) {
		return fmt.Sprintf("Method(%d)", i)
	}
	return _Method_name[_Method_index[i]:_Method_index[i+1]]
}
