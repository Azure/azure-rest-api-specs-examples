package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/ListAccountsBySubscription.json
func ExampleAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListBySubscriptionPager(nil)
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
		// page = armmaps.AccountsClientListBySubscriptionResponse{
		// 	Accounts: armmaps.Accounts{
		// 		Value: []*armmaps.Account{
		// 			{
		// 				Name: to.Ptr("myMapsAccount2"),
		// 				Type: to.Ptr("Microsoft.Maps/accounts"),
		// 				ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount2"),
		// 				Kind: to.Ptr(armmaps.KindGen2),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armmaps.AccountProperties{
		// 					DisableLocalAuth: to.Ptr(false),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					UniqueID: to.Ptr("b2e763e6-d6f3-4858-9e2b-7cf8df85c593"),
		// 				},
		// 				SKU: &armmaps.SKU{
		// 					Name: to.Ptr(armmaps.NameG2),
		// 					Tier: to.Ptr("Standard"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"test": to.Ptr("true"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myMapsAccount"),
		// 				Type: to.Ptr("Microsoft.Maps/accounts"),
		// 				ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount"),
		// 				Kind: to.Ptr(armmaps.KindGen2),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armmaps.AccountProperties{
		// 					DisableLocalAuth: to.Ptr(true),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					UniqueID: to.Ptr("b2e763e6-d6f3-4858-9e2b-7cf8df85c592"),
		// 				},
		// 				SKU: &armmaps.SKU{
		// 					Name: to.Ptr(armmaps.NameG2),
		// 					Tier: to.Ptr("Standard"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"test": to.Ptr("true"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
