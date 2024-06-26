package armplaywrighttesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwrighttesting/armplaywrighttesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3eb9ec8e9c8f717c6b461c4c0f49a4662fb948fd/specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Quotas_ListBySubscription.json
func ExampleQuotasClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywrighttesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQuotasClient().NewListBySubscriptionPager("eastus", nil)
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
		// page.QuotaListResult = armplaywrighttesting.QuotaListResult{
		// 	Value: []*armplaywrighttesting.Quota{
		// 		{
		// 			Name: to.Ptr("ScalableExecution"),
		// 			Type: to.Ptr("Microsoft.AzurePlaywrightService/Locations/Quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.AzurePlaywrightService/locations/eastus/quotas/ScalableExecution"),
		// 			Properties: &armplaywrighttesting.QuotaProperties{
		// 				FreeTrial: &armplaywrighttesting.FreeTrialProperties{
		// 					AccountID: to.Ptr("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		// 					AllocatedValue: to.Ptr[int32](0),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
		// 					ExpiryAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
		// 					PercentageUsed: to.Ptr[float64](100),
		// 					State: to.Ptr(armplaywrighttesting.FreeTrialStateActive),
		// 					UsedValue: to.Ptr[int32](0),
		// 				},
		// 				ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
