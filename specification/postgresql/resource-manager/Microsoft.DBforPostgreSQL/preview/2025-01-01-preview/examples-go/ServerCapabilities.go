package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/ServerCapabilities.json
func ExampleServerCapabilitiesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerCapabilitiesClient().NewListPager("testrg", "pgtestsvc4", nil)
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
		// page.CapabilitiesListResult = armpostgresqlflexibleservers.CapabilitiesListResult{
		// 	Value: []*armpostgresqlflexibleservers.FlexibleServerCapability{
		// 		{
		// 			Name: to.Ptr("FlexibleServerCapabilities"),
		// 			FastProvisioningSupported: to.Ptr(armpostgresqlflexibleservers.FastProvisioningSupportedEnumEnabled),
		// 			GeoBackupSupported: to.Ptr(armpostgresqlflexibleservers.GeoBackupSupportedEnumEnabled),
		// 			OnlineResizeSupported: to.Ptr(armpostgresqlflexibleservers.OnlineResizeSupportedEnumEnabled),
		// 			Restricted: to.Ptr(armpostgresqlflexibleservers.RestrictedEnumDisabled),
		// 			StorageAutoGrowthSupported: to.Ptr(armpostgresqlflexibleservers.StorageAutoGrowthSupportedEnumEnabled),
		// 			SupportedFastProvisioningEditions: []*armpostgresqlflexibleservers.FastProvisioningEditionCapability{
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("12"),
		// 					SupportedSKU: to.Ptr("standard_b1ms"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("12"),
		// 					SupportedSKU: to.Ptr("standard_b2s"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("12"),
		// 					SupportedSKU: to.Ptr("standard_d2s_v3"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("12"),
		// 					SupportedSKU: to.Ptr("standard_d2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("12"),
		// 					SupportedSKU: to.Ptr("standard_e2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](512),
		// 					SupportedTier: to.Ptr("MemoryOptimized"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("13"),
		// 					SupportedSKU: to.Ptr("standard_b1ms"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("13"),
		// 					SupportedSKU: to.Ptr("standard_b2s"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("13"),
		// 					SupportedSKU: to.Ptr("standard_d2s_v3"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("13"),
		// 					SupportedSKU: to.Ptr("standard_d2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("13"),
		// 					SupportedSKU: to.Ptr("standard_e2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](512),
		// 					SupportedTier: to.Ptr("MemoryOptimized"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("14"),
		// 					SupportedSKU: to.Ptr("standard_b1ms"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("14"),
		// 					SupportedSKU: to.Ptr("standard_b2s"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("14"),
		// 					SupportedSKU: to.Ptr("standard_d2s_v3"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("14"),
		// 					SupportedSKU: to.Ptr("standard_d2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("14"),
		// 					SupportedSKU: to.Ptr("standard_e2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](512),
		// 					SupportedTier: to.Ptr("MemoryOptimized"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("15"),
		// 					SupportedSKU: to.Ptr("standard_b1ms"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("15"),
		// 					SupportedSKU: to.Ptr("standard_b2s"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("15"),
		// 					SupportedSKU: to.Ptr("standard_d2s_v3"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("15"),
		// 					SupportedSKU: to.Ptr("standard_d2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("15"),
		// 					SupportedSKU: to.Ptr("standard_e2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](512),
		// 					SupportedTier: to.Ptr("MemoryOptimized"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("16"),
		// 					SupportedSKU: to.Ptr("standard_b1ms"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("16"),
		// 					SupportedSKU: to.Ptr("standard_b2s"),
		// 					SupportedStorageGb: to.Ptr[int32](32),
		// 					SupportedTier: to.Ptr("Burstable"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("16"),
		// 					SupportedSKU: to.Ptr("standard_d2s_v3"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("16"),
		// 					SupportedSKU: to.Ptr("standard_d2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](128),
		// 					SupportedTier: to.Ptr("GeneralPurpose"),
		// 				},
		// 				{
		// 					ServerCount: to.Ptr[int32](0),
		// 					SupportedServerVersions: to.Ptr("16"),
		// 					SupportedSKU: to.Ptr("standard_e2ds_v4"),
		// 					SupportedStorageGb: to.Ptr[int32](512),
		// 					SupportedTier: to.Ptr("MemoryOptimized"),
		// 			}},
		// 			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 				{
		// 					Name: to.Ptr("FastProvisioning"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("ZoneRedundantHa"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("GeoBackup"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("ZoneRedundantHaAndGeoBackup"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("StorageAutoGrowth"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("OnlineResize"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("OfferRestricted"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumDisabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("IndexTuning"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 				},
		// 				{
		// 					Name: to.Ptr("Clusters"),
		// 					Status: to.Ptr(armpostgresqlflexibleservers.SupportedFeatureStatusEnumEnabled),
		// 			}},
		// 			SupportedServerEditions: []*armpostgresqlflexibleservers.FlexibleServerEditionCapability{
		// 				{
		// 					Name: to.Ptr("Burstable"),
		// 					DefaultSKUName: to.Ptr("Standard_B1ms"),
		// 					SupportedServerSKUs: []*armpostgresqlflexibleservers.ServerSKUCapability{
		// 						{
		// 							Name: to.Ptr("Standard_B1ms"),
		// 							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 							},
		// 							SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 								to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 								to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 								SupportedIops: to.Ptr[int32](640),
		// 								SupportedMemoryPerVcoreMb: to.Ptr[int64](2048),
		// 								SupportedZones: []*string{
		// 									to.Ptr("1"),
		// 									to.Ptr("2"),
		// 									to.Ptr("3")},
		// 									VCores: to.Ptr[int32](1),
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard_B2s"),
		// 									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 									},
		// 									SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 										to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 										to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 										SupportedIops: to.Ptr[int32](1280),
		// 										SupportedMemoryPerVcoreMb: to.Ptr[int64](2048),
		// 										SupportedZones: []*string{
		// 											to.Ptr("1"),
		// 											to.Ptr("2"),
		// 											to.Ptr("3")},
		// 											VCores: to.Ptr[int32](2),
		// 										},
		// 										{
		// 											Name: to.Ptr("Standard_B2ms"),
		// 											SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 											},
		// 											SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 												to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 												to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 												SupportedIops: to.Ptr[int32](1920),
		// 												SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 												SupportedZones: []*string{
		// 													to.Ptr("1"),
		// 													to.Ptr("2"),
		// 													to.Ptr("3")},
		// 													VCores: to.Ptr[int32](2),
		// 												},
		// 												{
		// 													Name: to.Ptr("Standard_B4ms"),
		// 													SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 													},
		// 													SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 														to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 														to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 														SupportedIops: to.Ptr[int32](2880),
		// 														SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 														SupportedZones: []*string{
		// 															to.Ptr("1"),
		// 															to.Ptr("2"),
		// 															to.Ptr("3")},
		// 															VCores: to.Ptr[int32](4),
		// 														},
		// 														{
		// 															Name: to.Ptr("Standard_B8ms"),
		// 															SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 															},
		// 															SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																SupportedIops: to.Ptr[int32](4320),
		// 																SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																SupportedZones: []*string{
		// 																	to.Ptr("1"),
		// 																	to.Ptr("2"),
		// 																	to.Ptr("3")},
		// 																	VCores: to.Ptr[int32](8),
		// 																},
		// 																{
		// 																	Name: to.Ptr("Standard_B12ms"),
		// 																	SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																	},
		// 																	SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																		to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																		to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																		SupportedIops: to.Ptr[int32](4320),
		// 																		SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																		SupportedZones: []*string{
		// 																			to.Ptr("1"),
		// 																			to.Ptr("2"),
		// 																			to.Ptr("3")},
		// 																			VCores: to.Ptr[int32](12),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("Standard_B16ms"),
		// 																			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																			},
		// 																			SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																				to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																				to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																				SupportedIops: to.Ptr[int32](4320),
		// 																				SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																				SupportedZones: []*string{
		// 																					to.Ptr("1"),
		// 																					to.Ptr("2"),
		// 																					to.Ptr("3")},
		// 																					VCores: to.Ptr[int32](16),
		// 																				},
		// 																				{
		// 																					Name: to.Ptr("Standard_B20ms"),
		// 																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																					},
		// 																					SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																						to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																						to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																						SupportedIops: to.Ptr[int32](4320),
		// 																						SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																						SupportedZones: []*string{
		// 																							to.Ptr("1"),
		// 																							to.Ptr("2"),
		// 																							to.Ptr("3")},
		// 																							VCores: to.Ptr[int32](20),
		// 																					}},
		// 																					SupportedStorageEditions: []*armpostgresqlflexibleservers.StorageEditionCapability{
		// 																						{
		// 																							Name: to.Ptr("ManagedDisk"),
		// 																							DefaultStorageSizeMb: to.Ptr[int64](32768),
		// 																							SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P4"),
		// 																									StorageSizeMb: to.Ptr[int64](32768),
		// 																									SupportedIops: to.Ptr[int32](120),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P4"),
		// 																											Iops: to.Ptr[int32](120),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P6"),
		// 																											Iops: to.Ptr[int32](240),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P10"),
		// 																											Iops: to.Ptr[int32](500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P15"),
		// 																											Iops: to.Ptr[int32](1100),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P20"),
		// 																											Iops: to.Ptr[int32](2300),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P30"),
		// 																											Iops: to.Ptr[int32](5000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P40"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P6"),
		// 																									StorageSizeMb: to.Ptr[int64](65536),
		// 																									SupportedIops: to.Ptr[int32](240),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P6"),
		// 																											Iops: to.Ptr[int32](240),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P10"),
		// 																											Iops: to.Ptr[int32](500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P15"),
		// 																											Iops: to.Ptr[int32](1100),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P20"),
		// 																											Iops: to.Ptr[int32](2300),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P30"),
		// 																											Iops: to.Ptr[int32](5000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P40"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P10"),
		// 																									StorageSizeMb: to.Ptr[int64](131072),
		// 																									SupportedIops: to.Ptr[int32](500),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P10"),
		// 																											Iops: to.Ptr[int32](500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P15"),
		// 																											Iops: to.Ptr[int32](1100),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P20"),
		// 																											Iops: to.Ptr[int32](2300),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P30"),
		// 																											Iops: to.Ptr[int32](5000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P40"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P15"),
		// 																									StorageSizeMb: to.Ptr[int64](262144),
		// 																									SupportedIops: to.Ptr[int32](1100),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P15"),
		// 																											Iops: to.Ptr[int32](1100),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P20"),
		// 																											Iops: to.Ptr[int32](2300),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P30"),
		// 																											Iops: to.Ptr[int32](5000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P40"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P20"),
		// 																									StorageSizeMb: to.Ptr[int64](524288),
		// 																									SupportedIops: to.Ptr[int32](2300),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P20"),
		// 																											Iops: to.Ptr[int32](2300),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P30"),
		// 																											Iops: to.Ptr[int32](5000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P40"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P30"),
		// 																									StorageSizeMb: to.Ptr[int64](1048576),
		// 																									SupportedIops: to.Ptr[int32](5000),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P30"),
		// 																											Iops: to.Ptr[int32](5000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P40"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P40"),
		// 																									StorageSizeMb: to.Ptr[int64](2097152),
		// 																									SupportedIops: to.Ptr[int32](7500),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P40"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P50"),
		// 																									StorageSizeMb: to.Ptr[int64](4193280),
		// 																									SupportedIops: to.Ptr[int32](7500),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P50"),
		// 																									StorageSizeMb: to.Ptr[int64](4194304),
		// 																									SupportedIops: to.Ptr[int32](7500),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P50"),
		// 																											Iops: to.Ptr[int32](7500),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P60"),
		// 																									StorageSizeMb: to.Ptr[int64](8388608),
		// 																									SupportedIops: to.Ptr[int32](16000),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P60"),
		// 																											Iops: to.Ptr[int32](16000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P70"),
		// 																											Iops: to.Ptr[int32](18000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P80"),
		// 																											Iops: to.Ptr[int32](20000),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P70"),
		// 																									StorageSizeMb: to.Ptr[int64](16777216),
		// 																									SupportedIops: to.Ptr[int32](18000),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P70"),
		// 																											Iops: to.Ptr[int32](18000),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("P80"),
		// 																											Iops: to.Ptr[int32](20000),
		// 																									}},
		// 																								},
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("P80"),
		// 																									StorageSizeMb: to.Ptr[int64](33553408),
		// 																									SupportedIops: to.Ptr[int32](20000),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("P80"),
		// 																											Iops: to.Ptr[int32](20000),
		// 																									}},
		// 																							}},
		// 																						},
		// 																						{
		// 																							Name: to.Ptr("ManagedDiskV2"),
		// 																							DefaultStorageSizeMb: to.Ptr[int64](32768),
		// 																							SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																								{
		// 																									DefaultIopsTier: to.Ptr("None"),
		// 																									MaximumStorageSizeMb: to.Ptr[int64](67108864),
		// 																									StorageSizeMb: to.Ptr[int64](32768),
		// 																									SupportedIops: to.Ptr[int32](3000),
		// 																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																										{
		// 																											Name: to.Ptr("None"),
		// 																											Iops: to.Ptr[int32](0),
		// 																									}},
		// 																									SupportedMaximumIops: to.Ptr[int32](80000),
		// 																									SupportedMaximumThroughput: to.Ptr[int32](1200),
		// 																									SupportedThroughput: to.Ptr[int32](125),
		// 																							}},
		// 																					}},
		// 																				},
		// 																				{
		// 																					Name: to.Ptr("GeneralPurpose"),
		// 																					DefaultSKUName: to.Ptr("Standard_D4ds_v5"),
		// 																					SupportedServerSKUs: []*armpostgresqlflexibleservers.ServerSKUCapability{
		// 																						{
		// 																							Name: to.Ptr("Standard_D2s_v3"),
		// 																							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																							},
		// 																							SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																								to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																								to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																								SupportedIops: to.Ptr[int32](3200),
		// 																								SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																								SupportedZones: []*string{
		// 																									to.Ptr("1"),
		// 																									to.Ptr("2"),
		// 																									to.Ptr("3")},
		// 																									VCores: to.Ptr[int32](2),
		// 																								},
		// 																								{
		// 																									Name: to.Ptr("Standard_D4ds_v5"),
		// 																									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																									},
		// 																									SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																										to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																										to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																										SupportedIops: to.Ptr[int32](6400),
		// 																										SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																										SupportedZones: []*string{
		// 																											to.Ptr("1"),
		// 																											to.Ptr("2"),
		// 																											to.Ptr("3")},
		// 																											VCores: to.Ptr[int32](4),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("Standard_D8s_v3"),
		// 																											SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																											},
		// 																											SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																												to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																												to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																												SupportedIops: to.Ptr[int32](12800),
		// 																												SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																												SupportedZones: []*string{
		// 																													to.Ptr("1"),
		// 																													to.Ptr("2"),
		// 																													to.Ptr("3")},
		// 																													VCores: to.Ptr[int32](8),
		// 																												},
		// 																												{
		// 																													Name: to.Ptr("Standard_D16s_v3"),
		// 																													SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																													},
		// 																													SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																														to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																														to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																														SupportedIops: to.Ptr[int32](25600),
		// 																														SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																														SupportedZones: []*string{
		// 																															to.Ptr("1"),
		// 																															to.Ptr("2"),
		// 																															to.Ptr("3")},
		// 																															VCores: to.Ptr[int32](16),
		// 																														},
		// 																														{
		// 																															Name: to.Ptr("Standard_D32s_v3"),
		// 																															SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																															},
		// 																															SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																SupportedIops: to.Ptr[int32](51200),
		// 																																SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																SupportedZones: []*string{
		// 																																	to.Ptr("1"),
		// 																																	to.Ptr("2"),
		// 																																	to.Ptr("3")},
		// 																																	VCores: to.Ptr[int32](32),
		// 																																},
		// 																																{
		// 																																	Name: to.Ptr("Standard_D48s_v3"),
		// 																																	SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																	},
		// 																																	SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																		to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																		to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																		SupportedIops: to.Ptr[int32](76800),
		// 																																		SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																		SupportedZones: []*string{
		// 																																			to.Ptr("1"),
		// 																																			to.Ptr("2"),
		// 																																			to.Ptr("3")},
		// 																																			VCores: to.Ptr[int32](48),
		// 																																		},
		// 																																		{
		// 																																			Name: to.Ptr("Standard_D64s_v3"),
		// 																																			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																			},
		// 																																			SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																				to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																				to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																				SupportedIops: to.Ptr[int32](80000),
		// 																																				SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																				SupportedZones: []*string{
		// 																																					to.Ptr("1"),
		// 																																					to.Ptr("2"),
		// 																																					to.Ptr("3")},
		// 																																					VCores: to.Ptr[int32](64),
		// 																																				},
		// 																																				{
		// 																																					Name: to.Ptr("Standard_D2ds_v4"),
		// 																																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																					},
		// 																																					SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																						to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																						to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																						SupportedIops: to.Ptr[int32](3200),
		// 																																						SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																						SupportedZones: []*string{
		// 																																							to.Ptr("1"),
		// 																																							to.Ptr("2"),
		// 																																							to.Ptr("3")},
		// 																																							VCores: to.Ptr[int32](2),
		// 																																						},
		// 																																						{
		// 																																							Name: to.Ptr("Standard_D4ds_v4"),
		// 																																							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																							},
		// 																																							SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																								to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																								to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																								SupportedIops: to.Ptr[int32](6400),
		// 																																								SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																								SupportedZones: []*string{
		// 																																									to.Ptr("1"),
		// 																																									to.Ptr("2"),
		// 																																									to.Ptr("3")},
		// 																																									VCores: to.Ptr[int32](4),
		// 																																								},
		// 																																								{
		// 																																									Name: to.Ptr("Standard_D8ds_v4"),
		// 																																									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																									},
		// 																																									SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																										to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																										to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																										SupportedIops: to.Ptr[int32](12800),
		// 																																										SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																										SupportedZones: []*string{
		// 																																											to.Ptr("1"),
		// 																																											to.Ptr("2"),
		// 																																											to.Ptr("3")},
		// 																																											VCores: to.Ptr[int32](8),
		// 																																										},
		// 																																										{
		// 																																											Name: to.Ptr("Standard_D16ds_v4"),
		// 																																											SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																											},
		// 																																											SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																												to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																												to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																												SupportedIops: to.Ptr[int32](25600),
		// 																																												SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																												SupportedZones: []*string{
		// 																																													to.Ptr("1"),
		// 																																													to.Ptr("2"),
		// 																																													to.Ptr("3")},
		// 																																													VCores: to.Ptr[int32](16),
		// 																																												},
		// 																																												{
		// 																																													Name: to.Ptr("Standard_D32ds_v4"),
		// 																																													SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																													},
		// 																																													SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																														to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																														to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																														SupportedIops: to.Ptr[int32](51200),
		// 																																														SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																														SupportedZones: []*string{
		// 																																															to.Ptr("1"),
		// 																																															to.Ptr("2"),
		// 																																															to.Ptr("3")},
		// 																																															VCores: to.Ptr[int32](32),
		// 																																														},
		// 																																														{
		// 																																															Name: to.Ptr("Standard_D48ds_v4"),
		// 																																															SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																															},
		// 																																															SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																SupportedIops: to.Ptr[int32](76800),
		// 																																																SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																SupportedZones: []*string{
		// 																																																	to.Ptr("1"),
		// 																																																	to.Ptr("2"),
		// 																																																	to.Ptr("3")},
		// 																																																	VCores: to.Ptr[int32](48),
		// 																																																},
		// 																																																{
		// 																																																	Name: to.Ptr("Standard_D64ds_v4"),
		// 																																																	SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																	},
		// 																																																	SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																		to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																		to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																		SupportedIops: to.Ptr[int32](80000),
		// 																																																		SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																		SupportedZones: []*string{
		// 																																																			to.Ptr("1"),
		// 																																																			to.Ptr("2"),
		// 																																																			to.Ptr("3")},
		// 																																																			VCores: to.Ptr[int32](64),
		// 																																																		},
		// 																																																		{
		// 																																																			Name: to.Ptr("Standard_D2ds_v5"),
		// 																																																			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																			},
		// 																																																			SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																				to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																				to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																				SupportedIops: to.Ptr[int32](3750),
		// 																																																				SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																				SupportedZones: []*string{
		// 																																																					to.Ptr("1"),
		// 																																																					to.Ptr("2"),
		// 																																																					to.Ptr("3")},
		// 																																																					VCores: to.Ptr[int32](2),
		// 																																																				},
		// 																																																				{
		// 																																																					Name: to.Ptr("Standard_D4ds_v5"),
		// 																																																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																					},
		// 																																																					SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																						to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																						to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																						SupportedIops: to.Ptr[int32](6400),
		// 																																																						SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																						SupportedZones: []*string{
		// 																																																							to.Ptr("1"),
		// 																																																							to.Ptr("2"),
		// 																																																							to.Ptr("3")},
		// 																																																							VCores: to.Ptr[int32](4),
		// 																																																						},
		// 																																																						{
		// 																																																							Name: to.Ptr("Standard_D8ds_v5"),
		// 																																																							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																							},
		// 																																																							SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																								to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																								to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																								SupportedIops: to.Ptr[int32](12800),
		// 																																																								SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																								SupportedZones: []*string{
		// 																																																									to.Ptr("1"),
		// 																																																									to.Ptr("2"),
		// 																																																									to.Ptr("3")},
		// 																																																									VCores: to.Ptr[int32](8),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("Standard_D16ds_v5"),
		// 																																																									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																									},
		// 																																																									SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																										to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																										to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																										SupportedIops: to.Ptr[int32](25600),
		// 																																																										SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																										SupportedZones: []*string{
		// 																																																											to.Ptr("1"),
		// 																																																											to.Ptr("2"),
		// 																																																											to.Ptr("3")},
		// 																																																											VCores: to.Ptr[int32](16),
		// 																																																										},
		// 																																																										{
		// 																																																											Name: to.Ptr("Standard_D32ds_v5"),
		// 																																																											SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																											},
		// 																																																											SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																												to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																												to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																												SupportedIops: to.Ptr[int32](51200),
		// 																																																												SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																												SupportedZones: []*string{
		// 																																																													to.Ptr("1"),
		// 																																																													to.Ptr("2"),
		// 																																																													to.Ptr("3")},
		// 																																																													VCores: to.Ptr[int32](32),
		// 																																																												},
		// 																																																												{
		// 																																																													Name: to.Ptr("Standard_D48ds_v5"),
		// 																																																													SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																													},
		// 																																																													SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																														to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																														to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																														SupportedIops: to.Ptr[int32](76800),
		// 																																																														SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																														SupportedZones: []*string{
		// 																																																															to.Ptr("1"),
		// 																																																															to.Ptr("2"),
		// 																																																															to.Ptr("3")},
		// 																																																															VCores: to.Ptr[int32](48),
		// 																																																														},
		// 																																																														{
		// 																																																															Name: to.Ptr("Standard_D64ds_v5"),
		// 																																																															SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																															},
		// 																																																															SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																SupportedIops: to.Ptr[int32](80000),
		// 																																																																SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																																SupportedZones: []*string{
		// 																																																																	to.Ptr("1"),
		// 																																																																	to.Ptr("2"),
		// 																																																																	to.Ptr("3")},
		// 																																																																	VCores: to.Ptr[int32](64),
		// 																																																																},
		// 																																																																{
		// 																																																																	Name: to.Ptr("Standard_D96ds_v5"),
		// 																																																																	SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																	},
		// 																																																																	SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																		SupportedIops: to.Ptr[int32](80000),
		// 																																																																		SupportedMemoryPerVcoreMb: to.Ptr[int64](4096),
		// 																																																																		SupportedZones: []*string{
		// 																																																																			to.Ptr("1"),
		// 																																																																			to.Ptr("2"),
		// 																																																																			to.Ptr("3")},
		// 																																																																			VCores: to.Ptr[int32](96),
		// 																																																																	}},
		// 																																																																	SupportedStorageEditions: []*armpostgresqlflexibleservers.StorageEditionCapability{
		// 																																																																		{
		// 																																																																			Name: to.Ptr("ManagedDisk"),
		// 																																																																			DefaultStorageSizeMb: to.Ptr[int64](65536),
		// 																																																																			SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P4"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](32768),
		// 																																																																					SupportedIops: to.Ptr[int32](120),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P4"),
		// 																																																																							Iops: to.Ptr[int32](120),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P6"),
		// 																																																																							Iops: to.Ptr[int32](240),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P10"),
		// 																																																																							Iops: to.Ptr[int32](500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P15"),
		// 																																																																							Iops: to.Ptr[int32](1100),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P20"),
		// 																																																																							Iops: to.Ptr[int32](2300),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P30"),
		// 																																																																							Iops: to.Ptr[int32](5000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P40"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P6"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](65536),
		// 																																																																					SupportedIops: to.Ptr[int32](240),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P6"),
		// 																																																																							Iops: to.Ptr[int32](240),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P10"),
		// 																																																																							Iops: to.Ptr[int32](500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P15"),
		// 																																																																							Iops: to.Ptr[int32](1100),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P20"),
		// 																																																																							Iops: to.Ptr[int32](2300),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P30"),
		// 																																																																							Iops: to.Ptr[int32](5000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P40"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P10"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](131072),
		// 																																																																					SupportedIops: to.Ptr[int32](500),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P10"),
		// 																																																																							Iops: to.Ptr[int32](500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P15"),
		// 																																																																							Iops: to.Ptr[int32](1100),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P20"),
		// 																																																																							Iops: to.Ptr[int32](2300),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P30"),
		// 																																																																							Iops: to.Ptr[int32](5000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P40"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P15"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](262144),
		// 																																																																					SupportedIops: to.Ptr[int32](1100),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P15"),
		// 																																																																							Iops: to.Ptr[int32](1100),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P20"),
		// 																																																																							Iops: to.Ptr[int32](2300),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P30"),
		// 																																																																							Iops: to.Ptr[int32](5000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P40"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P20"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](524288),
		// 																																																																					SupportedIops: to.Ptr[int32](2300),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P20"),
		// 																																																																							Iops: to.Ptr[int32](2300),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P30"),
		// 																																																																							Iops: to.Ptr[int32](5000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P40"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P30"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](1048576),
		// 																																																																					SupportedIops: to.Ptr[int32](5000),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P30"),
		// 																																																																							Iops: to.Ptr[int32](5000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P40"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P40"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](2097152),
		// 																																																																					SupportedIops: to.Ptr[int32](7500),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P40"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P50"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](4193280),
		// 																																																																					SupportedIops: to.Ptr[int32](7500),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P50"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](4194304),
		// 																																																																					SupportedIops: to.Ptr[int32](7500),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P50"),
		// 																																																																							Iops: to.Ptr[int32](7500),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P60"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](8388608),
		// 																																																																					SupportedIops: to.Ptr[int32](16000),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P60"),
		// 																																																																							Iops: to.Ptr[int32](16000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P70"),
		// 																																																																							Iops: to.Ptr[int32](18000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P80"),
		// 																																																																							Iops: to.Ptr[int32](20000),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P70"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](16777216),
		// 																																																																					SupportedIops: to.Ptr[int32](18000),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P70"),
		// 																																																																							Iops: to.Ptr[int32](18000),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P80"),
		// 																																																																							Iops: to.Ptr[int32](20000),
		// 																																																																					}},
		// 																																																																				},
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("P80"),
		// 																																																																					StorageSizeMb: to.Ptr[int64](33553408),
		// 																																																																					SupportedIops: to.Ptr[int32](20000),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("P80"),
		// 																																																																							Iops: to.Ptr[int32](20000),
		// 																																																																					}},
		// 																																																																			}},
		// 																																																																		},
		// 																																																																		{
		// 																																																																			Name: to.Ptr("ManagedDiskV2"),
		// 																																																																			DefaultStorageSizeMb: to.Ptr[int64](65536),
		// 																																																																			SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("None"),
		// 																																																																					MaximumStorageSizeMb: to.Ptr[int64](67108864),
		// 																																																																					StorageSizeMb: to.Ptr[int64](32768),
		// 																																																																					SupportedIops: to.Ptr[int32](3000),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("None"),
		// 																																																																							Iops: to.Ptr[int32](0),
		// 																																																																					}},
		// 																																																																					SupportedMaximumIops: to.Ptr[int32](80000),
		// 																																																																					SupportedMaximumThroughput: to.Ptr[int32](1200),
		// 																																																																					SupportedThroughput: to.Ptr[int32](125),
		// 																																																																			}},
		// 																																																																		},
		// 																																																																		{
		// 																																																																			Name: to.Ptr("UltraDisk"),
		// 																																																																			DefaultStorageSizeMb: to.Ptr[int64](65536),
		// 																																																																			SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																																																																				{
		// 																																																																					DefaultIopsTier: to.Ptr("None"),
		// 																																																																					MaximumStorageSizeMb: to.Ptr[int64](67108864),
		// 																																																																					StorageSizeMb: to.Ptr[int64](32768),
		// 																																																																					SupportedIops: to.Ptr[int32](4800),
		// 																																																																					SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																						{
		// 																																																																							Name: to.Ptr("None"),
		// 																																																																							Iops: to.Ptr[int32](0),
		// 																																																																					}},
		// 																																																																					SupportedMaximumIops: to.Ptr[int32](400000),
		// 																																																																					SupportedMaximumThroughput: to.Ptr[int32](10000),
		// 																																																																					SupportedThroughput: to.Ptr[int32](1200),
		// 																																																																			}},
		// 																																																																	}},
		// 																																																																},
		// 																																																																{
		// 																																																																	Name: to.Ptr("MemoryOptimized"),
		// 																																																																	DefaultSKUName: to.Ptr("Standard_E4ds_v5"),
		// 																																																																	SupportedServerSKUs: []*armpostgresqlflexibleservers.ServerSKUCapability{
		// 																																																																		{
		// 																																																																			Name: to.Ptr("Standard_E2s_v3"),
		// 																																																																			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																			},
		// 																																																																			SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																				SupportedIops: to.Ptr[int32](3200),
		// 																																																																				SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																				SupportedZones: []*string{
		// 																																																																					to.Ptr("1"),
		// 																																																																					to.Ptr("2"),
		// 																																																																					to.Ptr("3")},
		// 																																																																					VCores: to.Ptr[int32](2),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("Standard_E4s_v3"),
		// 																																																																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																					},
		// 																																																																					SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																						SupportedIops: to.Ptr[int32](6400),
		// 																																																																						SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																						SupportedZones: []*string{
		// 																																																																							to.Ptr("1"),
		// 																																																																							to.Ptr("2"),
		// 																																																																							to.Ptr("3")},
		// 																																																																							VCores: to.Ptr[int32](4),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("Standard_E8s_v3"),
		// 																																																																							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																							},
		// 																																																																							SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																								to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																								to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																								SupportedIops: to.Ptr[int32](12800),
		// 																																																																								SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																								SupportedZones: []*string{
		// 																																																																									to.Ptr("1"),
		// 																																																																									to.Ptr("2"),
		// 																																																																									to.Ptr("3")},
		// 																																																																									VCores: to.Ptr[int32](8),
		// 																																																																								},
		// 																																																																								{
		// 																																																																									Name: to.Ptr("Standard_E16s_v3"),
		// 																																																																									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																									},
		// 																																																																									SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																										to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																										to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																										SupportedIops: to.Ptr[int32](25600),
		// 																																																																										SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																										SupportedZones: []*string{
		// 																																																																											to.Ptr("1"),
		// 																																																																											to.Ptr("2"),
		// 																																																																											to.Ptr("3")},
		// 																																																																											VCores: to.Ptr[int32](16),
		// 																																																																										},
		// 																																																																										{
		// 																																																																											Name: to.Ptr("Standard_E32s_v3"),
		// 																																																																											SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																											},
		// 																																																																											SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																												to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																												to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																												SupportedIops: to.Ptr[int32](32000),
		// 																																																																												SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																												SupportedZones: []*string{
		// 																																																																													to.Ptr("1"),
		// 																																																																													to.Ptr("2"),
		// 																																																																													to.Ptr("3")},
		// 																																																																													VCores: to.Ptr[int32](32),
		// 																																																																												},
		// 																																																																												{
		// 																																																																													Name: to.Ptr("Standard_E48s_v3"),
		// 																																																																													SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																													},
		// 																																																																													SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																														to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																														to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																														SupportedIops: to.Ptr[int32](51200),
		// 																																																																														SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																														SupportedZones: []*string{
		// 																																																																															to.Ptr("1"),
		// 																																																																															to.Ptr("2"),
		// 																																																																															to.Ptr("3")},
		// 																																																																															VCores: to.Ptr[int32](48),
		// 																																																																														},
		// 																																																																														{
		// 																																																																															Name: to.Ptr("Standard_E64s_v3"),
		// 																																																																															SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																															},
		// 																																																																															SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																SupportedIops: to.Ptr[int32](76800),
		// 																																																																																SupportedMemoryPerVcoreMb: to.Ptr[int64](6912),
		// 																																																																																SupportedZones: []*string{
		// 																																																																																	to.Ptr("1"),
		// 																																																																																	to.Ptr("2"),
		// 																																																																																	to.Ptr("3")},
		// 																																																																																	VCores: to.Ptr[int32](64),
		// 																																																																																},
		// 																																																																																{
		// 																																																																																	Name: to.Ptr("Standard_E2ds_v4"),
		// 																																																																																	SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																	},
		// 																																																																																	SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																		SupportedIops: to.Ptr[int32](3200),
		// 																																																																																		SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																		SupportedZones: []*string{
		// 																																																																																			to.Ptr("1"),
		// 																																																																																			to.Ptr("2"),
		// 																																																																																			to.Ptr("3")},
		// 																																																																																			VCores: to.Ptr[int32](2),
		// 																																																																																		},
		// 																																																																																		{
		// 																																																																																			Name: to.Ptr("Standard_E4ds_v4"),
		// 																																																																																			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																			},
		// 																																																																																			SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																				SupportedIops: to.Ptr[int32](6400),
		// 																																																																																				SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																				SupportedZones: []*string{
		// 																																																																																					to.Ptr("1"),
		// 																																																																																					to.Ptr("2"),
		// 																																																																																					to.Ptr("3")},
		// 																																																																																					VCores: to.Ptr[int32](4),
		// 																																																																																				},
		// 																																																																																				{
		// 																																																																																					Name: to.Ptr("Standard_E8ds_v4"),
		// 																																																																																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																					},
		// 																																																																																					SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																						SupportedIops: to.Ptr[int32](12800),
		// 																																																																																						SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																						SupportedZones: []*string{
		// 																																																																																							to.Ptr("1"),
		// 																																																																																							to.Ptr("2"),
		// 																																																																																							to.Ptr("3")},
		// 																																																																																							VCores: to.Ptr[int32](8),
		// 																																																																																						},
		// 																																																																																						{
		// 																																																																																							Name: to.Ptr("Standard_E16ds_v4"),
		// 																																																																																							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																							},
		// 																																																																																							SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																								to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																								to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																								SupportedIops: to.Ptr[int32](25600),
		// 																																																																																								SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																								SupportedZones: []*string{
		// 																																																																																									to.Ptr("1"),
		// 																																																																																									to.Ptr("2"),
		// 																																																																																									to.Ptr("3")},
		// 																																																																																									VCores: to.Ptr[int32](16),
		// 																																																																																								},
		// 																																																																																								{
		// 																																																																																									Name: to.Ptr("Standard_E20ds_v4"),
		// 																																																																																									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																									},
		// 																																																																																									SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																										to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																										to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																										SupportedIops: to.Ptr[int32](32000),
		// 																																																																																										SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																										SupportedZones: []*string{
		// 																																																																																											to.Ptr("1"),
		// 																																																																																											to.Ptr("2"),
		// 																																																																																											to.Ptr("3")},
		// 																																																																																											VCores: to.Ptr[int32](20),
		// 																																																																																										},
		// 																																																																																										{
		// 																																																																																											Name: to.Ptr("Standard_E32ds_v4"),
		// 																																																																																											SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																											},
		// 																																																																																											SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																												to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																												to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																												SupportedIops: to.Ptr[int32](51200),
		// 																																																																																												SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																												SupportedZones: []*string{
		// 																																																																																													to.Ptr("1"),
		// 																																																																																													to.Ptr("2"),
		// 																																																																																													to.Ptr("3")},
		// 																																																																																													VCores: to.Ptr[int32](32),
		// 																																																																																												},
		// 																																																																																												{
		// 																																																																																													Name: to.Ptr("Standard_E48ds_v4"),
		// 																																																																																													SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																													},
		// 																																																																																													SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																														to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																														to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																														SupportedIops: to.Ptr[int32](76800),
		// 																																																																																														SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																														SupportedZones: []*string{
		// 																																																																																															to.Ptr("1"),
		// 																																																																																															to.Ptr("2"),
		// 																																																																																															to.Ptr("3")},
		// 																																																																																															VCores: to.Ptr[int32](48),
		// 																																																																																														},
		// 																																																																																														{
		// 																																																																																															Name: to.Ptr("Standard_E64ds_v4"),
		// 																																																																																															SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																															},
		// 																																																																																															SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																SupportedIops: to.Ptr[int32](80000),
		// 																																																																																																SupportedMemoryPerVcoreMb: to.Ptr[int64](6912),
		// 																																																																																																SupportedZones: []*string{
		// 																																																																																																	to.Ptr("1"),
		// 																																																																																																	to.Ptr("2"),
		// 																																																																																																	to.Ptr("3")},
		// 																																																																																																	VCores: to.Ptr[int32](64),
		// 																																																																																																},
		// 																																																																																																{
		// 																																																																																																	Name: to.Ptr("Standard_M64"),
		// 																																																																																																	SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																	},
		// 																																																																																																	SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																		SupportedIops: to.Ptr[int32](40000),
		// 																																																																																																		SupportedMemoryPerVcoreMb: to.Ptr[int64](16384),
		// 																																																																																																		SupportedZones: []*string{
		// 																																																																																																			to.Ptr("1"),
		// 																																																																																																			to.Ptr("2"),
		// 																																																																																																			to.Ptr("3")},
		// 																																																																																																			VCores: to.Ptr[int32](64),
		// 																																																																																																		},
		// 																																																																																																		{
		// 																																																																																																			Name: to.Ptr("Standard_M128"),
		// 																																																																																																			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																			},
		// 																																																																																																			SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																				SupportedIops: to.Ptr[int32](80000),
		// 																																																																																																				SupportedMemoryPerVcoreMb: to.Ptr[int64](16384),
		// 																																																																																																				SupportedZones: []*string{
		// 																																																																																																					to.Ptr("1"),
		// 																																																																																																					to.Ptr("2"),
		// 																																																																																																					to.Ptr("3")},
		// 																																																																																																					VCores: to.Ptr[int32](128),
		// 																																																																																																				},
		// 																																																																																																				{
		// 																																																																																																					Name: to.Ptr("Standard_E2ds_v5"),
		// 																																																																																																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																					},
		// 																																																																																																					SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																						SupportedIops: to.Ptr[int32](3750),
		// 																																																																																																						SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																						SupportedZones: []*string{
		// 																																																																																																							to.Ptr("1"),
		// 																																																																																																							to.Ptr("2"),
		// 																																																																																																							to.Ptr("3")},
		// 																																																																																																							VCores: to.Ptr[int32](2),
		// 																																																																																																						},
		// 																																																																																																						{
		// 																																																																																																							Name: to.Ptr("Standard_E4ds_v5"),
		// 																																																																																																							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																							},
		// 																																																																																																							SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																								to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																								to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																								SupportedIops: to.Ptr[int32](6400),
		// 																																																																																																								SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																								SupportedZones: []*string{
		// 																																																																																																									to.Ptr("1"),
		// 																																																																																																									to.Ptr("2"),
		// 																																																																																																									to.Ptr("3")},
		// 																																																																																																									VCores: to.Ptr[int32](4),
		// 																																																																																																								},
		// 																																																																																																								{
		// 																																																																																																									Name: to.Ptr("Standard_E8ds_v5"),
		// 																																																																																																									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																									},
		// 																																																																																																									SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																										to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																										to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																										SupportedIops: to.Ptr[int32](12800),
		// 																																																																																																										SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																										SupportedZones: []*string{
		// 																																																																																																											to.Ptr("1"),
		// 																																																																																																											to.Ptr("2"),
		// 																																																																																																											to.Ptr("3")},
		// 																																																																																																											VCores: to.Ptr[int32](8),
		// 																																																																																																										},
		// 																																																																																																										{
		// 																																																																																																											Name: to.Ptr("Standard_E16ds_v5"),
		// 																																																																																																											SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																											},
		// 																																																																																																											SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																												to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																												to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																												SupportedIops: to.Ptr[int32](25600),
		// 																																																																																																												SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																												SupportedZones: []*string{
		// 																																																																																																													to.Ptr("1"),
		// 																																																																																																													to.Ptr("2"),
		// 																																																																																																													to.Ptr("3")},
		// 																																																																																																													VCores: to.Ptr[int32](16),
		// 																																																																																																												},
		// 																																																																																																												{
		// 																																																																																																													Name: to.Ptr("Standard_E20ds_v5"),
		// 																																																																																																													SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																													},
		// 																																																																																																													SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																														to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																														to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																														SupportedIops: to.Ptr[int32](32000),
		// 																																																																																																														SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																														SupportedZones: []*string{
		// 																																																																																																															to.Ptr("1"),
		// 																																																																																																															to.Ptr("2"),
		// 																																																																																																															to.Ptr("3")},
		// 																																																																																																															VCores: to.Ptr[int32](20),
		// 																																																																																																														},
		// 																																																																																																														{
		// 																																																																																																															Name: to.Ptr("Standard_E32ds_v5"),
		// 																																																																																																															SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																															},
		// 																																																																																																															SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																																to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																																to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																																SupportedIops: to.Ptr[int32](51200),
		// 																																																																																																																SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																																SupportedZones: []*string{
		// 																																																																																																																	to.Ptr("1"),
		// 																																																																																																																	to.Ptr("2"),
		// 																																																																																																																	to.Ptr("3")},
		// 																																																																																																																	VCores: to.Ptr[int32](32),
		// 																																																																																																																},
		// 																																																																																																																{
		// 																																																																																																																	Name: to.Ptr("Standard_E48ds_v5"),
		// 																																																																																																																	SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																	},
		// 																																																																																																																	SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																																		to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																																		SupportedIops: to.Ptr[int32](76800),
		// 																																																																																																																		SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																																		SupportedZones: []*string{
		// 																																																																																																																			to.Ptr("1"),
		// 																																																																																																																			to.Ptr("2"),
		// 																																																																																																																			to.Ptr("3")},
		// 																																																																																																																			VCores: to.Ptr[int32](48),
		// 																																																																																																																		},
		// 																																																																																																																		{
		// 																																																																																																																			Name: to.Ptr("Standard_E64ds_v5"),
		// 																																																																																																																			SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																			},
		// 																																																																																																																			SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																																				to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																																				SupportedIops: to.Ptr[int32](80000),
		// 																																																																																																																				SupportedMemoryPerVcoreMb: to.Ptr[int64](8192),
		// 																																																																																																																				SupportedZones: []*string{
		// 																																																																																																																					to.Ptr("1"),
		// 																																																																																																																					to.Ptr("2"),
		// 																																																																																																																					to.Ptr("3")},
		// 																																																																																																																					VCores: to.Ptr[int32](64),
		// 																																																																																																																				},
		// 																																																																																																																				{
		// 																																																																																																																					Name: to.Ptr("Standard_E96ds_v5"),
		// 																																																																																																																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																					},
		// 																																																																																																																					SupportedHaMode: []*armpostgresqlflexibleservers.HaMode{
		// 																																																																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeSameZone),
		// 																																																																																																																						to.Ptr(armpostgresqlflexibleservers.HaModeZoneRedundant)},
		// 																																																																																																																						SupportedIops: to.Ptr[int32](80000),
		// 																																																																																																																						SupportedMemoryPerVcoreMb: to.Ptr[int64](7168),
		// 																																																																																																																						SupportedZones: []*string{
		// 																																																																																																																							to.Ptr("1"),
		// 																																																																																																																							to.Ptr("2"),
		// 																																																																																																																							to.Ptr("3")},
		// 																																																																																																																							VCores: to.Ptr[int32](96),
		// 																																																																																																																					}},
		// 																																																																																																																					SupportedStorageEditions: []*armpostgresqlflexibleservers.StorageEditionCapability{
		// 																																																																																																																						{
		// 																																																																																																																							Name: to.Ptr("ManagedDisk"),
		// 																																																																																																																							DefaultStorageSizeMb: to.Ptr[int64](131072),
		// 																																																																																																																							SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P4"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](32768),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](120),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P4"),
		// 																																																																																																																											Iops: to.Ptr[int32](120),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P6"),
		// 																																																																																																																											Iops: to.Ptr[int32](240),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P10"),
		// 																																																																																																																											Iops: to.Ptr[int32](500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P15"),
		// 																																																																																																																											Iops: to.Ptr[int32](1100),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P20"),
		// 																																																																																																																											Iops: to.Ptr[int32](2300),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P30"),
		// 																																																																																																																											Iops: to.Ptr[int32](5000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P40"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P6"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](65536),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](240),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P6"),
		// 																																																																																																																											Iops: to.Ptr[int32](240),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P10"),
		// 																																																																																																																											Iops: to.Ptr[int32](500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P15"),
		// 																																																																																																																											Iops: to.Ptr[int32](1100),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P20"),
		// 																																																																																																																											Iops: to.Ptr[int32](2300),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P30"),
		// 																																																																																																																											Iops: to.Ptr[int32](5000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P40"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P10"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](131072),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](500),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P10"),
		// 																																																																																																																											Iops: to.Ptr[int32](500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P15"),
		// 																																																																																																																											Iops: to.Ptr[int32](1100),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P20"),
		// 																																																																																																																											Iops: to.Ptr[int32](2300),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P30"),
		// 																																																																																																																											Iops: to.Ptr[int32](5000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P40"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P15"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](262144),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](1100),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P15"),
		// 																																																																																																																											Iops: to.Ptr[int32](1100),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P20"),
		// 																																																																																																																											Iops: to.Ptr[int32](2300),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P30"),
		// 																																																																																																																											Iops: to.Ptr[int32](5000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P40"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P20"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](524288),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](2300),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P20"),
		// 																																																																																																																											Iops: to.Ptr[int32](2300),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P30"),
		// 																																																																																																																											Iops: to.Ptr[int32](5000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P40"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P30"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](1048576),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](5000),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P30"),
		// 																																																																																																																											Iops: to.Ptr[int32](5000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P40"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P40"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](2097152),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](7500),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P40"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P50"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](4193280),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](7500),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P50"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](4194304),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](7500),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P50"),
		// 																																																																																																																											Iops: to.Ptr[int32](7500),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P60"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](8388608),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](16000),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P60"),
		// 																																																																																																																											Iops: to.Ptr[int32](16000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P70"),
		// 																																																																																																																											Iops: to.Ptr[int32](18000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P80"),
		// 																																																																																																																											Iops: to.Ptr[int32](20000),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P70"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](16777216),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](18000),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P70"),
		// 																																																																																																																											Iops: to.Ptr[int32](18000),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P80"),
		// 																																																																																																																											Iops: to.Ptr[int32](20000),
		// 																																																																																																																									}},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("P80"),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](33553408),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](20000),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("P80"),
		// 																																																																																																																											Iops: to.Ptr[int32](20000),
		// 																																																																																																																									}},
		// 																																																																																																																							}},
		// 																																																																																																																						},
		// 																																																																																																																						{
		// 																																																																																																																							Name: to.Ptr("ManagedDiskV2"),
		// 																																																																																																																							DefaultStorageSizeMb: to.Ptr[int64](131072),
		// 																																																																																																																							SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("None"),
		// 																																																																																																																									MaximumStorageSizeMb: to.Ptr[int64](67108864),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](32768),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](3000),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("None"),
		// 																																																																																																																											Iops: to.Ptr[int32](0),
		// 																																																																																																																									}},
		// 																																																																																																																									SupportedMaximumIops: to.Ptr[int32](80000),
		// 																																																																																																																									SupportedMaximumThroughput: to.Ptr[int32](1200),
		// 																																																																																																																									SupportedThroughput: to.Ptr[int32](125),
		// 																																																																																																																							}},
		// 																																																																																																																						},
		// 																																																																																																																						{
		// 																																																																																																																							Name: to.Ptr("UltraDisk"),
		// 																																																																																																																							DefaultStorageSizeMb: to.Ptr[int64](131072),
		// 																																																																																																																							SupportedStorageMb: []*armpostgresqlflexibleservers.StorageMbCapability{
		// 																																																																																																																								{
		// 																																																																																																																									DefaultIopsTier: to.Ptr("None"),
		// 																																																																																																																									MaximumStorageSizeMb: to.Ptr[int64](67108864),
		// 																																																																																																																									StorageSizeMb: to.Ptr[int64](32768),
		// 																																																																																																																									SupportedIops: to.Ptr[int32](4800),
		// 																																																																																																																									SupportedIopsTiers: []*armpostgresqlflexibleservers.StorageTierCapability{
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("None"),
		// 																																																																																																																											Iops: to.Ptr[int32](0),
		// 																																																																																																																									}},
		// 																																																																																																																									SupportedMaximumIops: to.Ptr[int32](400000),
		// 																																																																																																																									SupportedMaximumThroughput: to.Ptr[int32](10000),
		// 																																																																																																																									SupportedThroughput: to.Ptr[int32](1200),
		// 																																																																																																																							}},
		// 																																																																																																																					}},
		// 																																																																																																																			}},
		// 																																																																																																																			SupportedServerVersions: []*armpostgresqlflexibleservers.ServerVersionCapability{
		// 																																																																																																																				{
		// 																																																																																																																					Name: to.Ptr("11"),
		// 																																																																																																																					SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																					},
		// 																																																																																																																					SupportedVersionsToUpgrade: []*string{
		// 																																																																																																																						to.Ptr("12"),
		// 																																																																																																																						to.Ptr("13"),
		// 																																																																																																																						to.Ptr("14"),
		// 																																																																																																																						to.Ptr("15"),
		// 																																																																																																																						to.Ptr("16")},
		// 																																																																																																																					},
		// 																																																																																																																					{
		// 																																																																																																																						Name: to.Ptr("12"),
		// 																																																																																																																						SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																						},
		// 																																																																																																																						SupportedVersionsToUpgrade: []*string{
		// 																																																																																																																							to.Ptr("13"),
		// 																																																																																																																							to.Ptr("14"),
		// 																																																																																																																							to.Ptr("15"),
		// 																																																																																																																							to.Ptr("16")},
		// 																																																																																																																						},
		// 																																																																																																																						{
		// 																																																																																																																							Name: to.Ptr("13"),
		// 																																																																																																																							SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																							},
		// 																																																																																																																							SupportedVersionsToUpgrade: []*string{
		// 																																																																																																																								to.Ptr("14"),
		// 																																																																																																																								to.Ptr("15"),
		// 																																																																																																																								to.Ptr("16")},
		// 																																																																																																																							},
		// 																																																																																																																							{
		// 																																																																																																																								Name: to.Ptr("14"),
		// 																																																																																																																								SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																								},
		// 																																																																																																																								SupportedVersionsToUpgrade: []*string{
		// 																																																																																																																									to.Ptr("15"),
		// 																																																																																																																									to.Ptr("16")},
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									Name: to.Ptr("15"),
		// 																																																																																																																									SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																									},
		// 																																																																																																																									SupportedVersionsToUpgrade: []*string{
		// 																																																																																																																										to.Ptr("16")},
		// 																																																																																																																									},
		// 																																																																																																																									{
		// 																																																																																																																										Name: to.Ptr("16"),
		// 																																																																																																																										SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																										},
		// 																																																																																																																										SupportedVersionsToUpgrade: []*string{
		// 																																																																																																																										},
		// 																																																																																																																									},
		// 																																																																																																																									{
		// 																																																																																																																										Name: to.Ptr("17"),
		// 																																																																																																																										SupportedFeatures: []*armpostgresqlflexibleservers.SupportedFeature{
		// 																																																																																																																										},
		// 																																																																																																																										SupportedVersionsToUpgrade: []*string{
		// 																																																																																																																										},
		// 																																																																																																																								}},
		// 																																																																																																																								ZoneRedundantHaAndGeoBackupSupported: to.Ptr(armpostgresqlflexibleservers.ZoneRedundantHaAndGeoBackupSupportedEnumEnabled),
		// 																																																																																																																								ZoneRedundantHaSupported: to.Ptr(armpostgresqlflexibleservers.ZoneRedundantHaSupportedEnumEnabled),
		// 																																																																																																																						}},
		// 																																																																																																																					}
	}
}
