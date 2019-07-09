package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
  "net/http"
  "fmt"
  "strings"
  "io/ioutil"
)

// base reesource declaration and schema
func buildJob() *schema.Resource {
  return &schema.Resource {
    Create: buildJobCreate,
    Read:   buildJobRead,
    Update: buildJobUpdate,
    Delete: buildJobDelete,

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
      "commercial": &schema.Schema {
        Type:        schema.TypeBool,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("TRAVIS_BRANCH", false),
        Description: "Whether to use the commercial or free version of TravisCI.",
      },
    },
  }
}

// create a build job; TODO: .com for commercial, validate repo string
func buildJobCreate(data *schema.ResourceData, m interface{}) error {
  // convert repository value
  repository := strings.Replace(data.Get("repository").(string), "/", "%2F", 0)
  // construct endpoint
  endpoint := fmt.Sprintf("https://api.travis-ci.org/repo/%s/requests", repository)
  // access travis endpoint to initiate build job
  client := &http.Client{}
  request, err := client.NewRequest("POST", endpoint, nil)//, body='{"request": {"branch":"master"}}') TODO: and use branch from schema
  client.Header.Add("Content-Type", "application/json")
  client.Header.Add("Accept", "application/json")
  client.Header.Add("Travis-API-Version", "3")
  client.Header.Add("Authorization", travisOpts.token)
  response, err := client.Do(request)
  // TODO: err handle
  // read body of response
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  // set resource id to response body
  data.SetId(body)

  return buildJobRead(data, m)
}

// read a build job
func buildJobRead(data *schema.ResourceData, m interface{}) error {
  return nil
}

// update a build job
func buildJobUpdate(data *schema.ResourceData, m interface{}) error {
  return buildJobRead(data, m)
}

// delete a build job
func buildJobDelete(data *schema.ResourceData, m interface{}) error {
  return nil
}
