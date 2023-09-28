package armplaywrighttesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwrighttesting/armplaywrighttesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3eb9ec8e9c8f717c6b461c4c0f49a4662fb948fd/specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Accounts_ListBySubscription.json
func ExampleAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywrighttesting.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.AccountListResult = armplaywrighttesting.AccountListResult{
		// 	Value: []*armplaywrighttesting.Account{
		// 		{
		// 			Name: to.Ptr("myPlaywrightAccount"),
		// 			Type: to.Ptr("Microsoft.AzurePlaywrightService/accounts"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.AzurePlaywrightService/accounts/myPlaywrightAccount"),
		// 			SystemData: &armplaywrighttesting.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33Z"); return t}()),
		// 				CreatedBy: to.Ptr("userId1001"),
		// 				CreatedByType: to.Ptr(armplaywrighttesting.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("userId1001"),
		// 				LastModifiedByType: to.Ptr(armplaywrighttesting.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"Team": to.Ptr("Dev Exp"),
		// 			},
		// 			Properties: &armplaywrighttesting.AccountProperties{
		// 				DashboardURI: to.Ptr("https://dashboard.00000000-0000-0000-0000-000000000000.domain.com"),
		// 				ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
		// 				RegionalAffinity: to.Ptr(armplaywrighttesting.EnablementStatusEnabled),
		// 				Reporting: to.Ptr(armplaywrighttesting.EnablementStatusEnabled),
		// 				ScalableExecution: to.Ptr(armplaywrighttesting.EnablementStatusEnabled),
		// 			},
		// 	}},
		// }
	}
}
