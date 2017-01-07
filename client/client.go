package client

import (
	"fmt"
	"net/http"
	"twittersearch/constants"
	"twittersearch/credentials"
	"twittersearch/request"
	"twittersearch/response"

	"twittersearch/_libs/github.com/kurrik/oauth1a"
)

type TwitterClient struct {
	OAuth      *oauth1a.Service
	User       *oauth1a.UserConfig
	HTTPClient *http.Client
}

func New() *TwitterClient {
	return new(TwitterClient)
}

// Load user and API credentials in Twitter client using oauth1a library
func (tClient *TwitterClient) Load(crdnls credentials.Credentials) {
	tClient.User = oauth1a.NewAuthorizedConfig(crdnls.OauthAccessToken,
		crdnls.OauthAccessTokenSecret)
	tClient.OAuth = &oauth1a.Service{
		RequestURL:   constants.REQUEST_TOKEN,
		AuthorizeURL: constants.AUTHORIZE,
		AccessURL:    constants.ACCESS_TOKEN,
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:    crdnls.ConsumerKey,
			ConsumerSecret: crdnls.ConsumerSecret,
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}
	tClient.HTTPClient = new(http.Client)
}

// Sign and send request to Twitter API and return API response
func (tClient *TwitterClient) SendRequest(tReq *request.TwitterRequest) (
	response.TwitterResponse, error) {

	tClient.OAuth.Sign(tReq.Request, tClient.User)
	tResp := response.TwitterResponse{}
	if httpResponse, err := tClient.HTTPClient.Do(tReq.Request); err != nil {
		return tResp, fmt.Errorf("%v: %v", constants.SEND_REQUEST_ERROR, err.Error())
	} else {
		tResp.Response = httpResponse
		return tResp, nil
	}
}
