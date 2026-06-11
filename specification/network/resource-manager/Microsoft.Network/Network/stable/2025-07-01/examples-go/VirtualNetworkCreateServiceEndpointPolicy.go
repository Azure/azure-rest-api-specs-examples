package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkCreateServiceEndpointPolicy.json
func ExampleVirtualNetworksClient_BeginCreateOrUpdate_createVirtualNetworkWithServiceEndpointsAndServiceEndpointPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworksClient().BeginCreateOrUpdate(ctx, "vnetTest", "vnet1", armnetwork.VirtualNetwork{
		Location: to.Ptr("eastus2euap"),
		Properties: &armnetwork.VirtualNetworkPropertiesFormat{
			AddressSpace: &armnetwork.AddressSpace{
				AddressPrefixes: []*string{
					to.Ptr("10.0.0.0/16"),
				},
			},
			Subnets: []*armnetwork.Subnet{
				{
					Name: to.Ptr("test-1"),
					Properties: &armnetwork.SubnetPropertiesFormat{
						AddressPrefix: to.Ptr("10.0.0.0/16"),
						ServiceEndpointPolicies: []*armnetwork.ServiceEndpointPolicy{
							{
								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetTest/providers/Microsoft.Network/serviceEndpointPolicies/ServiceEndpointPolicy1"),
							},
						},
						ServiceEndpoints: []*armnetwork.ServiceEndpointPropertiesFormat{
							{
								Service: to.Ptr("Microsoft.Storage"),
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.VirtualNetworksClientCreateOrUpdateResponse{
	// 	VirtualNetwork: armnetwork.VirtualNetwork{
	// 		Name: to.Ptr("vnet1"),
	// 		Type: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetTest/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Properties: &armnetwork.VirtualNetworkPropertiesFormat{
	// 			AddressSpace: &armnetwork.AddressSpace{
	// 				AddressPrefixes: []*string{
	// 					to.Ptr("10.0.0.0/16"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Subnets: []*armnetwork.Subnet{
	// 				{
	// 					Name: to.Ptr("test-1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetTest/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/test-1"),
	// 					Properties: &armnetwork.SubnetPropertiesFormat{
	// 						AddressPrefix: to.Ptr("10.0.0.0/16"),
	// 						IPConfigurations: []*armnetwork.IPConfiguration{
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						ResourceNavigationLinks: []*armnetwork.ResourceNavigationLink{
	// 						},
	// 						ServiceEndpointPolicies: []*armnetwork.ServiceEndpointPolicy{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetTest/providers/Microsoft.Network/serviceEndpointPolicies/ServiceEndpointPolicy1"),
	// 							},
	// 						},
	// 						ServiceEndpoints: []*armnetwork.ServiceEndpointPropertiesFormat{
	// 							{
	// 								Locations: []*string{
	// 									to.Ptr("eastus2(stage)"),
	// 									to.Ptr("usnorth(stage)"),
	// 								},
	// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 								Service: to.Ptr("Microsoft.Storage"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			VirtualNetworkPeerings: []*armnetwork.VirtualNetworkPeering{
	// 			},
	// 		},
	// 	},
	// }
}
