package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkCreateWithIpamPool.json
func ExampleVirtualNetworksClient_BeginCreateOrUpdate_createVirtualNetworkWithIpamPool() {
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
				IpamPoolPrefixAllocations: []*armnetwork.IpamPoolPrefixAllocation{
					{
						NumberOfIPAddresses: to.Ptr("65536"),
						Pool: &armnetwork.IpamPoolPrefixAllocationPool{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool"),
						},
					}},
			},
			Subnets: []*armnetwork.Subnet{
				{
					Name: to.Ptr("test-1"),
					Properties: &armnetwork.SubnetPropertiesFormat{
						IpamPoolPrefixAllocations: []*armnetwork.IpamPoolPrefixAllocation{
							{
								NumberOfIPAddresses: to.Ptr("80"),
								Pool: &armnetwork.IpamPoolPrefixAllocationPool{
									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool"),
								},
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armnetwork.VirtualNetwork{
	// 	Name: to.Ptr("test-vnet"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armnetwork.VirtualNetworkPropertiesFormat{
	// 		AddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.0.0.0/22")},
	// 				IpamPoolPrefixAllocations: []*armnetwork.IpamPoolPrefixAllocation{
	// 					{
	// 						AllocatedAddressPrefixes: []*string{
	// 							to.Ptr("10.0.0.0/22")},
	// 							NumberOfIPAddresses: to.Ptr("800"),
	// 							Pool: &armnetwork.IpamPoolPrefixAllocationPool{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool"),
	// 							},
	// 					}},
	// 				},
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				Subnets: []*armnetwork.Subnet{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
	// 						Name: to.Ptr("subnet1"),
	// 						Properties: &armnetwork.SubnetPropertiesFormat{
	// 							AddressPrefixes: []*string{
	// 								to.Ptr("10.0.0.0/26"),
	// 								to.Ptr("10.0.0.64/28")},
	// 								IpamPoolPrefixAllocations: []*armnetwork.IpamPoolPrefixAllocation{
	// 									{
	// 										AllocatedAddressPrefixes: []*string{
	// 											to.Ptr("10.0.0.0/26"),
	// 											to.Ptr("10.0.0.64/28")},
	// 											NumberOfIPAddresses: to.Ptr("80"),
	// 											Pool: &armnetwork.IpamPoolPrefixAllocationPool{
	// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool"),
	// 											},
	// 									}},
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 								},
	// 						}},
	// 					},
	// 				}
}
