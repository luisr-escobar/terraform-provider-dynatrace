package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDynatraceDataSourceManagementZone_basic(t *testing.T) {
	name := "Test Management Zone"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceDataSourceManagementZoneBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("dynatrace_management_zone.test", "name", name),
				),
			},
			{
				Config: testAccDynatraceDataSourceManagementZoneBasic(name) +
					testAccDynatraceDataSourceManagementZoneRead(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.dynatrace_management_zone.test", "name", name),
				),
			},
		},
	})
}

func testAccDynatraceDataSourceManagementZoneBasic(name string) string {
	return fmt.Sprintf(`resource "dynatrace_management_zone" "test" {
		name = "%s"
	}
	  `, name)
}

func testAccDynatraceDataSourceManagementZoneRead() string {
	return fmt.Sprintf(`data "dynatrace_management_zone" "test" {
    	name = "${dynatrace_management_zone.test.name}"
}
`)
}
