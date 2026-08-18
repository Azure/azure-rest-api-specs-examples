package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/AddressPrefixSetList.json
func ExampleAddressPrefixSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAddressPrefixSetsClient().NewListPager("rg1", "test-asg", nil)
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
		// page = armnetwork.AddressPrefixSetsClientListResponse{
		// 	AddressPrefixSetListResult: armnetwork.AddressPrefixSetListResult{
		// 		Value: []*armnetwork.AddressPrefixSet{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationSecurityGroups/test-asg/addressPrefixSets/test-prefix-set1"),
		// 				Name: to.Ptr("test-prefix-set1"),
		// 				Type: to.Ptr("Microsoft.Network/applicationSecurityGroups/addressPrefixSets"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000000\""),
		// 				Properties: &armnetwork.AddressPrefixSetPropertiesFormat{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("10.0.0.0/16"),
		// 						to.Ptr("192.168.1.0/24"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/applicationSecurityGroups/test-asg/addressPrefixSets/test-prefix-set2"),
		// 				Name: to.Ptr("test-prefix-set2"),
		// 				Type: to.Ptr("Microsoft.Network/applicationSecurityGroups/addressPrefixSets"),
		// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-000000000001\""),
		// 				Properties: &armnetwork.AddressPrefixSetPropertiesFormat{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("172.16.0.0/12"),
		// 						to.Ptr("2001:db8::/32"),
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
