package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2024-10-01-preview/examples/CapabilitiesByLocationList.json
func ExampleLocationBasedCapabilitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationBasedCapabilitiesClient().NewListPager("WestUS", nil)
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
		// page.CapabilitiesListResult = armmysqlflexibleservers.CapabilitiesListResult{
		// 	Value: []*armmysqlflexibleservers.CapabilityProperties{
		// 		{
		// 			SupportedFlexibleServerEditions: []*armmysqlflexibleservers.ServerEditionCapability{
		// 				{
		// 					Name: to.Ptr("Burstable"),
		// 					SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 						{
		// 							Name: to.Ptr("5.7"),
		// 							SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 								{
		// 									Name: to.Ptr("Standard_B1s"),
		// 									SupportedIops: to.Ptr[int64](400),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 									VCores: to.Ptr[int64](1),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_B1ms"),
		// 									SupportedIops: to.Ptr[int64](640),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 									VCores: to.Ptr[int64](1),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_B2s"),
		// 									SupportedIops: to.Ptr[int64](1280),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 									VCores: to.Ptr[int64](2),
		// 							}},
		// 						},
		// 						{
		// 							Name: to.Ptr("8.0.21"),
		// 							SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 								{
		// 									Name: to.Ptr("Standard_B1s"),
		// 									SupportedIops: to.Ptr[int64](400),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 									VCores: to.Ptr[int64](1),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_B1ms"),
		// 									SupportedIops: to.Ptr[int64](640),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 									VCores: to.Ptr[int64](1),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_B2s"),
		// 									SupportedIops: to.Ptr[int64](1280),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 									VCores: to.Ptr[int64](2),
		// 							}},
		// 					}},
		// 					SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 						{
		// 							Name: to.Ptr("Premium"),
		// 							MaxBackupRetentionDays: to.Ptr[int64](35),
		// 							MaxStorageSize: to.Ptr[int64](16777216),
		// 							MinBackupRetentionDays: to.Ptr[int64](7),
		// 							MinStorageSize: to.Ptr[int64](20480),
		// 					}},
		// 				},
		// 				{
		// 					Name: to.Ptr("GeneralPurpose"),
		// 					SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 						{
		// 							Name: to.Ptr("5.7"),
		// 							SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 								{
		// 									Name: to.Ptr("Standard_D2ds_v4"),
		// 									SupportedIops: to.Ptr[int64](3200),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](2),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D4ds_v4"),
		// 									SupportedIops: to.Ptr[int64](6400),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](4),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D8ds_v4"),
		// 									SupportedIops: to.Ptr[int64](12800),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](8),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D16ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](16),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D32ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](32),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D48ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](48),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D64ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](64),
		// 							}},
		// 						},
		// 						{
		// 							Name: to.Ptr("8.0.21"),
		// 							SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 								{
		// 									Name: to.Ptr("Standard_D2ds_v4"),
		// 									SupportedIops: to.Ptr[int64](3200),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](2),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D4ds_v4"),
		// 									SupportedIops: to.Ptr[int64](6400),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](4),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D8ds_v4"),
		// 									SupportedIops: to.Ptr[int64](12800),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](8),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D16ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](16),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D32ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](32),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D48ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](48),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_D64ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 									VCores: to.Ptr[int64](64),
		// 							}},
		// 					}},
		// 					SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 						{
		// 							Name: to.Ptr("Premium"),
		// 							MaxBackupRetentionDays: to.Ptr[int64](35),
		// 							MaxStorageSize: to.Ptr[int64](16777216),
		// 							MinBackupRetentionDays: to.Ptr[int64](7),
		// 							MinStorageSize: to.Ptr[int64](20480),
		// 					}},
		// 				},
		// 				{
		// 					Name: to.Ptr("MemoryOptimized"),
		// 					SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 						{
		// 							Name: to.Ptr("5.7"),
		// 							SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 								{
		// 									Name: to.Ptr("Standard_E2ds_v4"),
		// 									SupportedIops: to.Ptr[int64](3200),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](2),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E4ds_v4"),
		// 									SupportedIops: to.Ptr[int64](6400),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](4),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E8ds_v4"),
		// 									SupportedIops: to.Ptr[int64](12800),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](8),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E16ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](16),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E32ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](32),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E48ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](48),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E64ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](64),
		// 							}},
		// 						},
		// 						{
		// 							Name: to.Ptr("8.0.21"),
		// 							SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 								{
		// 									Name: to.Ptr("Standard_E2ds_v4"),
		// 									SupportedIops: to.Ptr[int64](3200),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](2),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E4ds_v4"),
		// 									SupportedIops: to.Ptr[int64](6400),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](4),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E8ds_v4"),
		// 									SupportedIops: to.Ptr[int64](12800),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](8),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E16ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](16),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E32ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](32),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E48ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](48),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_E64ds_v4"),
		// 									SupportedIops: to.Ptr[int64](20000),
		// 									SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 									VCores: to.Ptr[int64](64),
		// 							}},
		// 					}},
		// 					SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 						{
		// 							Name: to.Ptr("Premium"),
		// 							MaxBackupRetentionDays: to.Ptr[int64](35),
		// 							MaxStorageSize: to.Ptr[int64](16777216),
		// 							MinBackupRetentionDays: to.Ptr[int64](7),
		// 							MinStorageSize: to.Ptr[int64](20480),
		// 					}},
		// 			}},
		// 			SupportedGeoBackupRegions: []*string{
		// 			},
		// 			SupportedHAMode: []*string{
		// 				to.Ptr("SameZone"),
		// 				to.Ptr("ZoneRedundant")},
		// 				Zone: to.Ptr("none"),
		// 			},
		// 			{
		// 				SupportedFlexibleServerEditions: []*armmysqlflexibleservers.ServerEditionCapability{
		// 					{
		// 						Name: to.Ptr("Burstable"),
		// 						SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 							{
		// 								Name: to.Ptr("5.7"),
		// 								SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 									{
		// 										Name: to.Ptr("Standard_B1s"),
		// 										SupportedIops: to.Ptr[int64](400),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 										VCores: to.Ptr[int64](1),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_B1ms"),
		// 										SupportedIops: to.Ptr[int64](640),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 										VCores: to.Ptr[int64](1),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_B2s"),
		// 										SupportedIops: to.Ptr[int64](1280),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 										VCores: to.Ptr[int64](2),
		// 								}},
		// 							},
		// 							{
		// 								Name: to.Ptr("8.0.21"),
		// 								SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 									{
		// 										Name: to.Ptr("Standard_B1s"),
		// 										SupportedIops: to.Ptr[int64](400),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 										VCores: to.Ptr[int64](1),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_B1ms"),
		// 										SupportedIops: to.Ptr[int64](640),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 										VCores: to.Ptr[int64](1),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_B2s"),
		// 										SupportedIops: to.Ptr[int64](1280),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 										VCores: to.Ptr[int64](2),
		// 								}},
		// 						}},
		// 						SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 							{
		// 								Name: to.Ptr("Premium"),
		// 								MaxBackupRetentionDays: to.Ptr[int64](35),
		// 								MaxStorageSize: to.Ptr[int64](16777216),
		// 								MinBackupRetentionDays: to.Ptr[int64](7),
		// 								MinStorageSize: to.Ptr[int64](20480),
		// 						}},
		// 					},
		// 					{
		// 						Name: to.Ptr("GeneralPurpose"),
		// 						SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 							{
		// 								Name: to.Ptr("5.7"),
		// 								SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 									{
		// 										Name: to.Ptr("Standard_D2ds_v4"),
		// 										SupportedIops: to.Ptr[int64](3200),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](2),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D4ds_v4"),
		// 										SupportedIops: to.Ptr[int64](6400),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](4),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D8ds_v4"),
		// 										SupportedIops: to.Ptr[int64](12800),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](8),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D16ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](16),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D32ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](32),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D48ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](48),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D64ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](64),
		// 								}},
		// 							},
		// 							{
		// 								Name: to.Ptr("8.0.21"),
		// 								SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 									{
		// 										Name: to.Ptr("Standard_D2ds_v4"),
		// 										SupportedIops: to.Ptr[int64](3200),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](2),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D4ds_v4"),
		// 										SupportedIops: to.Ptr[int64](6400),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](4),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D8ds_v4"),
		// 										SupportedIops: to.Ptr[int64](12800),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](8),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D16ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](16),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D32ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](32),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D48ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](48),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_D64ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 										VCores: to.Ptr[int64](64),
		// 								}},
		// 						}},
		// 						SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 							{
		// 								Name: to.Ptr("Premium"),
		// 								MaxBackupRetentionDays: to.Ptr[int64](35),
		// 								MaxStorageSize: to.Ptr[int64](16777216),
		// 								MinBackupRetentionDays: to.Ptr[int64](7),
		// 								MinStorageSize: to.Ptr[int64](20480),
		// 						}},
		// 					},
		// 					{
		// 						Name: to.Ptr("MemoryOptimized"),
		// 						SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 							{
		// 								Name: to.Ptr("5.7"),
		// 								SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 									{
		// 										Name: to.Ptr("Standard_E2ds_v4"),
		// 										SupportedIops: to.Ptr[int64](3200),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](2),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E4ds_v4"),
		// 										SupportedIops: to.Ptr[int64](6400),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](4),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E8ds_v4"),
		// 										SupportedIops: to.Ptr[int64](12800),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](8),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E16ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](16),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E32ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](32),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E48ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](48),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E64ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](64),
		// 								}},
		// 							},
		// 							{
		// 								Name: to.Ptr("8.0.21"),
		// 								SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 									{
		// 										Name: to.Ptr("Standard_E2ds_v4"),
		// 										SupportedIops: to.Ptr[int64](3200),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](2),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E4ds_v4"),
		// 										SupportedIops: to.Ptr[int64](6400),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](4),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E8ds_v4"),
		// 										SupportedIops: to.Ptr[int64](12800),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](8),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E16ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](16),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E32ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](32),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E48ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](48),
		// 									},
		// 									{
		// 										Name: to.Ptr("Standard_E64ds_v4"),
		// 										SupportedIops: to.Ptr[int64](20000),
		// 										SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 										VCores: to.Ptr[int64](64),
		// 								}},
		// 						}},
		// 						SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 							{
		// 								Name: to.Ptr("Premium"),
		// 								MaxBackupRetentionDays: to.Ptr[int64](35),
		// 								MaxStorageSize: to.Ptr[int64](16777216),
		// 								MinBackupRetentionDays: to.Ptr[int64](7),
		// 								MinStorageSize: to.Ptr[int64](20480),
		// 						}},
		// 				}},
		// 				SupportedGeoBackupRegions: []*string{
		// 				},
		// 				SupportedHAMode: []*string{
		// 					to.Ptr("SameZone"),
		// 					to.Ptr("ZoneRedundant")},
		// 					Zone: to.Ptr("1"),
		// 				},
		// 				{
		// 					SupportedFlexibleServerEditions: []*armmysqlflexibleservers.ServerEditionCapability{
		// 						{
		// 							Name: to.Ptr("Burstable"),
		// 							SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 								{
		// 									Name: to.Ptr("5.7"),
		// 									SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 										{
		// 											Name: to.Ptr("Standard_B1s"),
		// 											SupportedIops: to.Ptr[int64](400),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 											VCores: to.Ptr[int64](1),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_B1ms"),
		// 											SupportedIops: to.Ptr[int64](640),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 											VCores: to.Ptr[int64](1),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_B2s"),
		// 											SupportedIops: to.Ptr[int64](1280),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 											VCores: to.Ptr[int64](2),
		// 									}},
		// 								},
		// 								{
		// 									Name: to.Ptr("8.0.21"),
		// 									SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 										{
		// 											Name: to.Ptr("Standard_B1s"),
		// 											SupportedIops: to.Ptr[int64](400),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 											VCores: to.Ptr[int64](1),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_B1ms"),
		// 											SupportedIops: to.Ptr[int64](640),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 											VCores: to.Ptr[int64](1),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_B2s"),
		// 											SupportedIops: to.Ptr[int64](1280),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 											VCores: to.Ptr[int64](2),
		// 									}},
		// 							}},
		// 							SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 								{
		// 									Name: to.Ptr("Premium"),
		// 									MaxBackupRetentionDays: to.Ptr[int64](35),
		// 									MaxStorageSize: to.Ptr[int64](16777216),
		// 									MinBackupRetentionDays: to.Ptr[int64](7),
		// 									MinStorageSize: to.Ptr[int64](20480),
		// 							}},
		// 						},
		// 						{
		// 							Name: to.Ptr("GeneralPurpose"),
		// 							SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 								{
		// 									Name: to.Ptr("5.7"),
		// 									SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 										{
		// 											Name: to.Ptr("Standard_D2ds_v4"),
		// 											SupportedIops: to.Ptr[int64](3200),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](2),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D4ds_v4"),
		// 											SupportedIops: to.Ptr[int64](6400),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](4),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D8ds_v4"),
		// 											SupportedIops: to.Ptr[int64](12800),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](8),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D16ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](16),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D32ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](32),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D48ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](48),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D64ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](64),
		// 									}},
		// 								},
		// 								{
		// 									Name: to.Ptr("8.0.21"),
		// 									SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 										{
		// 											Name: to.Ptr("Standard_D2ds_v4"),
		// 											SupportedIops: to.Ptr[int64](3200),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](2),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D4ds_v4"),
		// 											SupportedIops: to.Ptr[int64](6400),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](4),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D8ds_v4"),
		// 											SupportedIops: to.Ptr[int64](12800),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](8),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D16ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](16),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D32ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](32),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D48ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](48),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_D64ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 											VCores: to.Ptr[int64](64),
		// 									}},
		// 							}},
		// 							SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 								{
		// 									Name: to.Ptr("Premium"),
		// 									MaxBackupRetentionDays: to.Ptr[int64](35),
		// 									MaxStorageSize: to.Ptr[int64](16777216),
		// 									MinBackupRetentionDays: to.Ptr[int64](7),
		// 									MinStorageSize: to.Ptr[int64](20480),
		// 							}},
		// 						},
		// 						{
		// 							Name: to.Ptr("MemoryOptimized"),
		// 							SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 								{
		// 									Name: to.Ptr("5.7"),
		// 									SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 										{
		// 											Name: to.Ptr("Standard_E2ds_v4"),
		// 											SupportedIops: to.Ptr[int64](3200),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](2),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E4ds_v4"),
		// 											SupportedIops: to.Ptr[int64](6400),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](4),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E8ds_v4"),
		// 											SupportedIops: to.Ptr[int64](12800),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](8),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E16ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](16),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E32ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](32),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E48ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](48),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E64ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](64),
		// 									}},
		// 								},
		// 								{
		// 									Name: to.Ptr("8.0.21"),
		// 									SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 										{
		// 											Name: to.Ptr("Standard_E2ds_v4"),
		// 											SupportedIops: to.Ptr[int64](3200),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](2),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E4ds_v4"),
		// 											SupportedIops: to.Ptr[int64](6400),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](4),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E8ds_v4"),
		// 											SupportedIops: to.Ptr[int64](12800),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](8),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E16ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](16),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E32ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](32),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E48ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](48),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_E64ds_v4"),
		// 											SupportedIops: to.Ptr[int64](20000),
		// 											SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 											VCores: to.Ptr[int64](64),
		// 									}},
		// 							}},
		// 							SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 								{
		// 									Name: to.Ptr("Premium"),
		// 									MaxBackupRetentionDays: to.Ptr[int64](35),
		// 									MaxStorageSize: to.Ptr[int64](16777216),
		// 									MinBackupRetentionDays: to.Ptr[int64](7),
		// 									MinStorageSize: to.Ptr[int64](20480),
		// 							}},
		// 					}},
		// 					SupportedGeoBackupRegions: []*string{
		// 					},
		// 					SupportedHAMode: []*string{
		// 						to.Ptr("SameZone"),
		// 						to.Ptr("ZoneRedundant")},
		// 						Zone: to.Ptr("2"),
		// 					},
		// 					{
		// 						SupportedFlexibleServerEditions: []*armmysqlflexibleservers.ServerEditionCapability{
		// 							{
		// 								Name: to.Ptr("Burstable"),
		// 								SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 									{
		// 										Name: to.Ptr("5.7"),
		// 										SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 											{
		// 												Name: to.Ptr("Standard_B1s"),
		// 												SupportedIops: to.Ptr[int64](400),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 												VCores: to.Ptr[int64](1),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_B1ms"),
		// 												SupportedIops: to.Ptr[int64](640),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 												VCores: to.Ptr[int64](1),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_B2s"),
		// 												SupportedIops: to.Ptr[int64](1280),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 												VCores: to.Ptr[int64](2),
		// 										}},
		// 									},
		// 									{
		// 										Name: to.Ptr("8.0.21"),
		// 										SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 											{
		// 												Name: to.Ptr("Standard_B1s"),
		// 												SupportedIops: to.Ptr[int64](400),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](1024),
		// 												VCores: to.Ptr[int64](1),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_B1ms"),
		// 												SupportedIops: to.Ptr[int64](640),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 												VCores: to.Ptr[int64](1),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_B2s"),
		// 												SupportedIops: to.Ptr[int64](1280),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](2048),
		// 												VCores: to.Ptr[int64](2),
		// 										}},
		// 								}},
		// 								SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 									{
		// 										Name: to.Ptr("Premium"),
		// 										MaxBackupRetentionDays: to.Ptr[int64](35),
		// 										MaxStorageSize: to.Ptr[int64](16777216),
		// 										MinBackupRetentionDays: to.Ptr[int64](7),
		// 										MinStorageSize: to.Ptr[int64](20480),
		// 								}},
		// 							},
		// 							{
		// 								Name: to.Ptr("GeneralPurpose"),
		// 								SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 									{
		// 										Name: to.Ptr("5.7"),
		// 										SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 											{
		// 												Name: to.Ptr("Standard_D2ds_v4"),
		// 												SupportedIops: to.Ptr[int64](3200),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](2),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D4ds_v4"),
		// 												SupportedIops: to.Ptr[int64](6400),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](4),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D8ds_v4"),
		// 												SupportedIops: to.Ptr[int64](12800),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](8),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D16ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](16),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D32ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](32),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D48ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](48),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D64ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](64),
		// 										}},
		// 									},
		// 									{
		// 										Name: to.Ptr("8.0.21"),
		// 										SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 											{
		// 												Name: to.Ptr("Standard_D2ds_v4"),
		// 												SupportedIops: to.Ptr[int64](3200),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](2),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D4ds_v4"),
		// 												SupportedIops: to.Ptr[int64](6400),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](4),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D8ds_v4"),
		// 												SupportedIops: to.Ptr[int64](12800),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](8),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D16ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](16),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D32ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](32),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D48ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](48),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_D64ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](4096),
		// 												VCores: to.Ptr[int64](64),
		// 										}},
		// 								}},
		// 								SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 									{
		// 										Name: to.Ptr("Premium"),
		// 										MaxBackupRetentionDays: to.Ptr[int64](35),
		// 										MaxStorageSize: to.Ptr[int64](16777216),
		// 										MinBackupRetentionDays: to.Ptr[int64](7),
		// 										MinStorageSize: to.Ptr[int64](20480),
		// 								}},
		// 							},
		// 							{
		// 								Name: to.Ptr("MemoryOptimized"),
		// 								SupportedServerVersions: []*armmysqlflexibleservers.ServerVersionCapability{
		// 									{
		// 										Name: to.Ptr("5.7"),
		// 										SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 											{
		// 												Name: to.Ptr("Standard_E2ds_v4"),
		// 												SupportedIops: to.Ptr[int64](3200),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](2),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E4ds_v4"),
		// 												SupportedIops: to.Ptr[int64](6400),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](4),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E8ds_v4"),
		// 												SupportedIops: to.Ptr[int64](12800),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](8),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E16ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](16),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E32ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](32),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E48ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](48),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E64ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](64),
		// 										}},
		// 									},
		// 									{
		// 										Name: to.Ptr("8.0.21"),
		// 										SupportedSKUs: []*armmysqlflexibleservers.SKUCapability{
		// 											{
		// 												Name: to.Ptr("Standard_E2ds_v4"),
		// 												SupportedIops: to.Ptr[int64](3200),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](2),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E4ds_v4"),
		// 												SupportedIops: to.Ptr[int64](6400),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](4),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E8ds_v4"),
		// 												SupportedIops: to.Ptr[int64](12800),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](8),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E16ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](16),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E32ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](32),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E48ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](48),
		// 											},
		// 											{
		// 												Name: to.Ptr("Standard_E64ds_v4"),
		// 												SupportedIops: to.Ptr[int64](20000),
		// 												SupportedMemoryPerVCoreMB: to.Ptr[int64](8192),
		// 												VCores: to.Ptr[int64](64),
		// 										}},
		// 								}},
		// 								SupportedStorageEditions: []*armmysqlflexibleservers.StorageEditionCapability{
		// 									{
		// 										Name: to.Ptr("Premium"),
		// 										MaxBackupRetentionDays: to.Ptr[int64](35),
		// 										MaxStorageSize: to.Ptr[int64](16777216),
		// 										MinBackupRetentionDays: to.Ptr[int64](7),
		// 										MinStorageSize: to.Ptr[int64](20480),
		// 								}},
		// 						}},
		// 						SupportedGeoBackupRegions: []*string{
		// 						},
		// 						SupportedHAMode: []*string{
		// 							to.Ptr("SameZone"),
		// 							to.Ptr("ZoneRedundant")},
		// 							Zone: to.Ptr("3"),
		// 					}},
		// 				}
	}
}
