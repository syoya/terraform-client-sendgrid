package sendgrid_client

import (
	"encoding/json"
	"fmt"
)

// Client is the object that handles talking to the Datadog API. This maintains
// state information for a particular application connection.
type Client struct {
	apiKey string
}

// NewClient returns a new datadog.Client which can be used to access the API
// methods. The expected argument is the API key.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (client *Client) Validate() (bool, error) {
	return client.GetAllTemplates()
}

func (client *Client) GetBody(reqbody interface{}) ([]byte, error) {
	if reqbody != nil {
		bjson, err := json.Marshal(reqbody)
		if err != nil {
			return nil, err
		}
		return bjson, nil
	}
	return nil, fmt.Errorf("body structure is nil")
}
