package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
)

// struct for passing opts from provider to resources
type travisOpts struct {
  token string
}

// init provider block
func Provider() terraform.ResourceProvider {
  return &schema.Provider {
    Schema: map[string]*schema.Schema {
      "token": &schema.Schema {
        Type:        schema.TypeString,
        Required:    true,
        DefaultFunc: schema.EnvDefaultFunc("TRAVIS_TOKEN", nil),
        Description: "TravisCI API Token",
      },
    },
    ResourcesMap: map[string]*schema.Resource {
      "build_job": buildJob(),
    },
    ConfigureFunc: configureProvider,
  }
}

// configure provider options; TODO: token validity check of some kind
func configureProvider(data *schema.ResourceData) (travisOpts, error) {
  opt := &travisOpts {
    token: data.Get("token").(string),
  }

  return opt, err
}
