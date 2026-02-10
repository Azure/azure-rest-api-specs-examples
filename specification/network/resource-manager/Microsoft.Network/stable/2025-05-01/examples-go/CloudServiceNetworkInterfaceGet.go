package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CloudServiceNetworkInterfaceGet.json
func ExampleInterfacesClient_GetCloudServiceNetworkInterface() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInterfacesClient().GetCloudServiceNetworkInterface(ctx, "rg1", "cs1", "TestVMRole_IN_0", "nic1", &armnetwork.InterfacesClientGetCloudServiceNetworkInterfaceOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Interface = armnetwork.Interface{
	// 	Name: to.Ptr("nic1"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/TestVMRole_IN_0/networkInterfaces/nic1"),
	// 	Properties: &armnetwork.InterfacePropertiesFormat{
	// 		DNSSettings: &armnetwork.InterfaceDNSSettings{
	// 			AppliedDNSServers: []*string{
	// 			},
	// 			DNSServers: []*string{
	// 			},
	// 			InternalDomainNameSuffix: to.Ptr("dns.cdmx.internal.cloudapp.net"),
	// 		},
	// 		EnableAcceleratedNetworking: to.Ptr(false),
	// 		EnableIPForwarding: to.Ptr(false),
	// 		IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/TestVMRole_IN_0/networkInterfaces/nic1/ipConfigurations/ip1"),
	// 				Name: to.Ptr("ip1"),
	// 				Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
	// 					LoadBalancerBackendAddressPools: []*armnetwork.BackendAddressPool{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/addressPool1"),
	// 					}},
	// 					LoadBalancerInboundNatRules: []*armnetwork.InboundNatRule{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/natPool1.1"),
	// 					}},
	// 					Primary: to.Ptr(true),
	// 					PrivateIPAddress: to.Ptr("10.0.0.5"),
	// 					PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.PublicIPAddress{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/TestVMRole_IN_0/networkInterfaces/nic1/ipConfigurations/ip1/publicIPAddresses/pub1"),
	// 					},
	// 					Subnet: &armnetwork.Subnet{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
	// 					},
	// 				},
	// 		}},
	// 		MacAddress: to.Ptr("00-00-00-00-00-00"),
	// 		NetworkSecurityGroup: &armnetwork.SecurityGroup{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
	// 		},
	// 		Primary: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
