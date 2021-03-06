package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequest(t *testing.T) {
	request := CreateRepoRequest{
		Name:            "create repo request test",
		Description:     "create repo request test description",
		Private:         false,
		LicenseTemplate: "mit",
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"create repo request test","description":"create repo request test description","private":false,"license_template":"mit"}`, string(bytes))

	var target CreateRepoRequest

	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Private, request.Private)
	assert.EqualValues(t, target.LicenseTemplate, request.LicenseTemplate)
}
