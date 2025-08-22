package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/WorkflowVersions_ListByWorkflow_MaximumSet_Gen.json
func ExampleWorkflowVersionsClient_NewListByWorkflowPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("9D54FE4C-00AF-4836-8F48-B6A9C4E47192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowVersionsClient().NewListByWorkflowPager("rgconfigurationmanager", "testname", "testname", nil)
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
		// page = armworkloadorchestration.WorkflowVersionsClientListByWorkflowResponse{
		// 	WorkflowVersionListResult: armworkloadorchestration.WorkflowVersionListResult{
		// 		Value: []*armworkloadorchestration.WorkflowVersion{
		// 			{
		// 				Properties: &armworkloadorchestration.WorkflowVersionProperties{
		// 					Revision: to.Ptr[int32](4),
		// 					Configuration: to.Ptr("czvrwkzex"),
		// 					StageSpec: []*armworkloadorchestration.StageSpec{
		// 						{
		// 							Name: to.Ptr("amrbjd"),
		// 							Specification: map[string]any{
		// 							},
		// 							Tasks: []*armworkloadorchestration.TaskSpec{
		// 								{
		// 									Name: to.Ptr("xxmeyvmgydbcwxqwjhadjxjod"),
		// 									TargetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 									Specification: map[string]any{
		// 									},
		// 								},
		// 							},
		// 							TaskOption: &armworkloadorchestration.TaskOption{
		// 								Concurrency: to.Ptr[int32](3),
		// 								ErrorAction: &armworkloadorchestration.ErrorAction{
		// 									Mode: to.Ptr(armworkloadorchestration.ErrorActionModeStopOnAnyFailure),
		// 									MaxToleratedFailures: to.Ptr[int32](0),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ReviewID: to.Ptr("hcnzmdpszvvtn"),
		// 					State: to.Ptr(armworkloadorchestration.StateInReview),
		// 					Specification: map[string]any{
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
		// 				},
		// 				ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
		// 					Name: to.Ptr("szjrwimeqyiue"),
		// 					Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
		// 				},
		// 				ETag: to.Ptr("v"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 				Name: to.Ptr("rsghgriwbtckmabbcxmqbvshybhw"),
		// 				Type: to.Ptr("jahxsjfhxhpsobltqfdlfn"),
		// 				SystemData: &armworkloadorchestration.SystemData{
		// 					CreatedBy: to.Ptr("nvjczgdguyvllp"),
		// 					CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("uzbznzjgvaspvtqhyg"),
		// 					LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-09T10:11:50.747Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
