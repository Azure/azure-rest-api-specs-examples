package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/WorkflowVersions_Update_MaximumSet_Gen.json
func ExampleWorkflowVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkflowVersionsClient().BeginUpdate(ctx, "rgconfigurationmanager", "testname", "testname", "testname", armworkloadorchestration.WorkflowVersion{
		Properties: &armworkloadorchestration.WorkflowVersionProperties{
			StageSpec: []*armworkloadorchestration.StageSpec{
				{
					Name:          to.Ptr("amrbjd"),
					Specification: map[string]any{},
					Tasks: []*armworkloadorchestration.TaskSpec{
						{
							Name:          to.Ptr("xxmeyvmgydbcwxqwjhadjxjod"),
							TargetID:      to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
							Specification: map[string]any{},
						},
					},
					TaskOption: &armworkloadorchestration.TaskOption{
						Concurrency: to.Ptr[int32](3),
						ErrorAction: &armworkloadorchestration.ErrorAction{
							Mode:                 to.Ptr(armworkloadorchestration.ErrorActionModeStopOnAnyFailure),
							MaxToleratedFailures: to.Ptr[int32](0),
						},
					},
				},
			},
			Specification: map[string]any{},
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
	// res = armworkloadorchestration.WorkflowVersionsClientUpdateResponse{
	// 	WorkflowVersion: &armworkloadorchestration.WorkflowVersion{
	// 		Properties: &armworkloadorchestration.WorkflowVersionProperties{
	// 			StageSpec: []*armworkloadorchestration.StageSpec{
	// 				{
	// 					Name: to.Ptr("amrbjd"),
	// 					Specification: map[string]any{
	// 					},
	// 					Tasks: []*armworkloadorchestration.TaskSpec{
	// 						{
	// 							Name: to.Ptr("xxmeyvmgydbcwxqwjhadjxjod"),
	// 							TargetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 							Specification: map[string]any{
	// 							},
	// 						},
	// 					},
	// 					TaskOption: &armworkloadorchestration.TaskOption{
	// 						Concurrency: to.Ptr[int32](3),
	// 						ErrorAction: &armworkloadorchestration.ErrorAction{
	// 							Mode: to.Ptr(armworkloadorchestration.ErrorActionModeStopOnAnyFailure),
	// 							MaxToleratedFailures: to.Ptr[int32](0),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			State: to.Ptr(armworkloadorchestration.StateInReview),
	// 			Specification: map[string]any{
	// 			},
	// 			Revision: to.Ptr[int32](4),
	// 			Configuration: to.Ptr("czvrwkzex"),
	// 			ReviewID: to.Ptr("hcnzmdpszvvtn"),
	// 			ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
	// 			Name: to.Ptr("szjrwimeqyiue"),
	// 			Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
	// 		},
	// 		ETag: to.Ptr("v"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("rsghgriwbtckmabbcxmqbvshybhw"),
	// 		Type: to.Ptr("jahxsjfhxhpsobltqfdlfn"),
	// 		SystemData: &armworkloadorchestration.SystemData{
	// 			CreatedBy: to.Ptr("nvjczgdguyvllp"),
	// 			CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
	// 			LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
	// 		},
	// 	},
	// }
}
