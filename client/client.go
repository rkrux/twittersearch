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

func New() *TwitterClient {
	return new(TwitterClient)
}

// Load user and API credentials in Twitter client using oauth1a library
func (tClient *TwitterClient) Load(crdnls credentials.Credentials) {
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
}

// Sign and send request to Twitter API and return API response
func (tClient *TwitterClient) SendRequest(tReq *request.TwitterRequest) (
	response.TwitterResponse, error) {

	tClient.oAuth.Sign(tReq.Request, tClient.user)
	tResp := response.TwitterResponse{}
	if httpResponse, err := tClient.httpClient.Do(tReq.Request); err != nil {
		return tResp, fmt.Errorf("%v: %v", constants.SendRequestError, err.Error())
	} else {
		tResp.Response = httpResponse
		return tResp, nil
	}
}
