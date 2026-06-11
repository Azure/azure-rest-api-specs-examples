package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkProfileUpdateTags.json
func ExampleProfilesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().UpdateTags(ctx, "rg1", "test-np", armnetwork.TagsObject{
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
	// res = armnetwork.ProfilesClientUpdateTagsResponse{
	// 	Profile: armnetwork.Profile{
	// 		Name: to.Ptr("test-np"),
	// 		Type: to.Ptr("Microsoft.Network/networkProfiles"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/test-np"),
	// 		Location: to.Ptr("centraluseuap"),
	// 		Properties: &armnetwork.ProfilePropertiesFormat{
	// 			ContainerNetworkInterfaceConfigurations: []*armnetwork.ContainerNetworkInterfaceConfiguration{
	// 				{
	// 					Name: to.Ptr("eth0"),
	// 					Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations"),
	// 					Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0"),
	// 					Properties: &armnetwork.ContainerNetworkInterfaceConfigurationPropertiesFormat{
	// 						IPConfigurations: []*armnetwork.IPConfigurationProfile{
	// 							{
	// 								Name: to.Ptr("ipconfigprofile1"),
	// 								Type: to.Ptr("Microsoft.Network/networkProfiles/containerNetworkInterfaceConfigurations/ipConfigurations"),
	// 								Etag: to.Ptr("W/\"4d705a71-752f-4e0a-8057-c02b125b1c08\""),
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkProfiles/networkProfile1/containerNetworkInterfaceConfigurations/eth0/ipConfigurations/ipconfigprofile1"),
	// 								Properties: &armnetwork.IPConfigurationProfilePropertiesFormat{
	// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 									Subnet: &armnetwork.Subnet{
	// 										ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 			ContainerNetworkInterfaces: []*armnetwork.ContainerNetworkInterface{
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			ResourceGUID: to.Ptr("1570d8b6-ab8a-4ad2-81d6-d2799b429cbf"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
