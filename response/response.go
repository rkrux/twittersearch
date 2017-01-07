package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rushilkapoor/twittersearch/constants"
)

type twitterUser struct {
	Name          string `json:"name"`
	ScreenName    string `json:"screen_name"`
	Location      string `json:"location"`
	FollwersCount int    `json:"followers_count"`
}

type tweet struct {
	Text         string      `json:"text"`
	RetweetCount int         `json:"retweet_count"`
	CreatedAt    string      `json:"created_at"`
	User         twitterUser `json:"user"`
}

type TwitterResponseBody struct {
	ResponseBody []tweet `json:"statuses"`
}

type TwitterResponse struct {
	Response *http.Response
}

// Parse Twitter API response body to retrieve array of tweets
func (tResp TwitterResponse) ParseBody() (TwitterResponseBody, error) {
	defer tResp.Response.Body.Close()
	respBody := TwitterResponseBody{}
	if buffer, err := ioutil.ReadAll(tResp.Response.Body); err == nil {
		switch tResp.Response.StatusCode {
		case constants.StatusOk:
			if err = json.Unmarshal(buffer, &respBody); err != nil {
				return respBody,
					fmt.Errorf("%v: %v", constants.ParseResponseError, err.Error())
			}
			return respBody, nil
		default:
			return respBody,
				fmt.Errorf("%v: Status Code %v",
					constants.ParseResponseError, tResp.Response.StatusCode)
		}
	} else {
		return respBody, fmt.Errorf("%v: %v", constants.ParseResponseError, err.Error())
	}
}

// Display tweet details based on re-tweet condition
func (tRespBody TwitterResponseBody) FilterTweets() {
	index := 0
	for _, tweet := range tRespBody.ResponseBody {
		if tweet.RetweetCount < 1 {
			continue
		}
		index++
		fmt.Printf("Tweet %v) %v\n", index, tweet.Text)
		fmt.Printf(" - by %v (@%v) with %v followers, ", tweet.User.Name,
			tweet.User.ScreenName, tweet.User.FollwersCount)
		if tweet.User.Location != "" {
			fmt.Printf("from %v ", tweet.User.Location)
		}
		fmt.Printf("at %v\n\n", tweet.CreatedAt)
	}
}
