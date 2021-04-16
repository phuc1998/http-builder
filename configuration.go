package builder

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)



// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	Url         string
	Description string
	Variables   map[string]ServerVariable
}

// Configuration stores the configuration of the API client
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	Servers       []ServerConfiguration
	HTTPClient    *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "http://localhost",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: []ServerConfiguration{
			{
				Url:         "http://localhost",
				Description: "No description provided",
			},
		},
		HTTPClient: http.DefaultClient,
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP authorizationType to the default authorizationType in the request
func (c *Configuration) AddDefaultHeader(key string, value string) *Configuration {
	c.DefaultHeader[key] = value
	return c
}

// AddHTTPClient adds a custom HTTP client into builder
func (c *Configuration) AddHTTPClient(client *http.Client) *Configuration{
	if client != nil {
		c.HTTPClient = client
	}
	return c
}

// AddBasePath adds a base path into builder
func (c *Configuration) AddBasePath(basePath string) *Configuration{
	if len(basePath) > 0 {
		if _, err := url.Parse(basePath); err == nil {
			c.BasePath = basePath
		}
	}
	return c
}

// ServerUrl returns URL based on server settings
func (c *Configuration) ServerUrl(index int, variables map[string]string) (string, error) {
	if index < 0 || len(c.Servers) <= index {
		return "", fmt.Errorf("Index %v out of range %v", index, len(c.Servers)-1)
	}
	server := c.Servers[index]
	url := server.Url

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("The variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}
