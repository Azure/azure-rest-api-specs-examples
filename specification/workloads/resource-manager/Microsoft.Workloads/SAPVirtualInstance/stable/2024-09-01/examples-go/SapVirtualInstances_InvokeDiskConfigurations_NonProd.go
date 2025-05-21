package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapVirtualInstances_InvokeDiskConfigurations_NonProd.json
func ExampleSAPVirtualInstancesClient_GetDiskConfigurations_sapDiskConfigurationsForInputEnvironmentNonProd() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPVirtualInstancesClient().GetDiskConfigurations(ctx, "centralus", armworkloadssapvirtualinstance.SAPDiskConfigurationsRequest{
		AppLocation:    to.Ptr("eastus"),
		SapProduct:     to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		Environment:    to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeNonProd),
		DatabaseType:   to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
		DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
		DbVMSKU:        to.Ptr("Standard_M32ts"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadssapvirtualinstance.SAPVirtualInstancesClientGetDiskConfigurationsResponse{
	// 	SAPDiskConfigurationsResult: &armworkloadssapvirtualinstance.SAPDiskConfigurationsResult{
	// 		VolumeConfigurations: map[string]*armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 			"hana/data": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 				RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 					Count: to.Ptr[int64](4),
	// 					SizeGB: to.Ptr[int64](128),
	// 				},
	// 				SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](4),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](5),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](100),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 				},
	// 			},
	// 			"hana/log": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 				RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 					Count: to.Ptr[int64](3),
	// 					SizeGB: to.Ptr[int64](128),
	// 				},
	// 				SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](3),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](5),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](100),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 				},
	// 			},
	// 			"hana/shared": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 				RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 					},
	// 					Count: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](256),
	// 				},
	// 				SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](256),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](1100),
	// 						MbpsReadWrite: to.Ptr[int64](125),
	// 						DiskTier: to.Ptr("P15"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](512),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](2300),
	// 						MbpsReadWrite: to.Ptr[int64](150),
	// 						DiskTier: to.Ptr("P20"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](1024),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](5000),
	// 						MbpsReadWrite: to.Ptr[int64](200),
	// 						DiskTier: to.Ptr("P30"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](2048),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](7500),
	// 						MbpsReadWrite: to.Ptr[int64](250),
	// 						DiskTier: to.Ptr("P40"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](4096),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](7500),
	// 						MbpsReadWrite: to.Ptr[int64](250),
	// 						DiskTier: to.Ptr("P50"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](256),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E15"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](512),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E20"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](1024),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E30"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](2048),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E40"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](4096),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E50"),
	// 					},
	// 				},
	// 			},
	// 			"usr/sap": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 				RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 					},
	// 					Count: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](128),
	// 				},
	// 				SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](100),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](256),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](1100),
	// 						MbpsReadWrite: to.Ptr[int64](125),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](512),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](2300),
	// 						MbpsReadWrite: to.Ptr[int64](150),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E10"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](256),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E15"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](512),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E20"),
	// 					},
	// 				},
	// 			},
	// 			"os": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 				RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 					},
	// 					Count: to.Ptr[int64](1),
	// 					SizeGB: to.Ptr[int64](64),
	// 				},
	// 				SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](64),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](240),
	// 						MbpsReadWrite: to.Ptr[int64](50),
	// 						DiskTier: to.Ptr("P6"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](100),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](64),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("P6"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](1),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](1),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 				},
	// 			},
	// 			"backup": &armworkloadssapvirtualinstance.SAPDiskConfiguration{
	// 				RecommendedConfiguration: &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 					SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 						Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 					},
	// 					Count: to.Ptr[int64](2),
	// 					SizeGB: to.Ptr[int64](256),
	// 				},
	// 				SupportedConfigurations: []*armworkloadssapvirtualinstance.DiskDetails{
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E10"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](256),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E15"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](512),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](60),
	// 						DiskTier: to.Ptr("E20"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](128),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 						IopsReadWrite: to.Ptr[int64](500),
	// 						MbpsReadWrite: to.Ptr[int64](100),
	// 						DiskTier: to.Ptr("P10"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](256),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 						IopsReadWrite: to.Ptr[int64](1100),
	// 						MbpsReadWrite: to.Ptr[int64](125),
	// 						DiskTier: to.Ptr("P15"),
	// 					},
	// 					{
	// 						SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 							Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 						},
	// 						SizeGB: to.Ptr[int64](512),
	// 						MinimumSupportedDiskCount: to.Ptr[int64](0),
	// 						MaximumSupportedDiskCount: to.Ptr[int64](6),
	// 						IopsReadWrite: to.Ptr[int64](2300),
	// 						MbpsReadWrite: to.Ptr[int64](150),
	// 						DiskTier: to.Ptr("P20"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
