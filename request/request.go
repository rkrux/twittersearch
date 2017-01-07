package request

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/rushilkapoor/twittersearch/constants"
)

// Query parameters used in request
var queryParams = map[string]string{
	constants.SearchKey: constants.SearchValue,
	constants.CountKey:  constants.CountValue,
}

type TwitterRequest struct {
	Request *http.Request
}

func New() *TwitterRequest {
	return new(TwitterRequest)
}

// Creates Twitter API request with query parameters
func (tReq *TwitterRequest) Load() error {
	query := url.Values{}
	for key, value := range queryParams {
		query.Set(key, value)
	}
	apiURL := fmt.Sprintf("%v%v", constants.BaseRequestUrl, query.Encode())

	if httpRequest, err := http.NewRequest(constants.RequestMethod, apiURL, nil); err != nil {
		return fmt.Errorf("%v: %v", constants.CreateRequestError, err.Error())
	} else {
		tReq.Request = httpRequest
	}
	return nil
}
