package dynatrace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDynatraceWebApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDynatraceWebApplicationCreate,
		ReadContext:   resourceDynatraceWebApplicationRead,
		UpdateContext: resourceDynatraceWebApplicationUpdate,
		DeleteContext: resourceDynatraceWebApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the web application, displayed in the UI.",
				Required:    true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The type of the web application.",
				Optional:    true,
			},
			"real_user_monitoring_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Real user monitoring enabled/disabled.",
				Required:    true,
			},
			"cost_control_user_session_percentage": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Analize X% of user sessions.",
				Required:    true,
			},
			"load_action_key_performance_metric": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The key performance metric of load actions.",
				Required:    true,
			},
			"xhr_action_key_performance_metric": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The key performance metric of XHR actions.",
				Required:    true,
			},
			"url_injection_pattern": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Url injection pattern for manual web application.",
				Optional:    true,
			},
			"session_replay_config": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Session replay settings",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "You can exclude some actions from becoming XHR actions. Put a regular expression, matching all the required URLs, here. If noting specified the feature is disabled.",
							Required:    true,
						},
						"cost_control_percentage": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Session replay sampling rating in percentage.",
							Required:    true,
						},
					},
				},
			},
			"load_action_apdex_settings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Defines the Apdex settings of an application.",
				Required:    true,
				Elem:        apdexSchema(),
			},
			"xhr_action_apdex_settings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Defines the Apdex settings of an application.",
				Required:    true,
				Elem:        apdexSchema(),
			},
			"custom_action_apdex_settings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Defines the Apdex settings of an application.",
				Required:    true,
				Elem:        apdexSchema(),
			},
			"waterfall_settings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "These settings influence the monitoring data you receive for 3rd party, CDN, and 1st party resources.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uncompressed_resources_threshold": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Warn about uncompressed resources larger than X bytes.",
							Required:    true,
						},
						"resources_threshold": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Warn about resources larger than X bytes.",
							Required:    true,
						},
						"resources_browser_caching_threshold": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Warn about resources with a lower browser cache rate above X%.",
							Required:    true,
						},
						"slow_first_party_resources_threshold": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Warn about slow 1st party resources with a response time above X ms.",
							Required:    true,
						},
						"slow_third_party_resources_threshold": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Warn about slow 3rd party resources with a response time above X ms.",
							Required:    true,
						},
						"slow_cdn_resources_threshold": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Warn about slow CDN resources with a response time above X ms.",
							Required:    true,
						},
						"speed_index_visually_complete_ratio_threshold": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Warn if Speed index exceeds X % of Visually complete.",
							Required:    true,
						},
					},
				},
			},
			"monitoring_settings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Real user monitoring settings.",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fetch_requests": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "fetch() request capture enabled/disabled.",
							Required:    true,
						},
						"xml_http_request": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "XmlHttpRequest support enabled/disabled.",
							Required:    true,
						},
						"exclude_xhr_regex": &schema.Schema{
							Type:        schema.TypeString,
							Description: "You can exclude some actions from becoming XHR actions. Put a regular expression, matching all the required URLs, here. If noting specified the feature is disabled.",
							Required:    true,
						},
						"correlation_header_inclusion_regex": &schema.Schema{
							Type:        schema.TypeString,
							Description: "To enable RUM for XHR calls to AWS Lambda, define a regular expression matching these calls, Dynatrace can then automatically add a custom header (x-dtc) to each such request to the respective endpoints in AWS. Important: These endpoints must accept the x-dtc header, or the requests will fail.",
							Optional:    true,
						},
						"injection_mode": &schema.Schema{
							Type:        schema.TypeString,
							Description: "JavaScript injection mode.",
							Required:    true,
						},
						"add_cross_origin_anonymous_attribute": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Add the cross origin = anonymous attribute to capture JavaScript error messages and W3C resource timings.",
							Optional:    true,
						},
						"script_tag_cache_duration_in_hours": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Time duration for the cache settings.",
							Optional:    true,
						},
						"library_file_location": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The location of your application’s custom JavaScript library file. If nothing specified the root directory of your web server is used. Only supported by auto-injected applications.",
							Optional:    true,
						},
						"monitoring_data_path": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The location to send monitoring data from the JavaScript tag. Specify either a relative or an absolute URL. If you use an absolute URL, data will be sent using CORS.",
							Optional:    true,
						},
						"custom_configuration_properties": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Additional JavaScript tag properties that are specific to your application. To do this, type key=value pairs separated using a (|) symbol.",
							Required:    true,
						},
						"server_request_path_id": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Path to identify the server’s request ID.",
							Required:    true,
						},
						"secure_cookie_attribute": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Secure attribute usage for Dynatrace cookies enabled/disabled.",
							Required:    true,
						},
						"cookie_placement_domain": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Domain for cookie placement.",
							Required:    true,
						},
						"cache_control_header_optimizations": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled.",
							Required:    true,
						},
						"javascript_framework_support": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Support of various JavaScript frameworks.",
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"angular": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "AngularJS and Angular support enabled/disabled.",
										Required:    true,
									},
									"dojo": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Dojo support enabled/disabled.",
										Required:    true,
									},
									"ext_js": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "ExtJS, Sencha Touch support enabled/disabled.",
										Required:    true,
									},
									"icefaces": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "ICEfaces support enabled/disabled.",
										Required:    true,
									},
									"jquery": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "jQuery, Backbone.js support enabled/disabled.",
										Required:    true,
									},
									"moo_tools": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "MooTools support enabled/disabled.",
										Required:    true,
									},
									"prototype": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Prototype support enabled/disabled.",
										Required:    true,
									},
									"activex_object": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "ActiveXObject detection support enabled/disabled.",
										Required:    true,
									},
								},
							},
						},
						"content_capture": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Settings for content capture.",
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"javascript_errors": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "JavaScript errors monitoring enabled/disabled.",
										Required:    true,
									},
									"visually_complete_and_speed_index": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Visually complete and Speed index support enabled/disabled.",
										Required:    true,
									},
									"resource_timing_settings": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Settings for resource timings capture.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"w3c_resource_timings": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "W3C resource timings for third party/CDN enabled/disabled.",
													Required:    true,
												},
												"non_w3c_resource_timings": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Timing for JavaScript files and images on non-W3C supported browsers enabled/disabled.",
													Required:    true,
												},
												"non_w3c_resource_timings_instrumentation_delay": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "Instrumentation delay for monitoring resource and image resource impact in browsers that don't offer W3C resource timings. Valid values range from 0 to 9999. Only effective if nonW3cResourceTimings is enabled",
													Required:    true,
												},
												"resource_timing_capture_type": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Defines how detailed resource timings are captured.",
													Required:    true,
												},
												"resource_timings_domain_limit": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "Limits the number of domains for which W3C resource timings are captured. Only effective if resourceTimingCaptureType is CAPTURE_LIMITED_SUMMARIES.",
													Required:    true,
												},
											},
										},
									},
									"timeout_settings": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Settings for timed action capture.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"timed_action_support": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Timed action support enabled/disabled. Enable to detect actions that trigger sending of XHRs via setTimout methods.",
													Required:    true,
												},
												"temporary_action_limit": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "Defines how deep temporary actions may cascade. 0 disables temporary actions completely. Recommended value if enabled is 3.",
													Required:    true,
												},
												"temporary_action_total_timeout": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The total timeout of all cascaded timeouts that should still be able to create a temporary action.",
													Required:    true,
												},
											},
										},
									},
									"visually_complete_2_settings": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Settings for VisuallyComplete2",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"image_url_blacklist": &schema.Schema{
													Type:        schema.TypeString,
													Description: "A RegularExpression used to exclude images and iframes from being detected by the VC module.",
													Optional:    true,
												},
												"mutation_blacklist": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Query selector for mutation nodes to ignore in VC and SI calculation.",
													Optional:    true,
												},
												"mutation_timeout": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "Determines the time in ms VC waits after an action closes to start calculation. Defaults to 50.",
													Optional:    true,
												},
												"inactivity_timeout": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The time in ms the VC module waits for no mutations happening on the page after the load action. Defaults to 1000.",
													Optional:    true,
												},
												"threshold": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "Minimum visible area in pixels of elements to be counted towards VC and SI. Defaults to 50.",
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"advanced_javascript_tag_settings": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Advanced JavaScript tag settings.",
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sync_beacon_firefox": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Send the beacon signal as a synchronous XMLHttpRequest using Firefox enabled/disabled.",
										Optional:    true,
									},
									"sync_beacon_internet_explorer": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Send the beacon signal as a synchronous XMLHttpRequest using Internet Explorer enabled/disabled.",
										Optional:    true,
									},
									"instrument_unsupported_ajax_frameworks": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Instrumentation of unsupported Ajax frameworks enabled/disabled.",
										Required:    true,
									},
									"special_characters_to_escape": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Additional special characters that are to be escaped using non-alphanumeric characters in HTML escape format.",
										Required:    true,
									},
									"max_action_name_length": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum character length for action names. Valid values range from 5 to 10000.",
										Required:    true,
									},
									"max_errors_to_capture": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Maximum number of errors to be captured per page. Valid values range from 0 to 50.",
										Required:    true,
									},
									"additional_event_handlers": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Additional event handlers and wrappers.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_mouseup_event_for_clicks": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Use mouseup event for clicks enabled/disabled.",
													Required:    true,
												},
												"click_event_handler": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Click event handler enabled/disabled.",
													Required:    true,
												},
												"mouseup_event_handler": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Mouseup event handler enabled/disabled.",
													Required:    true,
												},
												"blur_event_handler": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Blur event handler enabled/disabled.",
													Required:    true,
												},
												"change_event_handler": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Change event handler enabled/disabled.",
													Required:    true,
												},
												"to_string_method": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "toString method enabled/disabled.",
													Required:    true,
												},
												"max_dom_nodes_to_instrument": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "Max. number of DOM nodes to instrument. Valid values range from 0 to 100000.",
													Required:    true,
												},
											},
										},
									},
									"event_wrapper_settings": &schema.Schema{
										Type:        schema.TypeList,
										Description: "In addition to the event handlers, events called using addEventListener or attachEvent can be captured. Be careful with this option! Event wrappers can conflict with the JavaScript code on a web pag.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"click": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Click enabled/disabled",
													Required:    true,
												},
												"mouse_up": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "MouseUp enabled/disabled.",
													Required:    true,
												},
												"change": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Change enabled/disabled.",
													Required:    true,
												},
												"blur": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Blur enabled/disabled.",
													Required:    true,
												},
												"touch_start": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "TouchStart enabled/disabled.",
													Required:    true,
												},
												"touch_end": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "TouchEnd enabled/disabled.",
													Required:    true,
												},
											},
										},
									},
									"global_event_capture_settings": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Global event capture settings.",
										Required:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mouse_up": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "MouseUp enabled/disabled.",
													Required:    true,
												},
												"mouse_down": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "MouseDown enabled/disabled.",
													Required:    true,
												},
												"click": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Click enabled/disabled.",
													Required:    true,
												},
												"double_click": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "DoubleClick enabled/disabled.",
													Required:    true,
												},
												"key_up": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "KeyUp enabled/disabled.",
													Required:    true,
												},
												"key_down": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "KeyDown enabled/disabled.",
													Required:    true,
												},
												"scroll": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "Scroll enabled/disabled.",
													Required:    true,
												},
												"additional_event_captured_as_user_input": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Additional events to be captured globally as user input. For example, DragStart or DragEnd.",
													Required:    true,
												},
											},
										},
									},
								},
							},
						},
						"browser_restriction_settings": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Settings for restricting certain browser type, version, platform and, comparator. It also restricts the mode.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The mode of the list of browser restrictions.",
										Required:    true,
									},
									"browser_restriction": &schema.Schema{
										Type:        schema.TypeList,
										Description: "A list of browser restrictions.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"browser_version": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The version of the browser that is used.",
													Optional:    true,
												},
												"browser_type": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The type of the browser that is used.",
													Required:    true,
												},
												"platform": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The platform on which the browser is being used.",
													Optional:    true,
												},
												"comparator": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Compares different browsers together.",
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"ip_address_restriction_settings": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Settings for restricting certain ip addresses and for introducing subnet mask. It also restricts the mode.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The mode of the list of ip address restrictions.",
										Optional:    true,
									},
									"ip_address_restriction": &schema.Schema{
										Type:        schema.TypeList,
										Description: "The IP address or the IP address range to be mapped to the location.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"subnet_mask": &schema.Schema{
													Type:        schema.TypeInt,
													Description: "The subnet mask of the IP address range.",
													Optional:    true,
												},
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The IP address to be mapped.",
													Required:    true,
												},
												"address_to": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The to address of the IP address range.",
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"javascript_injection_rule": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Rules for javascript injection.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "The enable or disable rule of the java script injection.",
										Required:    true,
									},
									"url_operator": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The url operator of the java script injection.",
										Required:    true,
									},
									"url_pattern": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The url pattern of the java script injection.",
										Optional:    true,
									},
									"rule": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The url rule of the java script injection.",
										Required:    true,
									},
									"html_pattern": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The html pattern of the java script injection.",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"user_tag": &schema.Schema{
				Type:        schema.TypeList,
				Description: "User Tags settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"unique_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "UniqueId, unique among all userTags and properties of this application.",
							Required:    true,
						},
						"metadata_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "If it's of type metaData, metaData id of the userTag.",
							Optional:    true,
						},
						"cleanup_rule": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Cleanup rule expression of the userTag",
							Optional:    true,
						},
						"server_side_request_attribute": &schema.Schema{
							Type:        schema.TypeString,
							Description: "requestAttribute Id of the userTag",
							Optional:    true,
						},
						"ignore_case": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "If true, the value of this tag will always be stored in lower case. Defaults to false.",
							Optional:    true,
						},
					},
				},
			},
			"user_action_and_session_property": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Defines userAction and session custom defined properties settings of an application.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The display name of the property.",
							Optional:    true,
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The data type of the property.",
							Required:    true,
						},
						"origin": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The origin of the property.",
							Required:    true,
						},
						"aggregation": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The aggregation type of the property.",
							Optional:    true,
						},
						"store_as_user_action_property": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "If true, the property is stored as a user action property.",
							Optional:    true,
						},
						"store_as_session_property": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "If true, the property is stored as a session property.",
							Optional:    true,
						},
						"cleanup_rule": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The cleanup rule of the property.",
							Optional:    true,
						},
						"server_side_request_attribute": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The ID of the request attribute.",
							Optional:    true,
						},
						"unique_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The ID of the request attribute.",
							Required:    true,
						},
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Key of the property.",
							Required:    true,
						},
						"metadata_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "If the origin is META_DATA, metaData id of the property.",
							Optional:    true,
						},
						"ignore_case": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "If true, the value of this property will always be stored in lower case. Defaults to false.",
							Required:    true,
						},
					},
				},
			},
			"user_action_naming_settings": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Defines userAction and session custom defined properties settings of an application.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ignore_case": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Case insensitive naming.",
							Optional:    true,
						},
						"use_first_detected_load_action": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "First load action found under an XHR action should be used when true. Else the deepest one under the xhr action is used.",
							Optional:    true,
						},
						"split_user_actions_by_domain": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Deactivate this setting if different domains should not result in separate user actions.",
							Optional:    true,
						},
						"placeholder": &schema.Schema{
							Type:        schema.TypeList,
							Description: "User action placeholders",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Placeholder name.",
										Required:    true,
									},
									"input": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Input.",
										Required:    true,
									},
									"processing_part": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Part.",
										Required:    true,
									},
									"metadata_id": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "Id of the metadata.",
										Optional:    true,
									},
									"use_guessed_element_identifier": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "Use the element identifier that was selected by Dynatrace.",
										Required:    true,
									},
									"processing_steps": &schema.Schema{
										Type:        schema.TypeList,
										Description: "Processing actions",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Description: "An action to be taken by the processing.",
													Required:    true,
												},
												"pattern_before": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The pattern before the required value. It will be removed.",
													Optional:    true,
												},
												"pattern_before_search_type": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The required occurrence of patternBefore.",
													Optional:    true,
												},
												"pattern_after": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The pattern after the required value. It will be removed.",
													Optional:    true,
												},
												"pattern_after_search_type": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The required occurrence of patternAfter.",
													Optional:    true,
												},
												"replacement": &schema.Schema{
													Type:        schema.TypeString,
													Description: "Replacement for the original value.",
													Optional:    true,
												},
												"pattern_to_replace": &schema.Schema{
													Type:        schema.TypeString,
													Description: "The pattern to be replaced. Only applicable if the type is REPLACE_WITH_PATTERN.",
													Optional:    true,
												},
												"regular_expression": &schema.Schema{
													Type:        schema.TypeString,
													Description: "A regular expression for the string to be extracted or replaced. Only applicable if the type is EXTRACT_BY_REGULAR_EXPRESSION or REPLACE_WITH_REGULAR_EXPRESSION.",
													Optional:    true,
												},
												"fallback_to_input": &schema.Schema{
													Type:        schema.TypeBool,
													Description: "If set to true: Returns the input if patternBefore or patternAfter cannot be found and the type is SUBSTRING. Returns the input if regularExpression doesn't match and type is EXTRACT_BY_REGULAR_EXPRESSION. Otherwise null is returned.",
													Optional:    true,
												},
											},
										},
									},
								},
							},
						},
						"load_action_naming_rule": &schema.Schema{
							Type:        schema.TypeList,
							Description: "User action naming rules for loading actions",
							Optional:    true,
							Elem:        userActionNamingRuleSchema(),
						},
						"xhr_action_naming_rule": &schema.Schema{
							Type:        schema.TypeList,
							Description: "User action naming rules for xhr actions.",
							Optional:    true,
							Elem:        userActionNamingRuleSchema(),
						},
						"custom_action_naming_rule": &schema.Schema{
							Type:        schema.TypeList,
							Description: "User action naming rules for custom actions.",
							Optional:    true,
							Elem:        userActionNamingRuleSchema(),
						},
					},
				},
			},
			"metadata_capture_setting": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Java script agent meta data capture settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The type of the meta data to capture.",
							Required:    true,
						},
						"capturing_name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The name of the meta data to capture.",
							Required:    true,
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Name for displaying the captured values in Dynatrace.",
							Required:    true,
						},
						"unique_id": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "The unique id of the meta data to capture.",
							Optional:    true,
						},
						"public_metadata": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "True if this metadata should be captured regardless of the privacy settings.",
							Optional:    true,
						},
					},
				},
			},
			"conversion_goal": &schema.Schema{
				Type:        schema.TypeList,
				Description: "A list of conversion goals of the application.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The name of the conversion goal.",
							Required:    true,
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The ID of conversion goal. Omit it while creating a new conversion goal.",
							Optional:    true,
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The type of the conversion goal.",
							Optional:    true,
						},
						"destination_details": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Configuration of a destination-based conversion goal.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url_or_path": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The path to be reached to hit the conversion goal.",
										Required:    true,
									},
									"match_type": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The operator of the match.",
										Optional:    true,
									},
									"case_sensitive": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "The match is case-sensitive (true) or (false).",
										Optional:    true,
									},
								},
							},
						},
						"user_action_details": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Configuration of a user action-based conversion goal.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The value to be matched to hit the conversion goal.",
										Optional:    true,
									},
									"case_sensitive": &schema.Schema{
										Type:        schema.TypeBool,
										Description: "The match is case-sensitive (true) or (false).",
										Optional:    true,
									},
									"match_type": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The operator of the match.",
										Optional:    true,
									},
									"match_entity": &schema.Schema{
										Type:        schema.TypeString,
										Description: "The type of the entity to which the rule applies.",
										Optional:    true,
									},
									"action_type": &schema.Schema{
										Type:        schema.TypeString,
										Description: "Type of the action to which the rule applies.",
										Optional:    true,
									},
								},
							},
						},
						"visit_duration_details": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Configuration of a visit duration-based conversion goal.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"duration_in_millis": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The duration of session to hit the conversion goal, in milliseconds.",
										Required:    true,
									},
								},
							},
						},
						"visit_num_action_details": &schema.Schema{
							Type:        schema.TypeList,
							Description: "Configuration of a number of user actions-based conversion goal.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_user_actions": &schema.Schema{
										Type:        schema.TypeInt,
										Description: "The number of user actions to hit the conversion goal.",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func apdexSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Maximal value of apdex, which is considered as satisfied user experience.",
				Optional:    true,
			},
			"tolerated_threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Maximal value of apdex, which is considered as satisfied user experience.",
				Optional:    true,
			},
			"frustrating_threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Maximal value of apdex, which is considered as tolerable user experience.",
				Optional:    true,
			},
			"tolerated_fallback_threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Fallback threshold of an XHR action, defining a satisfied user experience, when the configured KPM is not available.",
				Optional:    true,
			},
			"frustrating_fallback_threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Fallback threshold of an XHR action, defining a tolerable user experience, when the configured KPM is not available.",
				Optional:    true,
			},
		},
	}
}

func userActionNamingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"template": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Naming pattern. Use Curly brackets {} to select placeholders.",
				Required:    true,
			},
			"condition": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Defines the conditions when the naming rule should apply.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operand1": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Must be a defined placeholder wrapped in curly braces.",
							Required:    true,
						},
						"operand2": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Must be null if operator is \"IS_EMPTY\", a regex if operator is \"MATCHES_REGULAR_ERPRESSION\". In all other cases the value can be a freetext or a placeholder wrapped in curly braces.",
							Optional:    true,
						},
						"operator": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The operator of the condition.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func resourceDynatraceWebApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	wa, err := expandWebApplication(d)
	if err != nil {
		return diag.FromErr(err)
	}

	webApplication, _, err := dynatraceConfigClientV1.RUMWebApplicationConfigurationApi.CreateWebApplicationConfig(authConfigV1).WebApplicationConfig(*wa).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create dynatrace web app",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId(webApplication.Id)

	resourceDynatraceWebApplicationRead(ctx, d, m)

	return diags

}

func resourceDynatraceWebApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	webApplicationID := d.Id()

	webApplication, _, err := dynatraceConfigClientV1.RUMWebApplicationConfigurationApi.GetWebApplicationConfig(authConfigV1, webApplicationID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to read dynatrace web app",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	flattenDynatraceWebApplication(webApplication, d)

	return diags

}

func resourceDynatraceWebApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	webApplicationID := d.Id()

	wa, err := expandWebApplication(d)
	if err != nil {
		return diag.FromErr(err)
	}

	_, _, err = dynatraceConfigClientV1.RUMWebApplicationConfigurationApi.UpdateWebApplicationConfig(authConfigV1, webApplicationID).WebApplicationConfig(*wa).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update dynatrace web app",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	return resourceDynatraceWebApplicationRead(ctx, d, m)

}

func resourceDynatraceWebApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConf := m.(*ProviderConfiguration)
	dynatraceConfigClientV1 := providerConf.DynatraceConfigClientV1
	authConfigV1 := providerConf.AuthConfigV1

	var diags diag.Diagnostics

	webApplicationID := d.Id()

	_, err := dynatraceConfigClientV1.RUMWebApplicationConfigurationApi.DeleteWebApplicationConfig(authConfigV1, webApplicationID).Execute()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete dynatrace web app",
			Detail:   getErrorMessage(err),
		})
		return diags
	}

	d.SetId("")

	return diags
}
