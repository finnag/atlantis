// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v3"
	"reflect"

	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsPullRequest() models.PullRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullRequest))(nil)).Elem()))
	var nullValue models.PullRequest
	return nullValue
}

func EqModelsPullRequest(value models.PullRequest) models.PullRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullRequest
	return nullValue
}

func NotEqModelsPullRequest(value models.PullRequest) models.PullRequest {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.PullRequest
	return nullValue
}

func ModelsPullRequestThat(matcher pegomock.ArgumentMatcher) models.PullRequest {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.PullRequest
	return nullValue
}
