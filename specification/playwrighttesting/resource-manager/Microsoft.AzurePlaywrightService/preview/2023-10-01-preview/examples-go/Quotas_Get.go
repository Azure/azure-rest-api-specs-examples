package armplaywrighttesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwrighttesting/armplaywrighttesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3eb9ec8e9c8f717c6b461c4c0f49a4662fb948fd/specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Quotas_Get.json
func ExampleQuotasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywrighttesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotasClient().Get(ctx, "eastus", armplaywrighttesting.QuotaNamesScalableExecution, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Quota = armplaywrighttesting.Quota{
	// 	Name: to.Ptr("ScalableExecution"),
	// 	Type: to.Ptr("Microsoft.AzurePlaywrightService/Locations/Quotas"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.AzurePlaywrightService/locations/eastus/quotas/ScalableExecution"),
	// 	Properties: &armplaywrighttesting.QuotaProperties{
	// 		FreeTrial: &armplaywrighttesting.FreeTrialProperties{
	// 			AccountID: to.Ptr("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
	// 			AllocatedValue: to.Ptr[int32](0),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
	// 			ExpiryAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-31T10:19:36.081Z"); return t}()),
	// 			PercentageUsed: to.Ptr[float64](100),
	// 			State: to.Ptr(armplaywrighttesting.FreeTrialStateActive),
	// 			UsedValue: to.Ptr[int32](0),
	// 		},
	// 		ProvisioningState: to.Ptr(armplaywrighttesting.ProvisioningStateSucceeded),
	// 	},
	// }
}
