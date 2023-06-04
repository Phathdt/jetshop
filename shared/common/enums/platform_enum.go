// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.6
// Revision: 97611fddaa414f53713597918c5e954646cb8623
// Build Date: 2023-03-26T21:38:06Z
// Built By: goreleaser

package enums

import (
	"errors"
	"fmt"
)

const (
	// PlatformCodeLazada is a PlatformCode of type lazada.
	PlatformCodeLazada PlatformCode = "lazada"
	// PlatformCodeShopee is a PlatformCode of type shopee.
	PlatformCodeShopee PlatformCode = "shopee"
	// PlatformCodeFacebook is a PlatformCode of type facebook.
	PlatformCodeFacebook PlatformCode = "facebook"
)

var ErrInvalidPlatformCode = errors.New("not a valid PlatformCode")

// String implements the Stringer interface.
func (x PlatformCode) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PlatformCode) IsValid() bool {
	_, err := ParsePlatformCode(string(x))
	return err == nil
}

var _PlatformCodeValue = map[string]PlatformCode{
	"lazada":   PlatformCodeLazada,
	"shopee":   PlatformCodeShopee,
	"facebook": PlatformCodeFacebook,
}

// ParsePlatformCode attempts to convert a string to a PlatformCode.
func ParsePlatformCode(name string) (PlatformCode, error) {
	if x, ok := _PlatformCodeValue[name]; ok {
		return x, nil
	}
	return PlatformCode(""), fmt.Errorf("%s is %w", name, ErrInvalidPlatformCode)
}