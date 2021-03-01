package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDynatraceMaintenanceWindow_basic(t *testing.T) {
	rName := fmt.Sprintf("tf-acc-test-%s.dynatrace", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	name := fmt.Sprintf("%s", rName)
	resourceName := "dynatrace_maintenance_window.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDynatraceMaintenanceWindowDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceMaintenanceWindowConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceMaintenanceWindowExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
				),
			},
			{
				Config: testAccDynatraceMaintenanceWindowConfigModified(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceMaintenanceWindowExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttrSet(resourceName, "scope.0.match.0.type"),
					resource.TestCheckResourceAttrSet(resourceName, "schedule.0.recurrence_type"),
					resource.TestCheckResourceAttr(resourceName, "type", "UNPLANNED"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDynatraceMaintenanceWindowDestroy(s *terraform.State) error {
	providerConf := testAccProvider.Meta().(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dynatrace_maintenance_window" {
			continue
		}

		maintenanceWindowID := rs.Primary.ID

		maintenanceWindow, _, err := dynatraceConfigClientV1.MaintenanceWindowsApi.GetMaintenanceWindow(authConfigV1, maintenanceWindowID).Execute()
		if err == nil {
			if maintenanceWindow.Id == &rs.Primary.ID {
				return fmt.Errorf("Maintenance Window still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckDynatraceMaintenanceWindowExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testAccProvider.Meta().(*ProviderConfiguration)
		dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
		authConfigV1 := providerConf.AuthConfigV1

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		mainteanceWindowID := rs.Primary.ID

		_, _, err := dynatraceConfigClientV1.MaintenanceWindowsApi.GetMaintenanceWindow(authConfigV1, mainteanceWindowID).Execute()
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccDynatraceMaintenanceWindowConfig(name string) string {
	return fmt.Sprintf(`resource "dynatrace_maintenance_window" "test" {
		name = "%s"
		description = "Weekly udpate of windows servers"
		type = "UNPLANNED" 
		suppression = "DETECT_PROBLEMS_DONT_ALERT"
		scope {
			match {
			type = "HOST"
			tags {
				context = "CONTEXTLESS"
				key = "OS"  
				value = "windows"
			}
			}
		}
		schedule {
			recurrence_type = "ONCE"
			start = "2020-10-20 15:38"
			end = "2020-10-25 15:38"
			zone_id = "America/Chicago"
		}
		}
`, name)
}

func testAccDynatraceMaintenanceWindowConfigModified(name string) string {
	return fmt.Sprintf(`resource "dynatrace_maintenance_window" "test" {
		name = "%s"
		description = "Weekly udpate of windows servers"
		type = "UNPLANNED" 
		suppression = "DETECT_PROBLEMS_DONT_ALERT"
		scope {
		  match {
			type = "HOST"
			tags {
			  context = "CONTEXTLESS"
			  key = "OS"  
			  value = "windows"
			}
		  }
		}
		schedule {
		  recurrence_type = "WEEKLY"
		  recurrence {
			day_of_week = "FRIDAY"
			start_time = "19:21"
			duration_minutes = 60
		  }
		  start = "2025-10-20 15:38"
		  end = "2025-10-25 15:38"
		  zone_id = "America/Chicago"
		}
	  }
`, name)
}
