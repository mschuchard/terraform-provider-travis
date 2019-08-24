package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/helper/validation"
  "fmt"
  "strings"
  "encoding/json"
)

// base reesource declaration and schema
func buildJob() *schema.Resource {
  return &schema.Resource {
    Create: buildJobCreate,
    Read:   buildJobRead,
    Update: buildJobUpdate,
    Delete: buildJobDelete,
    Exists: buildJobExists,

    Schema: map[string]*schema.Schema {
      "repository": &schema.Schema {
        Type:         schema.TypeString,
        Required:     true,
        DefaultFunc:  schema.EnvDefaultFunc("TRAVIS_REPO", nil),
        // TODO: needs to be *regexp.Regexp and not string type
        ValidateFunc: validation.StringMatch(`^[a-zA-Z0-9]+/[a-zA-Z0-9]$`, "The repository argument value must be of the form organization_name/repository_name."),
        Description:  "Repository to trigger the build job for.",
      },
      "branch": &schema.Schema {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("TRAVIS_BRANCH", "master"),
        Description: "Branch of the repository for the build job.",
      },
      "message": &schema.Schema {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("TRAVIS_MESSAGE", nil),
        Description: "The commit message for the build job.",
      },
    },
  }
}

// create a build job
func buildJobCreate(data *schema.ResourceData, meta interface{}) error {
  // convert repository value
  repository := strings.Replace(data.Get("repository").(string), "/", "%2F", 0)

  // construct endpoint
  endpoint := fmt.Sprintf("%s/requests", repository)

  // construct headers
  headers := map[string]string{}
  // construct request body; TODO: custom message; nested map not allowed for type declare in go
  requestMap := map[string]string{"request": {"branch": data.Get("branch").(string)}}
  requestBody, err := json.Marshal(requestMap)

  // receive response body
  responseBody, err := apiClient("POST", endpoint, headers, requestBody)

  // set resource id to response body
  data.SetId(responseBody)

  // TODO: error handle
  if err != nil {
    // stuff
  }

  return buildJobRead(data, meta)
}

// read a build job
func buildJobRead(data *schema.ResourceData, meta interface{}) error {
  return nil
}

// update a build job
func buildJobUpdate(data *schema.ResourceData, meta interface{}) error {
  return buildJobRead(data, meta)
}

// delete a build job
func buildJobDelete(data *schema.ResourceData, meta interface{}) error {
  return nil
}

// check for build job existence
func buildJobExists(data *schema.ResourceData, meta interface{}) (exists bool, err error) {
  return true, nil
}
