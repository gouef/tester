package tester

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var TestingT assert.TestingT

type Tester struct {
	T *testing.T
}

func (t *Tester) FailNow(failureMessage string, msgAndArgs ...interface{}) bool {
	return assert.FailNowf()
}

func (t *Tester) Fail(failureMessage string, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Empty(object any, msgAndArgs ...any) bool {
	return assert.Empty(t.T, object, msgAndArgs)
}

func (t *Tester) NotEmpty(object any, msgAndArgs ...any) bool {
	return assert.NotEmpty(t.T, object, msgAndArgs)
}

func (t *Tester) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotImplements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Equal(expected, actual any, msgAndArgs ...any) bool {
	return assert.Equal(t.T, expected, actual, msgAndArgs)
}

func (t *Tester) NotEqual(expected, actual any, msgAndArgs ...any) bool {
	return assert.NotEqual(t.T, expected, actual, msgAndArgs)
}

func (t *Tester) Same(expected, actual interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotSame(expected, actual interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) EqualValues(expected, actual interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) EqualExportedValues(expected, actual interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Exactly(expected, actual interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Nil(object any, msgAndArgs ...any) bool {
	return assert.Nil(t.T, object, msgAndArgs)
}

func (t *Tester) NotNil(object any, msgAndArgs ...any) bool {
	return assert.NotNil(t.T, object, msgAndArgs)
}

func (t *Tester) Len(object interface{}, length int, msgAndArgs ...interface{}) bool {

}

func (t *Tester) True(value bool, msgAndArgs ...any) bool {
	return assert.True(t.T, value, msgAndArgs)
}

func (t *Tester) False(value bool, msgAndArgs ...any) bool {
	return assert.False(t.T, value, msgAndArgs)
}

func (t *Tester) NotEqualValues(expected, actual interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Contains(s, contains interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotContains(s, contains interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Subset(list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {

}

func (t *Tester) NotSubset(list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {

}

func (t *Tester) ElementsMatch(listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {

}

func (t *Tester) NotElementsMatch(listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {

}

func (t *Tester) Condition(comp Comparison, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Panics(f PanicTestFunc, msgAndArgs ...interface{}) bool {

}

func (t *Tester) PanicsWithValue(expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) bool {

}

func (t *Tester) PanicsWithError(errString string, f PanicTestFunc, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) bool {

}

func (t *Tester) WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {

}

func (t *Tester) WithinRange(actual, start, end time.Time, msgAndArgs ...interface{}) bool {

}

func (t *Tester) InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {

}

func (t *Tester) InDeltaSlice(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {

}

func (t *Tester) InDeltaMapValues(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {

}

func (t *Tester) InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {

}

func (t *Tester) InEpsilonSlice(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {

}

// Errors

func (t *Tester) Error(err error, msgAndArgs ...any) bool {
	return assert.Error(t.T, err, msgAndArgs...)
}

func (t *Tester) NoError(err error, msgAndArgs ...any) bool {
	return assert.NoError(t.T, err, msgAndArgs...)
}

func (t *Tester) EqualError(theError error, errString string, msgAndArgs ...interface{}) bool {
	return assert.EqualError(t.T, theError, errString, msgAndArgs)
}

func (t *Tester) ErrorContains(theError error, contains string, msgAndArgs ...interface{}) bool {
	return assert.ErrorContains(t.T, theError, contains, msgAndArgs)
}

func (t *Tester) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Zero(i interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotZero(i interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) FileExists(path string, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NoFileExists(path string, msgAndArgs ...interface{}) bool {

}

func (t *Tester) DirExists(path string, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NoDirExists(path string, msgAndArgs ...interface{}) bool {

}

func (t *Tester) JSONEq(expected string, actual string, msgAndArgs ...interface{}) bool {

}

func (t *Tester) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {

}

func (t *Tester) EventuallyWithT(condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {

}

func (t *Tester) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {

}

func (t *Tester) ErrorIs(err, target error, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotErrorIs(err, target error, msgAndArgs ...interface{}) bool {

}

func (t *Tester) ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) bool {

}

func (t *Tester) NotErrorAs(err error, target interface{}, msgAndArgs ...interface{}) bool {

}
