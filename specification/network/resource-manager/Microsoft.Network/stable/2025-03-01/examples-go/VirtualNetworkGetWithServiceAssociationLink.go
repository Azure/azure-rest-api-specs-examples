package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGetWithServiceAssociationLink.json
func ExampleVirtualNetworksClient_Get_getVirtualNetworkWithServiceAssociationLinks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworksClient().Get(ctx, "rg1", "test-vnet", &armnetwork.VirtualNetworksClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armnetwork.VirtualNetwork{
	// 	Name: to.Ptr("test-vnet"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworks"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.VirtualNetworkPropertiesFormat{
	// 		AddressSpace: &armnetwork.AddressSpace{
	// 			AddressPrefixes: []*string{
	// 				to.Ptr("10.0.0.0/16")},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			Subnets: []*armnetwork.Subnet{
	// 				{
	// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"),
	// 					Name: to.Ptr("subnet1"),
	// 					Etag: to.Ptr("W/\"4d3e91b4-f67f-48be-880b-e4a8abdd019e\""),
	// 					Properties: &armnetwork.SubnetPropertiesFormat{
	// 						AddressPrefix: to.Ptr("10.0.214.0/24"),
	// 						Delegations: []*armnetwork.Delegation{
	// 							{
	// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1/delegations/aciDelegation"),
	// 								Name: to.Ptr("aciDelegation"),
	// 								Etag: to.Ptr("W/\"4d3e91b4-f67f-48be-880b-e4a8abdd019e\""),
	// 								Properties: &armnetwork.ServiceDelegationPropertiesFormat{
	// 									Actions: []*string{
	// 										to.Ptr("Microsoft.Network/virtualNetworks/subnets/action")},
	// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 										ServiceName: to.Ptr("Microsoft.Provider/resourceType"),
	// 									},
	// 							}},
	// 							IPConfigurationProfiles: []*armnetwork.IPConfigurationProfile{
	// 								{
	// 									ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0/ipConfigurations/ipconfigprofile1"),
	// 							}},
	// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 							ServiceAssociationLinks: []*armnetwork.ServiceAssociationLink{
	// 								{
	// 									ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1/serviceAssociationLinks/serviceAssociationLink1"),
	// 									Name: to.Ptr("serviceAssociationLink1"),
	// 									Etag: to.Ptr("W/\"4d3e91b4-f67f-48be-880b-e4a8abdd019e\""),
	// 									Properties: &armnetwork.ServiceAssociationLinkPropertiesFormat{
	// 										LinkedResourceType: to.Ptr("Microsoft.Provider/resourceType"),
	// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 									},
	// 							}},
	// 							ServiceEndpoints: []*armnetwork.ServiceEndpointPropertiesFormat{
	// 							},
	// 						},
	// 				}},
	// 				VirtualNetworkPeerings: []*armnetwork.VirtualNetworkPeering{
	// 				},
	// 			},
	// 		}
}
