package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSplitCutset(t *testing.T) {
	{
		input := "aaa,bbb;    ccc；         \nddd"
		want := []string{"aaa", "bbb", "ccc", "ddd"}
		recv := SplitCutset(input, ",;；\n")
		assert.Equal(t, want, recv)
	}
}
