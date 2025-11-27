package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetFunctionAppStacksForLocation.json
func ExampleProviderClient_NewGetFunctionAppStacksForLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProviderClient().NewGetFunctionAppStacksForLocationPager("westus", &armappservice.ProviderClientGetFunctionAppStacksForLocationOptions{StackOsType: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FunctionAppStackCollection = armappservice.FunctionAppStackCollection{
		// 	Value: []*armappservice.FunctionAppStack{
		// 		{
		// 			Name: to.Ptr("dotnet"),
		// 			Type: to.Ptr("Microsoft.Web/locations/functionAppStacks"),
		// 			ID: to.Ptr("/providers/Microsoft.Web/locations/functionAppStacks/dotnet"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armappservice.FunctionAppStackProperties{
		// 				DisplayText: to.Ptr(".NET"),
		// 				MajorVersions: []*armappservice.FunctionAppMajorVersion{
		// 					{
		// 						DisplayText: to.Ptr(".NET Core 3"),
		// 						MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 3.1"),
		// 								StackSettings: &armappservice.FunctionAppRuntimes{
		// 									LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 										},
		// 										AppSettingsDictionary: map[string]*string{
		// 											"FUNCTIONS_WORKER_RUNTIME": to.Ptr("dotnet"),
		// 										},
		// 										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 											IsSupported: to.Ptr(true),
		// 											SupportedVersion: to.Ptr("3.1.301"),
		// 										},
		// 										RemoteDebuggingSupported: to.Ptr(false),
		// 										RuntimeVersion: to.Ptr("dotnet|3.1"),
		// 										SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 											LinuxFxVersion: to.Ptr("dotnet|3.1"),
		// 											Use32BitWorkerProcess: to.Ptr(false),
		// 										},
		// 										SupportedFunctionsExtensionVersions: []*string{
		// 											to.Ptr("~3")},
		// 										},
		// 										WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 											AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 												IsSupported: to.Ptr(true),
		// 											},
		// 											AppSettingsDictionary: map[string]*string{
		// 												"FUNCTIONS_WORKER_RUNTIME": to.Ptr("dotnet"),
		// 											},
		// 											GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 												IsSupported: to.Ptr(true),
		// 												SupportedVersion: to.Ptr("3.1.301"),
		// 											},
		// 											RemoteDebuggingSupported: to.Ptr(false),
		// 											RuntimeVersion: to.Ptr("3.1"),
		// 											SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 												Use32BitWorkerProcess: to.Ptr(true),
		// 											},
		// 											SupportedFunctionsExtensionVersions: []*string{
		// 												to.Ptr("~3")},
		// 											},
		// 										},
		// 										Value: to.Ptr("3.1"),
		// 								}},
		// 								Value: to.Ptr("3"),
		// 							},
		// 							{
		// 								DisplayText: to.Ptr(".NET Core 2"),
		// 								MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 									{
		// 										DisplayText: to.Ptr(".NET Core 2.2"),
		// 										StackSettings: &armappservice.FunctionAppRuntimes{
		// 											LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 												AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 													IsSupported: to.Ptr(true),
		// 												},
		// 												AppSettingsDictionary: map[string]*string{
		// 													"FUNCTIONS_WORKER_RUNTIME": to.Ptr("dotnet"),
		// 												},
		// 												GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 													IsSupported: to.Ptr(true),
		// 													SupportedVersion: to.Ptr("2.2.207"),
		// 												},
		// 												RemoteDebuggingSupported: to.Ptr(false),
		// 												RuntimeVersion: to.Ptr("dotnet|2.2"),
		// 												SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 													LinuxFxVersion: to.Ptr("dotnet|2.2"),
		// 													Use32BitWorkerProcess: to.Ptr(false),
		// 												},
		// 												SupportedFunctionsExtensionVersions: []*string{
		// 													to.Ptr("~2")},
		// 												},
		// 												WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 													AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 														IsSupported: to.Ptr(true),
		// 													},
		// 													AppSettingsDictionary: map[string]*string{
		// 														"FUNCTIONS_WORKER_RUNTIME": to.Ptr("dotnet"),
		// 													},
		// 													GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 														IsSupported: to.Ptr(true),
		// 														SupportedVersion: to.Ptr("2.2.207"),
		// 													},
		// 													RemoteDebuggingSupported: to.Ptr(false),
		// 													RuntimeVersion: to.Ptr("2.2"),
		// 													SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 														Use32BitWorkerProcess: to.Ptr(true),
		// 													},
		// 													SupportedFunctionsExtensionVersions: []*string{
		// 														to.Ptr("~2")},
		// 													},
		// 												},
		// 												Value: to.Ptr("2.2"),
		// 										}},
		// 										Value: to.Ptr("2"),
		// 									},
		// 									{
		// 										DisplayText: to.Ptr(".NET Framework 4"),
		// 										MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 											{
		// 												DisplayText: to.Ptr(".NET Framework 4.7"),
		// 												StackSettings: &armappservice.FunctionAppRuntimes{
		// 													WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 														AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 															IsSupported: to.Ptr(true),
		// 														},
		// 														AppSettingsDictionary: map[string]*string{
		// 														},
		// 														GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 															IsSupported: to.Ptr(false),
		// 														},
		// 														RemoteDebuggingSupported: to.Ptr(false),
		// 														RuntimeVersion: to.Ptr("4.7"),
		// 														SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 															Use32BitWorkerProcess: to.Ptr(true),
		// 														},
		// 														SupportedFunctionsExtensionVersions: []*string{
		// 															to.Ptr("~1")},
		// 														},
		// 													},
		// 													Value: to.Ptr("4.7"),
		// 											}},
		// 											Value: to.Ptr("2"),
		// 									}},
		// 									PreferredOs: to.Ptr(armappservice.StackPreferredOsWindows),
		// 									Value: to.Ptr("dotnet"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("node"),
		// 								Type: to.Ptr("Microsoft.Web/locations/functionAppStacks"),
		// 								ID: to.Ptr("/providers/Microsoft.Web/locations/functionAppStacks/node"),
		// 								Location: to.Ptr("westus"),
		// 								Properties: &armappservice.FunctionAppStackProperties{
		// 									DisplayText: to.Ptr("Node.js"),
		// 									MajorVersions: []*armappservice.FunctionAppMajorVersion{
		// 										{
		// 											DisplayText: to.Ptr("Node.js 14"),
		// 											MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 												{
		// 													DisplayText: to.Ptr("Node.js 14 LTS"),
		// 													StackSettings: &armappservice.FunctionAppRuntimes{
		// 														LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 															AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																IsSupported: to.Ptr(true),
		// 															},
		// 															AppSettingsDictionary: map[string]*string{
		// 																"FUNCTIONS_WORKER_RUNTIME": to.Ptr("node"),
		// 															},
		// 															GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																IsSupported: to.Ptr(true),
		// 																SupportedVersion: to.Ptr("14.x"),
		// 															},
		// 															IsHidden: to.Ptr(true),
		// 															IsPreview: to.Ptr(true),
		// 															RemoteDebuggingSupported: to.Ptr(false),
		// 															RuntimeVersion: to.Ptr("Node|14"),
		// 															SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																LinuxFxVersion: to.Ptr("Node|14"),
		// 																Use32BitWorkerProcess: to.Ptr(false),
		// 															},
		// 															SupportedFunctionsExtensionVersions: []*string{
		// 																to.Ptr("~3")},
		// 															},
		// 															WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																	IsSupported: to.Ptr(true),
		// 																},
		// 																AppSettingsDictionary: map[string]*string{
		// 																	"FUNCTIONS_WORKER_RUNTIME": to.Ptr("node"),
		// 																	"WEBSITE_NODE_DEFAULT_VERSION": to.Ptr("~14"),
		// 																},
		// 																GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																	IsSupported: to.Ptr(true),
		// 																	SupportedVersion: to.Ptr("14.x"),
		// 																},
		// 																IsHidden: to.Ptr(true),
		// 																IsPreview: to.Ptr(true),
		// 																RemoteDebuggingSupported: to.Ptr(false),
		// 																RuntimeVersion: to.Ptr("~14"),
		// 																SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																	Use32BitWorkerProcess: to.Ptr(true),
		// 																},
		// 																SupportedFunctionsExtensionVersions: []*string{
		// 																	to.Ptr("~3")},
		// 																},
		// 															},
		// 															Value: to.Ptr("14 LTS"),
		// 													}},
		// 													Value: to.Ptr("14"),
		// 												},
		// 												{
		// 													DisplayText: to.Ptr("Node.js 12"),
		// 													MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 														{
		// 															DisplayText: to.Ptr("Node.js 12 LTS"),
		// 															StackSettings: &armappservice.FunctionAppRuntimes{
		// 																LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																	AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																		IsSupported: to.Ptr(true),
		// 																	},
		// 																	AppSettingsDictionary: map[string]*string{
		// 																		"FUNCTIONS_WORKER_RUNTIME": to.Ptr("node"),
		// 																	},
		// 																	GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																		IsSupported: to.Ptr(true),
		// 																		SupportedVersion: to.Ptr("12.x"),
		// 																	},
		// 																	RemoteDebuggingSupported: to.Ptr(false),
		// 																	RuntimeVersion: to.Ptr("Node|12"),
		// 																	SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																		LinuxFxVersion: to.Ptr("Node|12"),
		// 																		Use32BitWorkerProcess: to.Ptr(false),
		// 																	},
		// 																	SupportedFunctionsExtensionVersions: []*string{
		// 																		to.Ptr("~3")},
		// 																	},
		// 																	WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																		AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																			IsSupported: to.Ptr(true),
		// 																		},
		// 																		AppSettingsDictionary: map[string]*string{
		// 																			"FUNCTIONS_WORKER_RUNTIME": to.Ptr("node"),
		// 																			"WEBSITE_NODE_DEFAULT_VERSION": to.Ptr("~12"),
		// 																		},
		// 																		GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																			IsSupported: to.Ptr(true),
		// 																			SupportedVersion: to.Ptr("12.x"),
		// 																		},
		// 																		RemoteDebuggingSupported: to.Ptr(false),
		// 																		RuntimeVersion: to.Ptr("~12"),
		// 																		SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																			Use32BitWorkerProcess: to.Ptr(true),
		// 																		},
		// 																		SupportedFunctionsExtensionVersions: []*string{
		// 																			to.Ptr("~3")},
		// 																		},
		// 																	},
		// 																	Value: to.Ptr("12 LTS"),
		// 															}},
		// 															Value: to.Ptr("12"),
		// 														},
		// 														{
		// 															DisplayText: to.Ptr("Node.js 10"),
		// 															MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																{
		// 																	DisplayText: to.Ptr("Node.js 10 LTS"),
		// 																	StackSettings: &armappservice.FunctionAppRuntimes{
		// 																		LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																			AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																				IsSupported: to.Ptr(true),
		// 																			},
		// 																			AppSettingsDictionary: map[string]*string{
		// 																				"FUNCTIONS_WORKER_RUNTIME": to.Ptr("node"),
		// 																			},
		// 																			GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																				IsSupported: to.Ptr(true),
		// 																				SupportedVersion: to.Ptr("10.x"),
		// 																			},
		// 																			RemoteDebuggingSupported: to.Ptr(false),
		// 																			RuntimeVersion: to.Ptr("Node|10"),
		// 																			SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																				LinuxFxVersion: to.Ptr("Node|10"),
		// 																				Use32BitWorkerProcess: to.Ptr(false),
		// 																			},
		// 																			SupportedFunctionsExtensionVersions: []*string{
		// 																				to.Ptr("~2"),
		// 																				to.Ptr("~3")},
		// 																			},
		// 																			WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																				AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																					IsSupported: to.Ptr(true),
		// 																				},
		// 																				AppSettingsDictionary: map[string]*string{
		// 																					"FUNCTIONS_WORKER_RUNTIME": to.Ptr("node"),
		// 																					"WEBSITE_NODE_DEFAULT_VERSION": to.Ptr("~10"),
		// 																				},
		// 																				GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																					IsSupported: to.Ptr(true),
		// 																					SupportedVersion: to.Ptr("10.x"),
		// 																				},
		// 																				RemoteDebuggingSupported: to.Ptr(false),
		// 																				RuntimeVersion: to.Ptr("~10"),
		// 																				SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																					Use32BitWorkerProcess: to.Ptr(true),
		// 																				},
		// 																				SupportedFunctionsExtensionVersions: []*string{
		// 																					to.Ptr("~2"),
		// 																					to.Ptr("~3")},
		// 																				},
		// 																			},
		// 																			Value: to.Ptr("10 LTS"),
		// 																	}},
		// 																	Value: to.Ptr("10"),
		// 																},
		// 																{
		// 																	DisplayText: to.Ptr("Node.js 8"),
		// 																	MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																		{
		// 																			DisplayText: to.Ptr("Node.js 8 LTS"),
		// 																			StackSettings: &armappservice.FunctionAppRuntimes{
		// 																				WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																					AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																						IsSupported: to.Ptr(true),
		// 																					},
		// 																					AppSettingsDictionary: map[string]*string{
		// 																						"FUNCTIONS_WORKER_RUNTIME": to.Ptr("node"),
		// 																						"WEBSITE_NODE_DEFAULT_VERSION": to.Ptr("~8"),
		// 																					},
		// 																					GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																						IsSupported: to.Ptr(true),
		// 																						SupportedVersion: to.Ptr("8.x"),
		// 																					},
		// 																					RemoteDebuggingSupported: to.Ptr(false),
		// 																					RuntimeVersion: to.Ptr("~8"),
		// 																					SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																						Use32BitWorkerProcess: to.Ptr(true),
		// 																					},
		// 																					SupportedFunctionsExtensionVersions: []*string{
		// 																						to.Ptr("~2")},
		// 																					},
		// 																				},
		// 																				Value: to.Ptr("8 LTS"),
		// 																		}},
		// 																		Value: to.Ptr("8"),
		// 																	},
		// 																	{
		// 																		DisplayText: to.Ptr("Node.js 6"),
		// 																		MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																			{
		// 																				DisplayText: to.Ptr("Node.js 6 LTS"),
		// 																				StackSettings: &armappservice.FunctionAppRuntimes{
		// 																					WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																						AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																							IsSupported: to.Ptr(true),
		// 																						},
		// 																						AppSettingsDictionary: map[string]*string{
		// 																							"WEBSITE_NODE_DEFAULT_VERSION": to.Ptr("~6"),
		// 																						},
		// 																						GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																							IsSupported: to.Ptr(false),
		// 																						},
		// 																						RemoteDebuggingSupported: to.Ptr(false),
		// 																						RuntimeVersion: to.Ptr("~6"),
		// 																						SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																							Use32BitWorkerProcess: to.Ptr(true),
		// 																						},
		// 																						SupportedFunctionsExtensionVersions: []*string{
		// 																							to.Ptr("~1")},
		// 																						},
		// 																					},
		// 																					Value: to.Ptr("6 LTS"),
		// 																			}},
		// 																			Value: to.Ptr("6"),
		// 																	}},
		// 																	PreferredOs: to.Ptr(armappservice.StackPreferredOsWindows),
		// 																	Value: to.Ptr("node"),
		// 																},
		// 															},
		// 															{
		// 																Name: to.Ptr("python"),
		// 																Type: to.Ptr("Microsoft.Web/locations/functionAppStacks"),
		// 																ID: to.Ptr("/providers/Microsoft.Web/locations/functionAppStacks/python"),
		// 																Location: to.Ptr("westus"),
		// 																Properties: &armappservice.FunctionAppStackProperties{
		// 																	DisplayText: to.Ptr("Python"),
		// 																	MajorVersions: []*armappservice.FunctionAppMajorVersion{
		// 																		{
		// 																			DisplayText: to.Ptr("Python 3"),
		// 																			MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																				{
		// 																					DisplayText: to.Ptr("Python 3.8"),
		// 																					StackSettings: &armappservice.FunctionAppRuntimes{
		// 																						LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																							AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																								IsSupported: to.Ptr(true),
		// 																							},
		// 																							AppSettingsDictionary: map[string]*string{
		// 																								"FUNCTIONS_WORKER_RUNTIME": to.Ptr("python"),
		// 																							},
		// 																							GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																								IsSupported: to.Ptr(true),
		// 																								SupportedVersion: to.Ptr("3.8"),
		// 																							},
		// 																							RemoteDebuggingSupported: to.Ptr(false),
		// 																							RuntimeVersion: to.Ptr("Python|3.8"),
		// 																							SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																								LinuxFxVersion: to.Ptr("Python|3.8"),
		// 																								Use32BitWorkerProcess: to.Ptr(false),
		// 																							},
		// 																							SupportedFunctionsExtensionVersions: []*string{
		// 																								to.Ptr("~3")},
		// 																							},
		// 																						},
		// 																						Value: to.Ptr("3.8"),
		// 																					},
		// 																					{
		// 																						DisplayText: to.Ptr("Python 3.7"),
		// 																						StackSettings: &armappservice.FunctionAppRuntimes{
		// 																							LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																								AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																									IsSupported: to.Ptr(true),
		// 																								},
		// 																								AppSettingsDictionary: map[string]*string{
		// 																									"FUNCTIONS_WORKER_RUNTIME": to.Ptr("python"),
		// 																								},
		// 																								GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																									IsSupported: to.Ptr(true),
		// 																									SupportedVersion: to.Ptr("3.7"),
		// 																								},
		// 																								RemoteDebuggingSupported: to.Ptr(false),
		// 																								RuntimeVersion: to.Ptr("Python|3.7"),
		// 																								SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																									LinuxFxVersion: to.Ptr("Python|3.7"),
		// 																									Use32BitWorkerProcess: to.Ptr(false),
		// 																								},
		// 																								SupportedFunctionsExtensionVersions: []*string{
		// 																									to.Ptr("~2"),
		// 																									to.Ptr("~3")},
		// 																								},
		// 																							},
		// 																							Value: to.Ptr("3.7"),
		// 																						},
		// 																						{
		// 																							DisplayText: to.Ptr("Python 3.6"),
		// 																							StackSettings: &armappservice.FunctionAppRuntimes{
		// 																								LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																									AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																										IsSupported: to.Ptr(true),
		// 																									},
		// 																									AppSettingsDictionary: map[string]*string{
		// 																										"FUNCTIONS_WORKER_RUNTIME": to.Ptr("python"),
		// 																									},
		// 																									GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																										IsSupported: to.Ptr(true),
		// 																										SupportedVersion: to.Ptr("3.6"),
		// 																									},
		// 																									RemoteDebuggingSupported: to.Ptr(false),
		// 																									RuntimeVersion: to.Ptr("Python|3.6"),
		// 																									SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																										LinuxFxVersion: to.Ptr("Python|3.6"),
		// 																										Use32BitWorkerProcess: to.Ptr(false),
		// 																									},
		// 																									SupportedFunctionsExtensionVersions: []*string{
		// 																										to.Ptr("~2"),
		// 																										to.Ptr("~3")},
		// 																									},
		// 																								},
		// 																								Value: to.Ptr("3.6"),
		// 																						}},
		// 																						Value: to.Ptr("3"),
		// 																				}},
		// 																				PreferredOs: to.Ptr(armappservice.StackPreferredOsLinux),
		// 																				Value: to.Ptr("python"),
		// 																			},
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("java"),
		// 																			Type: to.Ptr("Microsoft.Web/locations/functionAppStacks"),
		// 																			ID: to.Ptr("/providers/Microsoft.Web/locations/functionAppStacks/java"),
		// 																			Location: to.Ptr("westus"),
		// 																			Properties: &armappservice.FunctionAppStackProperties{
		// 																				DisplayText: to.Ptr("Java"),
		// 																				MajorVersions: []*armappservice.FunctionAppMajorVersion{
		// 																					{
		// 																						DisplayText: to.Ptr("Java 11"),
		// 																						MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																							{
		// 																								DisplayText: to.Ptr("Java 11"),
		// 																								StackSettings: &armappservice.FunctionAppRuntimes{
		// 																									LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																										AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																											IsSupported: to.Ptr(true),
		// 																										},
		// 																										AppSettingsDictionary: map[string]*string{
		// 																											"FUNCTIONS_WORKER_RUNTIME": to.Ptr("java"),
		// 																										},
		// 																										GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																											IsSupported: to.Ptr(true),
		// 																											SupportedVersion: to.Ptr("11"),
		// 																										},
		// 																										IsAutoUpdate: to.Ptr(true),
		// 																										RemoteDebuggingSupported: to.Ptr(false),
		// 																										RuntimeVersion: to.Ptr("Java|11"),
		// 																										SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																											LinuxFxVersion: to.Ptr("Java|11"),
		// 																											Use32BitWorkerProcess: to.Ptr(false),
		// 																										},
		// 																										SupportedFunctionsExtensionVersions: []*string{
		// 																											to.Ptr("~3")},
		// 																										},
		// 																										WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																											AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																												IsSupported: to.Ptr(true),
		// 																											},
		// 																											AppSettingsDictionary: map[string]*string{
		// 																												"FUNCTIONS_WORKER_RUNTIME": to.Ptr("java"),
		// 																											},
		// 																											GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																												IsSupported: to.Ptr(true),
		// 																												SupportedVersion: to.Ptr("11"),
		// 																											},
		// 																											IsAutoUpdate: to.Ptr(true),
		// 																											RemoteDebuggingSupported: to.Ptr(false),
		// 																											RuntimeVersion: to.Ptr("11"),
		// 																											SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																												JavaVersion: to.Ptr("11"),
		// 																												Use32BitWorkerProcess: to.Ptr(true),
		// 																											},
		// 																											SupportedFunctionsExtensionVersions: []*string{
		// 																												to.Ptr("~3")},
		// 																											},
		// 																										},
		// 																										Value: to.Ptr("11.0"),
		// 																								}},
		// 																								Value: to.Ptr("11"),
		// 																							},
		// 																							{
		// 																								DisplayText: to.Ptr("Java 8"),
		// 																								MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																									{
		// 																										DisplayText: to.Ptr("Java 8"),
		// 																										StackSettings: &armappservice.FunctionAppRuntimes{
		// 																											LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																												AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																													IsSupported: to.Ptr(true),
		// 																												},
		// 																												AppSettingsDictionary: map[string]*string{
		// 																													"FUNCTIONS_WORKER_RUNTIME": to.Ptr("java"),
		// 																												},
		// 																												GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																													IsSupported: to.Ptr(true),
		// 																													SupportedVersion: to.Ptr("8"),
		// 																												},
		// 																												IsAutoUpdate: to.Ptr(true),
		// 																												IsDefault: to.Ptr(true),
		// 																												RemoteDebuggingSupported: to.Ptr(false),
		// 																												RuntimeVersion: to.Ptr("Java|8"),
		// 																												SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																													LinuxFxVersion: to.Ptr("Java|8"),
		// 																													Use32BitWorkerProcess: to.Ptr(false),
		// 																												},
		// 																												SupportedFunctionsExtensionVersions: []*string{
		// 																													to.Ptr("~3")},
		// 																												},
		// 																												WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																													AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																														IsSupported: to.Ptr(true),
		// 																													},
		// 																													AppSettingsDictionary: map[string]*string{
		// 																														"FUNCTIONS_WORKER_RUNTIME": to.Ptr("java"),
		// 																													},
		// 																													GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																														IsSupported: to.Ptr(true),
		// 																														SupportedVersion: to.Ptr("8"),
		// 																													},
		// 																													IsAutoUpdate: to.Ptr(true),
		// 																													IsDefault: to.Ptr(true),
		// 																													RemoteDebuggingSupported: to.Ptr(false),
		// 																													RuntimeVersion: to.Ptr("1.8"),
		// 																													SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																														JavaVersion: to.Ptr("1.8"),
		// 																														Use32BitWorkerProcess: to.Ptr(true),
		// 																													},
		// 																													SupportedFunctionsExtensionVersions: []*string{
		// 																														to.Ptr("~2"),
		// 																														to.Ptr("~3")},
		// 																													},
		// 																												},
		// 																												Value: to.Ptr("8.0"),
		// 																										}},
		// 																										Value: to.Ptr("8"),
		// 																								}},
		// 																								PreferredOs: to.Ptr(armappservice.StackPreferredOsWindows),
		// 																								Value: to.Ptr("java"),
		// 																							},
		// 																						},
		// 																						{
		// 																							Name: to.Ptr("powershell"),
		// 																							Type: to.Ptr("Microsoft.Web/locations/functionAppStacks"),
		// 																							ID: to.Ptr("/providers/Microsoft.Web/locations/functionAppStacks/powershell"),
		// 																							Location: to.Ptr("westus"),
		// 																							Properties: &armappservice.FunctionAppStackProperties{
		// 																								DisplayText: to.Ptr("PowerShell Core"),
		// 																								MajorVersions: []*armappservice.FunctionAppMajorVersion{
		// 																									{
		// 																										DisplayText: to.Ptr("PowerShell 7"),
		// 																										MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																											{
		// 																												DisplayText: to.Ptr("PowerShell 7.0"),
		// 																												StackSettings: &armappservice.FunctionAppRuntimes{
		// 																													WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																														AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																															IsSupported: to.Ptr(true),
		// 																														},
		// 																														AppSettingsDictionary: map[string]*string{
		// 																															"FUNCTIONS_WORKER_RUNTIME": to.Ptr("powershell"),
		// 																														},
		// 																														GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																															IsSupported: to.Ptr(false),
		// 																														},
		// 																														RemoteDebuggingSupported: to.Ptr(false),
		// 																														RuntimeVersion: to.Ptr("~7"),
		// 																														SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																															PowerShellVersion: to.Ptr("~7"),
		// 																															Use32BitWorkerProcess: to.Ptr(true),
		// 																														},
		// 																														SupportedFunctionsExtensionVersions: []*string{
		// 																															to.Ptr("~3")},
		// 																														},
		// 																													},
		// 																													Value: to.Ptr("7.0"),
		// 																											}},
		// 																											Value: to.Ptr("7"),
		// 																										},
		// 																										{
		// 																											DisplayText: to.Ptr("PowerShell Core 6"),
		// 																											MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																												{
		// 																													DisplayText: to.Ptr("PowerShell Core 6.2"),
		// 																													StackSettings: &armappservice.FunctionAppRuntimes{
		// 																														WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																															AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																																IsSupported: to.Ptr(true),
		// 																															},
		// 																															AppSettingsDictionary: map[string]*string{
		// 																																"FUNCTIONS_WORKER_RUNTIME": to.Ptr("powershell"),
		// 																															},
		// 																															EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-04T00:00:00.000Z"); return t}()),
		// 																															GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																																IsSupported: to.Ptr(false),
		// 																															},
		// 																															IsDeprecated: to.Ptr(true),
		// 																															RemoteDebuggingSupported: to.Ptr(false),
		// 																															RuntimeVersion: to.Ptr("~6"),
		// 																															SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																																PowerShellVersion: to.Ptr("~6"),
		// 																																Use32BitWorkerProcess: to.Ptr(true),
		// 																															},
		// 																															SupportedFunctionsExtensionVersions: []*string{
		// 																																to.Ptr("~2"),
		// 																																to.Ptr("~3")},
		// 																															},
		// 																														},
		// 																														Value: to.Ptr("6.2"),
		// 																												}},
		// 																												Value: to.Ptr("6"),
		// 																										}},
		// 																										PreferredOs: to.Ptr(armappservice.StackPreferredOsWindows),
		// 																										Value: to.Ptr("powershell"),
		// 																									},
		// 																								},
		// 																								{
		// 																									Name: to.Ptr("custom"),
		// 																									Type: to.Ptr("Microsoft.Web/locations/functionAppStacks"),
		// 																									ID: to.Ptr("/providers/Microsoft.Web/locations/functionAppStacks/custom"),
		// 																									Location: to.Ptr("westus"),
		// 																									Properties: &armappservice.FunctionAppStackProperties{
		// 																										DisplayText: to.Ptr("Custom"),
		// 																										MajorVersions: []*armappservice.FunctionAppMajorVersion{
		// 																											{
		// 																												DisplayText: to.Ptr("Custom"),
		// 																												MinorVersions: []*armappservice.FunctionAppMinorVersion{
		// 																													{
		// 																														DisplayText: to.Ptr("Custom Handler"),
		// 																														StackSettings: &armappservice.FunctionAppRuntimes{
		// 																															LinuxRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																																AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																																	IsSupported: to.Ptr(true),
		// 																																},
		// 																																AppSettingsDictionary: map[string]*string{
		// 																																	"FUNCTIONS_WORKER_RUNTIME": to.Ptr("custom"),
		// 																																},
		// 																																GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																																	IsSupported: to.Ptr(false),
		// 																																},
		// 																																IsPreview: to.Ptr(true),
		// 																																RemoteDebuggingSupported: to.Ptr(false),
		// 																																RuntimeVersion: to.Ptr(""),
		// 																																SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																																	LinuxFxVersion: to.Ptr(""),
		// 																																	Use32BitWorkerProcess: to.Ptr(false),
		// 																																},
		// 																																SupportedFunctionsExtensionVersions: []*string{
		// 																																	to.Ptr("~3"),
		// 																																	to.Ptr("~2")},
		// 																																},
		// 																																WindowsRuntimeSettings: &armappservice.FunctionAppRuntimeSettings{
		// 																																	AppInsightsSettings: &armappservice.AppInsightsWebAppStackSettings{
		// 																																		IsSupported: to.Ptr(true),
		// 																																	},
		// 																																	AppSettingsDictionary: map[string]*string{
		// 																																		"FUNCTIONS_WORKER_RUNTIME": to.Ptr("custom"),
		// 																																	},
		// 																																	GitHubActionSettings: &armappservice.GitHubActionWebAppStackSettings{
		// 																																		IsSupported: to.Ptr(false),
		// 																																	},
		// 																																	IsPreview: to.Ptr(true),
		// 																																	RemoteDebuggingSupported: to.Ptr(false),
		// 																																	RuntimeVersion: to.Ptr("custom"),
		// 																																	SiteConfigPropertiesDictionary: &armappservice.SiteConfigPropertiesDictionary{
		// 																																		Use32BitWorkerProcess: to.Ptr(true),
		// 																																	},
		// 																																	SupportedFunctionsExtensionVersions: []*string{
		// 																																		to.Ptr("~3"),
		// 																																		to.Ptr("~2")},
		// 																																	},
		// 																																},
		// 																																Value: to.Ptr("custom"),
		// 																														}},
		// 																														Value: to.Ptr("custom"),
		// 																												}},
		// 																												PreferredOs: to.Ptr(armappservice.StackPreferredOsWindows),
		// 																												Value: to.Ptr("custom"),
		// 																											},
		// 																									}},
		// 																								}
	}
}
