package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2026-07-01/SandboxGroupsCreateOrUpdate.json
func ExampleSandboxGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSandboxGroupsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "mySandboxGroup", armcontainerinstance.SandboxGroup{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"environment": to.Ptr("test"),
		},
		Properties: &armcontainerinstance.SandboxGroupProperties{
			NetworkProfile: &armcontainerinstance.SandboxGroupNetworkProfile{
				Subnets: []*armcontainerinstance.SubnetReference{
					{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerinstance.SandboxGroupsClientCreateOrUpdateResponse{
	// 	SandboxGroup: armcontainerinstance.SandboxGroup{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerInstance/sandboxGroups/mySandboxGroup"),
	// 		Name: to.Ptr("mySandboxGroup"),
	// 		Type: to.Ptr("Microsoft.ContainerInstance/sandboxGroups"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("test"),
	// 		},
	// 		Properties: &armcontainerinstance.SandboxGroupProperties{
	// 			ProvisioningState: to.Ptr(armcontainerinstance.SandboxGroupProvisioningStateSucceeded),
	// 			NetworkProfile: &armcontainerinstance.SandboxGroupNetworkProfile{
	// 				Subnets: []*armcontainerinstance.SubnetReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
	// 					},
	// 				},
	// 			},
	// 			ManagementResourceGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mySandboxGroup-mrg"),
	// 		},
	// 	},
	// }
}
