package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkProfileList.json
func ExampleProfilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProfilesClient().NewListPager("rg1", nil)
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
		// page.ProfileListResult = armnetwork.ProfileListResult{
		// 	Value: []*armnetwork.Profile{
		// 		{
		// 			Name: to.Ptr("networkProfile1"),
		// 			Type: to.Ptr("Microsoft.Network/networkProfiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.ProfilePropertiesFormat{
		// 				ContainerNetworkInterfaceConfigurations: []*armnetwork.ContainerNetworkInterfaceConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0"),
		// 						Name: to.Ptr("eth0"),
		// 						Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations"),
		// 						Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
		// 						Properties: &armnetwork.ContainerNetworkInterfaceConfigurationPropertiesFormat{
		// 							IPConfigurations: []*armnetwork.IPConfigurationProfile{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0/ipConfigurations/ipconfigprofile1"),
		// 									Name: to.Ptr("ipconfigprofile1"),
		// 									Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations/ipConfigurations"),
		// 									Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
		// 									Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										Subnet: &armnetwork.Subnet{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0/ipConfigurations/ipconfigprofile2"),
		// 									Name: to.Ptr("ipconfigprofile2"),
		// 									Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations/ipConfigurations"),
		// 									Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
		// 									Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										Subnet: &armnetwork.Subnet{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
		// 										},
		// 									},
		// 							}},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				ContainerNetworkInterfaces: []*armnetwork.ContainerNetworkInterface{
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("1570d8b6-ab8a-4ad2-81d6-d2799b429cbf"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("networkProfile2"),
		// 			Type: to.Ptr("Microsoft.Network/networkProfiles"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armnetwork.ProfilePropertiesFormat{
		// 				ContainerNetworkInterfaceConfigurations: []*armnetwork.ContainerNetworkInterfaceConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth1"),
		// 						Name: to.Ptr("eth1"),
		// 						Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations"),
		// 						Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
		// 						Properties: &armnetwork.ContainerNetworkInterfaceConfigurationPropertiesFormat{
		// 							IPConfigurations: []*armnetwork.IPConfigurationProfile{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth1/ipConfigurations/ipconfigprofile3"),
		// 									Name: to.Ptr("ipconfigprofile3"),
		// 									Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations/ipConfigurations"),
		// 									Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
		// 									Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										Subnet: &armnetwork.Subnet{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
		// 										},
		// 									},
		// 							}},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				ContainerNetworkInterfaces: []*armnetwork.ContainerNetworkInterface{
		// 				},
		// 			},
		// 	}},
		// }
	}
}
