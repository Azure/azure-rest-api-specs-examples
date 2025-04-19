package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2025-04-01-preview/RetentionPolicies_Update_MaximumSet_Gen.json
func ExampleRetentionPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("194D3C1E-462F-4738-9025-092A628C06EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRetentionPoliciesClient().BeginUpdate(ctx, "rgdurabletask", "testscheduler", armdurabletask.RetentionPolicy{
		Properties: &armdurabletask.RetentionPolicyProperties{
			RetentionPolicies: []*armdurabletask.RetentionPolicyDetails{
				{
					RetentionPeriodInDays: to.Ptr[int32](30),
				},
				{
					RetentionPeriodInDays: to.Ptr[int32](10),
					OrchestrationState:    to.Ptr(armdurabletask.PurgeableOrchestrationStateFailed),
				},
				{
					RetentionPeriodInDays: to.Ptr[int32](24),
					OrchestrationState:    to.Ptr(armdurabletask.PurgeableOrchestrationStateCompleted),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdurabletask.RetentionPoliciesClientUpdateResponse{
	// 	RetentionPolicy: &armdurabletask.RetentionPolicy{
	// 		Properties: &armdurabletask.RetentionPolicyProperties{
	// 			RetentionPolicies: []*armdurabletask.RetentionPolicyDetails{
	// 				{
	// 					RetentionPeriodInDays: to.Ptr[int32](30),
	// 				},
	// 				{
	// 					RetentionPeriodInDays: to.Ptr[int32](10),
	// 					OrchestrationState: to.Ptr(armdurabletask.PurgeableOrchestrationStateFailed),
	// 				},
	// 				{
	// 					RetentionPeriodInDays: to.Ptr[int32](24),
	// 					OrchestrationState: to.Ptr(armdurabletask.PurgeableOrchestrationStateCompleted),
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/194D3C1E-462F-4738-9025-092A628C06EB/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/retentionPolicies/default"),
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.DurableTask/schedulers/retentionPolicies"),
	// 		SystemData: &armdurabletask.SystemData{
	// 			CreatedBy: to.Ptr("tenmbevaunjzikxowqexrsx"),
	// 			CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.365Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xfvdcegtj"),
	// 			LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.366Z"); return t}()),
	// 		},
	// 	},
	// }
}
