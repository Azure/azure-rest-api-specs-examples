package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkCreateServiceEndpoints.json
func ExampleVirtualNetworksClient_BeginCreateOrUpdate_virtualNetworkCreateServiceEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewVirtualNetworksClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "vnetTest", "vnet1", armnetwork.VirtualNetwork{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.VirtualNetworkPropertiesFormat{
			AddressSpace: &armnetwork.AddressSpace{
				AddressPrefixes: []*string{
					to.Ptr("10.0.0.0/16")},
			},
			Subnets: []*armnetwork.Subnet{
				{
					Name: to.Ptr("test-1"),
					Properties: &armnetwork.SubnetPropertiesFormat{
						AddressPrefix: to.Ptr("10.0.0.0/16"),
						ServiceEndpoints: []*armnetwork.ServiceEndpointPropertiesFormat{
							{
								Service: to.Ptr("Microsoft.Storage"),
							}},
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
