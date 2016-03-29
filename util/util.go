package util

import (
	"regexp"
)

type StrArr []string

var IdentifierRe = regexp.MustCompile(`^[_a-zA-Z][_0-9a-zA-Z]*$`)
var AddressRe = regexp.MustCompile(`^(0[xX])?[0-9A-Fa-f]{40}$`)

func NewStrArr() *StrArr {
	return &StrArr{}
}

func (strArr *StrArr) Add(str string) {
	*strArr = append(*strArr, str)
}
