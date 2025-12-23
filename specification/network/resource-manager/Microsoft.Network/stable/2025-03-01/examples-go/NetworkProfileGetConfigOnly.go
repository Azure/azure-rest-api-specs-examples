package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkProfileGetConfigOnly.json
func ExampleProfilesClient_Get_getNetworkProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().Get(ctx, "rg1", "networkProfile1", &armnetwork.ProfilesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Profile = armnetwork.Profile{
	// 	Name: to.Ptr("networkProfile1"),
	// 	Type: to.Ptr("Microsoft.Network/networkProfiles"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 	Properties: &armnetwork.ProfilePropertiesFormat{
	// 		ContainerNetworkInterfaceConfigurations: []*armnetwork.ContainerNetworkInterfaceConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0"),
	// 				Name: to.Ptr("eth0"),
	// 				Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations"),
	// 				Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 				Properties: &armnetwork.ContainerNetworkInterfaceConfigurationPropertiesFormat{
	// 					IPConfigurations: []*armnetwork.IPConfigurationProfile{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0/ipConfigurations/ipconfigprofile1"),
	// 							Name: to.Ptr("ipconfigprofile1"),
	// 							Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations/ipConfigurations"),
	// 							Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 							Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
	// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 								Subnet: &armnetwork.Subnet{
	// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
	// 								},
	// 							},
	// 						},
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0/ipConfigurations/ipconfigprofile2"),
	// 							Name: to.Ptr("ipconfigprofile2"),
	// 							Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations/ipConfigurations"),
	// 							Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 							Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
	// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 								Subnet: &armnetwork.Subnet{
	// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
	// 								},
	// 							},
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 			},
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth1"),
	// 				Name: to.Ptr("eth1"),
	// 				Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations"),
	// 				Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 				Properties: &armnetwork.ContainerNetworkInterfaceConfigurationPropertiesFormat{
	// 					IPConfigurations: []*armnetwork.IPConfigurationProfile{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth1/ipConfigurations/ipconfigprofile3"),
	// 							Name: to.Ptr("ipconfigprofile3"),
	// 							Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations/ipConfigurations"),
	// 							Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 							Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
	// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 								Subnet: &armnetwork.Subnet{
	// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
	// 								},
	// 							},
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		ContainerNetworkInterfaces: []*armnetwork.ContainerNetworkInterface{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("1570d8b6-ab8a-4ad2-81d6-d2799b429cbf"),
	// 	},
	// }
}
