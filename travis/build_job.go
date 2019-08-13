package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
  "fmt"
  "strings"
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
        Type:        schema.TypeString,
        Required:    true,
        DefaultFunc: schema.EnvDefaultFunc("TRAVIS_REPO", nil),
        Description: "Repository to trigger the build job for.",
      },
      "branch": &schema.Schema {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("TRAVIS_BRANCH", "master"),
        Description: "Branch of the repository for the build job.",
      },
    },
  }
}

// create a build job; TODO: validate repo string
func buildJobCreate(data *schema.ResourceData, meta interface{}) error {
  // convert repository value
  repository := strings.Replace(data.Get("repository").(string), "/", "%2F", 0)

  // construct endpoint
  endpoint := fmt.Sprintf("%s/requests", repository)

  // construct headers
  headers := map[string]string{}
  //, body='{"request": {"branch":"master"}}') TODO: and use branch from schema

  // receive response body
  body, err := apiClient("POST", endpoint, headers, "")

  // set resource id to response body
  data.SetId(body)

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
