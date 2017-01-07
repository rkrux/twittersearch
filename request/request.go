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

// TwitterRequest to be sent to API
type TwitterRequest struct {
	Request *http.Request
}

// New creates Twitter API request with query parameters
func New() (*TwitterRequest, error) {

	tReq := new(TwitterRequest)
	var (
		httpRequest *http.Request
		err         error
	)

	query := url.Values{}
	for key, value := range queryParams {
		query.Set(key, value)
	}
	apiURL := fmt.Sprintf("%v%v", constants.BaseRequestUrl, query.Encode())

	if httpRequest, err = http.NewRequest(constants.RequestMethod, apiURL, nil); err != nil {
		return tReq, fmt.Errorf("%v: %v", constants.CreateRequestError, err.Error())
	}
	tReq.Request = httpRequest
	return tReq, nil
}
