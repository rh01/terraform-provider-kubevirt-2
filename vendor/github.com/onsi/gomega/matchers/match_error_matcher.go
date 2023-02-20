package matchers

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
)

type MatchErrorMatcher struct {
	Expected interface{}
}

func (matcher *MatchErrorMatcher) Match(actual interface{}) (success bool, err error) {
	if isNil(actual) {
		return false, fmt.Errorf("Expected an error, got nil")
	}

	if !isError(actual) {
		return false, fmt.Errorf("Expected an error.  Got:\n%s", format.Object(actual, 1))
	}

	actualErr := actual.(error)
<<<<<<< HEAD
	expected := matcher.Expected

	if isError(expected) {
		return reflect.DeepEqual(actualErr, expected) || errors.Is(actualErr, expected.(error)), nil
	}

	if isString(expected) {
		return actualErr.Error() == expected, nil
=======

	if isError(matcher.Expected) {
		return reflect.DeepEqual(actualErr, matcher.Expected), nil
	}

	if isString(matcher.Expected) {
		return actualErr.Error() == matcher.Expected, nil
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
	}

	var subMatcher omegaMatcher
	var hasSubMatcher bool
<<<<<<< HEAD
	if expected != nil {
		subMatcher, hasSubMatcher = (expected).(omegaMatcher)
=======
	if matcher.Expected != nil {
		subMatcher, hasSubMatcher = (matcher.Expected).(omegaMatcher)
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
		if hasSubMatcher {
			return subMatcher.Match(actualErr.Error())
		}
	}

<<<<<<< HEAD
	return false, fmt.Errorf(
		"MatchError must be passed an error, a string, or a Matcher that can match on strings. Got:\n%s",
		format.Object(expected, 1))
=======
	return false, fmt.Errorf("MatchError must be passed an error, string, or Matcher that can match on strings.  Got:\n%s", format.Object(matcher.Expected, 1))
>>>>>>> 0faf8ce (Revert "Upgrade go mod and dependencies")
}

func (matcher *MatchErrorMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to match error", matcher.Expected)
}

func (matcher *MatchErrorMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to match error", matcher.Expected)
}
