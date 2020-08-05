package github_provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	headerAuthorizationFormat := getAuthorizationHeader("123")
	assert.EqualValues(t, "token 123", headerAuthorizationFormat)
}

func TestCreateRepoErrorRestClient(t *testing.T) {

}
