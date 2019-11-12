package travis

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
  "github.com/hashicorp/terraform/helper/validation"

  "regexp"
)

// init provider block
func Provider() terraform.ResourceProvider {
  return &schema.Provider {
    Schema: map[string]*schema.Schema {
      "token": &schema.Schema {
        Type:         schema.TypeString,
        Required:     true,
        DefaultFunc:  schema.EnvDefaultFunc("TRAVIS_TOKEN", nil),
        ValidateFunc: validation.StringMatch(regexpValidate("`^[a-zA-Z0-9]+$`"), "The token argument value must conform to characters and integers."),
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

// regexp validator helper
func regexpValidate(expression string) *Regexp {
  // strip away error from regexp.compile for use in validation
  regexpStruct, _ := regexp.Compile(expression)
  return regexpStruct
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
