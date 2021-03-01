package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDynatraceDataSourceAlertingProfile_basic(t *testing.T) {
	name := "Test Alerting Profile"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceDataSourceAlertingProfileBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("dynatrace_alerting_profile.test", "display_name", name),
				),
			},
			{
				Config: testAccDynatraceDataSourceAlertingProfileBasic(name) +
					testAccDynatraceDataSourceAlertingProfileRead(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.dynatrace_alerting_profile.test", "display_name", name),
				),
			},
		},
	})
}

func testAccDynatraceDataSourceAlertingProfileBasic(name string) string {
	return fmt.Sprintf(`resource "dynatrace_alerting_profile" "test" {
		display_name = "%s"
	}
	  `, name)
}

func testAccDynatraceDataSourceAlertingProfileRead() string {
	return fmt.Sprintf(`data "dynatrace_alerting_profile" "test" {
    	display_name = "${dynatrace_alerting_profile.test.display_name}"
}
`)
}
