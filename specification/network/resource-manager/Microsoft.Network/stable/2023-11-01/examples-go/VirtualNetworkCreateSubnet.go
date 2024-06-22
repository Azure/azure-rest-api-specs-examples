package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/VirtualNetworkCreateSubnet.json
func ExampleVirtualNetworksClient_BeginCreateOrUpdate_createVirtualNetworkWithSubnet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworksClient().BeginCreateOrUpdate(ctx, "rg1", "test-vnet", armnetwork.VirtualNetwork{
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
						AddressPrefix: to.Ptr("10.0.0.0/24"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armnetwork.VirtualNetwork{
	// 	Name: to.Ptr("test-vnet"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.VirtualNetworkPropertiesFormat{
	// 		AddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.0.0.0/16")},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Subnets: []*armnetwork.Subnet{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-1"),
	// 					Name: to.Ptr("test-1"),
	// 					Properties: &armnetwork.SubnetPropertiesFormat{
	// 						AddressPrefix: to.Ptr("10.0.0.0/24"),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 			}},
	// 			VirtualNetworkPeerings: []*armnetwork.VirtualNetworkPeering{
	// 			},
	// 		},
	// 	}
}
