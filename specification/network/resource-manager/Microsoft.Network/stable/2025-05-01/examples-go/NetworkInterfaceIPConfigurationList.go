package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkInterfaceIPConfigurationList.json
func ExampleInterfaceIPConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.InterfaceIPConfigurationListResult = armnetwork.InterfaceIPConfigurationListResult{
		// 	Value: []*armnetwork.InterfaceIPConfiguration{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
		// 			Name: to.Ptr("ipconfig1"),
		// 			Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 			Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
		// 				Primary: to.Ptr(true),
		// 				PrivateIPAddress: to.Ptr("10.0.0.4"),
		// 				PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				Subnet: &armnetwork.Subnet{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet12/subnets/subnet12"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
