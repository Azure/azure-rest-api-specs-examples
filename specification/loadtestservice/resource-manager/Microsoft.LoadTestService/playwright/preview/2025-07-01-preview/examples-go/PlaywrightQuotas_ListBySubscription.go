package armplaywright_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwright/armplaywright"
)

// Generated from example definition: 2025-07-01-preview/PlaywrightQuotas_ListBySubscription.json
func ExampleQuotasClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywright.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armplaywright.QuotasClientListBySubscriptionResponse{
		// 	QuotaListResult: armplaywright.QuotaListResult{
		// 		Value: []*armplaywright.Quota{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/eastus/PlaywrightQuotas/ExecutionMinutes"),
		// 				Name: to.Ptr(armplaywright.QuotaNameExecutionMinutes),
		// 				Type: to.Ptr("Microsoft.LoadTestService/Locations/PlaywrightQuotas"),
		// 				Properties: &armplaywright.QuotaProperties{
		// 					FreeTrial: &armplaywright.FreeTrialProperties{
		// 						WorkspaceID: to.Ptr("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
		// 						State: to.Ptr(armplaywright.FreeTrialStateActive),
		// 					},
		// 					ProvisioningState: to.Ptr(armplaywright.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
