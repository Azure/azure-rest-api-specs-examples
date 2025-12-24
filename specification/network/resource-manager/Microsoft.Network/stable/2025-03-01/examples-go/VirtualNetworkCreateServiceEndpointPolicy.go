package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkCreateServiceEndpointPolicy.json
func ExampleVirtualNetworksClient_BeginCreateOrUpdate_createVirtualNetworkWithServiceEndpointsAndServiceEndpointPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworksClient().BeginCreateOrUpdate(ctx, "vnetTest", "vnet1", armnetwork.VirtualNetwork{
		Location: to.Ptr("eastus2euap"),
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
						ServiceEndpointPolicies: []*armnetwork.ServiceEndpointPolicy{
							{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/vnetTest/providers/Microsoft.Network/serviceEndpointPolicies/ServiceEndpointPolicy1"),
							}},
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armnetwork.VirtualNetwork{
	// 	Name: to.Ptr("vnet1"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/vnetTest/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armnetwork.VirtualNetworkPropertiesFormat{
	// 		AddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.0.0.0/16")},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Subnets: []*armnetwork.Subnet{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/vnetTest/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/test-1"),
	// 					Name: to.Ptr("test-1"),
	// 					Properties: &armnetwork.SubnetPropertiesFormat{
	// 						AddressPrefix: to.Ptr("10.0.0.0/16"),
	// 						IPConfigurations: []*armnetwork.IPConfiguration{
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						ResourceNavigationLinks: []*armnetwork.ResourceNavigationLink{
	// 						},
	// 						ServiceEndpointPolicies: []*armnetwork.ServiceEndpointPolicy{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/vnetTest/providers/Microsoft.Network/serviceEndpointPolicies/ServiceEndpointPolicy1"),
	// 						}},
	// 						ServiceEndpoints: []*armnetwork.ServiceEndpointPropertiesFormat{
	// 							{
	// 								Locations: []*string{
	// 									to.Ptr("eastus2(stage)"),
	// 									to.Ptr("usnorth(stage)")},
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 									Service: to.Ptr("Microsoft.Storage"),
	// 							}},
	// 						},
	// 				}},
	// 				VirtualNetworkPeerings: []*armnetwork.VirtualNetworkPeering{
	// 				},
	// 			},
	// 		}
}
