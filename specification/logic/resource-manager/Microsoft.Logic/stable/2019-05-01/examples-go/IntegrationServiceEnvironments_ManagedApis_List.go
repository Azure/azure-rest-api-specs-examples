package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_List.json
func ExampleIntegrationServiceEnvironmentManagedApisClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationServiceEnvironmentManagedApisClient().NewListPager("testResourceGroup", "testIntegrationServiceEnvironment", nil)
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
		// page.IntegrationServiceEnvironmentManagedAPIListResult = armlogic.IntegrationServiceEnvironmentManagedAPIListResult{
		// 	Value: []*armlogic.IntegrationServiceEnvironmentManagedAPI{
		// 		{
		// 			Name: to.Ptr("azurefile"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 			ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/azurefile"),
		// 			Location: to.Ptr("brazilsouth"),
		// 			Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
		// 				Name: to.Ptr("azurefile"),
		// 				Capabilities: []*string{
		// 					to.Ptr("blob")},
		// 					ConnectionParameters: map[string]any{
		// 						"accessKey": map[string]any{
		// 							"type": "securestring",
		// 							"uiDefinition":map[string]any{
		// 								"description": "Specify a valid primary/secondary storage account access key.",
		// 								"constraints":map[string]any{
		// 									"required": "true",
		// 								},
		// 								"displayName": "Azure Storage Account Access Key",
		// 								"tooltip": "Specify a valid primary/secondary storage account access key.",
		// 							},
		// 						},
		// 						"accountName": map[string]any{
		// 							"type": "string",
		// 							"uiDefinition":map[string]any{
		// 								"description": "Name of the storage account the connector should use.",
		// 								"constraints":map[string]any{
		// 									"required": "true",
		// 								},
		// 								"displayName": "Azure Storage Account name",
		// 								"tooltip": "Provide the storage account name",
		// 							},
		// 						},
		// 					},
		// 					GeneralInformation: &armlogic.APIResourceGeneralInformation{
		// 						Description: to.Ptr("Microsoft Azure Storage provides a massively scalable, durable, and highly available storage for data on the cloud, and serves as the data storage solution for modern applications. Connect to File Storage to perform various operations such as create, update, get and delete on files in your Azure Storage account."),
		// 						DisplayName: to.Ptr("Azure File Storage"),
		// 						IconURL: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/azurefile/icon_1.0.1216.1605.png"),
		// 						ReleaseTag: to.Ptr("Preview"),
		// 						Tier: to.Ptr(armlogic.APITierStandard),
		// 					},
		// 					IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 						Name: to.Ptr("tes-ise-ga"),
		// 						Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 						ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 					},
		// 					Metadata: &armlogic.APIResourceMetadata{
		// 						BrandColor: to.Ptr("#804998"),
		// 						Source: to.Ptr("marketplace"),
		// 					},
		// 					ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 					RuntimeUrls: []*string{
		// 						to.Ptr("https://flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net/apim/azurefile")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("azuretables"),
		// 					Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 					ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/azuretables"),
		// 					Location: to.Ptr("brazilsouth"),
		// 					Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
		// 						Name: to.Ptr("azuretables"),
		// 						Capabilities: []*string{
		// 							to.Ptr("actions")},
		// 							ConnectionParameters: map[string]any{
		// 								"sharedkey": map[string]any{
		// 									"type": "securestring",
		// 									"uiDefinition":map[string]any{
		// 										"description": "The shared storage key of your storage account",
		// 										"constraints":map[string]any{
		// 											"required": "true",
		// 										},
		// 										"displayName": "Shared Storage Key",
		// 										"tooltip": "Provide a shared storage key for the storage account used for tables as it appears in the Azure portal",
		// 									},
		// 								},
		// 								"storageaccount": map[string]any{
		// 									"type": "string",
		// 									"uiDefinition":map[string]any{
		// 										"description": "The name of your storage account",
		// 										"constraints":map[string]any{
		// 											"required": "true",
		// 										},
		// 										"displayName": "Storage Account Name",
		// 										"tooltip": "Provide the name of the storage account used for tables as it appears in the Azure portal",
		// 									},
		// 								},
		// 							},
		// 							GeneralInformation: &armlogic.APIResourceGeneralInformation{
		// 								Description: to.Ptr("Azure Table storage is a service that stores structured NoSQL data in the cloud, providing a key/attribute store with a schemaless design. Sign into your Storage account to create, update, and query tables and more."),
		// 								DisplayName: to.Ptr("Azure Table Storage"),
		// 								IconURL: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/azuretables/icon_1.0.1216.1605.png"),
		// 								ReleaseTag: to.Ptr("Production"),
		// 								Tier: to.Ptr(armlogic.APITierStandard),
		// 							},
		// 							IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 								Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 								Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 								ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 							},
		// 							Metadata: &armlogic.APIResourceMetadata{
		// 								BrandColor: to.Ptr("#804998"),
		// 								Source: to.Ptr("marketplace"),
		// 							},
		// 							ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 							RuntimeUrls: []*string{
		// 								to.Ptr("https://flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net/apim/azuretables")},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("eventhubs"),
		// 							Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 							ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/eventhubs"),
		// 							Location: to.Ptr("brazilsouth"),
		// 							Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
		// 								Name: to.Ptr("eventhubs"),
		// 								Capabilities: []*string{
		// 									to.Ptr("actions")},
		// 									ConnectionParameters: map[string]any{
		// 										"connectionString": map[string]any{
		// 											"type": "securestring",
		// 											"uiDefinition":map[string]any{
		// 												"description": "Azure Event Hubs Connection String",
		// 												"constraints":map[string]any{
		// 													"required": "true",
		// 												},
		// 												"displayName": "Connection String",
		// 												"tooltip": "Provide Azure Event Hubs Connection String",
		// 											},
		// 										},
		// 									},
		// 									GeneralInformation: &armlogic.APIResourceGeneralInformation{
		// 										Description: to.Ptr("Connect to Azure Event Hubs to send and receive events."),
		// 										DisplayName: to.Ptr("Event Hubs"),
		// 										IconURL: to.Ptr("https://connectoricons-df.azureedge.net/eventhubs/icon_1.0.1228.1637.png"),
		// 										ReleaseTag: to.Ptr("Production"),
		// 										Tier: to.Ptr(armlogic.APITierStandard),
		// 									},
		// 									IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 										Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 										Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 										ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 									},
		// 									Metadata: &armlogic.APIResourceMetadata{
		// 										BrandColor: to.Ptr("#c4d5ff"),
		// 										Source: to.Ptr("marketplace"),
		// 									},
		// 									ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 									RuntimeUrls: []*string{
		// 										to.Ptr("https://flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net/apim/eventhubs")},
		// 									},
		// 								},
		// 								{
		// 									Name: to.Ptr("mq"),
		// 									Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 									ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/mq"),
		// 									Location: to.Ptr("brazilsouth"),
		// 									Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
		// 										Name: to.Ptr("mq"),
		// 										Capabilities: []*string{
		// 											to.Ptr("actions"),
		// 											to.Ptr("cloud")},
		// 											ConnectionParameters: map[string]any{
		// 												"authType": map[string]any{
		// 													"type": "string",
		// 													"allowedValues":[]any{
		// 														map[string]any{
		// 															"value": "basic",
		// 														},
		// 														map[string]any{
		// 															"value": "anonymous",
		// 														},
		// 													},
		// 													"uiDefinition":map[string]any{
		// 														"description": "Authentication type to connect to your Queue Manager",
		// 														"constraints":map[string]any{
		// 															"allowedValues":[]any{
		// 																map[string]any{
		// 																	"text": "Basic",
		// 																	"value": "basic",
		// 																},
		// 																map[string]any{
		// 																	"text": "Anonymous",
		// 																	"value": "anonymous",
		// 																},
		// 															},
		// 															"capability":[]any{
		// 															},
		// 															"required": "true",
		// 															"tabIndex": float64(6),
		// 														},
		// 														"displayName": "Authentication",
		// 														"tooltip": "Authentication type to connect to your Queue Manager",
		// 													},
		// 												},
		// 												"channel": map[string]any{
		// 													"type": "string",
		// 													"uiDefinition":map[string]any{
		// 														"description": "The channel to connect to the Queue Manager on.",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 																"cloud",
		// 															},
		// 															"required": "true",
		// 															"tabIndex": float64(3),
		// 														},
		// 														"displayName": "Channel name",
		// 														"tooltip": "Channel name",
		// 													},
		// 												},
		// 												"effectiveUserName": map[string]any{
		// 													"type": "string",
		// 													"uiDefinition":map[string]any{
		// 														"description": "The user name to use when connecting to MQ.",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 																"cloud",
		// 															},
		// 															"clearText": true,
		// 															"required": "true",
		// 															"tabIndex": float64(5),
		// 														},
		// 														"displayName": "Connect As",
		// 														"tooltip": "Connect As",
		// 													},
		// 												},
		// 												"enableSSL": map[string]any{
		// 													"type": "bool",
		// 													"uiDefinition":map[string]any{
		// 														"description": "Enable SSL? (True/False)",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 															},
		// 															"required": "false",
		// 															"tabIndex": float64(9),
		// 														},
		// 														"displayName": "Enable SSL?",
		// 														"tooltip": "Specify if SSL needs to be enabled",
		// 													},
		// 												},
		// 												"password": map[string]any{
		// 													"type": "securestring",
		// 													"uiDefinition":map[string]any{
		// 														"description": "Password credential",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 																"cloud",
		// 															},
		// 															"required": "false",
		// 															"tabIndex": float64(8),
		// 														},
		// 														"displayName": "Password",
		// 														"tooltip": "Password credential",
		// 													},
		// 												},
		// 												"queue": map[string]any{
		// 													"type": "string",
		// 													"uiDefinition":map[string]any{
		// 														"description": "The name of the default queue, may be overriden.",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 																"cloud",
		// 															},
		// 															"required": "true",
		// 															"tabIndex": float64(4),
		// 														},
		// 														"displayName": "Default queue name",
		// 														"tooltip": "Default queue name",
		// 													},
		// 												},
		// 												"queueManager": map[string]any{
		// 													"type": "string",
		// 													"uiDefinition":map[string]any{
		// 														"description": "The name of the Queue Manager to connect to.",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 																"cloud",
		// 															},
		// 															"required": "true",
		// 															"tabIndex": float64(2),
		// 														},
		// 														"displayName": "Queue Manager name",
		// 														"tooltip": "Queue Manager name",
		// 													},
		// 												},
		// 												"server": map[string]any{
		// 													"type": "string",
		// 													"uiDefinition":map[string]any{
		// 														"description": "The host name and port the MQ server is running on (host:port).",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 																"cloud",
		// 															},
		// 															"required": "true",
		// 															"tabIndex": float64(1),
		// 														},
		// 														"displayName": "Server",
		// 														"tooltip": "Server:Port",
		// 													},
		// 												},
		// 												"username": map[string]any{
		// 													"type": "securestring",
		// 													"uiDefinition":map[string]any{
		// 														"description": "Username credential",
		// 														"constraints":map[string]any{
		// 															"capability":[]any{
		// 																"cloud",
		// 															},
		// 															"clearText": true,
		// 															"required": "false",
		// 															"tabIndex": float64(7),
		// 														},
		// 														"displayName": "Username",
		// 														"tooltip": "Username credential",
		// 													},
		// 												},
		// 											},
		// 											GeneralInformation: &armlogic.APIResourceGeneralInformation{
		// 												Description: to.Ptr("Microsoft Connector for MQ integrates Azure applications with resources stored in an IBM WebSphere MQ server. Supports message send and receive operations."),
		// 												DisplayName: to.Ptr("MQ"),
		// 												IconURL: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/mq/icon_1.0.1216.1605.png"),
		// 												ReleaseTag: to.Ptr("Production"),
		// 												Tier: to.Ptr(armlogic.APITierEnterprise),
		// 											},
		// 											IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 												Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 												Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 												ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 											},
		// 											Metadata: &armlogic.APIResourceMetadata{
		// 												BrandColor: to.Ptr("#466BB0"),
		// 												Source: to.Ptr("marketplace"),
		// 											},
		// 											ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 											RuntimeUrls: []*string{
		// 												to.Ptr("https://flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net/apim/mq")},
		// 											},
		// 										},
		// 										{
		// 											Name: to.Ptr("sap"),
		// 											Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 											ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/sap"),
		// 											Location: to.Ptr("brazilsouth"),
		// 											Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
		// 												Name: to.Ptr("sap"),
		// 												Capabilities: []*string{
		// 													to.Ptr("actions"),
		// 													to.Ptr("triggers"),
		// 													to.Ptr("integrationServiceEnvironment")},
		// 													ConnectionParameters: map[string]any{
		// 														"appServerHost": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The hostname of the SAP Application Server. ",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"LogonType":map[string]any{
		// 																			"values":[]any{
		// 																				"ApplicationServer",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(7),
		// 																},
		// 																"displayName": "AS Host",
		// 																"tooltip": "The hostname of the specific SAP Application Server to connect to.",
		// 															},
		// 														},
		// 														"appServerService": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The service name or port number of the SAP Application Server.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"LogonType":map[string]any{
		// 																			"values":[]any{
		// 																				"ApplicationServer",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(8),
		// 																},
		// 																"displayName": "AS Service",
		// 																"tooltip": "The service name or port numbe of the specific SAP Application Server to connect to (Optional for connection type (Logon) A - Application Server).",
		// 															},
		// 														},
		// 														"authType": map[string]any{
		// 															"type": "string",
		// 															"allowedValues":[]any{
		// 																map[string]any{
		// 																	"value": "basic",
		// 																},
		// 															},
		// 															"defaultValue": "basic",
		// 															"uiDefinition":map[string]any{
		// 																"description": "Authentication type to connect to the SAP System. Must be basic (username and password).",
		// 																"constraints":map[string]any{
		// 																	"allowedValues":[]any{
		// 																		map[string]any{
		// 																			"text": "Basic",
		// 																			"value": "basic",
		// 																		},
		// 																	},
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"required": "true",
		// 																	"tabIndex": float64(3),
		// 																},
		// 																"displayName": "Authentication Type",
		// 																"tooltip": "Authentication type to connect to the SAP System.",
		// 															},
		// 														},
		// 														"client": map[string]any{
		// 															"type": "int",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The SAP client ID to connect to the SAP system.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"required": "true",
		// 																	"tabIndex": float64(2),
		// 																},
		// 																"displayName": "Client",
		// 																"tooltip": "The SAP backend's client (or 'Mandant') into which to log in. It is a number ranging from 000 to 999.",
		// 															},
		// 														},
		// 														"logonGroup": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The Logon Group for the SAP System.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"LogonType":map[string]any{
		// 																			"values":[]any{
		// 																				"Group",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(13),
		// 																},
		// 																"displayName": "MS Logon Group",
		// 																"tooltip": "The Logon Group for the SAP System, from which the Message Server shall select an Application Server (Only available if connection type (Logon) is B - Message Server (Group))",
		// 															},
		// 														},
		// 														"logonType": map[string]any{
		// 															"type": "string",
		// 															"allowedValues":[]any{
		// 																map[string]any{
		// 																	"uiDefinition":map[string]any{
		// 																		"description": "Application Server Logon (Type A)",
		// 																		"displayName": "Application Server (Type A)",
		// 																	},
		// 																	"value": "ApplicationServer",
		// 																},
		// 																map[string]any{
		// 																	"uiDefinition":map[string]any{
		// 																		"description": "Group Logon (Type B aka Message Server)",
		// 																		"displayName": "Group (Type B aka Message Server)",
		// 																	},
		// 																	"value": "Group",
		// 																},
		// 															},
		// 															"defaultValue": "Group",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The type of logon to the SAP System, either Application Server Logon (Type A) or Group Logon (Type B aka Message Server).",
		// 																"constraints":map[string]any{
		// 																	"allowedValues":[]any{
		// 																		map[string]any{
		// 																			"text": "Application Server",
		// 																			"value": "ApplicationServer",
		// 																		},
		// 																		map[string]any{
		// 																			"text": "Group",
		// 																			"value": "Group",
		// 																		},
		// 																	},
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"required": "true",
		// 																	"tabIndex": float64(6),
		// 																},
		// 																"displayName": "Logon Type",
		// 																"tooltip": "The type of logon to the SAP System.",
		// 															},
		// 														},
		// 														"messageServerHost": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The hostname of the SAP Message Server aka R3 System Name.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"LogonType":map[string]any{
		// 																			"values":[]any{
		// 																				"Group",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(10),
		// 																},
		// 																"displayName": "MS Server Host",
		// 																"tooltip": "The hostname of the SAP System’s Message Server (central instance) aka R3 System Name (Mandatory if connection type (Logon) is B - Message Server (Group)).",
		// 															},
		// 														},
		// 														"messageServerService": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The service name or port number of the Message Server.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"LogonType":map[string]any{
		// 																			"values":[]any{
		// 																				"Group",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(11),
		// 																},
		// 																"displayName": "MS Service Name or Port Number",
		// 																"tooltip": "The Service Name (as defined in etc/services) or the Port Number under which the Message Server is listening for load-balancing requests (Mandatory if connection type (Logon) is B - Message Server (Group) and System ID is not present)",
		// 															},
		// 														},
		// 														"password": map[string]any{
		// 															"type": "securestring",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The password to be used for log in to the SAP System.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"clearText": false,
		// 																	"required": "false",
		// 																	"tabIndex": float64(5),
		// 																},
		// 																"displayName": "SAP Password",
		// 																"tooltip": "The password to be used for log in to the SAP System.",
		// 															},
		// 														},
		// 														"safeTyping": map[string]any{
		// 															"type": "bool",
		// 															"uiDefinition":map[string]any{
		// 																"description": "Use strong types if false, otherwise use string as safe type.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(14),
		// 																},
		// 																"displayName": "Safe Typing",
		// 																"tooltip": "Enables safe typing, which treats types as string instead of strong types (e.g. Date, Time, Datetime), relaxing type and value validations.",
		// 															},
		// 														},
		// 														"sncLibraryPath": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "Name or path of the SNC library to be used.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"UseSnc":map[string]any{
		// 																			"values":[]any{
		// 																				"true",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(16),
		// 																},
		// 																"displayName": "SNC Library",
		// 																"tooltip": "Name or path of the SNC library to be used. Path can be absolute path or relative to the NCo library",
		// 															},
		// 														},
		// 														"sncMyName": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "Identity to be used for this particular destination/server (optional).",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"UseSnc":map[string]any{
		// 																			"values":[]any{
		// 																				"true",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(18),
		// 																},
		// 																"displayName": "SNC My Name",
		// 																"tooltip": "The installed SNC solution usually knows its own SNC name. Only for solutions supporting 'multiple identities', you may need to specify the identity to be used for this particular destination/server (optional)",
		// 															},
		// 														},
		// 														"sncPartnerName": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The backend's SNC name (Required when Logon Type is Application Server).",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"UseSnc":map[string]any{
		// 																			"values":[]any{
		// 																				"true",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(19),
		// 																},
		// 																"displayName": "SNC Partner Name",
		// 																"tooltip": "The backend's SNC name (Required when Logon Type is Application Server)",
		// 															},
		// 														},
		// 														"sncQop": map[string]any{
		// 															"type": "string",
		// 															"allowedValues":[]any{
		// 																map[string]any{
		// 																	"value": "Authentication",
		// 																},
		// 																map[string]any{
		// 																	"value": "Integrity",
		// 																},
		// 																map[string]any{
		// 																	"value": "Privacy",
		// 																},
		// 																map[string]any{
		// 																	"value": "Default",
		// 																},
		// 																map[string]any{
		// 																	"value": "Maximum",
		// 																},
		// 															},
		// 															"defaultValue": "Default",
		// 															"uiDefinition":map[string]any{
		// 																"description": "Quality of Service to be used for SNC communication of this particular destination/server.",
		// 																"constraints":map[string]any{
		// 																	"allowedValues":[]any{
		// 																		map[string]any{
		// 																			"text": "Authentication only",
		// 																			"value": "Authentication",
		// 																		},
		// 																		map[string]any{
		// 																			"text": "Authentication and integrity protection",
		// 																			"value": "Integrity",
		// 																		},
		// 																		map[string]any{
		// 																			"text": "Authentication, integrity and privacy protection (encryption)",
		// 																			"value": "Privacy",
		// 																		},
		// 																		map[string]any{
		// 																			"text": "Default protection",
		// 																			"value": "Default",
		// 																		},
		// 																		map[string]any{
		// 																			"text": "Maximum protection",
		// 																			"value": "Maximum",
		// 																		},
		// 																	},
		// 																	"capability":[]any{
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"UseSnc":map[string]any{
		// 																			"values":[]any{
		// 																				"true",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(20),
		// 																},
		// 																"displayName": "SNC Quality of Protection",
		// 																"tooltip": "Quality of Service to be used for SNC communication of this particular destination/server",
		// 															},
		// 														},
		// 														"sncSso": map[string]any{
		// 															"type": "string",
		// 															"allowedValues":[]any{
		// 																map[string]any{
		// 																	"value": "On",
		// 																},
		// 																map[string]any{
		// 																	"value": "Off",
		// 																},
		// 															},
		// 															"defaultValue": "On",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The SNC SSO specifies whether to use SNC identity or credentials provided on RFC level.",
		// 																"constraints":map[string]any{
		// 																	"allowedValues":[]any{
		// 																		map[string]any{
		// 																			"text": "Logon using the SNC identity",
		// 																			"value": "On",
		// 																		},
		// 																		map[string]any{
		// 																			"text": "Logon with the username/password information provided on RFC level",
		// 																			"value": "Off",
		// 																		},
		// 																	},
		// 																	"capability":[]any{
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"UseSnc":map[string]any{
		// 																			"values":[]any{
		// 																				"true",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(17),
		// 																},
		// 																"displayName": "SNC SSO",
		// 																"tooltip": "The SNC SSO specifies whether to use SNC identity or credentials provided on RFC level.",
		// 															},
		// 														},
		// 														"systemID": map[string]any{
		// 															"type": "string",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The System ID of the SAP system.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"LogonType":map[string]any{
		// 																			"values":[]any{
		// 																				"Group",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(12),
		// 																},
		// 																"displayName": "MS System ID",
		// 																"tooltip": "The SAP system’s three-letter system ID (Mandatory if connection type (Logon) is B - Message Server (Group) and Message Server Service is not present)",
		// 															},
		// 														},
		// 														"systemNumber": map[string]any{
		// 															"type": "int",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The SAP System’s System Number. It is a number ranging from 00 to 99.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"dependencies":map[string]any{
		// 																		"LogonType":map[string]any{
		// 																			"values":[]any{
		// 																				"ApplicationServer",
		// 																			},
		// 																		},
		// 																	},
		// 																	"required": "false",
		// 																	"tabIndex": float64(9),
		// 																},
		// 																"displayName": "AS System Number",
		// 																"tooltip": "The SAP System’s System Number. It is a number ranging from 00 to 99 (Mandatory if connection type (Logon) is A - Application Server).",
		// 															},
		// 														},
		// 														"useSnc": map[string]any{
		// 															"type": "bool",
		// 															"defaultValue": "false",
		// 															"uiDefinition":map[string]any{
		// 																"description": "When selected, the connections will be secured with SNC.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																	},
		// 																	"tabIndex": float64(15),
		// 																},
		// 																"displayName": "Use SNC",
		// 																"tooltip": "When selected, the connections will be secured with SNC",
		// 															},
		// 														},
		// 														"username": map[string]any{
		// 															"type": "securestring",
		// 															"uiDefinition":map[string]any{
		// 																"description": "The username to be used for log in to the SAP System.",
		// 																"constraints":map[string]any{
		// 																	"capability":[]any{
		// 																		"integrationServiceEnvironment",
		// 																	},
		// 																	"clearText": true,
		// 																	"required": "false",
		// 																	"tabIndex": float64(4),
		// 																},
		// 																"displayName": "SAP Username",
		// 																"tooltip": "The username to be used for log in to the SAP System.",
		// 															},
		// 														},
		// 													},
		// 													GeneralInformation: &armlogic.APIResourceGeneralInformation{
		// 														Description: to.Ptr("SAP connector is part of Microsoft’s hybrid integration platform. Use this connector to run SAP based applications side by side with Logic Apps to connect with On-Premise Line of Business systems."),
		// 														DisplayName: to.Ptr("SAP"),
		// 														IconURL: to.Ptr("https://connectoricons-df.azureedge.net/sap/icon_1.0.1224.1629.png"),
		// 														ReleaseTag: to.Ptr("Production"),
		// 														Tier: to.Ptr(armlogic.APITierEnterprise),
		// 													},
		// 													IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 														Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 														Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 														ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 													},
		// 													Metadata: &armlogic.APIResourceMetadata{
		// 														BrandColor: to.Ptr("#99e7ff"),
		// 														DeploymentParameters: &armlogic.APIDeploymentParameterMetadataSet{
		// 															PackageContentLink: &armlogic.APIDeploymentParameterMetadata{
		// 																Type: to.Ptr("blob"),
		// 																Description: to.Ptr("Upload package containing SAP NCo components and optionally SNC components."),
		// 																DisplayName: to.Ptr("SAP package"),
		// 																IsRequired: to.Ptr(true),
		// 																Visibility: to.Ptr(armlogic.APIDeploymentParameterVisibilityDefault),
		// 															},
		// 															RedisCacheConnectionString: &armlogic.APIDeploymentParameterMetadata{
		// 																Type: to.Ptr("string"),
		// 																Description: to.Ptr("Redis connection string"),
		// 																DisplayName: to.Ptr("Redis connection string"),
		// 																IsRequired: to.Ptr(true),
		// 																Visibility: to.Ptr(armlogic.APIDeploymentParameterVisibilityInternal),
		// 															},
		// 														},
		// 														Source: to.Ptr("marketplace"),
		// 													},
		// 													ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 													RuntimeUrls: []*string{
		// 														to.Ptr("https://flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net/apim/sap")},
		// 													},
		// 												},
		// 												{
		// 													Name: to.Ptr("servicebus"),
		// 													Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 													ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/servicebus"),
		// 													Location: to.Ptr("brazilsouth"),
		// 													Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
		// 														Name: to.Ptr("servicebus"),
		// 														Capabilities: []*string{
		// 															to.Ptr("actions")},
		// 															ConnectionParameters: map[string]any{
		// 																"connectionString": map[string]any{
		// 																	"type": "securestring",
		// 																	"uiDefinition":map[string]any{
		// 																		"description": "Azure Service Bus Connection String",
		// 																		"constraints":map[string]any{
		// 																			"required": "true",
		// 																		},
		// 																		"displayName": "Connection String",
		// 																		"tooltip": "Provide Azure Service Bus Connection String",
		// 																	},
		// 																},
		// 															},
		// 															GeneralInformation: &armlogic.APIResourceGeneralInformation{
		// 																Description: to.Ptr("Connect to Azure Service Bus to send and receive messages. You can perform actions such as send to queue, send to topic, receive from queue, receive from subscription, etc."),
		// 																DisplayName: to.Ptr("Service Bus"),
		// 																IconURL: to.Ptr("https://powerappsconnectorsdf.blob.core.windows.net/officialicons/servicebus/icon_1.0.1216.1605.png"),
		// 																ReleaseTag: to.Ptr("Production"),
		// 																Tier: to.Ptr(armlogic.APITierStandard),
		// 															},
		// 															IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 																Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 																Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 																ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 															},
		// 															Metadata: &armlogic.APIResourceMetadata{
		// 																BrandColor: to.Ptr("#c4d5ff"),
		// 																Source: to.Ptr("marketplace"),
		// 															},
		// 															ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 															RuntimeUrls: []*string{
		// 																to.Ptr("https://flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net/apim/servicebus")},
		// 															},
		// 														},
		// 														{
		// 															Name: to.Ptr("si3270"),
		// 															Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments/managedApis"),
		// 															ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment/managedApis/si3270"),
		// 															Location: to.Ptr("brazilsouth"),
		// 															Properties: &armlogic.IntegrationServiceEnvironmentManagedAPIProperties{
		// 																Name: to.Ptr("si3270"),
		// 																Capabilities: []*string{
		// 																	to.Ptr("actions")},
		// 																	ConnectionParameters: map[string]any{
		// 																		"codePage": map[string]any{
		// 																			"type": "int",
		// 																			"defaultValue": float64(37),
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "Specifies the code page of the host.",
		// 																				"constraints":map[string]any{
		// 																					"required": "false",
		// 																					"tabIndex": float64(4),
		// 																				},
		// 																				"displayName": "Code Page",
		// 																				"tooltip": "Code Page",
		// 																			},
		// 																		},
		// 																		"deviceType": map[string]any{
		// 																			"type": "string",
		// 																			"allowedValues":[]any{
		// 																				map[string]any{
		// 																					"value": "IBM-3278-2",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3278-2-E",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-2",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-2-E",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3278-3",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3278-3-E",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-3",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-3-E",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3278-4",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3278-4-E",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-4",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-4-E",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3278-5",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3278-5-E",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-5",
		// 																				},
		// 																				map[string]any{
		// 																					"value": "IBM-3279-5-E",
		// 																				},
		// 																			},
		// 																			"defaultValue": "IBM-3278-2",
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "The 3270 model to use.",
		// 																				"constraints":map[string]any{
		// 																					"allowedValues":[]any{
		// 																						map[string]any{
		// 																							"text": "IBM-3278-2",
		// 																							"value": "IBM-3278-2",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3278-2-E",
		// 																							"value": "IBM-3278-2-E",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-2",
		// 																							"value": "IBM-3279-2",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-2-E",
		// 																							"value": "IBM-3279-2-E",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3278-3",
		// 																							"value": "IBM-3278-3",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3278-3-E",
		// 																							"value": "IBM-3278-3-E",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-3",
		// 																							"value": "IBM-3279-3",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-3-E",
		// 																							"value": "IBM-3279-3-E",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3278-4",
		// 																							"value": "IBM-3278-4",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3278-4-E",
		// 																							"value": "IBM-3278-4-E",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-4",
		// 																							"value": "IBM-3279-4",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-4-E",
		// 																							"value": "IBM-3279-4-E",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3278-5",
		// 																							"value": "IBM-3278-5",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3278-5-E",
		// 																							"value": "IBM-3278-5-E",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-5",
		// 																							"value": "IBM-3279-5",
		// 																						},
		// 																						map[string]any{
		// 																							"text": "IBM-3279-5-E",
		// 																							"value": "IBM-3279-5-E",
		// 																						},
		// 																					},
		// 																					"required": "false",
		// 																					"tabIndex": float64(3),
		// 																				},
		// 																				"displayName": "Device Type",
		// 																				"tooltip": "Device Type",
		// 																			},
		// 																		},
		// 																		"integrationAccountId": map[string]any{
		// 																			"type": "string",
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "Integration Account ID",
		// 																				"constraints":map[string]any{
		// 																					"required": "true",
		// 																				},
		// 																				"displayName": "Integration Account ID",
		// 																				"tooltip": "Provide Resource ID for the integration account",
		// 																			},
		// 																		},
		// 																		"integrationAccountUrl": map[string]any{
		// 																			"type": "securestring",
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "Integration Account SAS URL",
		// 																				"constraints":map[string]any{
		// 																					"required": "true",
		// 																				},
		// 																				"displayName": "Integration Account SAS URL",
		// 																				"tooltip": "Provide SAS URL for the integration account",
		// 																			},
		// 																		},
		// 																		"logicalUnitName": map[string]any{
		// 																			"type": "string",
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "The name of specific logical unit to request from the TN 3270 server.",
		// 																				"constraints":map[string]any{
		// 																					"required": "false",
		// 																					"tabIndex": float64(5),
		// 																				},
		// 																				"displayName": "Logical Unit Name",
		// 																				"tooltip": "Logical Unit Name",
		// 																			},
		// 																		},
		// 																		"port": map[string]any{
		// 																			"type": "int",
		// 																			"defaultValue": float64(23),
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "Specifies the port for the TN 3270 server.",
		// 																				"constraints":map[string]any{
		// 																					"required": "false",
		// 																					"tabIndex": float64(2),
		// 																				},
		// 																				"displayName": "Port",
		// 																				"tooltip": "Specifies the port for the TN 3270 server",
		// 																			},
		// 																		},
		// 																		"server": map[string]any{
		// 																			"type": "string",
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "The name of the host running a TN 3270 server.",
		// 																				"constraints":map[string]any{
		// 																					"required": "true",
		// 																					"tabIndex": float64(1),
		// 																				},
		// 																				"displayName": "Server",
		// 																				"tooltip": "Server",
		// 																			},
		// 																		},
		// 																		"useSsl": map[string]any{
		// 																			"type": "bool",
		// 																			"defaultValue": "false",
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "Enable SSL? (True/False)",
		// 																				"constraints":map[string]any{
		// 																					"required": "false",
		// 																					"tabIndex": float64(6),
		// 																				},
		// 																				"displayName": "Enable SSL?",
		// 																				"tooltip": "Specify if SSL needs to be enabled",
		// 																			},
		// 																		},
		// 																		"validateCertificate": map[string]any{
		// 																			"type": "bool",
		// 																			"defaultValue": "false",
		// 																			"uiDefinition":map[string]any{
		// 																				"description": "Validate host ssl certificate? (True/False)",
		// 																				"constraints":map[string]any{
		// 																					"required": "false",
		// 																					"tabIndex": float64(7),
		// 																				},
		// 																				"displayName": "Validate host ssl certificate?",
		// 																				"tooltip": "Specify if host ssl certificate should be validated",
		// 																			},
		// 																		},
		// 																	},
		// 																	GeneralInformation: &armlogic.APIResourceGeneralInformation{
		// 																		Description: to.Ptr("Microsoft Connector for 3270 integrates Azure applications with resources available via TN 3270 Screen Scraping."),
		// 																		DisplayName: to.Ptr("IBM 3270"),
		// 																		IconURL: to.Ptr("https://connectoricons-df.azureedge.net/si3270/icon_1.0.1228.1636.png"),
		// 																		ReleaseTag: to.Ptr("Preview"),
		// 																		Tier: to.Ptr(armlogic.APITierEnterprise),
		// 																	},
		// 																	IntegrationServiceEnvironment: &armlogic.ResourceReference{
		// 																		Name: to.Ptr("testIntegrationServiceEnvironment"),
		// 																		Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
		// 																		ID: to.Ptr("/subscriptions/f34b22a3-2202-4fb1-b040-1332bd928c84/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationServiceEnvironments/testIntegrationServiceEnvironment"),
		// 																	},
		// 																	Metadata: &armlogic.APIResourceMetadata{
		// 																		BrandColor: to.Ptr("#59b4d9"),
		// 																		ConnectionType: to.Ptr("integrationAccount"),
		// 																		Source: to.Ptr("marketplace"),
		// 																	},
		// 																	ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
		// 																	RuntimeUrls: []*string{
		// 																		to.Ptr("https://flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net/apim/si3270")},
		// 																	},
		// 															}},
		// 														}
	}
}
