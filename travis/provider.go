package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
  "github.com/hashicorp/terraform/helper/validation"
)

// init provider block
func Provider() terraform.ResourceProvider {
  return &schema.Provider {
    Schema: map[string]*schema.Schema {
      "token": &schema.Schema {
        Type:         schema.TypeString,
        Required:     true,
        DefaultFunc:  schema.EnvDefaultFunc("TRAVIS_TOKEN", nil),
        // TODO: needs to be *regexp.Regexp and not string type
        ValidateFunc: validation.StringMatch(`^[a-zA-Z0-9]+$`, "The token argument value must conform to characters and integers."),
        Description:  "TravisCI API Token",
      },
      "commercial": &schema.Schema {
        Type:        schema.TypeBool,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("TRAVIS_COMMERCIAL", false),
        Description: "Whether to use the commercial or free version of TravisCI.",
      },
    },
    ResourcesMap: map[string]*schema.Resource {
      "build_job": buildJob(),
    },
    ConfigureFunc: configureProvider,
  }
}

// configure provider options; TODO: token validity check of some kind
func configureProvider(data *schema.ResourceData) (interface{}, error) {
  // store input options in opts struct
  opts := &travisOpts {
    token: data.Get("token").(string),
    commercial: data.Get("commercial").(bool),
  }

  // TODO: err handle

  return opts, nil
}
