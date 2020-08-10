package test_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMockContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080", nil)
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	request.Header.Add("X-Mock", "true")
	c := GetMockContext(request, response)

	assert.EqualValues(t, http.MethodPost, c.Request.Method)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.Nil(t, nil, request.Body)

	assert.EqualValues(t, "127.0.0.1:8080", c.Request.Host)
	assert.EqualValues(t, "8080", c.Request.URL.Port())
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.Request.Header.Get("X-Mock"))
	assert.EqualValues(t, "true", c.Request.Header.Get("x-mock"))
}
