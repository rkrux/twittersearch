package constants

const (
	// FILES
	CredentialsFile = "credentials.json"

	// ERRORS
	CredentialsReadError     = "Error reading credentials"
	CredentialsValidateError = "Error validating credentials"
	IncompleteCredentials    = "Credentials are incomplete"
	CreateRequestError       = "Error forming HTTP request"
	SendRequestError         = "Error sending HTTP request"
	ParseResponseError       = "Error parsing response body"

	// REQUEST PARAMS
	BaseUrl        = "https://api.twitter.com/oauth"
	RequestToken   = BaseUrl + "/request_token"
	Authorize      = BaseUrl + "/authorize"
	AccessToken    = BaseUrl + "/access_token"
	BaseRequestUrl = "https://api.twitter.com/1.1/search/tweets.json?"
	RequestMethod  = "GET"

	// QUERY PARAMS
	SearchKey   = "q"
	SearchValue = "#custserv"
	CountKey    = "count"
	CountValue  = "100"

	// RESPONSE STATUS
	StatusOk = 200
)
