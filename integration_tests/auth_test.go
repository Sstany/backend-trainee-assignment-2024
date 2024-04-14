package integrationtests

import (
	"encoding/json"
	"net/http"
	"strings"

	"banney/sdk/models"
)

func NewAdminToken() string {
	bodyJSON := `{"login":"admin","password":"","is_admin":true}`

	body := strings.NewReader(bodyJSON)

	resp, err := http.Post("http://localhost:8090/auth/register", "application/json", body)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	var token models.Access
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return ""
	}

	return token.Token
}
