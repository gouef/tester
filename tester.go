package tester

import "github.com/stretchr/testify/assert"

var t assert.TestingT

func Empty(object interface{}, msgAndArgs ...interface{}) bool {
	return assert.Empty(t, object, msgAndArgs)
}
