package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/VirtualNetworkAppliances_CreateOrUpdate.json
func ExampleVirtualNetworkAppliancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworkAppliancesClient().BeginCreateOrUpdate(ctx, "rg1", "test-vna", armnetwork.VirtualNetworkAppliance{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.VirtualNetworkAppliancePropertiesFormat{
			BandwidthInGbps: to.Ptr[float64](100),
			Subnet: &armnetwork.Subnet{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
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
	// res = armnetwork.VirtualNetworkAppliancesClientCreateOrUpdateResponse{
	// 	VirtualNetworkAppliance: armnetwork.VirtualNetworkAppliance{
	// 		Name: to.Ptr("test-vna"),
	// 		Type: to.Ptr("Microsoft.Network/virtualNetworkAppliances"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetwork.VirtualNetworkAppliancePropertiesFormat{
	// 			IPConfigurations: []*armnetwork.VirtualNetworkApplianceIPConfiguration{
	// 				{
	// 					Name: to.Ptr("ipconfig1"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna/ipConfigurations/ipconfig1"),
	// 					Properties: &armnetwork.VirtualNetworkApplianceIPConfigurationProperties{
	// 						Primary: to.Ptr(true),
	// 						PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 						PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("c9b1f5f4-5d6e-4f0a-89c1-6b3d3e3c3e3c"),
	// 			Subnet: &armnetwork.Subnet{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
	// 			},
	// 		},
	// 	},
	// }
}
