package sendgrid_client

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// Sendgrid  Template
type Template struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Sendgrid  Template Version
type TemplateVersion struct {
	Template     Template `json:"template,omitempty"`
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	HtmlContent  string   `json:"html_content,omitempty"`
	PlainContent string   `json:"plain_content,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	Active       bool     `json:"active,omitempty"`
}

///////////////////////////////////////////////////
// Create a transactional template.
// POST /templates

func (client *Client) CreateTemplate(template *Template) (*Template, error) {
	request := sendgrid.GetRequest(client.apiKey, "/v3/templates", "")
	request.Method = "POST"
	var err error
	request.Body, err = client.GetBody(template)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateResponse(response)
	}
}

///////////////////////////////////////////////////
// Get all transactional templates.
// GET /templates

func (client *Client) GetAllTemplates() (bool, error) {
	request := sendgrid.GetRequest(client.apiKey, "/v3/templates", "")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return false, err
	} else {
		err := processEmptyResponse(response)
		if err != nil {
			return false, nil
		} else {
			return true, nil
		}
	}
}

///////////////////////////////////////////////////
// Update a transactional template.
// PATCH /templates/{template_id}

func (client *Client) UpdateTemplate(id string, template *Template) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s", id), "")
	request.Method = "PATCH"
	var err error
	request.Body, err = client.GetBody(template)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Get a single transactional template.
// GET /templates/{template_id}

func (client *Client) GetTemplate(id string) (*Template, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s", id), "")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateResponse(response)
	}
}

///////////////////////////////////////////////////
// Delete a template.
// DELETE /templates/{template_id}

func (client *Client) DeleteTemplate(id string) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s", id), "")
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Create a new transactional template version.
// POST /templates/{template_id}/versions

func (client *Client) CreateTemplateVersion(version *TemplateVersion) (*TemplateVersion, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions", version.Template.Id), "")
	request.Method = "POST"
	var err error
	request.Body, err = client.GetBody(version)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateVersionResponse(response)
	}
}

///////////////////////////////////////////////////
// Update a transactional template version.
// PATCH /templates/{template_id}/versions/{version_id}

func (client *Client) UpdateTemplateVersion(version *TemplateVersion) error {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s", version.Template.Id, version.Id), "")
	request.Method = "PATCH"
	var err error
	request.Body, err = client.GetBody(version)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return processEmptyResponse(response)
	}
}

///////////////////////////////////////////////////
// Get a specific transactional template version.
// GET /templates/{template_id}/versions/{version_id}

func (client *Client) GetSpecificTemplateVersion(templateId, versionId string) (*TemplateVersion, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s", templateId, versionId), "")
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateVersionResponse(response)
	}
}

///////////////////////////////////////////////////
// Delete a transactional template version.
// DELETE /templates/{template_id}/versions/{version_id}

func (client *Client) DeleteTemplateVersion(templateId, versionId string) (*TemplateVersion, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s", templateId, versionId), "")
	request.Method = "DELETE"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateVersionResponse(response)
	}
}

///////////////////////////////////////////////////
// Activate a transactional template version.
// POST /templates/{template_id}/versions/{version_id}/activate

func (client *Client) ActivateTemplateVersion(templateId, versionId string) (*TemplateVersion, error) {
	request := sendgrid.GetRequest(client.apiKey, fmt.Sprintf("/v3/templates/%s/versions/%s/activate", templateId, versionId), "")
	request.Method = "POST"
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return processTemplateVersionResponse(response)
	}
}

func processTemplateResponse(response *rest.Response) (*Template, error) {
	err := processEmptyResponse(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var out Template
	err = json.Unmarshal([]byte(response.Body), &out)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &out, nil
}

func processTemplateVersionResponse(response *rest.Response) (*TemplateVersion, error) {
	err := processEmptyResponse(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var out TemplateVersion
	err = json.Unmarshal([]byte(response.Body), &out)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &out, nil
}

func processEmptyResponse(response *rest.Response) error {
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	if response.StatusCode >= 300 {
		return fmt.Errorf("Error calling API: status code: %d", response.StatusCode)
	}
	return nil
}
