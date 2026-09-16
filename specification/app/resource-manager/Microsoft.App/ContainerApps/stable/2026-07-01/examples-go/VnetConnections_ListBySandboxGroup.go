package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/VnetConnections_ListBySandboxGroup.json
func ExampleVnetConnectionsClient_NewListBySandboxGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVnetConnectionsClient().NewListBySandboxGroupPager("myRg", "testgroup", nil)
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
		// page = armappcontainers.VnetConnectionsClientListBySandboxGroupResponse{
		// 	VnetConnectionListResult: armappcontainers.VnetConnectionListResult{
		// 		Value: []*armappcontainers.VnetConnection{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRg/providers/Microsoft.App/sandboxGroups/testgroup/vnetConnections/myVnetConnection"),
		// 				Name: to.Ptr("myVnetConnection"),
		// 				Type: to.Ptr("Microsoft.App/sandboxGroups/vnetConnections"),
		// 				Properties: &armappcontainers.VnetConnectionProperties{
		// 					ProvisioningState: to.Ptr(armappcontainers.VnetConnectionProvisioningStateSucceeded),
		// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRg/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
