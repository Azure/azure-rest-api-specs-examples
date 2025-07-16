package armplaywright_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwright/armplaywright"
)

// Generated from example definition: 2025-07-01-preview/PlaywrightWorkspaceQuotas_Get.json
func ExampleWorkspaceQuotasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplaywright.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceQuotasClient().Get(ctx, "dummyrg", "myWorkspace", armplaywright.QuotaNameExecutionMinutes, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armplaywright.WorkspaceQuotasClientGetResponse{
	// 	WorkspaceQuota: &armplaywright.WorkspaceQuota{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.LoadTestService/PlaywrightWorkspaces/myWorkspace/quotas/ExecutionMinutes"),
	// 		Name: to.Ptr(armplaywright.QuotaNameExecutionMinutes),
	// 		Type: to.Ptr("Microsoft.LoadTestService/PlaywrightWorkspaces/Quotas"),
	// 		Properties: &armplaywright.WorkspaceQuotaProperties{
	// 			FreeTrial: &armplaywright.WorkspaceFreeTrialProperties{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-31T10:19:36.081Z"); return t}()),
	// 				ExpiryAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-31T10:19:36.081Z"); return t}()),
	// 				AllocatedValue: to.Ptr[int32](10),
	// 				UsedValue: to.Ptr[float32](10),
	// 				PercentageUsed: to.Ptr[float32](100),
	// 			},
	// 			ProvisioningState: to.Ptr(armplaywright.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
