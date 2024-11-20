package vspcapi

import (
	"encoding/base64"

	"github.com/simonbuckner/goquadac"
)

type VspcApi struct {
	*goquadac.ApiHelper
}

func NewVspcApi(baseUrl string) *VspcApi {
	api := goquadac.NewApiHelper(baseUrl).SetDefaultHeader("accept", "application/json")
	return &VspcApi{
		ApiHelper: api,
	}
}

func (api *VspcApi) Authenticate(apiKey, apiSecret string) error {
	api.SetAuthHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":"+apiSecret)))
	return nil
}
