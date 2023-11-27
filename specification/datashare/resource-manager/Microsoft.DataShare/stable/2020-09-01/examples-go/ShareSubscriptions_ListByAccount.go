package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ShareSubscriptions_ListByAccount.json
func ExampleShareSubscriptionsClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewShareSubscriptionsClient().NewListByAccountPager("SampleResourceGroup", "Account1", &armdatashare.ShareSubscriptionsClientListByAccountOptions{SkipToken: nil,
		Filter:  nil,
		Orderby: nil,
	})
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
		// page.ShareSubscriptionList = armdatashare.ShareSubscriptionList{
		// 	Value: []*armdatashare.ShareSubscription{
		// 		{
		// 			Name: to.Ptr("ShareSubscription1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/sharesubscriptions/ShareSubscription1"),
		// 			Properties: &armdatashare.ShareSubscriptionProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T06:15:15.681Z"); return t}()),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-26T22:33:24.578Z"); return t}()),
		// 				InvitationID: to.Ptr("12345678-1234-1234-12345678abd"),
		// 				ProviderEmail: to.Ptr("jack.rose@microsoft.com"),
		// 				ProviderName: to.Ptr("Jack Rose"),
		// 				ProviderTenantName: to.Ptr("Microsoft"),
		// 				ShareDescription: to.Ptr("Some share"),
		// 				ShareKind: to.Ptr(armdatashare.ShareKindCopyBased),
		// 				ShareName: to.Ptr("share1"),
		// 				ShareSubscriptionStatus: to.Ptr(armdatashare.ShareSubscriptionStatusActive),
		// 				ShareTerms: to.Ptr("Confidential"),
		// 				SourceShareLocation: to.Ptr("eastus2"),
		// 				UserEmail: to.Ptr("john.smith@microsoft.com"),
		// 				UserName: to.Ptr("johnsmith@microsoft.com"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ShareSubscription1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shareSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/sharesubscriptions/ShareSubscription1"),
		// 			Properties: &armdatashare.ShareSubscriptionProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-17T22:32:36.818Z"); return t}()),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-26T22:33:24.578Z"); return t}()),
		// 				InvitationID: to.Ptr("4256e2cf-0f82-4865-961b-12f83333f487"),
		// 				ProviderEmail: to.Ptr("jack.rose@microsoft.com"),
		// 				ProviderName: to.Ptr("Jack Rose"),
		// 				ProviderTenantName: to.Ptr("Microsoft"),
		// 				ShareDescription: to.Ptr("Some share"),
		// 				ShareKind: to.Ptr(armdatashare.ShareKindCopyBased),
		// 				ShareName: to.Ptr("share1"),
		// 				ShareSubscriptionStatus: to.Ptr(armdatashare.ShareSubscriptionStatusActive),
		// 				ShareTerms: to.Ptr("Confidential"),
		// 				SourceShareLocation: to.Ptr("eastus2"),
		// 				UserEmail: to.Ptr("john.smith@microsoft.com"),
		// 				UserName: to.Ptr("John Smith"),
		// 			},
		// 	}},
		// }
	}
}
