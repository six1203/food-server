package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("ToIZET+HX^a4Y$U3B%k{r{Yc=z4{%bi|", uint(1))
	spew.Dump(tk)
	panic(tk)
}
