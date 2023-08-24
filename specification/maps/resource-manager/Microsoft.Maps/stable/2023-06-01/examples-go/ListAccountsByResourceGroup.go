package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9403296f0b0e112b0d8222ad05fd1d79ee10e03/specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/ListAccountsByResourceGroup.json
func ExampleAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.Accounts = armmaps.Accounts{
		// 	Value: []*armmaps.Account{
		// 		{
		// 			Name: to.Ptr("myMapsAccount2"),
		// 			Type: to.Ptr("Microsoft.Maps/accounts"),
		// 			ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount2"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"test": to.Ptr("true"),
		// 			},
		// 			Kind: to.Ptr(armmaps.KindGen1),
		// 			Properties: &armmaps.AccountProperties{
		// 				DisableLocalAuth: to.Ptr(false),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				UniqueID: to.Ptr("b2e763e6-d6f3-4858-9e2b-7cf8df85c593"),
		// 			},
		// 			SKU: &armmaps.SKU{
		// 				Name: to.Ptr(armmaps.NameS0),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myMapsAccount"),
		// 			Type: to.Ptr("Microsoft.Maps/accounts"),
		// 			ID: to.Ptr("/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Maps/accounts/myMapsAccount"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"test": to.Ptr("true"),
		// 			},
		// 			Kind: to.Ptr(armmaps.KindGen2),
		// 			Properties: &armmaps.AccountProperties{
		// 				DisableLocalAuth: to.Ptr(true),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				UniqueID: to.Ptr("b2e763e6-d6f3-4858-9e2b-7cf8df85c592"),
		// 			},
		// 			SKU: &armmaps.SKU{
		// 				Name: to.Ptr(armmaps.NameG2),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
