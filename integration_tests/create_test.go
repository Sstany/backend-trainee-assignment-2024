package integrationtests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"banney/sdk"
	"banney/sdk/models"

	"gotest.tools/v3/assert"
)

func TestBannerCreate(t *testing.T) {
	bodyJSON := `{"tag_ids":[1,2],"feature_id":1,"content":{"foo":"bar"}, "is_active":true}`
	adminToken = NewAdminToken()
	body := strings.NewReader(bodyJSON)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8090/banner", body)
	assert.NilError(t, err)

	req.Header.Add(sdk.HeaderToken, adminToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	defer resp.Body.Close()

	var success models.BannerCreated

	err = json.NewDecoder(resp.Body).Decode(&success)
	assert.NilError(t, err)

	testDBClient.Cleanup()
}

func TestBannerCreateFailure(t *testing.T) {
	bodyJSON := `{"tag_ids":[1,2],"feature_id":"1","content":{"foo":"bar"}, "is_active":true}`

	body := strings.NewReader(bodyJSON)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8090/banner", body)
	assert.NilError(t, err)

	req.Header.Add(sdk.HeaderToken, adminToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	defer resp.Body.Close()

	var serverError models.ServerError

	err = json.NewDecoder(resp.Body).Decode(&serverError)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)

	testDBClient.Cleanup()
}

func TestBannerCreateAlreadyExists(t *testing.T) {
	bodyJSON := `{"tag_ids":[1,2],"feature_id":1,"content":{"foo":"bar"}, "is_active":true}`

	body := strings.NewReader(bodyJSON)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8090/banner", body)
	assert.NilError(t, err)

	req.Header.Add(sdk.HeaderToken, adminToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)

	defer resp.Body.Close()

	var success models.BannerCreated

	err = json.NewDecoder(resp.Body).Decode(&success)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, http.StatusCreated)

	body = strings.NewReader(bodyJSON)

	req, err = http.NewRequest(http.MethodPost, "http://localhost:8090/banner", body)
	assert.NilError(t, err)

	req.Header.Add(sdk.HeaderToken, adminToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	assert.NilError(t, err)

	defer resp.Body.Close()

	var serverError models.ServerError

	err = json.NewDecoder(resp.Body).Decode(&serverError)
	assert.NilError(t, err)

	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)

	testDBClient.Cleanup()

}
