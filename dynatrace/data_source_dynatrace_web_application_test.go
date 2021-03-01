package dynatrace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDynatraceDataSourceWebApplication_basic(t *testing.T) {
	name := "Test Web Application"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDynatraceDataSourceWebApplicationBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("dynatrace_web_application.test", "name", name),
				),
			},
			{
				Config: testAccDynatraceDataSourceWebApplicationBasic(name) +
					testAccDynatraceDataSourceWebApplicationRead(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.dynatrace_web_application.test", "name", name),
				),
			},
		},
	})
}

func testAccDynatraceDataSourceWebApplicationBasic(name string) string {
	return fmt.Sprintf(`resource "dynatrace_web_application" "test" {
		name = "%s"
		cost_control_user_session_percentage = 100
        load_action_key_performance_metric   = "VISUALLY_COMPLETE"
        real_user_monitoring_enabled         = true
        type                                 = "AUTO_INJECTED"
        xhr_action_key_performance_metric    = "VISUALLY_COMPLETE"

        custom_action_apdex_settings {
            frustrating_fallback_threshold = 12000
            frustrating_threshold          = 12000
            threshold                      = 0
            tolerated_fallback_threshold   = 3000
            tolerated_threshold            = 3000
        }

        load_action_apdex_settings {
            frustrating_fallback_threshold = 12000
            frustrating_threshold          = 12000
            threshold                      = 0
            tolerated_fallback_threshold   = 3000
            tolerated_threshold            = 3000
        }

        monitoring_settings {
            add_cross_origin_anonymous_attribute = true
            cache_control_header_optimizations   = true
            fetch_requests                       = false
            injection_mode                       = "JAVASCRIPT_TAG"
            script_tag_cache_duration_in_hours   = 1
            secure_cookie_attribute              = false
            xml_http_request                     = false
            exclude_xhr_regex                    = ""
            cookie_placement_domain              = ""
            custom_configuration_properties      = ""
            server_request_path_id               = ""

            advanced_javascript_tag_settings {
                instrument_unsupported_ajax_frameworks = false
                max_action_name_length                 = 100
                max_errors_to_capture                  = 10
                sync_beacon_firefox                    = false
                sync_beacon_internet_explorer          = false
                special_characters_to_escape           = ""

                additional_event_handlers {
                    blur_event_handler            = false
                    change_event_handler          = false
                    click_event_handler           = false
                    max_dom_nodes_to_instrument   = 5000
                    mouseup_event_handler         = false
                    to_string_method              = false
                    user_mouseup_event_for_clicks = false
                }

                event_wrapper_settings {
                    blur        = false
                    change      = false
                    click       = false
                    mouse_up    = false
                    touch_end   = false
                    touch_start = false
                }

                global_event_capture_settings {
                    click        = true
                    double_click = true
                    key_down     = true
                    key_up       = true
                    mouse_down   = true
                    mouse_up     = true
                    scroll       = true
                    additional_event_captured_as_user_input = ""
                }
            }


            content_capture {
                javascript_errors                 = true
                visually_complete_and_speed_index = true

                resource_timing_settings {
                    non_w3c_resource_timings                       = false
                    non_w3c_resource_timings_instrumentation_delay = 50
                    resource_timing_capture_type                   = "CAPTURE_FULL_DETAILS"
                    resource_timings_domain_limit                  = 10
                    w3c_resource_timings                           = true
                }

                timeout_settings {
                    temporary_action_limit         = 0
                    temporary_action_total_timeout = 100
                    timed_action_support           = false
                }

                visually_complete_2_settings {
                    inactivity_timeout = 1000
                    mutation_timeout   = 50
                    threshold          = 50
                }
            }


            javascript_framework_support {
                activex_object = false
                angular        = false
                dojo           = false
                ext_js         = false
                icefaces       = false
                jquery         = false
                moo_tools      = false
                prototype      = false
            }
        }

        session_replay_config {
            cost_control_percentage = 100
            enabled                 = false
        }

        user_action_naming_settings {
            ignore_case                    = true
            split_user_actions_by_domain   = true
            use_first_detected_load_action = true
        }

        waterfall_settings {
            resources_browser_caching_threshold           = 50
            resources_threshold                           = 100000
            slow_cdn_resources_threshold                  = 200000
            slow_first_party_resources_threshold          = 200000
            slow_third_party_resources_threshold          = 200000
            speed_index_visually_complete_ratio_threshold = 50
            uncompressed_resources_threshold              = 860
        }

        xhr_action_apdex_settings {
            frustrating_fallback_threshold = 12000
            frustrating_threshold          = 12000
            threshold                      = 0
            tolerated_fallback_threshold   = 3000
            tolerated_threshold            = 3000
        }
    }
	  `, name)
}

func testAccDynatraceDataSourceWebApplicationRead() string {
	return fmt.Sprintf(`data "dynatrace_web_application" "test" {
    	name = "${dynatrace_web_application.test.name}"

}
`)
}
