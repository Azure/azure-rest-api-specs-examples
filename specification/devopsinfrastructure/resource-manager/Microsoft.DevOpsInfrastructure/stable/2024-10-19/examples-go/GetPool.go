package armdevopsinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
)

// Generated from example definition: 2024-10-19/GetPool.json
func ExamplePoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("a2e95d27-c161-4b61-bda4-11512c14c2c2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolsClient().Get(ctx, "rg", "pool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevopsinfrastructure.PoolsClientGetResponse{
	// 	Pool: &armdevopsinfrastructure.Pool{
	// 		Properties: &armdevopsinfrastructure.PoolProperties{
	// 			ProvisioningState: to.Ptr(armdevopsinfrastructure.ProvisioningStateSucceeded),
	// 			MaximumConcurrency: to.Ptr[int32](10),
	// 			DevCenterProjectResourceID: to.Ptr("/subscriptions/222e81d0-cf38-4dab-baa5-289bf16baaa4/resourceGroups/rg-1es-devcenter/providers/Microsoft.DevCenter/projects/1ES"),
	// 			OrganizationProfile: &armdevopsinfrastructure.AzureDevOpsOrganizationProfile{
	// 				Kind: to.Ptr("AzureDevOps"),
	// 				Organizations: []*armdevopsinfrastructure.Organization{
	// 					{
	// 						URL: to.Ptr("https://mseng.visualstudio.com"),
	// 					},
	// 				},
	// 			},
	// 			AgentProfile: &armdevopsinfrastructure.StatelessAgentProfile{
	// 				Kind: to.Ptr("Stateless"),
	// 			},
	// 			FabricProfile: &armdevopsinfrastructure.VmssFabricProfile{
	// 				Kind: to.Ptr("Vmss"),
	// 				SKU: &armdevopsinfrastructure.DevOpsAzureSKU{
	// 					Name: to.Ptr("Standard_D4ads_v5"),
	// 				},
	// 				Images: []*armdevopsinfrastructure.PoolImage{
	// 					{
	// 						ResourceID: to.Ptr("/MicrosoftWindowsServer/WindowsServer/2019-Datacenter/latest"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/a2e95d27-c161-4b61-bda4-11512c14c2c2/resourceGroups/rg/providers/Microsoft.DevOpsInfrastructure/Pools/pool"),
	// 		Location: to.Ptr("eastus"),
	// 	},
	// }
}
