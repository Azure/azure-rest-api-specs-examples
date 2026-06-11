package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkAppliances_ListBySubscription.json
func ExampleVirtualNetworkAppliancesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkAppliancesClient().NewListAllPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armnetwork.VirtualNetworkAppliancesClientListAllResponse{
		// 	VirtualNetworkApplianceListResult: armnetwork.VirtualNetworkApplianceListResult{
		// 		Value: []*armnetwork.VirtualNetworkAppliance{
		// 			{
		// 				Name: to.Ptr("test-vna"),
		// 				Type: to.Ptr("Microsoft.Network/virtualNetworkAppliances"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetwork.VirtualNetworkAppliancePropertiesFormat{
		// 					IPConfigurations: []*armnetwork.VirtualNetworkApplianceIPConfiguration{
		// 						{
		// 							Name: to.Ptr("ipconfig1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna/ipConfigurations/ipconfig1"),
		// 							Properties: &armnetwork.VirtualNetworkApplianceIPConfigurationProperties{
		// 								Primary: to.Ptr(true),
		// 								PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 								PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 								PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("c9b1f5f4-5d6e-4f0a-89c1-6b3d3e3c3e3c"),
		// 					Subnet: &armnetwork.Subnet{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("test-vna2"),
		// 				Type: to.Ptr("Microsoft.Network/virtualNetworkAppliances"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna2"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetwork.VirtualNetworkAppliancePropertiesFormat{
		// 					IPConfigurations: []*armnetwork.VirtualNetworkApplianceIPConfiguration{
		// 						{
		// 							Name: to.Ptr("ipconfig1"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna/ipConfigurations/ipconfig1"),
		// 							Properties: &armnetwork.VirtualNetworkApplianceIPConfigurationProperties{
		// 								Primary: to.Ptr(true),
		// 								PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 								PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 								PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("c9b1f5f4-5d6e-4f0a-89c1-6b3d3e3c3e3d"),
		// 					Subnet: &armnetwork.Subnet{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
