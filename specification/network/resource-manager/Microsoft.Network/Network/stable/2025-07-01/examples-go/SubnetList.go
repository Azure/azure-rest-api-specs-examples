package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/SubnetList.json
func ExampleSubnetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubnetsClient().NewListPager("subnet-test", "vnetname", nil)
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
		// page = armnetwork.SubnetsClientListResponse{
		// 	SubnetListResult: armnetwork.SubnetListResult{
		// 		Value: []*armnetwork.Subnet{
		// 			{
		// 				Name: to.Ptr("subnet1"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/subnet-test/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/subnet1"),
		// 				Properties: &armnetwork.SubnetPropertiesFormat{
		// 					AddressPrefix: to.Ptr("10.0.0.0/16"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("subnet2"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/subnet-test/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/subnet2"),
		// 				Properties: &armnetwork.SubnetPropertiesFormat{
		// 					AddressPrefix: to.Ptr("10.0.0.0/16"),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
