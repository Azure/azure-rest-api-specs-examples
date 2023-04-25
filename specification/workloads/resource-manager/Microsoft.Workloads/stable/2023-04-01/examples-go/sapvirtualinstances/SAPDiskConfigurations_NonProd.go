package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPDiskConfigurations_NonProd.json
func ExampleClient_SAPDiskConfigurations_sapDiskConfigurationsNonProd() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().SAPDiskConfigurations(ctx, "centralus", &armworkloads.ClientSAPDiskConfigurationsOptions{SAPDiskConfigurations: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPDiskConfigurationsResult = armworkloads.SAPDiskConfigurationsResult{
	// 	VolumeConfigurations: map[string]*armworkloads.SAPDiskConfiguration{
	// 		"backup": &armworkloads.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloads.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](2),
	// 				SizeGB: to.Ptr[int64](256),
	// 				SKU: &armworkloads.DiskSKU{
	// 					Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloads.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("E10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E15"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E20"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P15"),
	// 					IopsReadWrite: to.Ptr[int64](1100),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](125),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P20"),
	// 					IopsReadWrite: to.Ptr[int64](2300),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](150),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"hana/data": &armworkloads.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloads.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](4),
	// 				SizeGB: to.Ptr[int64](128),
	// 				SKU: &armworkloads.DiskSKU{
	// 					Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloads.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](5),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](4),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"hana/log": &armworkloads.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloads.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](3),
	// 				SizeGB: to.Ptr[int64](128),
	// 				SKU: &armworkloads.DiskSKU{
	// 					Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloads.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](5),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](3),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"hana/shared": &armworkloads.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloads.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](1),
	// 				SizeGB: to.Ptr[int64](256),
	// 				SKU: &armworkloads.DiskSKU{
	// 					Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloads.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P15"),
	// 					IopsReadWrite: to.Ptr[int64](1100),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](125),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P20"),
	// 					IopsReadWrite: to.Ptr[int64](2300),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](150),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P30"),
	// 					IopsReadWrite: to.Ptr[int64](5000),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](200),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](1024),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P40"),
	// 					IopsReadWrite: to.Ptr[int64](7500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](250),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](2048),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P50"),
	// 					IopsReadWrite: to.Ptr[int64](7500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](250),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](4096),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E15"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E20"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E30"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](1024),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E40"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](2048),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E50"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](4096),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 			}},
	// 		},
	// 		"os": &armworkloads.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloads.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](1),
	// 				SizeGB: to.Ptr[int64](64),
	// 				SKU: &armworkloads.DiskSKU{
	// 					Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloads.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P6"),
	// 					IopsReadWrite: to.Ptr[int64](240),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](50),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](64),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P6"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](64),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 			}},
	// 		},
	// 		"usr/sap": &armworkloads.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloads.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](1),
	// 				SizeGB: to.Ptr[int64](128),
	// 				SKU: &armworkloads.DiskSKU{
	// 					Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloads.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](1100),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](125),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](2300),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](150),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E15"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("E20"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](60),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloads.DiskSKU{
	// 						Name: to.Ptr(armworkloads.DiskSKUNameStandardSSDLRS),
	// 					},
	// 			}},
	// 		},
	// 	},
	// }
}
