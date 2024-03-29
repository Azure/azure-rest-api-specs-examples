package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_NetworkHealth.json
func ExampleIntegrationServiceEnvironmentNetworkHealthClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationServiceEnvironmentNetworkHealthClient().Get(ctx, "testResourceGroup", "testIntegrationServiceEnvironment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Value = map[string]*armlogic.IntegrationServiceEnvironmentSubnetNetworkHealth{
	// 	"LA1": &armlogic.IntegrationServiceEnvironmentSubnetNetworkHealth{
	// 		NetworkDependencyHealthState: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityState("Healthy")),
	// 		OutboundNetworkDependencies: []*armlogic.IntegrationServiceEnvironmentNetworkDependency{
	// 			{
	// 				Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureStorage),
	// 				DisplayName: to.Ptr("Azure Storage"),
	// 				Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 					{
	// 						Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 						DomainName: to.Ptr("www.table.core.windows.net"),
	// 						Ports: []*string{
	// 							to.Ptr("80"),
	// 							to.Ptr("443")},
	// 						},
	// 						{
	// 							Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 							DomainName: to.Ptr("www.blob.core.windows.net"),
	// 							Ports: []*string{
	// 								to.Ptr("80"),
	// 								to.Ptr("443")},
	// 							},
	// 							{
	// 								Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 								DomainName: to.Ptr("www.queue.core.windows.net"),
	// 								Ports: []*string{
	// 									to.Ptr("80"),
	// 									to.Ptr("443")},
	// 							}},
	// 						},
	// 						{
	// 							Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureActiveDirectory),
	// 							DisplayName: to.Ptr("Azure Active Directory"),
	// 							Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 								{
	// 									Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 									DomainName: to.Ptr("graph.windows.net"),
	// 									Ports: []*string{
	// 										to.Ptr("80"),
	// 										to.Ptr("443")},
	// 								}},
	// 							},
	// 							{
	// 								Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSSLCertificateVerification),
	// 								DisplayName: to.Ptr("SSL Certificate Verification"),
	// 								Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 									{
	// 										Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 										DomainName: to.Ptr("ocsp.msocsp.com"),
	// 										Ports: []*string{
	// 											to.Ptr("80"),
	// 											to.Ptr("443")},
	// 										},
	// 										{
	// 											Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 											DomainName: to.Ptr("mscrl.microsoft.com"),
	// 											Ports: []*string{
	// 												to.Ptr("80"),
	// 												to.Ptr("443")},
	// 											},
	// 											{
	// 												Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 												DomainName: to.Ptr("crl.microsoft.com"),
	// 												Ports: []*string{
	// 													to.Ptr("80"),
	// 													to.Ptr("443")},
	// 												},
	// 												{
	// 													Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 													DomainName: to.Ptr("www.microsoft.com"),
	// 													Ports: []*string{
	// 														to.Ptr("80"),
	// 														to.Ptr("443")},
	// 													},
	// 													{
	// 														Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 														DomainName: to.Ptr("crl3.digicert.com"),
	// 														Ports: []*string{
	// 															to.Ptr("80"),
	// 															to.Ptr("443")},
	// 														},
	// 														{
	// 															Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 															DomainName: to.Ptr("ocsp.digicert.com"),
	// 															Ports: []*string{
	// 																to.Ptr("80"),
	// 																to.Ptr("443")},
	// 															},
	// 															{
	// 																Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																DomainName: to.Ptr("cacerts.digicert.com"),
	// 																Ports: []*string{
	// 																	to.Ptr("80"),
	// 																	to.Ptr("443")},
	// 																},
	// 																{
	// 																	Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																	DomainName: to.Ptr("www.thawte.com"),
	// 																	Ports: []*string{
	// 																		to.Ptr("80"),
	// 																		to.Ptr("443")},
	// 																}},
	// 															},
	// 															{
	// 																Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeDiagnosticLogsAndMetrics),
	// 																DisplayName: to.Ptr("Diagnostic Logs And Metrics"),
	// 																Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																	{
	// 																		Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																		DomainName: to.Ptr("az-prod.metrics.nsatc.net"),
	// 																		Ports: []*string{
	// 																			to.Ptr("443")},
	// 																	}},
	// 																},
	// 																{
	// 																	Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeIntegrationServiceEnvironmentConnectors),
	// 																	DisplayName: to.Ptr("Integration Service Environment Connectors"),
	// 																	Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																		{
	// 																			Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																			DomainName: to.Ptr("flow-yticlpbiumhig-cqsn-apim-runtime.brazilus.environments.microsoftazurelogicapps.net"),
	// 																			Ports: []*string{
	// 																				to.Ptr("443")},
	// 																		}},
	// 																	},
	// 																	{
	// 																		Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRedisCache),
	// 																		DisplayName: to.Ptr("Redis Cache"),
	// 																		Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																			{
	// 																				Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																				DomainName: to.Ptr("flowyticlpbiumhigcqsn.redis.cache.windows.net"),
	// 																				Ports: []*string{
	// 																					to.Ptr("6380")},
	// 																			}},
	// 																	}},
	// 																	OutboundNetworkHealth: &armlogic.IntegrationServiceEnvironmentNetworkDependencyHealth{
	// 																		State: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyHealthStateHealthy),
	// 																	},
	// 																},
	// 																"LA2": &armlogic.IntegrationServiceEnvironmentSubnetNetworkHealth{
	// 																	NetworkDependencyHealthState: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityState("Healthy")),
	// 																	OutboundNetworkDependencies: []*armlogic.IntegrationServiceEnvironmentNetworkDependency{
	// 																		{
	// 																			Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureStorage),
	// 																			DisplayName: to.Ptr("Azure Storage"),
	// 																			Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																				{
	// 																					Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																					DomainName: to.Ptr("blob.core.windows.net"),
	// 																					Ports: []*string{
	// 																						to.Ptr("80"),
	// 																						to.Ptr("443")},
	// 																					},
	// 																					{
	// 																						Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																						DomainName: to.Ptr("queue.core.windows.net"),
	// 																						Ports: []*string{
	// 																							to.Ptr("80"),
	// 																							to.Ptr("443")},
	// 																						},
	// 																						{
	// 																							Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																							DomainName: to.Ptr("table.core.windows.net"),
	// 																							Ports: []*string{
	// 																								to.Ptr("80"),
	// 																								to.Ptr("443")},
	// 																							},
	// 																							{
	// 																								Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																								DomainName: to.Ptr("file.core.windows.net"),
	// 																								Ports: []*string{
	// 																									to.Ptr("80"),
	// 																									to.Ptr("443")},
	// 																							}},
	// 																						},
	// 																						{
	// 																							Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSQL),
	// 																							DisplayName: to.Ptr("Azure SQL Database"),
	// 																							Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																								{
	// 																									Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																									DomainName: to.Ptr("database.windows.net"),
	// 																									Ports: []*string{
	// 																										to.Ptr("1433")},
	// 																								}},
	// 																							},
	// 																							{
	// 																								Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureManagement),
	// 																								DisplayName: to.Ptr("Azure Management"),
	// 																								Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																									{
	// 																										Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																										DomainName: to.Ptr("management.core.windows.net"),
	// 																										Ports: []*string{
	// 																											to.Ptr("443")},
	// 																										},
	// 																										{
	// 																											Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																											DomainName: to.Ptr("admin.core.windows.net"),
	// 																											Ports: []*string{
	// 																												to.Ptr("443")},
	// 																											},
	// 																											{
	// 																												Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																												DomainName: to.Ptr("management.azure.com"),
	// 																												Ports: []*string{
	// 																													to.Ptr("443")},
	// 																											}},
	// 																										},
	// 																										{
	// 																											Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureActiveDirectory),
	// 																											DisplayName: to.Ptr("Azure Active Directory"),
	// 																											Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																												{
	// 																													Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																													DomainName: to.Ptr("graph.windows.net"),
	// 																													Ports: []*string{
	// 																														to.Ptr("443")},
	// 																												}},
	// 																											},
	// 																											{
	// 																												Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRegionalService),
	// 																												DisplayName: to.Ptr("Regional Service"),
	// 																												Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																													{
	// 																														Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																														DomainName: to.Ptr("gr-prod-sn1.cloudapp.net"),
	// 																														Ports: []*string{
	// 																															to.Ptr("443")},
	// 																														},
	// 																														{
	// 																															Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																															DomainName: to.Ptr("az-prod.metrics.nsatc.net"),
	// 																															Ports: []*string{
	// 																																to.Ptr("443")},
	// 																														}},
	// 																													},
	// 																													{
	// 																														Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSSLCertificateVerification),
	// 																														DisplayName: to.Ptr("SSL Certificate Verification"),
	// 																														Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																															{
	// 																																Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																DomainName: to.Ptr("ocsp.msocsp.com"),
	// 																																Ports: []*string{
	// 																																	to.Ptr("80"),
	// 																																	to.Ptr("443")},
	// 																																},
	// 																																{
	// 																																	Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																	DomainName: to.Ptr("mscrl.microsoft.com"),
	// 																																	Ports: []*string{
	// 																																		to.Ptr("80"),
	// 																																		to.Ptr("443")},
	// 																																	},
	// 																																	{
	// 																																		Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																		DomainName: to.Ptr("crl.microsoft.com"),
	// 																																		Ports: []*string{
	// 																																			to.Ptr("80"),
	// 																																			to.Ptr("443")},
	// 																																		},
	// 																																		{
	// 																																			Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																			DomainName: to.Ptr("www.microsoft.com"),
	// 																																			Ports: []*string{
	// 																																				to.Ptr("80"),
	// 																																				to.Ptr("443")},
	// 																																			},
	// 																																			{
	// 																																				Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																				DomainName: to.Ptr("crl3.digicert.com"),
	// 																																				Ports: []*string{
	// 																																					to.Ptr("80"),
	// 																																					to.Ptr("443")},
	// 																																				},
	// 																																				{
	// 																																					Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																					DomainName: to.Ptr("ocsp.digicert.com"),
	// 																																					Ports: []*string{
	// 																																						to.Ptr("80"),
	// 																																						to.Ptr("443")},
	// 																																					},
	// 																																					{
	// 																																						Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																						DomainName: to.Ptr("cacerts.digicert.com"),
	// 																																						Ports: []*string{
	// 																																							to.Ptr("80"),
	// 																																							to.Ptr("443")},
	// 																																						},
	// 																																						{
	// 																																							Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																							DomainName: to.Ptr("www.thawte.com"),
	// 																																							Ports: []*string{
	// 																																								to.Ptr("80"),
	// 																																								to.Ptr("443")},
	// 																																						}},
	// 																																				}},
	// 																																				OutboundNetworkHealth: &armlogic.IntegrationServiceEnvironmentNetworkDependencyHealth{
	// 																																					State: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyHealthStateHealthy),
	// 																																				},
	// 																																			},
	// 																																			"LA3": &armlogic.IntegrationServiceEnvironmentSubnetNetworkHealth{
	// 																																				NetworkDependencyHealthState: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityState("Healthy")),
	// 																																				OutboundNetworkDependencies: []*armlogic.IntegrationServiceEnvironmentNetworkDependency{
	// 																																					{
	// 																																						Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeRecoveryService),
	// 																																						DisplayName: to.Ptr("Recovery Service"),
	// 																																						Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																																							{
	// 																																								Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																								DomainName: to.Ptr("https://prod3.metrics.nsatc.net:1886/RecoveryService"),
	// 																																								Ports: []*string{
	// 																																									to.Ptr("1886")},
	// 																																							}},
	// 																																						},
	// 																																						{
	// 																																							Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeSQL),
	// 																																							DisplayName: to.Ptr("SQL"),
	// 																																							Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																																								{
	// 																																									Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																									DomainName: to.Ptr("jnb6zdnqth.database.windows.net"),
	// 																																									Ports: []*string{
	// 																																										to.Ptr("1443")},
	// 																																								}},
	// 																																							},
	// 																																							{
	// 																																								Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeDiagnosticLogsAndMetrics),
	// 																																								DisplayName: to.Ptr("Diagnostic Logs And Metrics"),
	// 																																								Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																																									{
	// 																																										Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																										DomainName: to.Ptr("prod.warmpath.msftcloudes.com"),
	// 																																										Ports: []*string{
	// 																																											to.Ptr("443")},
	// 																																									}},
	// 																																								},
	// 																																								{
	// 																																									Category: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyCategoryTypeAzureStorage),
	// 																																									DisplayName: to.Ptr("Azure Storage"),
	// 																																									Endpoints: []*armlogic.IntegrationServiceEnvironmentNetworkEndpoint{
	// 																																										{
	// 																																											Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																											DomainName: to.Ptr("apimgmtstcfpq5vx3sh9wdme.blob.core.windows.net"),
	// 																																											Ports: []*string{
	// 																																												to.Ptr("443")},
	// 																																											},
	// 																																											{
	// 																																												Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																												DomainName: to.Ptr("apimgmtstcfpq5vx3sh9wdme.file.core.windows.net"),
	// 																																												Ports: []*string{
	// 																																													to.Ptr("445")},
	// 																																												},
	// 																																												{
	// 																																													Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																													DomainName: to.Ptr("apimgmtstcfpq5vx3sh9wdme.queue.core.windows.net"),
	// 																																													Ports: []*string{
	// 																																														to.Ptr("443")},
	// 																																													},
	// 																																													{
	// 																																														Accessibility: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityStateAvailable),
	// 																																														DomainName: to.Ptr("apimgmtstcfpq5vx3sh9wdme.table.core.windows.net"),
	// 																																														Ports: []*string{
	// 																																															to.Ptr("443")},
	// 																																													}},
	// 																																											}},
	// 																																											OutboundNetworkHealth: &armlogic.IntegrationServiceEnvironmentNetworkDependencyHealth{
	// 																																												State: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkDependencyHealthStateHealthy),
	// 																																											},
	// 																																										},
	// 																																										"LA4": &armlogic.IntegrationServiceEnvironmentSubnetNetworkHealth{
	// 																																											NetworkDependencyHealthState: to.Ptr(armlogic.IntegrationServiceEnvironmentNetworkEndPointAccessibilityState("Healthy")),
	// 																																										},
	// 																																									}
}
