package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/PrivateEndpointCreateWithASG.json
func ExamplePrivateEndpointsClient_BeginCreateOrUpdate_createPrivateEndpointWithApplicationSecurityGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewPrivateEndpointsClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "testPe", armnetwork.PrivateEndpoint{
		Location: to.Ptr("eastus2euap"),
		Properties: &armnetwork.PrivateEndpointProperties{
			ApplicationSecurityGroups: []*armnetwork.ApplicationSecurityGroup{
				{
					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/applicationSecurityGroup/asg1"),
				}},
			PrivateLinkServiceConnections: []*armnetwork.PrivateLinkServiceConnection{
				{
					Properties: &armnetwork.PrivateLinkServiceConnectionProperties{
						GroupIDs: []*string{
							to.Ptr("groupIdFromResource")},
						PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
						RequestMessage:       to.Ptr("Please approve my connection."),
					},
				}},
			Subnet: &armnetwork.Subnet{
				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
