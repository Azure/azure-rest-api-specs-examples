package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NetworkInterfaceIPConfigurationList.json
func ExampleInterfaceIPConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInterfaceIPConfigurationsClient().NewListPager("testrg", "nic1", nil)
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
		// page = armnetwork.InterfaceIPConfigurationsClientListResponse{
		// 	InterfaceIPConfigurationListResult: armnetwork.InterfaceIPConfigurationListResult{
		// 		Value: []*armnetwork.InterfaceIPConfiguration{
		// 			{
		// 				Name: to.Ptr("ipconfig1"),
		// 				Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
		// 				Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
		// 					Primary: to.Ptr(true),
		// 					PrivateIPAddress: to.Ptr("10.0.0.4"),
		// 					PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					Subnet: &armnetwork.Subnet{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet12/subnets/subnet12"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
