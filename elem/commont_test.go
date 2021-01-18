package elem

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandSlice(t *testing.T){
	elements := GenerateRandSlice(5,3)
	assert.True(t, len(elements)==3)
}
