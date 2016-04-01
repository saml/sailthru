package sailthru

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// BaseURL is Sailthru api url
	BaseURL = "https://api.sailthru.com"

	// ListURL is list api endpoint
	ListURL = BaseURL + "/list"

	// UserURL is user api endpiont
	UserURL = BaseURL + "/user"
)

var (
	// DefaultFormat is response format. Can be xml or json.
	DefaultFormat = "json"
	client        = &http.Client{}
)

// Client is a Sailthru client.
type Client struct {
	Name   string
	Key    string
	Secret string
	Lists  map[string]*List
}

// Data is any data sent to sailthru
type Data map[string]interface{}

// Init gives this client a name and fetches valid list names for it.
func (c *Client) Init(clientName string) error {
	c.Name = clientName

	lists, err := c.FetchLists()
	if err != nil {
		return err
	}

	m := make(map[string]*List)
	listNames := make([]string, 0, len(lists))
	for _, list := range lists {
		m[list.Name] = list
		listNames = append(listNames, list.Name)
	}

	c.Lists = m
	return nil
}

// Form form encodes data.
func (c *Client) Form(data Data) url.Values {
	q := url.Values{}
	q.Set("api_key", c.Key)
	q.Set("format", DefaultFormat)

	var s string
	if len(data) > 0 {
		b, err := json.Marshal(data)
		if err == nil {
			s = string(b)
			q.Set("json", s)
		}
	}

	q.Set("sig", c.Checksum(s))
	return q
}

// Get makes HTTP GET request to endpoint with data as query string.
func (c *Client) Get(endpoint string, data Data) (*http.Response, error) {
	q := c.Form(data)
	return client.Get(endpoint + "?" + q.Encode())
}

// Post makes HTTP POST (form post) request to endpoint with given data in the body.
func (c *Client) Post(endpoint string, data Data) (*http.Response, error) {
	q := c.Form(data)
	return client.PostForm(endpoint, q)
}
