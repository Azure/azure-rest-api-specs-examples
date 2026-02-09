package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkAppliances_Get.json
func ExampleVirtualNetworkAppliancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkAppliancesClient().Get(ctx, "rg1", "test-vna", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkAppliance = armnetwork.VirtualNetworkAppliance{
	// 	Name: to.Ptr("test-vna"),
	// 	Type: to.Ptr("Microsoft.Network/virtualNetworkAppliances"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.VirtualNetworkAppliancePropertiesFormat{
	// 		IPConfigurations: []*armnetwork.VirtualNetworkApplianceIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkAppliances/test-vna/ipConfigurations/ipconfig1"),
	// 				Name: to.Ptr("ipconfig1"),
	// 				Properties: &armnetwork.VirtualNetworkApplianceIPConfigurationProperties{
	// 					Primary: to.Ptr(true),
	// 					PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 					PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("c9b1f5f4-5d6e-4f0a-89c1-6b3d3e3c3e3c"),
	// 		Subnet: &armnetwork.Subnet{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
	// 		},
	// 	},
	// }
}
