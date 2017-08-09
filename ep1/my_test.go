// Package main provides ...
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeThing(t *testing.T) {
	assert.Equal(t, 1, 1, "1 과 1은 반드시 같습니다.")
	assert.Equal(t, 2, 2, "2 와 1은 같지 않습니다.")
}

func TestTheOther(t *testing.T) {
	assert.Equal(t, 1, 100, "의도된 Fail")
}

func TestTheOtherTwo(t *testing.T) {
	assert.Equal(t, 3, 100, "의도된 Fail")
}
