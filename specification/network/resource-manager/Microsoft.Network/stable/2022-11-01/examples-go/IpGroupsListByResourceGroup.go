package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/IpGroupsListByResourceGroup.json
func ExampleIPGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.IPGroupListResult = armnetwork.IPGroupListResult{
		// 	Value: []*armnetwork.IPGroup{
		// 		{
		// 			Name: to.Ptr("ipGroups1"),
		// 			Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.IPGroupPropertiesFormat{
		// 				FirewallPolicies: []*armnetwork.SubResource{
		// 				},
		// 				Firewalls: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 				}},
		// 				IPAddresses: []*string{
		// 					to.Ptr("13.64.39.16/32"),
		// 					to.Ptr("40.74.146.80/31"),
		// 					to.Ptr("40.74.147.32/28")},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("ipGroups2"),
		// 				Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups"),
		// 				Location: to.Ptr("centralus"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				Properties: &armnetwork.IPGroupPropertiesFormat{
		// 					FirewallPolicies: []*armnetwork.SubResource{
		// 					},
		// 					Firewalls: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 					}},
		// 					IPAddresses: []*string{
		// 						to.Ptr("14.64.39.16/32"),
		// 						to.Ptr("41.74.146.80/31"),
		// 						to.Ptr("42.74.147.32/28")},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					},
		// 			}},
		// 		}
	}
}
