package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPDiskConfigurations_Prod.json
func ExampleWorkloadsClient_SAPDiskConfigurations_sapDiskConfigurationsProd() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadsClient().SAPDiskConfigurations(ctx, "centralus", &armworkloadssapvirtualinstance.WorkloadsClientSAPDiskConfigurationsOptions{SAPDiskConfigurations: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPDiskConfigurationsResult = armworkloadssapvirtualinstance.SAPDiskConfigurationsResult{
	// 	VolumeConfigurations: map[string]*armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 		"backup": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](2),
	// 				SizeGB: to.Ptr[int64](256),
	// 				SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 					Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P15"),
	// 					IopsReadWrite: to.Ptr[int64](1100),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](125),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P20"),
	// 					IopsReadWrite: to.Ptr[int64](2300),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 					MbpsReadWrite: to.Ptr[int64](150),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"hana/data": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](4),
	// 				SizeGB: to.Ptr[int64](128),
	// 				SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 					Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](5),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](4),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"hana/log": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](3),
	// 				SizeGB: to.Ptr[int64](128),
	// 				SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 					Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](5),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](3),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"hana/shared": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](1),
	// 				SizeGB: to.Ptr[int64](256),
	// 				SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 					Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P15"),
	// 					IopsReadWrite: to.Ptr[int64](1100),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](125),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P20"),
	// 					IopsReadWrite: to.Ptr[int64](2300),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](150),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P30"),
	// 					IopsReadWrite: to.Ptr[int64](5000),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](200),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](1024),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P40"),
	// 					IopsReadWrite: to.Ptr[int64](7500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](250),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](2048),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P50"),
	// 					IopsReadWrite: to.Ptr[int64](7500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](250),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](4096),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"os": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](1),
	// 				SizeGB: to.Ptr[int64](64),
	// 				SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 					Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P6"),
	// 					IopsReadWrite: to.Ptr[int64](240),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](50),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](64),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 		"usr/sap": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 			RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 				Count: to.Ptr[int64](1),
	// 				SizeGB: to.Ptr[int64](128),
	// 				SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 					Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 				},
	// 			},
	// 			SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](500),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](100),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](128),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](1100),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](125),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](256),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 				},
	// 				{
	// 					DiskTier: to.Ptr("P10"),
	// 					IopsReadWrite: to.Ptr[int64](2300),
	// 					MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 					MbpsReadWrite: to.Ptr[int64](150),
	// 					MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](512),
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 			}},
	// 		},
	// 	},
	// }
}
