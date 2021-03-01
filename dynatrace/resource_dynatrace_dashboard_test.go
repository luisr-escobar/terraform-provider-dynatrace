package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDynatraceDashboard_basic(t *testing.T) {
	rName := fmt.Sprintf("tf-acc-test-%s.dynatrace", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	name := fmt.Sprintf("%s", rName)
	resourceName := "dynatrace_dashboard.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDynatraceDashboardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceDashboardConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceDashboardExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dashboard_metadata.0.name", name),
					resource.TestCheckResourceAttrSet(resourceName, "dashboard_metadata.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "tile.0.name"),
				),
			},
			{
				Config: testAccDynatraceDashboardConfigModified(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceDashboardExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "dashboard_metadata.0.name", name),
					resource.TestCheckResourceAttrSet(resourceName, "dashboard_metadata.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "tile.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "tile.1.name"),
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

func testAccCheckDynatraceDashboardDestroy(s *terraform.State) error {
	providerConf := testAccProvider.Meta().(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dynatrace_dashboard" {
			continue
		}

		dashboardID := rs.Primary.ID

		dashboard, _, err := dynatraceConfigClientV1.DashboardsApi.GetDashboard(authConfigV1, dashboardID).Execute()
		if err == nil {
			if dashboard.Id == &rs.Primary.ID {
				return fmt.Errorf("Dashboard still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckDynatraceDashboardExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testAccProvider.Meta().(*ProviderConfiguration)
		dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
		authConfigV1 := providerConf.AuthConfigV1

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		dashboardID := rs.Primary.ID

		_, _, err := dynatraceConfigClientV1.DashboardsApi.GetDashboard(authConfigV1, dashboardID).Execute()
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccDynatraceDashboardConfig(name string) string {
	return fmt.Sprintf(`resource "dynatrace_dashboard" "test" {
		dashboard_metadata {
			dashboard_filter {
			  timeframe = "l_7_DAYS"
			}
			name = "%s"
			tags = ["tag1", "tag2"]
			preset = false
			shared = true
			sharing_details {
			  link_shared = true
			  published = false
			}
		  }
		
		  tile {
			assigned_entities = []
			chart_visible = false
			configured = true
			custom_name = ""
			exclude_maintenance_windows = false
			limit = 0
			name = "Infrastructure"
			tile_type = "HEADER"
			tile_filter {}
			bounds {
			  top = 0
			  left = 0
			  width = 304
			  height = 38
			}
		  }	
	}
	  `, name)
}

func testAccDynatraceDashboardConfigModified(name string) string {
	return fmt.Sprintf(`resource "dynatrace_dashboard" "test" {
		dashboard_metadata {
			dashboard_filter {
			  timeframe = "l_7_DAYS"
			}
			name = "%s"
			tags = ["tag1", "tag2"]
			preset = false
			shared = true
			sharing_details {
			  link_shared = true
			  published = false
			}
		  }
		
		  tile {
			chart_visible = false
			configured = true
			custom_name = ""
			exclude_maintenance_windows = false
			limit = 0
			name = "Infrastructure"
			tile_type = "HEADER"
			tile_filter {}
			bounds {
			  top = 0
			  left = 0
			  width = 304
			  height = 38
			}
		  }

		  tile {
			name = "Host CPU Load"
			tile_filter {}
			tile_type = "CUSTOM_CHARTING"
			limit = 0
			bounds {
			  top = 38
			  left = 0
			  width = 456
			  height = 152
			}
			chart_visible = false
			configured = true
			exclude_maintenance_windows = false
			filter_config {
			  chart_config {
				legend_shown = true
				series {
				  aggregation = "AVG"
				  aggregation_rate = "TOTAL"
				  entity_type = "HOST"
				  metric = "builtin:host.cpu.load"
				  percentile = 0
				  sort_ascending = false
				  sort_column = true
				  type = "LINE"
				}
				type = "TIMESERIES"
			  }
			  custom_name = "CPU"
			  default_name = "Custom Chart"
			  type = "MIXED"
			  filters_per_entity_type = jsonencode(
				{
				  HOST = {
					AUTO_TAGS = [
					  "keptn_stage:production"
					  ]
				  }
				}
			  )
			}
		  }
	}
	  `, name)
}
