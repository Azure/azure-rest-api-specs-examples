package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ProviderShareSubscriptions_ListByShare.json
func ExampleProviderShareSubscriptionsClient_NewListBySharePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProviderShareSubscriptionsClient().NewListBySharePager("SampleResourceGroup", "Account1", "Share1", &armdatashare.ProviderShareSubscriptionsClientListByShareOptions{SkipToken: nil})
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
		// page.ProviderShareSubscriptionList = armdatashare.ProviderShareSubscriptionList{
		// 	Value: []*armdatashare.ProviderShareSubscription{
		// 		{
		// 			Name: to.Ptr("4256e2cf-0f82-4865-961b-12f83333f487"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shares/providerShareSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-12345678abc/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shares/Share1/providerShareSubscripitons/4256e2cf-0f82-4865-961b-12f83333f487"),
		// 			Properties: &armdatashare.ProviderShareSubscriptionProperties{
		// 				ConsumerEmail: to.Ptr("john.smith@microsoft.com"),
		// 				ConsumerName: to.Ptr("John Smith"),
		// 				ConsumerTenantName: to.Ptr("Microsoft"),
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-13T08:52:42.622Z"); return t}()),
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-08-26T22:33:24.578Z"); return t}()),
		// 				ProviderEmail: to.Ptr("john.smith@microsoft.com"),
		// 				ProviderName: to.Ptr("John Smith"),
		// 				ShareSubscriptionStatus: to.Ptr(armdatashare.ShareSubscriptionStatusActive),
		// 				SharedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-13T08:45:40.900Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
