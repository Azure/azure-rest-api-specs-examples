package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/IpGroupsListByResourceGroup.json
func ExampleIPGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPGroupsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armnetwork.IPGroupsClientListByResourceGroupResponse{
		// 	IPGroupListResult: armnetwork.IPGroupListResult{
		// 		Value: []*armnetwork.IPGroup{
		// 			{
		// 				Name: to.Ptr("ipGroups1"),
		// 				Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/myResourceGroup/ipGroups"),
		// 				Location: to.Ptr("westcentralus"),
		// 				Properties: &armnetwork.IPGroupPropertiesFormat{
		// 					FirewallPolicies: []*armnetwork.SubResource{
		// 					},
		// 					Firewalls: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 						},
		// 					},
		// 					IPAddresses: []*string{
		// 						to.Ptr("13.64.39.16/32"),
		// 						to.Ptr("40.74.146.80/31"),
		// 						to.Ptr("40.74.147.32/28"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("ipGroups2"),
		// 				Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/myResourceGroup/ipGroups"),
		// 				Location: to.Ptr("centralus"),
		// 				Properties: &armnetwork.IPGroupPropertiesFormat{
		// 					FirewallPolicies: []*armnetwork.SubResource{
		// 					},
		// 					Firewalls: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 						},
		// 					},
		// 					IPAddresses: []*string{
		// 						to.Ptr("14.64.39.16/32"),
		// 						to.Ptr("41.74.146.80/31"),
		// 						to.Ptr("42.74.147.32/28"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
