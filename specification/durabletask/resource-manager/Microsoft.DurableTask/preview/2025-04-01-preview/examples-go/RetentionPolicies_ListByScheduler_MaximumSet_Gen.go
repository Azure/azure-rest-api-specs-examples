package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2025-04-01-preview/RetentionPolicies_ListByScheduler_MaximumSet_Gen.json
func ExampleRetentionPoliciesClient_NewListBySchedulerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("194D3C1E-462F-4738-9025-092A628C06EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRetentionPoliciesClient().NewListBySchedulerPager("rgdurabletask", "myscheduler", nil)
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
		// page = armdurabletask.RetentionPoliciesClientListBySchedulerResponse{
		// 	RetentionPolicyListResult: armdurabletask.RetentionPolicyListResult{
		// 		Value: []*armdurabletask.RetentionPolicy{
		// 			{
		// 				Properties: &armdurabletask.RetentionPolicyProperties{
		// 					ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
		// 					RetentionPolicies: []*armdurabletask.RetentionPolicyDetails{
		// 						{
		// 							RetentionPeriodInDays: to.Ptr[int32](30),
		// 						},
		// 						{
		// 							RetentionPeriodInDays: to.Ptr[int32](10),
		// 							OrchestrationState: to.Ptr(armdurabletask.PurgeableOrchestrationStateFailed),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/194D3C1E-462F-4738-9025-092A628C06EB/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/retentionPolicies/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.DurableTask/schedulers/retentionPolicies"),
		// 				SystemData: &armdurabletask.SystemData{
		// 					CreatedBy: to.Ptr("zshkmc"),
		// 					CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-31T23:34:09.612Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ivqrae"),
		// 					LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-31T23:34:09.612Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/ai"),
		// 	},
		// }
	}
}
