package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/plugin"
  "github.com/hashicorp/terraform-plugin-sdk/terraform"
  "travis"
)

func main() {
  plugin.Serve(&plugin.ServeOpts {
    ProviderFunc: travis.Provider,
  })
}
