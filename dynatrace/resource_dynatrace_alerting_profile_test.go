package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDynatraceAlertingProfile_basic(t *testing.T) {
	rName := fmt.Sprintf("tf-acc-test-%s.dynatrace", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	name := fmt.Sprintf("%s", rName)
	resourceName := "dynatrace_alerting_profile.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDynatraceAlertingProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceAlertingProfileConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceAlertingProfileExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "display_name", name),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				),
			},
			{
				Config: testAccDynatraceAlertingProfileConfigModified(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceAlertingProfileExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "display_name", name),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "rule.0.severity_level"),
					resource.TestCheckResourceAttrSet(resourceName, "rule.0.tag_filters.0.include_mode"),
					resource.TestCheckResourceAttrSet(resourceName, "event_type_filter.0.predefined_event_filter.0.negate"),
					resource.TestCheckResourceAttr(resourceName, "rule.0.severity_level", "ERROR"),
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

func testAccCheckDynatraceAlertingProfileDestroy(s *terraform.State) error {
	providerConf := testAccProvider.Meta().(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dynatrace_alerting_profile" {
			continue
		}

		alertingProfileID := rs.Primary.ID

		alertingProfile, _, err := dynatraceConfigClientV1.AlertingProfilesApi.GetAlertingProfile(authConfigV1, alertingProfileID).Execute()
		if err == nil {
			if alertingProfile.Id == &rs.Primary.ID {
				return fmt.Errorf("Alerting profile still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckDynatraceAlertingProfileExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testAccProvider.Meta().(*ProviderConfiguration)
		dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
		authConfigV1 := providerConf.AuthConfigV1

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		alertingProfileID := rs.Primary.ID

		_, _, err := dynatraceConfigClientV1.AlertingProfilesApi.GetAlertingProfile(authConfigV1, alertingProfileID).Execute()
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccDynatraceAlertingProfileConfig(name string) string {
	return fmt.Sprintf(`resource "dynatrace_alerting_profile" "test" {
		display_name = "%s"
		rule{
			severity_level = "AVAILABILITY"
			tag_filters {
			include_mode = "INCLUDE_ALL"
			tag_filter {
				context = "CONTEXTLESS"
				key = "env"
				value = "prod"
			}
			}
			delay_in_minutes = 2
		}
	}
`, name)
}

func testAccDynatraceAlertingProfileConfigModified(name string) string {
	return fmt.Sprintf(`resource "dynatrace_alerting_profile" "test" {
		display_name = "%s"
		rule{
			severity_level = "ERROR"
			tag_filters {
				include_mode = "INCLUDE_ALL"
				tag_filter {
				context = "CONTEXTLESS"
				key = "env"
				value = "prod"
				}
			}
			delay_in_minutes = 2
		}
		event_type_filter{
			predefined_event_filter{
			  negate = true
			  event_type = "EC2_HIGH_CPU"
			}
		  }
		
		  event_type_filter{
			custom_event_filter{
			  custom_title_filter{
				enabled = true
				value = "sockshop"
				operator = "CONTAINS"
				negate = true
				case_insensitive = false
			  }
			}
		  }
	}
`, name)
}
