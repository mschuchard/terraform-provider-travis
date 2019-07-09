package main

import (
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/terraform"
  "travis"
)

func main() {
  plugin.Serve(&plugin.ServeOpts {
    ProviderFunc: func() terraform.ResourceProvider {
      return travis.Provider()
    },
  })
}
