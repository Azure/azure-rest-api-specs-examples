package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/LBBackendAddressPoolWithBackendAddressesPut.json
func ExampleLoadBalancerBackendAddressPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancerBackendAddressPoolsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "testrg", "lb", "backend", armnetwork.BackendAddressPool{
		Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
			LoadBalancerBackendAddresses: []*armnetwork.LoadBalancerBackendAddress{
				{
					Name: to.Ptr("address1"),
					Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
						IPAddress: to.Ptr("10.0.0.4"),
						VirtualNetwork: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
						},
					},
				},
				{
					Name: to.Ptr("address2"),
					Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
						IPAddress: to.Ptr("10.0.0.5"),
						VirtualNetwork: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
						},
					},
				}},
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
