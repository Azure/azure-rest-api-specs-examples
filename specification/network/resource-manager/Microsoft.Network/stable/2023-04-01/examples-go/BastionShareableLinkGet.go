package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/BastionShareableLinkGet.json
func ExampleManagementClient_NewGetBastionShareableLinkPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewGetBastionShareableLinkPager("rg1", "bastionhosttenant", armnetwork.BastionShareableLinkListRequest{
		VMs: []*armnetwork.BastionShareableLink{
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
				},
			},
			{
				VM: &armnetwork.VM{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
				},
			}},
	}, nil)
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
		// page.BastionShareableLinkListResult = armnetwork.BastionShareableLinkListResult{
		// 	Value: []*armnetwork.BastionShareableLink{
		// 		{
		// 			Bsl: to.Ptr("http://bst-bastionhostid.bastion.com/api/shareable-url/tokenvm1"),
		// 			CreatedAt: to.Ptr("2019-10-18T12:00:00.0000Z"),
		// 			VM: &armnetwork.VM{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 			},
		// 		},
		// 		{
		// 			Bsl: to.Ptr("http://bst-bastionhostid.bastion.com/api/shareable-url/tokenvm2"),
		// 			CreatedAt: to.Ptr("2019-10-17T12:00:00.0000Z"),
		// 			VM: &armnetwork.VM{
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rgx/providers/Microsoft.Compute/virtualMachines/vm2"),
		// 			},
		// 	}},
		// }
	}
}
