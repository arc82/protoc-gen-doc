package protoc_gen_doc_test

import (
	"github.com/arc82/protoc-gen-doc/parser"
	"github.com/arc82/protoc-gen-doc/test"
	"testing"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	codeGenRequest, _ := test.MakeCodeGeneratorRequest()

	for i := 0; i < b.N; i++ {
		parser.ParseCodeRequest(codeGenRequest)
	}
}
