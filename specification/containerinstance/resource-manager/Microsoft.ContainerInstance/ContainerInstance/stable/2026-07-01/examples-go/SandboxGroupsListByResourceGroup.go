package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v3"
)

// Generated from example definition: 2026-07-01/SandboxGroupsListByResourceGroup.json
func ExampleSandboxGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSandboxGroupsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armcontainerinstance.SandboxGroupsClientListByResourceGroupResponse{
		// 	SandboxGroupListResult: armcontainerinstance.SandboxGroupListResult{
		// 		Value: []*armcontainerinstance.SandboxGroup{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerInstance/sandboxGroups/mySandboxGroup"),
		// 				Name: to.Ptr("mySandboxGroup"),
		// 				Type: to.Ptr("Microsoft.ContainerInstance/sandboxGroups"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("test"),
		// 				},
		// 				Properties: &armcontainerinstance.SandboxGroupProperties{
		// 					ProvisioningState: to.Ptr(armcontainerinstance.SandboxGroupProvisioningStateSucceeded),
		// 					NetworkProfile: &armcontainerinstance.SandboxGroupNetworkProfile{
		// 						Subnets: []*armcontainerinstance.SubnetReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 							},
		// 						},
		// 					},
		// 					ManagementResourceGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mySandboxGroup-mrg"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
