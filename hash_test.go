package tokenizer

import (
	"fmt"
	"testing"

	"github.com/FrogoAI/testutils"
)

func TestMH(t *testing.T) {
	v, err := Normalize("test.èmcop")
	testutils.Equal(t, err, nil)
	fmt.Println(v)

	v2, err := Normalize("testèmCAP")
	testutils.Equal(t, err, nil)
	fmt.Println(v2)
}
