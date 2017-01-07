package constants

const (
	// FILES
	CREDENTIALS_FILE = "file.json"

	// ERRORS
	CREDENTIALS_READ_ERROR     = "Error reading credentials"
	CREDENTIALS_VALIDATE_ERROR = "Error validating credentials"
	INCOMPLETE_CREDENTIALS     = "Credentials are incomplete"
	CREATE_REQUEST_ERROR       = "Error forming HTTP request"
	SEND_REQUEST_ERROR         = "Error sending HTTP request"
	PARSE_RESPONSE_ERROR       = "Error parsing response body"

	// REQUEST PARAMS
	BASE_URL         = "https://api.twitter.com/oauth"
	REQUEST_TOKEN    = BASE_URL + "/request_token"
	AUTHORIZE        = BASE_URL + "/authorize"
	ACCESS_TOKEN     = BASE_URL + "/access_token"
	BASE_REQUEST_URL = "https://api.twitter.com/1.1/search/tweets.json?"
	REQUEST_METHOD   = "GET"

	// QUERY PARAMS
	SEARCH_KEY   = "q"
	SEARCH_VALUE = "#custserv"
	COUNT_KEY    = "count"
	COUNT_VALUE  = "100"

	// RESPONSE STATUS
	STATUS_OK = 200
)
