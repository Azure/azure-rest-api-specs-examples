package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2024-10-01-preview/examples/CapabilitySetByLocation.json
func ExampleLocationBasedCapabilitySetClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationBasedCapabilitySetClient().Get(ctx, "WestUS", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Capability = armmysqlflexibleservers.Capability{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/capabilities"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/providers/Microsoft.DBforMySQL/locations/WestUS/capabilitySets/default"),
	// 	Properties: &armmysqlflexibleservers.CapabilityPropertiesV2{
	// 		SupportedFeatures: []*armmysqlflexibleservers.FeatureProperty{
	// 			{
	// 				FeatureName: to.Ptr("SupportAcceleratedLogs"),
	// 				FeatureValue: to.Ptr("Enabled"),
	// 		}},
	// 		SupportedFlexibleServerEditions: []*armmysqlflexibleservers.ServerEditionCapabilityV2{
	// 			{
	// 				Name: to.Ptr("Burstable"),
	// 				DefaultSKU: to.Ptr("Standard_B1s"),
	// 				DefaultStorageSize: to.Ptr[int32](20480),
	// 				SupportedSKUs: []*armmysqlflexibleservers.SKUCapabilityV2{
	// 					{
	// 						Name: to.Ptr("Standard_B1s"),
	// 						SupportedHAMode: []*string{
	// 						},
	// 						SupportedIops: to.Ptr[int64](400),
	// 						SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
	// 						SupportedZones: []*string{
	// 							to.Ptr("1"),
	// 							to.Ptr("2"),
	// 							to.Ptr("3")},
	// 							VCores: to.Ptr[int64](1),
	// 						},
	// 						{
	// 							Name: to.Ptr("Standard_B1ms"),
	// 							SupportedHAMode: []*string{
	// 							},
	// 							SupportedIops: to.Ptr[int64](640),
	// 							SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
	// 							SupportedZones: []*string{
	// 								to.Ptr("1"),
	// 								to.Ptr("2"),
	// 								to.Ptr("3")},
	// 								VCores: to.Ptr[int64](1),
	// 							},
	// 							{
	// 								Name: to.Ptr("Standard_B2s"),
	// 								SupportedHAMode: []*string{
	// 								},
	// 								SupportedIops: to.Ptr[int64](1280),
	// 								SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
	// 								SupportedZones: []*string{
	// 									to.Ptr("1"),
	// 									to.Ptr("2"),
	// 									to.Ptr("3")},
	// 									VCores: to.Ptr[int64](2),
	// 								},
	// 								{
	// 									Name: to.Ptr("Standard_B2ms"),
	// 									SupportedHAMode: []*string{
	// 									},
	// 									SupportedIops: to.Ptr[int64](1700),
	// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 									SupportedZones: []*string{
	// 										to.Ptr("1"),
	// 										to.Ptr("2"),
	// 										to.Ptr("3")},
	// 										VCores: to.Ptr[int64](2),
	// 									},
	// 									{
	// 										Name: to.Ptr("Standard_B4ms"),
	// 										SupportedHAMode: []*string{
	// 										},
	// 										SupportedIops: to.Ptr[int64](2400),
	// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 										SupportedZones: []*string{
	// 											to.Ptr("1"),
	// 											to.Ptr("2"),
	// 											to.Ptr("3")},
	// 											VCores: to.Ptr[int64](4),
	// 										},
	// 										{
	// 											Name: to.Ptr("Standard_B8ms"),
	// 											SupportedHAMode: []*string{
	// 											},
	// 											SupportedIops: to.Ptr[int64](3100),
	// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 											SupportedZones: []*string{
	// 												to.Ptr("1"),
	// 												to.Ptr("2"),
	// 												to.Ptr("3")},
	// 												VCores: to.Ptr[int64](8),
	// 											},
	// 											{
	// 												Name: to.Ptr("Standard_B12ms"),
	// 												SupportedHAMode: []*string{
	// 												},
	// 												SupportedIops: to.Ptr[int64](3800),
	// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 												SupportedZones: []*string{
	// 													to.Ptr("1"),
	// 													to.Ptr("2"),
	// 													to.Ptr("3")},
	// 													VCores: to.Ptr[int64](12),
	// 												},
	// 												{
	// 													Name: to.Ptr("Standard_B16ms"),
	// 													SupportedHAMode: []*string{
	// 													},
	// 													SupportedIops: to.Ptr[int64](4300),
	// 													SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 													SupportedZones: []*string{
	// 														to.Ptr("1"),
	// 														to.Ptr("2"),
	// 														to.Ptr("3")},
	// 														VCores: to.Ptr[int64](16),
	// 													},
	// 													{
	// 														Name: to.Ptr("Standard_B20ms"),
	// 														SupportedHAMode: []*string{
	// 														},
	// 														SupportedIops: to.Ptr[int64](5000),
	// 														SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 														SupportedZones: []*string{
	// 															to.Ptr("1"),
	// 															to.Ptr("2"),
	// 															to.Ptr("3")},
	// 															VCores: to.Ptr[int64](20),
	// 													}},
	// 													SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
	// 														{
	// 															Name: to.Ptr("Premium"),
	// 															MaxBackupIntervalHours: to.Ptr[int64](24),
	// 															MaxBackupRetentionDays: to.Ptr[int64](35),
	// 															MaxStorageSize: to.Ptr[int64](16777216),
	// 															MinBackupIntervalHours: to.Ptr[int64](1),
	// 															MinBackupRetentionDays: to.Ptr[int64](7),
	// 															MinStorageSize: to.Ptr[int64](20480),
	// 													}},
	// 												},
	// 												{
	// 													Name: to.Ptr("GeneralPurpose"),
	// 													DefaultSKU: to.Ptr("Standard_D2ds_v4"),
	// 													DefaultStorageSize: to.Ptr[int32](65536),
	// 													SupportedSKUs: []*armmysqlflexibleservers.SKUCapabilityV2{
	// 														{
	// 															Name: to.Ptr("Standard_D2ds_v4"),
	// 															SupportedHAMode: []*string{
	// 																to.Ptr("SameZone"),
	// 																to.Ptr("ZoneRedundant")},
	// 																SupportedIops: to.Ptr[int64](3200),
	// 																SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																SupportedZones: []*string{
	// 																	to.Ptr("1"),
	// 																	to.Ptr("2"),
	// 																	to.Ptr("3")},
	// 																	VCores: to.Ptr[int64](2),
	// 																},
	// 																{
	// 																	Name: to.Ptr("Standard_D4ds_v4"),
	// 																	SupportedHAMode: []*string{
	// 																		to.Ptr("SameZone"),
	// 																		to.Ptr("ZoneRedundant")},
	// 																		SupportedIops: to.Ptr[int64](6400),
	// 																		SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																		SupportedZones: []*string{
	// 																			to.Ptr("1"),
	// 																			to.Ptr("2"),
	// 																			to.Ptr("3")},
	// 																			VCores: to.Ptr[int64](4),
	// 																		},
	// 																		{
	// 																			Name: to.Ptr("Standard_D8ds_v4"),
	// 																			SupportedHAMode: []*string{
	// 																				to.Ptr("SameZone"),
	// 																				to.Ptr("ZoneRedundant")},
	// 																				SupportedIops: to.Ptr[int64](12800),
	// 																				SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																				SupportedZones: []*string{
	// 																					to.Ptr("1"),
	// 																					to.Ptr("2"),
	// 																					to.Ptr("3")},
	// 																					VCores: to.Ptr[int64](8),
	// 																				},
	// 																				{
	// 																					Name: to.Ptr("Standard_D16ds_v4"),
	// 																					SupportedHAMode: []*string{
	// 																						to.Ptr("SameZone"),
	// 																						to.Ptr("ZoneRedundant")},
	// 																						SupportedIops: to.Ptr[int64](20000),
	// 																						SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																						SupportedZones: []*string{
	// 																							to.Ptr("1"),
	// 																							to.Ptr("2"),
	// 																							to.Ptr("3")},
	// 																							VCores: to.Ptr[int64](16),
	// 																						},
	// 																						{
	// 																							Name: to.Ptr("Standard_D32ds_v4"),
	// 																							SupportedHAMode: []*string{
	// 																								to.Ptr("SameZone"),
	// 																								to.Ptr("ZoneRedundant")},
	// 																								SupportedIops: to.Ptr[int64](20000),
	// 																								SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																								SupportedZones: []*string{
	// 																									to.Ptr("1"),
	// 																									to.Ptr("2"),
	// 																									to.Ptr("3")},
	// 																									VCores: to.Ptr[int64](32),
	// 																								},
	// 																								{
	// 																									Name: to.Ptr("Standard_D48ds_v4"),
	// 																									SupportedHAMode: []*string{
	// 																										to.Ptr("SameZone"),
	// 																										to.Ptr("ZoneRedundant")},
	// 																										SupportedIops: to.Ptr[int64](20000),
	// 																										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																										SupportedZones: []*string{
	// 																											to.Ptr("1"),
	// 																											to.Ptr("2"),
	// 																											to.Ptr("3")},
	// 																											VCores: to.Ptr[int64](48),
	// 																										},
	// 																										{
	// 																											Name: to.Ptr("Standard_D64ds_v4"),
	// 																											SupportedHAMode: []*string{
	// 																												to.Ptr("SameZone"),
	// 																												to.Ptr("ZoneRedundant")},
	// 																												SupportedIops: to.Ptr[int64](20000),
	// 																												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																												SupportedZones: []*string{
	// 																													to.Ptr("1"),
	// 																													to.Ptr("2"),
	// 																													to.Ptr("3")},
	// 																													VCores: to.Ptr[int64](64),
	// 																												},
	// 																												{
	// 																													Name: to.Ptr("Standard_D2ds_v5"),
	// 																													SupportedHAMode: []*string{
	// 																														to.Ptr("SameZone"),
	// 																														to.Ptr("ZoneRedundant")},
	// 																														SupportedIops: to.Ptr[int64](3200),
	// 																														SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																														SupportedZones: []*string{
	// 																															to.Ptr("1"),
	// 																															to.Ptr("2"),
	// 																															to.Ptr("3")},
	// 																															VCores: to.Ptr[int64](2),
	// 																														},
	// 																														{
	// 																															Name: to.Ptr("Standard_D4ds_v5"),
	// 																															SupportedHAMode: []*string{
	// 																																to.Ptr("SameZone"),
	// 																																to.Ptr("ZoneRedundant")},
	// 																																SupportedIops: to.Ptr[int64](6400),
	// 																																SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																SupportedZones: []*string{
	// 																																	to.Ptr("1"),
	// 																																	to.Ptr("2"),
	// 																																	to.Ptr("3")},
	// 																																	VCores: to.Ptr[int64](4),
	// 																																},
	// 																																{
	// 																																	Name: to.Ptr("Standard_D8ds_v5"),
	// 																																	SupportedHAMode: []*string{
	// 																																		to.Ptr("SameZone"),
	// 																																		to.Ptr("ZoneRedundant")},
	// 																																		SupportedIops: to.Ptr[int64](12800),
	// 																																		SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																		SupportedZones: []*string{
	// 																																			to.Ptr("1"),
	// 																																			to.Ptr("2"),
	// 																																			to.Ptr("3")},
	// 																																			VCores: to.Ptr[int64](8),
	// 																																		},
	// 																																		{
	// 																																			Name: to.Ptr("Standard_D16ds_v5"),
	// 																																			SupportedHAMode: []*string{
	// 																																				to.Ptr("SameZone"),
	// 																																				to.Ptr("ZoneRedundant")},
	// 																																				SupportedIops: to.Ptr[int64](20000),
	// 																																				SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																				SupportedZones: []*string{
	// 																																					to.Ptr("1"),
	// 																																					to.Ptr("2"),
	// 																																					to.Ptr("3")},
	// 																																					VCores: to.Ptr[int64](16),
	// 																																				},
	// 																																				{
	// 																																					Name: to.Ptr("Standard_D32ds_v5"),
	// 																																					SupportedHAMode: []*string{
	// 																																						to.Ptr("SameZone"),
	// 																																						to.Ptr("ZoneRedundant")},
	// 																																						SupportedIops: to.Ptr[int64](20000),
	// 																																						SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																						SupportedZones: []*string{
	// 																																							to.Ptr("1"),
	// 																																							to.Ptr("2"),
	// 																																							to.Ptr("3")},
	// 																																							VCores: to.Ptr[int64](32),
	// 																																						},
	// 																																						{
	// 																																							Name: to.Ptr("Standard_D48ds_v5"),
	// 																																							SupportedHAMode: []*string{
	// 																																								to.Ptr("SameZone"),
	// 																																								to.Ptr("ZoneRedundant")},
	// 																																								SupportedIops: to.Ptr[int64](20000),
	// 																																								SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																								SupportedZones: []*string{
	// 																																									to.Ptr("1"),
	// 																																									to.Ptr("2"),
	// 																																									to.Ptr("3")},
	// 																																									VCores: to.Ptr[int64](48),
	// 																																								},
	// 																																								{
	// 																																									Name: to.Ptr("Standard_D64ds_v5"),
	// 																																									SupportedHAMode: []*string{
	// 																																										to.Ptr("SameZone"),
	// 																																										to.Ptr("ZoneRedundant")},
	// 																																										SupportedIops: to.Ptr[int64](20000),
	// 																																										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																										SupportedZones: []*string{
	// 																																											to.Ptr("1"),
	// 																																											to.Ptr("2"),
	// 																																											to.Ptr("3")},
	// 																																											VCores: to.Ptr[int64](64),
	// 																																										},
	// 																																										{
	// 																																											Name: to.Ptr("Standard_D96ds_v5"),
	// 																																											SupportedHAMode: []*string{
	// 																																												to.Ptr("SameZone"),
	// 																																												to.Ptr("ZoneRedundant")},
	// 																																												SupportedIops: to.Ptr[int64](20000),
	// 																																												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																												SupportedZones: []*string{
	// 																																													to.Ptr("1"),
	// 																																													to.Ptr("2"),
	// 																																													to.Ptr("3")},
	// 																																													VCores: to.Ptr[int64](64),
	// 																																												},
	// 																																												{
	// 																																													Name: to.Ptr("Standard_D2ads_v5"),
	// 																																													SupportedHAMode: []*string{
	// 																																														to.Ptr("SameZone"),
	// 																																														to.Ptr("ZoneRedundant")},
	// 																																														SupportedIops: to.Ptr[int64](3200),
	// 																																														SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																														SupportedZones: []*string{
	// 																																															to.Ptr("1"),
	// 																																															to.Ptr("2"),
	// 																																															to.Ptr("3")},
	// 																																															VCores: to.Ptr[int64](2),
	// 																																														},
	// 																																														{
	// 																																															Name: to.Ptr("Standard_D4ads_v5"),
	// 																																															SupportedHAMode: []*string{
	// 																																																to.Ptr("SameZone"),
	// 																																																to.Ptr("ZoneRedundant")},
	// 																																																SupportedIops: to.Ptr[int64](6400),
	// 																																																SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																SupportedZones: []*string{
	// 																																																	to.Ptr("1"),
	// 																																																	to.Ptr("2"),
	// 																																																	to.Ptr("3")},
	// 																																																	VCores: to.Ptr[int64](4),
	// 																																																},
	// 																																																{
	// 																																																	Name: to.Ptr("Standard_D8ads_v5"),
	// 																																																	SupportedHAMode: []*string{
	// 																																																		to.Ptr("SameZone"),
	// 																																																		to.Ptr("ZoneRedundant")},
	// 																																																		SupportedIops: to.Ptr[int64](12800),
	// 																																																		SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																		SupportedZones: []*string{
	// 																																																			to.Ptr("1"),
	// 																																																			to.Ptr("2"),
	// 																																																			to.Ptr("3")},
	// 																																																			VCores: to.Ptr[int64](8),
	// 																																																		},
	// 																																																		{
	// 																																																			Name: to.Ptr("Standard_D16ads_v5"),
	// 																																																			SupportedHAMode: []*string{
	// 																																																				to.Ptr("SameZone"),
	// 																																																				to.Ptr("ZoneRedundant")},
	// 																																																				SupportedIops: to.Ptr[int64](20000),
	// 																																																				SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																				SupportedZones: []*string{
	// 																																																					to.Ptr("1"),
	// 																																																					to.Ptr("2"),
	// 																																																					to.Ptr("3")},
	// 																																																					VCores: to.Ptr[int64](16),
	// 																																																				},
	// 																																																				{
	// 																																																					Name: to.Ptr("Standard_D32ads_v5"),
	// 																																																					SupportedHAMode: []*string{
	// 																																																						to.Ptr("SameZone"),
	// 																																																						to.Ptr("ZoneRedundant")},
	// 																																																						SupportedIops: to.Ptr[int64](20000),
	// 																																																						SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																						SupportedZones: []*string{
	// 																																																							to.Ptr("1"),
	// 																																																							to.Ptr("2"),
	// 																																																							to.Ptr("3")},
	// 																																																							VCores: to.Ptr[int64](32),
	// 																																																						},
	// 																																																						{
	// 																																																							Name: to.Ptr("Standard_D48ads_v5"),
	// 																																																							SupportedHAMode: []*string{
	// 																																																								to.Ptr("SameZone"),
	// 																																																								to.Ptr("ZoneRedundant")},
	// 																																																								SupportedIops: to.Ptr[int64](20000),
	// 																																																								SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																								SupportedZones: []*string{
	// 																																																									to.Ptr("1"),
	// 																																																									to.Ptr("2"),
	// 																																																									to.Ptr("3")},
	// 																																																									VCores: to.Ptr[int64](48),
	// 																																																								},
	// 																																																								{
	// 																																																									Name: to.Ptr("Standard_D64ads_v5"),
	// 																																																									SupportedHAMode: []*string{
	// 																																																										to.Ptr("SameZone"),
	// 																																																										to.Ptr("ZoneRedundant")},
	// 																																																										SupportedIops: to.Ptr[int64](20000),
	// 																																																										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																										SupportedZones: []*string{
	// 																																																											to.Ptr("1"),
	// 																																																											to.Ptr("2"),
	// 																																																											to.Ptr("3")},
	// 																																																											VCores: to.Ptr[int64](64),
	// 																																																										},
	// 																																																										{
	// 																																																											Name: to.Ptr("Standard_D96ads_v5"),
	// 																																																											SupportedHAMode: []*string{
	// 																																																												to.Ptr("SameZone"),
	// 																																																												to.Ptr("ZoneRedundant")},
	// 																																																												SupportedIops: to.Ptr[int64](20000),
	// 																																																												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																												SupportedZones: []*string{
	// 																																																													to.Ptr("1"),
	// 																																																													to.Ptr("2"),
	// 																																																													to.Ptr("3")},
	// 																																																													VCores: to.Ptr[int64](96),
	// 																																																												},
	// 																																																												{
	// 																																																													Name: to.Ptr("Standard_D2s_v3"),
	// 																																																													SupportedHAMode: []*string{
	// 																																																														to.Ptr("SameZone"),
	// 																																																														to.Ptr("ZoneRedundant")},
	// 																																																														SupportedIops: to.Ptr[int64](3200),
	// 																																																														SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																														SupportedZones: []*string{
	// 																																																															to.Ptr("1"),
	// 																																																															to.Ptr("2"),
	// 																																																															to.Ptr("3")},
	// 																																																															VCores: to.Ptr[int64](2),
	// 																																																														},
	// 																																																														{
	// 																																																															Name: to.Ptr("Standard_D4s_v3"),
	// 																																																															SupportedHAMode: []*string{
	// 																																																																to.Ptr("SameZone"),
	// 																																																																to.Ptr("ZoneRedundant")},
	// 																																																																SupportedIops: to.Ptr[int64](6400),
	// 																																																																SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																																SupportedZones: []*string{
	// 																																																																	to.Ptr("1"),
	// 																																																																	to.Ptr("2"),
	// 																																																																	to.Ptr("3")},
	// 																																																																	VCores: to.Ptr[int64](4),
	// 																																																																},
	// 																																																																{
	// 																																																																	Name: to.Ptr("Standard_D8s_v3"),
	// 																																																																	SupportedHAMode: []*string{
	// 																																																																		to.Ptr("SameZone"),
	// 																																																																		to.Ptr("ZoneRedundant")},
	// 																																																																		SupportedIops: to.Ptr[int64](12800),
	// 																																																																		SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																																		SupportedZones: []*string{
	// 																																																																			to.Ptr("1"),
	// 																																																																			to.Ptr("2"),
	// 																																																																			to.Ptr("3")},
	// 																																																																			VCores: to.Ptr[int64](8),
	// 																																																																		},
	// 																																																																		{
	// 																																																																			Name: to.Ptr("Standard_D16s_v3"),
	// 																																																																			SupportedHAMode: []*string{
	// 																																																																				to.Ptr("SameZone"),
	// 																																																																				to.Ptr("ZoneRedundant")},
	// 																																																																				SupportedIops: to.Ptr[int64](20000),
	// 																																																																				SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																																				SupportedZones: []*string{
	// 																																																																					to.Ptr("1"),
	// 																																																																					to.Ptr("2"),
	// 																																																																					to.Ptr("3")},
	// 																																																																					VCores: to.Ptr[int64](16),
	// 																																																																				},
	// 																																																																				{
	// 																																																																					Name: to.Ptr("Standard_D32s_v3"),
	// 																																																																					SupportedHAMode: []*string{
	// 																																																																						to.Ptr("SameZone"),
	// 																																																																						to.Ptr("ZoneRedundant")},
	// 																																																																						SupportedIops: to.Ptr[int64](20000),
	// 																																																																						SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																																						SupportedZones: []*string{
	// 																																																																							to.Ptr("1"),
	// 																																																																							to.Ptr("2"),
	// 																																																																							to.Ptr("3")},
	// 																																																																							VCores: to.Ptr[int64](32),
	// 																																																																						},
	// 																																																																						{
	// 																																																																							Name: to.Ptr("Standard_D48s_v3"),
	// 																																																																							SupportedHAMode: []*string{
	// 																																																																								to.Ptr("SameZone"),
	// 																																																																								to.Ptr("ZoneRedundant")},
	// 																																																																								SupportedIops: to.Ptr[int64](20000),
	// 																																																																								SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																																								SupportedZones: []*string{
	// 																																																																									to.Ptr("1"),
	// 																																																																									to.Ptr("2"),
	// 																																																																									to.Ptr("3")},
	// 																																																																									VCores: to.Ptr[int64](48),
	// 																																																																								},
	// 																																																																								{
	// 																																																																									Name: to.Ptr("Standard_D64s_v3"),
	// 																																																																									SupportedHAMode: []*string{
	// 																																																																										to.Ptr("SameZone"),
	// 																																																																										to.Ptr("ZoneRedundant")},
	// 																																																																										SupportedIops: to.Ptr[int64](20000),
	// 																																																																										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
	// 																																																																										SupportedZones: []*string{
	// 																																																																											to.Ptr("1"),
	// 																																																																											to.Ptr("2"),
	// 																																																																											to.Ptr("3")},
	// 																																																																											VCores: to.Ptr[int64](64),
	// 																																																																									}},
	// 																																																																									SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
	// 																																																																										{
	// 																																																																											Name: to.Ptr("Premium"),
	// 																																																																											MaxBackupIntervalHours: to.Ptr[int64](24),
	// 																																																																											MaxBackupRetentionDays: to.Ptr[int64](35),
	// 																																																																											MaxStorageSize: to.Ptr[int64](16777216),
	// 																																																																											MinBackupIntervalHours: to.Ptr[int64](1),
	// 																																																																											MinBackupRetentionDays: to.Ptr[int64](7),
	// 																																																																											MinStorageSize: to.Ptr[int64](20480),
	// 																																																																									}},
	// 																																																																								},
	// 																																																																								{
	// 																																																																									Name: to.Ptr("MemoryOptimized"),
	// 																																																																									DefaultSKU: to.Ptr("Standard_E2ds_v4"),
	// 																																																																									DefaultStorageSize: to.Ptr[int32](131072),
	// 																																																																									SupportedSKUs: []*armmysqlflexibleservers.SKUCapabilityV2{
	// 																																																																										{
	// 																																																																											Name: to.Ptr("Standard_E2ds_v4"),
	// 																																																																											SupportedHAMode: []*string{
	// 																																																																												to.Ptr("SameZone"),
	// 																																																																												to.Ptr("ZoneRedundant")},
	// 																																																																												SupportedIops: to.Ptr[int64](5000),
	// 																																																																												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																												SupportedZones: []*string{
	// 																																																																													to.Ptr("1"),
	// 																																																																													to.Ptr("2"),
	// 																																																																													to.Ptr("3")},
	// 																																																																													VCores: to.Ptr[int64](2),
	// 																																																																												},
	// 																																																																												{
	// 																																																																													Name: to.Ptr("Standard_E4ds_v4"),
	// 																																																																													SupportedHAMode: []*string{
	// 																																																																														to.Ptr("SameZone"),
	// 																																																																														to.Ptr("ZoneRedundant")},
	// 																																																																														SupportedIops: to.Ptr[int64](10000),
	// 																																																																														SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																														SupportedZones: []*string{
	// 																																																																															to.Ptr("1"),
	// 																																																																															to.Ptr("2"),
	// 																																																																															to.Ptr("3")},
	// 																																																																															VCores: to.Ptr[int64](4),
	// 																																																																														},
	// 																																																																														{
	// 																																																																															Name: to.Ptr("Standard_E8ds_v4"),
	// 																																																																															SupportedHAMode: []*string{
	// 																																																																																to.Ptr("SameZone"),
	// 																																																																																to.Ptr("ZoneRedundant")},
	// 																																																																																SupportedIops: to.Ptr[int64](18000),
	// 																																																																																SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																SupportedZones: []*string{
	// 																																																																																	to.Ptr("1"),
	// 																																																																																	to.Ptr("2"),
	// 																																																																																	to.Ptr("3")},
	// 																																																																																	VCores: to.Ptr[int64](8),
	// 																																																																																},
	// 																																																																																{
	// 																																																																																	Name: to.Ptr("Standard_E16ds_v4"),
	// 																																																																																	SupportedHAMode: []*string{
	// 																																																																																		to.Ptr("SameZone"),
	// 																																																																																		to.Ptr("ZoneRedundant")},
	// 																																																																																		SupportedIops: to.Ptr[int64](28000),
	// 																																																																																		SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																		SupportedZones: []*string{
	// 																																																																																			to.Ptr("1"),
	// 																																																																																			to.Ptr("2"),
	// 																																																																																			to.Ptr("3")},
	// 																																																																																			VCores: to.Ptr[int64](16),
	// 																																																																																		},
	// 																																																																																		{
	// 																																																																																			Name: to.Ptr("Standard_E32ds_v4"),
	// 																																																																																			SupportedHAMode: []*string{
	// 																																																																																				to.Ptr("SameZone"),
	// 																																																																																				to.Ptr("ZoneRedundant")},
	// 																																																																																				SupportedIops: to.Ptr[int64](38000),
	// 																																																																																				SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																				SupportedZones: []*string{
	// 																																																																																					to.Ptr("1"),
	// 																																																																																					to.Ptr("2"),
	// 																																																																																					to.Ptr("3")},
	// 																																																																																					VCores: to.Ptr[int64](32),
	// 																																																																																				},
	// 																																																																																				{
	// 																																																																																					Name: to.Ptr("Standard_E48ds_v4"),
	// 																																																																																					SupportedHAMode: []*string{
	// 																																																																																						to.Ptr("SameZone"),
	// 																																																																																						to.Ptr("ZoneRedundant")},
	// 																																																																																						SupportedIops: to.Ptr[int64](48000),
	// 																																																																																						SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																						SupportedZones: []*string{
	// 																																																																																							to.Ptr("1"),
	// 																																																																																							to.Ptr("2"),
	// 																																																																																							to.Ptr("3")},
	// 																																																																																							VCores: to.Ptr[int64](48),
	// 																																																																																						},
	// 																																																																																						{
	// 																																																																																							Name: to.Ptr("Standard_E64ds_v4"),
	// 																																																																																							SupportedHAMode: []*string{
	// 																																																																																								to.Ptr("SameZone"),
	// 																																																																																								to.Ptr("ZoneRedundant")},
	// 																																																																																								SupportedIops: to.Ptr[int64](64000),
	// 																																																																																								SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																								SupportedZones: []*string{
	// 																																																																																									to.Ptr("1"),
	// 																																																																																									to.Ptr("2"),
	// 																																																																																									to.Ptr("3")},
	// 																																																																																									VCores: to.Ptr[int64](64),
	// 																																																																																								},
	// 																																																																																								{
	// 																																																																																									Name: to.Ptr("Standard_E80ids_v4"),
	// 																																																																																									SupportedHAMode: []*string{
	// 																																																																																										to.Ptr("SameZone"),
	// 																																																																																										to.Ptr("ZoneRedundant")},
	// 																																																																																										SupportedIops: to.Ptr[int64](72000),
	// 																																																																																										SupportedMemoryPerVCoreMB: to.Ptr[int64](6451),
	// 																																																																																										SupportedZones: []*string{
	// 																																																																																											to.Ptr("1"),
	// 																																																																																											to.Ptr("2"),
	// 																																																																																											to.Ptr("3")},
	// 																																																																																											VCores: to.Ptr[int64](80),
	// 																																																																																										},
	// 																																																																																										{
	// 																																																																																											Name: to.Ptr("Standard_E2ds_v5"),
	// 																																																																																											SupportedHAMode: []*string{
	// 																																																																																												to.Ptr("SameZone"),
	// 																																																																																												to.Ptr("ZoneRedundant")},
	// 																																																																																												SupportedIops: to.Ptr[int64](5000),
	// 																																																																																												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																												SupportedZones: []*string{
	// 																																																																																													to.Ptr("1"),
	// 																																																																																													to.Ptr("2"),
	// 																																																																																													to.Ptr("3")},
	// 																																																																																													VCores: to.Ptr[int64](2),
	// 																																																																																												},
	// 																																																																																												{
	// 																																																																																													Name: to.Ptr("Standard_E4ds_v5"),
	// 																																																																																													SupportedHAMode: []*string{
	// 																																																																																														to.Ptr("SameZone"),
	// 																																																																																														to.Ptr("ZoneRedundant")},
	// 																																																																																														SupportedIops: to.Ptr[int64](10000),
	// 																																																																																														SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																														SupportedZones: []*string{
	// 																																																																																															to.Ptr("1"),
	// 																																																																																															to.Ptr("2"),
	// 																																																																																															to.Ptr("3")},
	// 																																																																																															VCores: to.Ptr[int64](4),
	// 																																																																																														},
	// 																																																																																														{
	// 																																																																																															Name: to.Ptr("Standard_E8ds_v5"),
	// 																																																																																															SupportedHAMode: []*string{
	// 																																																																																																to.Ptr("SameZone"),
	// 																																																																																																to.Ptr("ZoneRedundant")},
	// 																																																																																																SupportedIops: to.Ptr[int64](18000),
	// 																																																																																																SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																SupportedZones: []*string{
	// 																																																																																																	to.Ptr("1"),
	// 																																																																																																	to.Ptr("2"),
	// 																																																																																																	to.Ptr("3")},
	// 																																																																																																	VCores: to.Ptr[int64](8),
	// 																																																																																																},
	// 																																																																																																{
	// 																																																																																																	Name: to.Ptr("Standard_E16ds_v5"),
	// 																																																																																																	SupportedHAMode: []*string{
	// 																																																																																																		to.Ptr("SameZone"),
	// 																																																																																																		to.Ptr("ZoneRedundant")},
	// 																																																																																																		SupportedIops: to.Ptr[int64](28000),
	// 																																																																																																		SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																		SupportedZones: []*string{
	// 																																																																																																			to.Ptr("1"),
	// 																																																																																																			to.Ptr("2"),
	// 																																																																																																			to.Ptr("3")},
	// 																																																																																																			VCores: to.Ptr[int64](16),
	// 																																																																																																		},
	// 																																																																																																		{
	// 																																																																																																			Name: to.Ptr("Standard_E32ds_v5"),
	// 																																																																																																			SupportedHAMode: []*string{
	// 																																																																																																				to.Ptr("SameZone"),
	// 																																																																																																				to.Ptr("ZoneRedundant")},
	// 																																																																																																				SupportedIops: to.Ptr[int64](38000),
	// 																																																																																																				SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																				SupportedZones: []*string{
	// 																																																																																																					to.Ptr("1"),
	// 																																																																																																					to.Ptr("2"),
	// 																																																																																																					to.Ptr("3")},
	// 																																																																																																					VCores: to.Ptr[int64](32),
	// 																																																																																																				},
	// 																																																																																																				{
	// 																																																																																																					Name: to.Ptr("Standard_E48ds_v5"),
	// 																																																																																																					SupportedHAMode: []*string{
	// 																																																																																																						to.Ptr("SameZone"),
	// 																																																																																																						to.Ptr("ZoneRedundant")},
	// 																																																																																																						SupportedIops: to.Ptr[int64](48000),
	// 																																																																																																						SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																						SupportedZones: []*string{
	// 																																																																																																							to.Ptr("1"),
	// 																																																																																																							to.Ptr("2"),
	// 																																																																																																							to.Ptr("3")},
	// 																																																																																																							VCores: to.Ptr[int64](48),
	// 																																																																																																						},
	// 																																																																																																						{
	// 																																																																																																							Name: to.Ptr("Standard_E64ds_v5"),
	// 																																																																																																							SupportedHAMode: []*string{
	// 																																																																																																								to.Ptr("SameZone"),
	// 																																																																																																								to.Ptr("ZoneRedundant")},
	// 																																																																																																								SupportedIops: to.Ptr[int64](64000),
	// 																																																																																																								SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																								SupportedZones: []*string{
	// 																																																																																																									to.Ptr("1"),
	// 																																																																																																									to.Ptr("2"),
	// 																																																																																																									to.Ptr("3")},
	// 																																																																																																									VCores: to.Ptr[int64](64),
	// 																																																																																																								},
	// 																																																																																																								{
	// 																																																																																																									Name: to.Ptr("Standard_E96ds_v5"),
	// 																																																																																																									SupportedHAMode: []*string{
	// 																																																																																																										to.Ptr("SameZone"),
	// 																																																																																																										to.Ptr("ZoneRedundant")},
	// 																																																																																																										SupportedIops: to.Ptr[int64](80000),
	// 																																																																																																										SupportedMemoryPerVCoreMB: to.Ptr[int64](7168),
	// 																																																																																																										SupportedZones: []*string{
	// 																																																																																																											to.Ptr("1"),
	// 																																																																																																											to.Ptr("2"),
	// 																																																																																																											to.Ptr("3")},
	// 																																																																																																											VCores: to.Ptr[int64](96),
	// 																																																																																																										},
	// 																																																																																																										{
	// 																																																																																																											Name: to.Ptr("Standard_E2ads_v5"),
	// 																																																																																																											SupportedHAMode: []*string{
	// 																																																																																																												to.Ptr("SameZone"),
	// 																																																																																																												to.Ptr("ZoneRedundant")},
	// 																																																																																																												SupportedIops: to.Ptr[int64](5000),
	// 																																																																																																												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																												SupportedZones: []*string{
	// 																																																																																																													to.Ptr("1"),
	// 																																																																																																													to.Ptr("2"),
	// 																																																																																																													to.Ptr("3")},
	// 																																																																																																													VCores: to.Ptr[int64](2),
	// 																																																																																																												},
	// 																																																																																																												{
	// 																																																																																																													Name: to.Ptr("Standard_E4ads_v5"),
	// 																																																																																																													SupportedHAMode: []*string{
	// 																																																																																																														to.Ptr("SameZone"),
	// 																																																																																																														to.Ptr("ZoneRedundant")},
	// 																																																																																																														SupportedIops: to.Ptr[int64](10000),
	// 																																																																																																														SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																														SupportedZones: []*string{
	// 																																																																																																															to.Ptr("1"),
	// 																																																																																																															to.Ptr("2"),
	// 																																																																																																															to.Ptr("3")},
	// 																																																																																																															VCores: to.Ptr[int64](4),
	// 																																																																																																														},
	// 																																																																																																														{
	// 																																																																																																															Name: to.Ptr("Standard_E8ads_v5"),
	// 																																																																																																															SupportedHAMode: []*string{
	// 																																																																																																																to.Ptr("SameZone"),
	// 																																																																																																																to.Ptr("ZoneRedundant")},
	// 																																																																																																																SupportedIops: to.Ptr[int64](18000),
	// 																																																																																																																SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																SupportedZones: []*string{
	// 																																																																																																																	to.Ptr("1"),
	// 																																																																																																																	to.Ptr("2"),
	// 																																																																																																																	to.Ptr("3")},
	// 																																																																																																																	VCores: to.Ptr[int64](8),
	// 																																																																																																																},
	// 																																																																																																																{
	// 																																																																																																																	Name: to.Ptr("Standard_E16ads_v5"),
	// 																																																																																																																	SupportedHAMode: []*string{
	// 																																																																																																																		to.Ptr("SameZone"),
	// 																																																																																																																		to.Ptr("ZoneRedundant")},
	// 																																																																																																																		SupportedIops: to.Ptr[int64](28000),
	// 																																																																																																																		SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																		SupportedZones: []*string{
	// 																																																																																																																			to.Ptr("1"),
	// 																																																																																																																			to.Ptr("2"),
	// 																																																																																																																			to.Ptr("3")},
	// 																																																																																																																			VCores: to.Ptr[int64](16),
	// 																																																																																																																		},
	// 																																																																																																																		{
	// 																																																																																																																			Name: to.Ptr("Standard_E32ads_v5"),
	// 																																																																																																																			SupportedHAMode: []*string{
	// 																																																																																																																				to.Ptr("SameZone"),
	// 																																																																																																																				to.Ptr("ZoneRedundant")},
	// 																																																																																																																				SupportedIops: to.Ptr[int64](38000),
	// 																																																																																																																				SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																				SupportedZones: []*string{
	// 																																																																																																																					to.Ptr("1"),
	// 																																																																																																																					to.Ptr("2"),
	// 																																																																																																																					to.Ptr("3")},
	// 																																																																																																																					VCores: to.Ptr[int64](32),
	// 																																																																																																																				},
	// 																																																																																																																				{
	// 																																																																																																																					Name: to.Ptr("Standard_E48ads_v5"),
	// 																																																																																																																					SupportedHAMode: []*string{
	// 																																																																																																																						to.Ptr("SameZone"),
	// 																																																																																																																						to.Ptr("ZoneRedundant")},
	// 																																																																																																																						SupportedIops: to.Ptr[int64](48000),
	// 																																																																																																																						SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																						SupportedZones: []*string{
	// 																																																																																																																							to.Ptr("1"),
	// 																																																																																																																							to.Ptr("2"),
	// 																																																																																																																							to.Ptr("3")},
	// 																																																																																																																							VCores: to.Ptr[int64](48),
	// 																																																																																																																						},
	// 																																																																																																																						{
	// 																																																																																																																							Name: to.Ptr("Standard_E64ads_v5"),
	// 																																																																																																																							SupportedHAMode: []*string{
	// 																																																																																																																								to.Ptr("SameZone"),
	// 																																																																																																																								to.Ptr("ZoneRedundant")},
	// 																																																																																																																								SupportedIops: to.Ptr[int64](64000),
	// 																																																																																																																								SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																								SupportedZones: []*string{
	// 																																																																																																																									to.Ptr("1"),
	// 																																																																																																																									to.Ptr("2"),
	// 																																																																																																																									to.Ptr("3")},
	// 																																																																																																																									VCores: to.Ptr[int64](64),
	// 																																																																																																																								},
	// 																																																																																																																								{
	// 																																																																																																																									Name: to.Ptr("Standard_E96ads_v5"),
	// 																																																																																																																									SupportedHAMode: []*string{
	// 																																																																																																																										to.Ptr("SameZone"),
	// 																																																																																																																										to.Ptr("ZoneRedundant")},
	// 																																																																																																																										SupportedIops: to.Ptr[int64](80000),
	// 																																																																																																																										SupportedMemoryPerVCoreMB: to.Ptr[int64](7168),
	// 																																																																																																																										SupportedZones: []*string{
	// 																																																																																																																											to.Ptr("1"),
	// 																																																																																																																											to.Ptr("2"),
	// 																																																																																																																											to.Ptr("3")},
	// 																																																																																																																											VCores: to.Ptr[int64](96),
	// 																																																																																																																										},
	// 																																																																																																																										{
	// 																																																																																																																											Name: to.Ptr("Standard_E2s_v3"),
	// 																																																																																																																											SupportedHAMode: []*string{
	// 																																																																																																																												to.Ptr("SameZone"),
	// 																																																																																																																												to.Ptr("ZoneRedundant")},
	// 																																																																																																																												SupportedIops: to.Ptr[int64](5000),
	// 																																																																																																																												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																												SupportedZones: []*string{
	// 																																																																																																																													to.Ptr("1"),
	// 																																																																																																																													to.Ptr("2"),
	// 																																																																																																																													to.Ptr("3")},
	// 																																																																																																																													VCores: to.Ptr[int64](2),
	// 																																																																																																																												},
	// 																																																																																																																												{
	// 																																																																																																																													Name: to.Ptr("Standard_E4s_v3"),
	// 																																																																																																																													SupportedHAMode: []*string{
	// 																																																																																																																														to.Ptr("SameZone"),
	// 																																																																																																																														to.Ptr("ZoneRedundant")},
	// 																																																																																																																														SupportedIops: to.Ptr[int64](10000),
	// 																																																																																																																														SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																														SupportedZones: []*string{
	// 																																																																																																																															to.Ptr("1"),
	// 																																																																																																																															to.Ptr("2"),
	// 																																																																																																																															to.Ptr("3")},
	// 																																																																																																																															VCores: to.Ptr[int64](4),
	// 																																																																																																																														},
	// 																																																																																																																														{
	// 																																																																																																																															Name: to.Ptr("Standard_E8s_v3"),
	// 																																																																																																																															SupportedHAMode: []*string{
	// 																																																																																																																																to.Ptr("SameZone"),
	// 																																																																																																																																to.Ptr("ZoneRedundant")},
	// 																																																																																																																																SupportedIops: to.Ptr[int64](18000),
	// 																																																																																																																																SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																																SupportedZones: []*string{
	// 																																																																																																																																	to.Ptr("1"),
	// 																																																																																																																																	to.Ptr("2"),
	// 																																																																																																																																	to.Ptr("3")},
	// 																																																																																																																																	VCores: to.Ptr[int64](8),
	// 																																																																																																																																},
	// 																																																																																																																																{
	// 																																																																																																																																	Name: to.Ptr("Standard_E16s_v3"),
	// 																																																																																																																																	SupportedHAMode: []*string{
	// 																																																																																																																																		to.Ptr("SameZone"),
	// 																																																																																																																																		to.Ptr("ZoneRedundant")},
	// 																																																																																																																																		SupportedIops: to.Ptr[int64](28000),
	// 																																																																																																																																		SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																																		SupportedZones: []*string{
	// 																																																																																																																																			to.Ptr("1"),
	// 																																																																																																																																			to.Ptr("2"),
	// 																																																																																																																																			to.Ptr("3")},
	// 																																																																																																																																			VCores: to.Ptr[int64](16),
	// 																																																																																																																																		},
	// 																																																																																																																																		{
	// 																																																																																																																																			Name: to.Ptr("Standard_E32s_v3"),
	// 																																																																																																																																			SupportedHAMode: []*string{
	// 																																																																																																																																				to.Ptr("SameZone"),
	// 																																																																																																																																				to.Ptr("ZoneRedundant")},
	// 																																																																																																																																				SupportedIops: to.Ptr[int64](38000),
	// 																																																																																																																																				SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																																				SupportedZones: []*string{
	// 																																																																																																																																					to.Ptr("1"),
	// 																																																																																																																																					to.Ptr("2"),
	// 																																																																																																																																					to.Ptr("3")},
	// 																																																																																																																																					VCores: to.Ptr[int64](32),
	// 																																																																																																																																				},
	// 																																																																																																																																				{
	// 																																																																																																																																					Name: to.Ptr("Standard_E48s_v3"),
	// 																																																																																																																																					SupportedHAMode: []*string{
	// 																																																																																																																																						to.Ptr("SameZone"),
	// 																																																																																																																																						to.Ptr("ZoneRedundant")},
	// 																																																																																																																																						SupportedIops: to.Ptr[int64](20000),
	// 																																																																																																																																						SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																																						SupportedZones: []*string{
	// 																																																																																																																																							to.Ptr("1"),
	// 																																																																																																																																							to.Ptr("2"),
	// 																																																																																																																																							to.Ptr("3")},
	// 																																																																																																																																							VCores: to.Ptr[int64](48),
	// 																																																																																																																																						},
	// 																																																																																																																																						{
	// 																																																																																																																																							Name: to.Ptr("Standard_E64s_v3"),
	// 																																																																																																																																							SupportedHAMode: []*string{
	// 																																																																																																																																								to.Ptr("SameZone"),
	// 																																																																																																																																								to.Ptr("ZoneRedundant")},
	// 																																																																																																																																								SupportedIops: to.Ptr[int64](64000),
	// 																																																																																																																																								SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
	// 																																																																																																																																								SupportedZones: []*string{
	// 																																																																																																																																									to.Ptr("1"),
	// 																																																																																																																																									to.Ptr("2"),
	// 																																																																																																																																									to.Ptr("3")},
	// 																																																																																																																																									VCores: to.Ptr[int64](64),
	// 																																																																																																																																							}},
	// 																																																																																																																																							SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
	// 																																																																																																																																								{
	// 																																																																																																																																									Name: to.Ptr("Premium"),
	// 																																																																																																																																									MaxBackupIntervalHours: to.Ptr[int64](24),
	// 																																																																																																																																									MaxBackupRetentionDays: to.Ptr[int64](35),
	// 																																																																																																																																									MaxStorageSize: to.Ptr[int64](16777216),
	// 																																																																																																																																									MinBackupIntervalHours: to.Ptr[int64](1),
	// 																																																																																																																																									MinBackupRetentionDays: to.Ptr[int64](7),
	// 																																																																																																																																									MinStorageSize: to.Ptr[int64](20480),
	// 																																																																																																																																							}},
	// 																																																																																																																																					}},
	// 																																																																																																																																					SupportedGeoBackupRegions: []*string{
	// 																																																																																																																																						to.Ptr("southcentralus")},
	// 																																																																																																																																						SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapabilityV2{
	// 																																																																																																																																							{
	// 																																																																																																																																								Name: to.Ptr("5.7"),
	// 																																																																																																																																							},
	// 																																																																																																																																							{
	// 																																																																																																																																								Name: to.Ptr("8.0.21"),
	// 																																																																																																																																						}},
	// 																																																																																																																																					},
	// 																																																																																																																																				}
}
