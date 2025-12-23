package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkInterfaceGet.json
func ExampleInterfacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInterfacesClient().Get(ctx, "rg1", "test-nic", &armnetwork.InterfacesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Interface = armnetwork.Interface{
	// 	Name: to.Ptr("test-nic"),
	// 	Type: to.Ptr("Microsoft.Network/networkInterfaces"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.InterfacePropertiesFormat{
	// 		DisableTCPStateTracking: to.Ptr(true),
	// 		DNSSettings: &armnetwork.InterfaceDNSSettings{
	// 			AppliedDNSServers: []*string{
	// 			},
	// 			DNSServers: []*string{
	// 			},
	// 			InternalDomainNameSuffix: to.Ptr("test.bx.internal.cloudapp.net"),
	// 		},
	// 		DscpConfiguration: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/dscpConfiguration/mydscpconfiguration"),
	// 		},
	// 		EnableAcceleratedNetworking: to.Ptr(true),
	// 		EnableIPForwarding: to.Ptr(false),
	// 		IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic/ipConfigurations/ipconfig1"),
	// 				Name: to.Ptr("ipconfig1"),
	// 				Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
	// 					Primary: to.Ptr(true),
	// 					PrivateIPAddress: to.Ptr("172.20.2.4"),
	// 					PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.PublicIPAddress{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"),
	// 					},
	// 					Subnet: &armnetwork.Subnet{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic/ipConfigurations/ipconfig2"),
	// 				Name: to.Ptr("ipconfig2"),
	// 				Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
	// 					Primary: to.Ptr(false),
	// 					PrivateIPAddress: to.Ptr("172.20.2.16/28"),
	// 					PrivateIPAddressPrefixLength: to.Ptr[int32](28),
	// 					PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Subnet: &armnetwork.Subnet{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
	// 					},
	// 				},
	// 		}},
	// 		MacAddress: to.Ptr("00-0D-3A-1B-C7-21"),
	// 		NetworkSecurityGroup: &armnetwork.SecurityGroup{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg"),
	// 		},
	// 		Primary: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VirtualMachine: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"),
	// 		},
	// 		VnetEncryptionSupported: to.Ptr(false),
	// 	},
	// }
}
