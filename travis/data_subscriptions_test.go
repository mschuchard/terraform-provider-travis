package travis

import (
  "testing"

  "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSubscriptions(test *testing.T) {
	resource.Test(test, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSubscriptionsConfig,
				Check: resource.ComposeTestCheckFunc(
					testDataSubscriptionsCheck("data.subscriptions.by_id", "subscriptions.test"),
				),
			},
		},
	})
}

func TestDataSubscriptionsCheck(name, reference string) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrPair(name, "subscriptions", reference, "subscriptions"),
	)
}

// this data source has no attribute arguments
var testDataSubscriptionsConfig = `data "subscriptions" "test" {}`
