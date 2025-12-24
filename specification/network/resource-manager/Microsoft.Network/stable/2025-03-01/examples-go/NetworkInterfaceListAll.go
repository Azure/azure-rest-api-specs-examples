package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkInterfaceListAll.json
func ExampleInterfacesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInterfacesClient().NewListAllPager(nil)
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
		// page.InterfaceListResult = armnetwork.InterfaceListResult{
		// 	Value: []*armnetwork.Interface{
		// 		{
		// 			Name: to.Ptr("test-nic"),
		// 			Type: to.Ptr("Microsoft.Network/networkInterfaces"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.InterfacePropertiesFormat{
		// 				DisableTCPStateTracking: to.Ptr(true),
		// 				DNSSettings: &armnetwork.InterfaceDNSSettings{
		// 					AppliedDNSServers: []*string{
		// 					},
		// 					DNSServers: []*string{
		// 					},
		// 					InternalDomainNameSuffix: to.Ptr("test.bx.internal.cloudapp.net"),
		// 				},
		// 				EnableAcceleratedNetworking: to.Ptr(true),
		// 				EnableIPForwarding: to.Ptr(false),
		// 				IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-nic/ipConfigurations/ipconfig1"),
		// 						Name: to.Ptr("ipconfig1"),
		// 						Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
		// 							Primary: to.Ptr(true),
		// 							PrivateIPAddress: to.Ptr("172.20.2.4"),
		// 							PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							PublicIPAddress: &armnetwork.PublicIPAddress{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"),
		// 							},
		// 							Subnet: &armnetwork.Subnet{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
		// 							},
		// 						},
		// 				}},
		// 				MacAddress: to.Ptr("00-0D-3A-1B-C7-21"),
		// 				NetworkSecurityGroup: &armnetwork.SecurityGroup{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg"),
		// 				},
		// 				Primary: to.Ptr(true),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				VirtualMachine: &armnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1"),
		// 				},
		// 				VnetEncryptionSupported: to.Ptr(false),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-nic2"),
		// 			Type: to.Ptr("Microsoft.Network/networkInterfaces"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rgnew/providers/Microsoft.Network/networkInterfaces/test-nic2"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.InterfacePropertiesFormat{
		// 				DisableTCPStateTracking: to.Ptr(true),
		// 				DNSSettings: &armnetwork.InterfaceDNSSettings{
		// 					AppliedDNSServers: []*string{
		// 					},
		// 					DNSServers: []*string{
		// 					},
		// 				},
		// 				EnableAcceleratedNetworking: to.Ptr(true),
		// 				EnableIPForwarding: to.Ptr(false),
		// 				IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rgnew/providers/Microsoft.Network/networkInterfaces/test-nic2/ipConfigurations/ipconfig1"),
		// 						Name: to.Ptr("ipconfig1"),
		// 						Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
		// 							Primary: to.Ptr(true),
		// 							PrivateIPAddress: to.Ptr("172.20.2.4"),
		// 							PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							PublicIPAddress: &armnetwork.PublicIPAddress{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rgnew/providers/Microsoft.Network/publicIPAddresses/test-ip2"),
		// 							},
		// 							Subnet: &armnetwork.Subnet{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rgnew/providers/Microsoft.Network/virtualNetworks/rgnew-vnet2/subnets/default"),
		// 							},
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				VnetEncryptionSupported: to.Ptr(false),
		// 			},
		// 	}},
		// }
	}
}
