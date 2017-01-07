package client

import (
	"fmt"
	"net/http"

	"github.com/rushilkapoor/twittersearch/constants"
	"github.com/rushilkapoor/twittersearch/credentials"
	"github.com/rushilkapoor/twittersearch/request"
	"github.com/rushilkapoor/twittersearch/response"

	"github.com/rushilkapoor/twittersearch/_libs/github.com/kurrik/oauth1a"
)

// TwitterClient to connect with API
type TwitterClient struct {
	oAuth      *oauth1a.Service
	user       *oauth1a.UserConfig
	httpClient *http.Client
}

// New loads user and API credentials in Twitter client using oauth1a library
func New(crdnls credentials.Credentials) *TwitterClient {
	tClient := new(TwitterClient)
	tClient.user = oauth1a.NewAuthorizedConfig(crdnls.OauthAccessToken,
		crdnls.OauthAccessTokenSecret)
	tClient.oAuth = &oauth1a.Service{
		RequestURL:   constants.RequestToken,
		AuthorizeURL: constants.Authorize,
		AccessURL:    constants.AccessToken,
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    crdnls.ConsumerKey,
			ConsumerSecret: crdnls.ConsumerSecret,
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}
	tClient.httpClient = new(http.Client)
	return tClient
}

// SendRequest sends oAuth signed request to Twitter API and return API response
func (tClient *TwitterClient) SendRequest(tReq *request.TwitterRequest) (
	response.TwitterResponse, error) {

	var (
		tResp        response.TwitterResponse
		httpResponse *http.Response
		err          error
	)

	tClient.oAuth.Sign(tReq.Request, tClient.user)
	if httpResponse, err = tClient.httpClient.Do(tReq.Request); err != nil {
		return tResp, fmt.Errorf("%v: %v", constants.SendRequestError, err.Error())
	}
	tResp.Response = httpResponse
	return tResp, nil
}
