package protobufs

type Test1 struct {
	Name   *string
	Nested *Test1Nested
}

type Test1Nested struct {
	F1 int
	F2 float32
	F3 []*Test1Repeated
}

type Test1Repeated struct {
	Age int
}
