package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Executions_ListByWorkflowVersion_MaximumSet_Gen.json
func ExampleExecutionsClient_NewListByWorkflowVersionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("EE6D9590-0D52-4B1C-935C-FE49DBF838EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExecutionsClient().NewListByWorkflowVersionPager("rgconfigurationmanager", "abcde", "abcde", "abcde", nil)
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
		// page = armworkloadorchestration.ExecutionsClientListByWorkflowVersionResponse{
		// 	ExecutionListResult: armworkloadorchestration.ExecutionListResult{
		// 		Value: []*armworkloadorchestration.Execution{
		// 			{
		// 				Properties: &armworkloadorchestration.ExecutionProperties{
		// 					WorkflowVersionID: to.Ptr("souenlqwltljsojdcbpc"),
		// 					Specification: map[string]any{
		// 					},
		// 					Status: &armworkloadorchestration.ExecutionStatus{
		// 						Status: to.Ptr[int32](999),
		// 						StageHistory: []*armworkloadorchestration.StageStatus{
		// 							{
		// 								Status: to.Ptr[int32](999),
		// 								StatusMessage: to.Ptr("fkmtwmdfdhbfkbjlkjsvyreoqtxpr"),
		// 								Stage: to.Ptr("hk"),
		// 								Nextstage: to.Ptr("lvatqeczrfvbl"),
		// 								ErrorMessage: to.Ptr("nyxzr"),
		// 								IsActive: to.Ptr(armworkloadorchestration.ActiveStateActive),
		// 								Inputs: map[string]any{
		// 								},
		// 								Outputs: map[string]any{
		// 								},
		// 							},
		// 						},
		// 						UpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-17T13:55:20.919Z"); return t}()),
		// 						StatusMessage: to.Ptr("dfmqxqonlrmrrmqkneij"),
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
		// 				},
		// 				ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
		// 					Name: to.Ptr("ugf"),
		// 					Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
		// 				},
		// 				ETag: to.Ptr("rybzaakessjpkwsprpclpdp"),
		// 				ID: to.Ptr("/subscriptions/329B3ADC-9319-4E19-9CF0-58102804FD70/resourceGroups/myresourceGroupName/providers/Microsoft.Edge/workflows/wfName/versions/1.0.0/executions/executionName"),
		// 				Name: to.Ptr("pljfijtnrx"),
		// 				Type: to.Ptr("ngvrxwzxtpqawwvhf"),
		// 				SystemData: &armworkloadorchestration.SystemData{
		// 					CreatedBy: to.Ptr("favedmahrbemfqzeuggazxzrvwugxw"),
		// 					CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-24T11:04:49.597Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("lywqfnyqrutroctdfbxzytel"),
		// 					LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-24T11:04:49.597Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
