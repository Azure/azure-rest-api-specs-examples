package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VmssNetworkInterfaceIpConfigGet.json
func ExampleInterfacesClient_GetVirtualMachineScaleSetIPConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInterfacesClient().GetVirtualMachineScaleSetIPConfiguration(ctx, "rg1", "vmss1", "2", "nic1", "ip1", &armnetwork.InterfacesClientGetVirtualMachineScaleSetIPConfigurationOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InterfaceIPConfiguration = armnetwork.InterfaceIPConfiguration{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachineScaleSets/vmss1/virtualMachines/2/networkInterfaces/nic1/ipConfigurations/ip1"),
	// 	Name: to.Ptr("ip1"),
	// 	Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
	// 		LoadBalancerBackendAddressPools: []*armnetwork.BackendAddressPool{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/addressPool1"),
	// 		}},
	// 		LoadBalancerInboundNatRules: []*armnetwork.InboundNatRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natPool1.2"),
	// 		}},
	// 		Primary: to.Ptr(true),
	// 		PrivateIPAddress: to.Ptr("10.0.0.6"),
	// 		PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Subnet: &armnetwork.Subnet{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 		},
	// 	},
	// }
}
