package test_tuils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func GetMockContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}
