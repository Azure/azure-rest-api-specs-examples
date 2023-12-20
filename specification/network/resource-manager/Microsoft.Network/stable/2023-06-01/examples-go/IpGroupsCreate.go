package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/IpGroupsCreate.json
func ExampleIPGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPGroupsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "ipGroups1", armnetwork.IPGroup{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.IPGroupPropertiesFormat{
			IPAddresses: []*string{
				to.Ptr("13.64.39.16/32"),
				to.Ptr("40.74.146.80/31"),
				to.Ptr("40.74.147.32/28")},
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
	// res.IPGroup = armnetwork.IPGroup{
	// 	Name: to.Ptr("ipGroups1"),
	// 	Type: to.Ptr("Microsoft.Network/ipGroups"),
	// 	ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups/ipGroups1"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.IPGroupPropertiesFormat{
	// 		FirewallPolicies: []*armnetwork.SubResource{
	// 		},
	// 		Firewalls: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 		}},
	// 		IPAddresses: []*string{
	// 			to.Ptr("13.64.39.16/32"),
	// 			to.Ptr("40.74.146.80/31"),
	// 			to.Ptr("40.74.147.32/28")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
