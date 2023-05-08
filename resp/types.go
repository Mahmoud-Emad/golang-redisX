package resp

import (
	"errors"
)

const (
    SIMPLE_STRING = '+'
    BULK_STRING   = '$'
    INTEGER       = ':'
    ARRAY         = '*'
    ERROR         = '-'
)

var (
    ErrInvalidSyntax 		= errors.New("resp: invalid syntax")
	arrayPrefixSlice      	= []byte{'*'}
	bulkStringPrefixSlice 	= []byte{'$'}
	lineEndingSlice       	= []byte{'\r', '\n'}
) 