package tester

import "github.com/stretchr/testify/assert"

func Empty(object any, msgAndArgs ...any) bool {
	return assert.Empty(TestingT, object, msgAndArgs)
}

func NotEmpty(object any, msgAndArgs ...any) bool {
	return assert.NotEmpty(TestingT, object, msgAndArgs)
}

func Equal(expected, actual any, msgAndArgs ...any) bool {
	return assert.Equal(TestingT, expected, actual, msgAndArgs)
}

func NotEqual(expected, actual any, msgAndArgs ...any) bool {
	return assert.NotEqual(TestingT, expected, actual, msgAndArgs)
}

func Nil(object any, msgAndArgs ...any) bool {
	return assert.Nil(TestingT, object, msgAndArgs)
}

func NotNil(object any, msgAndArgs ...any) bool {
	return assert.NotNil(TestingT, object, msgAndArgs)
}

func True(value bool, msgAndArgs ...any) bool {
	return assert.True(TestingT, value, msgAndArgs)
}

func False(value bool, msgAndArgs ...any) bool {
	return assert.False(TestingT, value, msgAndArgs)
}

func Error(err error, msgAndArgs ...any) bool {
	return assert.Error(TestingT, err, msgAndArgs...)
}

func NoError(err error, msgAndArgs ...any) bool {
	return assert.NoError(TestingT, err, msgAndArgs...)
}
