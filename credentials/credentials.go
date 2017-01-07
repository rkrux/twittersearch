package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"twittersearch/constants"
)

type Credentials struct {
	ConsumerKey            string `json:"consumerKey"`
	ConsumerSecret         string `json:"consumerSecret"`
	OauthAccessToken       string `json:"oauthAccessToken"`
	OauthAccessTokenSecret string `json:"oauthAccessTokenSecret"`
}

func New() *Credentials {
	return new(Credentials)
}

func (crdnls *Credentials) Load() error {
	if err := crdnls.Read(); err != nil {
		return err
	}
	if err := crdnls.Validate(); err != nil {
		return err
	}
	return nil
}

// Read credentials from JSON file and load them in struct
func (crdnls *Credentials) Read() error {
	var (
		fileData []byte
		err      error
	)
	if fileData, err = ioutil.ReadFile(constants.CREDENTIALS_FILE); err != nil {
		return fmt.Errorf("%v: %v", constants.CREDENTIALS_READ_ERROR, err.Error())
	}
	if err = json.Unmarshal(fileData, crdnls); err != nil {
		return fmt.Errorf("%v: %v", constants.CREDENTIALS_READ_ERROR, err.Error())
	}
	return nil
}

// Validate user and API credentials
func (crdnls *Credentials) Validate() error {
	if crdnls == nil || crdnls.ConsumerKey == "" || crdnls.ConsumerSecret == "" ||
		crdnls.OauthAccessToken == "" || crdnls.OauthAccessTokenSecret == "" {
		return fmt.Errorf("%v: %v", constants.CREDENTIALS_VALIDATE_ERROR,
			constants.INCOMPLETE_CREDENTIALS)
	}
	return nil
}
