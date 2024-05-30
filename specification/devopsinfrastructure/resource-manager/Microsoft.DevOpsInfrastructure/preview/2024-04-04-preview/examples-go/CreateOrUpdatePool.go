package armdevopsinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/CreateOrUpdatePool.json
func ExamplePoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoolsClient().BeginCreateOrUpdate(ctx, "rg", "pool", armdevopsinfrastructure.Pool{
		Location: to.Ptr("eastus"),
		Properties: &armdevopsinfrastructure.PoolProperties{
			AgentProfile: &armdevopsinfrastructure.StatelessAgentProfile{
				Kind: to.Ptr("Stateless"),
			},
			DevCenterProjectResourceID: to.Ptr("/subscriptions/222e81d0-cf38-4dab-baa5-289bf16baaa4/resourceGroups/rg-1es-devcenter/providers/Microsoft.DevCenter/projects/1ES"),
			FabricProfile: &armdevopsinfrastructure.VmssFabricProfile{
				Kind: to.Ptr("Vmss"),
				Images: []*armdevopsinfrastructure.PoolImage{
					{
						ResourceID: to.Ptr("/MicrosoftWindowsServer/WindowsServer/2019-Datacenter/latest"),
					}},
				SKU: &armdevopsinfrastructure.DevOpsAzureSKU{
					Name: to.Ptr("Standard_D4ads_v5"),
				},
			},
			MaximumConcurrency: to.Ptr[int32](10),
			OrganizationProfile: &armdevopsinfrastructure.AzureDevOpsOrganizationProfile{
				Kind: to.Ptr("AzureDevOps"),
				Organizations: []*armdevopsinfrastructure.Organization{
					{
						URL: to.Ptr("https://mseng.visualstudio.com"),
					}},
			},
			ProvisioningState: to.Ptr(armdevopsinfrastructure.ProvisioningStateSucceeded),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Pool = armdevopsinfrastructure.Pool{
	// 	ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/resourceGroups/rg/providers/Microsoft.DevOpsInfrastructure/Pools/pool"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armdevopsinfrastructure.PoolProperties{
	// 		AgentProfile: &armdevopsinfrastructure.StatelessAgentProfile{
	// 			Kind: to.Ptr("Stateless"),
	// 		},
	// 		DevCenterProjectResourceID: to.Ptr("/subscriptions/222e81d0-cf38-4dab-baa5-289bf16baaa4/resourceGroups/rg-1es-devcenter/providers/Microsoft.DevCenter/projects/1ES"),
	// 		FabricProfile: &armdevopsinfrastructure.VmssFabricProfile{
	// 			Kind: to.Ptr("Vmss"),
	// 			Images: []*armdevopsinfrastructure.PoolImage{
	// 				{
	// 					ResourceID: to.Ptr("/MicrosoftWindowsServer/WindowsServer/2019-Datacenter/latest"),
	// 			}},
	// 			SKU: &armdevopsinfrastructure.DevOpsAzureSKU{
	// 				Name: to.Ptr("Standard_D4ads_v5"),
	// 			},
	// 		},
	// 		MaximumConcurrency: to.Ptr[int32](10),
	// 		OrganizationProfile: &armdevopsinfrastructure.AzureDevOpsOrganizationProfile{
	// 			Kind: to.Ptr("AzureDevOps"),
	// 			Organizations: []*armdevopsinfrastructure.Organization{
	// 				{
	// 					URL: to.Ptr("https://mseng.visualstudio.com"),
	// 			}},
	// 		},
	// 		ProvisioningState: to.Ptr(armdevopsinfrastructure.ProvisioningStateSucceeded),
	// 	},
	// }
}
