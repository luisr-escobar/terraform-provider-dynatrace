package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDynatraceManagementZone_basic(t *testing.T) {
	rName := fmt.Sprintf("tf-acc-test-%s.dynatrace", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	name := fmt.Sprintf("%s", rName)
	resourceName := "dynatrace_management_zone.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDynatraceManagementZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceManagementZoneConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceManagementZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
				),
			},
			{
				Config: testAccDynatraceManagementZoneConfigModified(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceManagementZoneExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "rule.0.type"),
					resource.TestCheckResourceAttrSet(resourceName, "rule.0.enabled"),
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

func testAccCheckDynatraceManagementZoneDestroy(s *terraform.State) error {
	providerConf := testAccProvider.Meta().(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dynatrace_management_zone" {
			continue
		}

		managementZoneID := rs.Primary.ID

		managementZone, _, err := dynatraceConfigClientV1.ManagementZonesApi.GetManagementZone(authConfigV1, managementZoneID).Execute()
		if err == nil {
			if managementZone.Id == &rs.Primary.ID {
				return fmt.Errorf("Management Zone still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckDynatraceManagementZoneExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testAccProvider.Meta().(*ProviderConfiguration)
		dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
		authConfigV1 := providerConf.AuthConfigV1

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		managementZoneID := rs.Primary.ID

		_, _, err := dynatraceConfigClientV1.ManagementZonesApi.GetManagementZone(authConfigV1, managementZoneID).Execute()
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccDynatraceManagementZoneConfig(name string) string {
	return fmt.Sprintf(`resource "dynatrace_management_zone" "test" {
		name = "%s"
	}
	  `, name)
}

func testAccDynatraceManagementZoneConfigModified(name string) string {
	return fmt.Sprintf(`resource "dynatrace_management_zone" "test" {
		name = "%s"
		rule{
			type = "SERVICE"
			enabled = true
		
			condition {
			  key {
				attribute = "HOST_GROUP_NAME"
			  }
			  comparison_info {
				type = "STRING"
				operator = "BEGINS_WITH"
				value = jsonencode("simpleapp")
				negate = false
				case_sensitive = false
			  }
			}
		
		  }
	}
	  `, name)
}
