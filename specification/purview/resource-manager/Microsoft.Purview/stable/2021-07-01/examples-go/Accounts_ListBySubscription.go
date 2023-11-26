package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListBySubscription.json
func ExampleAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListBySubscriptionPager(&armpurview.AccountsClientListBySubscriptionOptions{SkipToken: nil})
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
		// page.AccountList = armpurview.AccountList{
		// 	Value: []*armpurview.Account{
		// 		{
		// 			Name: to.Ptr("account1"),
		// 			Type: to.Ptr("Microsoft.Purview/accounts"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
		// 			Location: to.Ptr("West US 2"),
		// 			SystemData: &armpurview.TrackedResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.692Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.343Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("client-name"),
		// 				LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armpurview.AccountProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.692Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByObjectID: to.Ptr("client-objectId"),
		// 				Endpoints: &armpurview.AccountPropertiesEndpoints{
		// 					Catalog: to.Ptr("https://account1.catalog.purview.azure-test.com"),
		// 					Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
		// 					Scan: to.Ptr("https://account1.scan.purview.azure-test.com"),
		// 				},
		// 				FriendlyName: to.Ptr("friendly-account1"),
		// 				ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
		// 			},
		// 			SKU: &armpurview.AccountSKU{
		// 				Name: to.Ptr(armpurview.NameStandard),
		// 				Capacity: to.Ptr[int32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("account2"),
		// 			Type: to.Ptr("Microsoft.Purview/accounts"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account2"),
		// 			Location: to.Ptr("West US 2"),
		// 			SystemData: &armpurview.TrackedResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.692Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.343Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("client-name"),
		// 				LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armpurview.AccountProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.692Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByObjectID: to.Ptr("client-objectId"),
		// 				Endpoints: &armpurview.AccountPropertiesEndpoints{
		// 					Catalog: to.Ptr("https://account2.catalog.purview.azure-test.com"),
		// 					Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
		// 					Scan: to.Ptr("https://account2.scan.purview.azure-test.com"),
		// 				},
		// 				FriendlyName: to.Ptr("friendly-account1"),
		// 				ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
		// 			},
		// 			SKU: &armpurview.AccountSKU{
		// 				Name: to.Ptr(armpurview.NameStandard),
		// 				Capacity: to.Ptr[int32](1),
		// 			},
		// 	}},
		// }
	}
}
