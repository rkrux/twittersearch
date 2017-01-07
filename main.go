package main

import (
	"fmt"
	"twittersearch/client"
	"twittersearch/credentials"
	"twittersearch/request"
)

func main() {

	// Load user and API credentials for authentication
	crdnls := credentials.New()
	if err := crdnls.Load(); err != nil {
		fmt.Println(err)
		return
	}

	// Get Twitter API client
	tClient := client.New()
	tClient.Load(*crdnls)

	// Create Twitter API request
	tReq := request.New()
	if err := tReq.Load(); err != nil {
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
