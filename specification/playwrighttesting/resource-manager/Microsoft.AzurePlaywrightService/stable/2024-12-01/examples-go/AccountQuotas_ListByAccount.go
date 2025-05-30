package armplaywrighttesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwrighttesting/armplaywrighttesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2bde125befabb21807a2021765901f20e3e74ec8/specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/stable/2024-12-01/examples/AccountQuotas_ListByAccount.json
func ExampleAccountQuotasClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywrighttesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountQuotasClient().NewListByAccountPager("dummyrg", "myPlaywrightAccount", nil)
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
		// page.AccountQuotaListResult = armplaywrighttesting.AccountQuotaListResult{
		// 	Value: []*armplaywrighttesting.AccountQuota{
		// 		{
		// 			Name: to.Ptr("ScalableExecution"),
		// 			Type: to.Ptr("Microsoft.AzurePlaywrightService/Accounts/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.AzurePlaywrightService/accounts/myPlaywrightAccount/quotas/ScalableExecution"),
		// 			Properties: &armplaywrighttesting.AccountQuotaProperties{
		// 				FreeTrial: &armplaywrighttesting.AccountFreeTrialProperties{
		// 					AllocatedValue: to.Ptr[int32](0),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
		// 					ExpiryAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
		// 					PercentageUsed: to.Ptr[float32](100),
		// 					UsedValue: to.Ptr[int32](0),
		// 				},
		// 				ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Reporting"),
		// 			Type: to.Ptr("Microsoft.AzurePlaywrightService/Accounts/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.AzurePlaywrightService/accounts/myPlaywrightAccount/quotas/Reporting"),
		// 			Properties: &armplaywrighttesting.AccountQuotaProperties{
		// 				FreeTrial: &armplaywrighttesting.AccountFreeTrialProperties{
		// 					AllocatedValue: to.Ptr[int32](0),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
		// 					ExpiryAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
		// 					PercentageUsed: to.Ptr[float32](100),
		// 					UsedValue: to.Ptr[int32](0),
		// 				},
		// 				ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
