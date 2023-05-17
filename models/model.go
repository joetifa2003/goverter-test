//go:generate go run github.com/jmattheis/goverter/cmd/goverter -packageName converter -output ../converter/generated.go .
package models

import (
	"test/protobufs"
)

// goverter:converter
// goverter:useZeroValueOnPointerInconsistency
type Converter interface {
	FromProto(*protobufs.Test1) Test1
	FromProtoMany([]*protobufs.Test1) []Test1
	FromProtoMap(map[string]*protobufs.Test1) map[string]Test1
}

type Test1 struct {
	Name   string
	Nested Test1Nested
}

type Test1Nested struct {
	F1 int
	F2 float32
	F3 []Test1Repeated
}

type Test1Repeated struct {
	Age int
}
