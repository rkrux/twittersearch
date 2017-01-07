package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/rushilkapoor/twittersearch/constants"
)

// Credentials to be used in Twitter Client
type Credentials struct {
	ConsumerKey            string `json:"consumerKey"`
	ConsumerSecret         string `json:"consumerSecret"`
	OauthAccessToken       string `json:"oauthAccessToken"`
	OauthAccessTokenSecret string `json:"oauthAccessTokenSecret"`
}

// New create validated credentials struct
func New() (*Credentials, error) {
	crdnls := new(Credentials)
	if err := crdnls.read(); err != nil {
		return crdnls, err
	}
	if err := crdnls.validate(); err != nil {
		return crdnls, err
	}
	return crdnls, nil
}

// Read credentials from JSON file and load them in struct
func (crdnls *Credentials) read() error {
	var (
		fileData []byte
		err      error
	)
	if fileData, err = ioutil.ReadFile(constants.CredentialsFile); err != nil {
		return fmt.Errorf("%v: %v", constants.CredentialsReadError, err.Error())
	}
	if err = json.Unmarshal(fileData, crdnls); err != nil {
		return fmt.Errorf("%v: %v", constants.CredentialsReadError, err.Error())
	}
	return nil
}

// Validate user and API credentials
func (crdnls *Credentials) validate() error {
	if crdnls == nil || crdnls.ConsumerKey == "" || crdnls.ConsumerSecret == "" ||
		crdnls.OauthAccessToken == "" || crdnls.OauthAccessTokenSecret == "" {
		return fmt.Errorf("%v: %v", constants.CredentialsValidateError,
			constants.IncompleteCredentials)
	}
	return nil
}
