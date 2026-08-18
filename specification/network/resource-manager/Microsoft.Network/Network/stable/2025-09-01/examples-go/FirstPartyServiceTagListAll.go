package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/FirstPartyServiceTagListAll.json
func ExampleFirstPartyServiceTagsClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFirstPartyServiceTagsClient().NewListAllPager(nil)
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
		// page = armnetwork.FirstPartyServiceTagsClientListAllResponse{
		// 	FirstPartyServiceTagListResult: armnetwork.FirstPartyServiceTagListResult{
		// 		Value: []*armnetwork.FirstPartyServiceTag{
		// 			{
		// 				Name: to.Ptr("myServiceTag"),
		// 				Type: to.Ptr("Microsoft.Network/firstPartyServiceTags"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firstPartyServiceTags/myServiceTag"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 				Properties: &armnetwork.FirstPartyServiceTagPropertiesFormat{
		// 					Value: to.Ptr("myServiceTagValue"),
		// 					FailedReason: to.Ptr(""),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myServiceTag2"),
		// 				Type: to.Ptr("Microsoft.Network/firstPartyServiceTags"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/firstPartyServiceTags/myServiceTag2"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 				Properties: &armnetwork.FirstPartyServiceTagPropertiesFormat{
		// 					Value: to.Ptr("myServiceTagValue2"),
		// 					FailedReason: to.Ptr(""),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
