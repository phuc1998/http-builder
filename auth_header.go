package builder

import "encoding/base64"

type authorizationType string

func (h authorizationType) String() string {
	return string(h)
}

var (
	// BasicAuthHeader takes BasicAuth as authentication for the request.
	BasicAuthHeader = authorizationType("Basic")

	// BearerHeader takes a string Bearer á authentication for the request.
	BearerHeader = authorizationType("Bearer")

	// APIKeyHeader takes an APIKey as authentication for the request
	APIKeyHeader = authorizationType("X-API-Key")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key   string
	Value string
}

var (
	headerFormatString = "%v %v"
	AuthorizationHeader = "Authorization"
)

// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
