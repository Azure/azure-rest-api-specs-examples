package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkInterfaceUpdateTags.json
func ExampleInterfacesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInterfacesClient().UpdateTags(ctx, "rg1", "test-nic", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
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
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
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
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VnetEncryptionSupported: to.Ptr(false),
	// 	},
	// }
}
