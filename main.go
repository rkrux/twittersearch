package main

import (
	"fmt"

	"github.com/rushilkapoor/twittersearch/client"
	"github.com/rushilkapoor/twittersearch/credentials"
	"github.com/rushilkapoor/twittersearch/request"
)

func main() {

	// Load user and API credentials for authentication
	crdnls, err := credentials.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get Twitter API client
	tClient := client.New(*crdnls)

	// Create Twitter API request
	tReq, err := request.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send request to Twitter API
	tResp, err := tClient.SendRequest(tReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse Twitter API response
	tRespBody, err := tResp.ParseBody()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display filtered tweets
	tRespBody.FilterTweets()
}
