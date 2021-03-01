package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDynatraceAutoTag_basic(t *testing.T) {
	rName := fmt.Sprintf("tf-acc-test-%s.dynatrace", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	name := fmt.Sprintf("%s", rName)
	resourceName := "dynatrace_auto_tag.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDynatraceAutoTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceAutoTagConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceAutoTagExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
				),
			},
			{
				Config: testAccDynatraceAutoTagConfigModified(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDynatraceAutoTagExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttrSet(resourceName, "rule.0.enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "rule.0.type"),
					resource.TestCheckResourceAttrSet(resourceName, "rule.0.condition.key.0.attribute"),
					resource.TestCheckResourceAttr(resourceName, "rule.0.enabled", "true"),
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

func testAccCheckDynatraceAutoTagDestroy(s *terraform.State) error {
	providerConf := testAccProvider.Meta().(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dynatrace_auto_tag" {
			continue
		}

		autoTagID := rs.Primary.ID

		autoTag, _, err := dynatraceConfigClientV1.AutomaticallyAppliedTagsApi.GetAutoTag(authConfigV1, autoTagID).Execute()
		if err == nil {
			if autoTag.Id == &rs.Primary.ID {
				return fmt.Errorf("Auto Tag still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckDynatraceAutoTagExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testAccProvider.Meta().(*ProviderConfiguration)
		dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
		authConfigV1 := providerConf.AuthConfigV1

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		autoTagID := rs.Primary.ID

		_, _, err := dynatraceConfigClientV1.AutomaticallyAppliedTagsApi.GetAutoTag(authConfigV1, autoTagID).Execute()
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccDynatraceAutoTagConfig(name string) string {
	return fmt.Sprintf(`resource "dynatrace_auto_tag" "test" {
		name                = "%s"
		rule {
			enabled           = false
			propagation_types = [
					"SERVICE_TO_HOST_LIKE",
					"SERVICE_TO_PROCESS_GROUP_LIKE"
			]
			type              = "SERVICE"
			value_format      = "{ProcessGroup:ExePath/:\\/(.*?)\\/www}"
			condition {
					comparison_info {
						negate         = false
						operator       = "EXISTS"
						type           = "STRING"
					}
	
					key {
						attribute   = "PROCESS_GROUP_PREDEFINED_METADATA"
						dynamic_key = jsonencode("EXE_PATH")
						type        = "PROCESS_PREDEFINED_METADATA_KEY"
					}
				}
			}
		}
`, name)
}

func testAccDynatraceAutoTagConfigModified(name string) string {
	return fmt.Sprintf(`resource "dynatrace_auto_tag" "test" {
		name = "%s"
		rule {
			enabled           = true
			propagation_types = [
					"SERVICE_TO_HOST_LIKE",
					"SERVICE_TO_PROCESS_GROUP_LIKE"
			]
			type              = "SERVICE"
			value_format      = "{ProcessGroup:ExePath/:\\/(.*?)\\/www}"
			condition {
					comparison_info {
						negate         = false
						operator       = "EXISTS"
						type           = "STRING"
					}
	
					key {
						attribute   = "PROCESS_GROUP_PREDEFINED_METADATA"
						dynamic_key = jsonencode("EXE_PATH")
						type        = "PROCESS_PREDEFINED_METADATA_KEY"
					}
			}
			condition {
				comparison_info {
					negate         = true
					operator       = "EXISTS"
					type           = "STRING"
				}
				key {
					attribute   = "PROCESS_GROUP_PREDEFINED_METADATA"
					dynamic_key = jsonencode("IIS_APP_POOL")
					type        = "PROCESS_PREDEFINED_METADATA_KEY"
				}
			}
		}
	}
`, name)
}
