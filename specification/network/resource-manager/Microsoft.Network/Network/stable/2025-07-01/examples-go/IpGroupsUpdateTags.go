package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/IpGroupsUpdateTags.json
func ExampleIPGroupsClient_UpdateGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPGroupsClient().UpdateGroups(ctx, "myResourceGroup", "ipGroups1", armnetwork.TagsObject{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.IPGroupsClientUpdateGroupsResponse{
	// 	IPGroup: armnetwork.IPGroup{
	// 		Name: to.Ptr("ipGroups1"),
	// 		Type: to.Ptr("Microsoft.Network/ipGroups"),
	// 		Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/resourceGroups/myResourceGroup/ipGroups/ipGroups1"),
	// 		Location: to.Ptr("westcentralus"),
	// 		Properties: &armnetwork.IPGroupPropertiesFormat{
	// 			FirewallPolicies: []*armnetwork.SubResource{
	// 			},
	// 			Firewalls: []*armnetwork.SubResource{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 				},
	// 			},
	// 			IPAddresses: []*string{
	// 				to.Ptr("13.64.39.16/32"),
	// 				to.Ptr("40.74.146.80/31"),
	// 				to.Ptr("40.74.147.32/28"),
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
