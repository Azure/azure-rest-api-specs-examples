package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkInterfaceCreate.json
func ExampleInterfacesClient_BeginCreateOrUpdate_createNetworkInterface() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInterfacesClient().BeginCreateOrUpdate(ctx, "rg1", "test-nic", armnetwork.Interface{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.InterfacePropertiesFormat{
			DisableTCPStateTracking:     to.Ptr(true),
			EnableAcceleratedNetworking: to.Ptr(true),
			IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
				{
					Name: to.Ptr("ipconfig1"),
					Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
						PublicIPAddress: &armnetwork.PublicIPAddress{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"),
						},
						Subnet: &armnetwork.Subnet{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
						},
					},
				},
				{
					Name: to.Ptr("ipconfig2"),
					Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
						PrivateIPAddressPrefixLength: to.Ptr[int32](28),
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
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VnetEncryptionSupported: to.Ptr(false),
	// 	},
	// }
}
