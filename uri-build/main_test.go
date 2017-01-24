package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuildQuery(t *testing.T) {
	gq := new(googleQuery)
	gq.Oe = "a"
	gq.Ie = "b"
	gq.GwsRd = "c"
	gq.Query = "d"

	actual := BuildQuery(gq)

	assert.Equal(t, "gws_rd=c&ie=b&oe=a&q=d", actual)
}
