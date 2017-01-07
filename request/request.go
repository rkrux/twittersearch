package request

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/rushilkapoor/twittersearch/constants"
)

// Query parameters used in request
var queryParams = map[string]string{
	constants.SEARCH_KEY: constants.SEARCH_VALUE,
	constants.COUNT_KEY:  constants.COUNT_VALUE,
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
	apiUrl := fmt.Sprintf("%v%v", constants.BASE_REQUEST_URL, query.Encode())

	if httpRequest, err := http.NewRequest(constants.REQUEST_METHOD, apiUrl, nil); err != nil {
		return fmt.Errorf("%v: %v", constants.CREATE_REQUEST_ERROR, err.Error())
	} else {
		tReq.Request = httpRequest
	}
	return nil
}
