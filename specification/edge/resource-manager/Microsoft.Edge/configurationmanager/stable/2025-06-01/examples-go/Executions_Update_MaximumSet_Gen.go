package armworkloadorchestration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadorchestration/armworkloadorchestration"
)

// Generated from example definition: 2025-06-01/Executions_Update_MaximumSet_Gen.json
func ExampleExecutionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadorchestration.NewClientFactory("FFA229AF-C1A3-4CB6-9E5D-62C25CFBE4D0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExecutionsClient().BeginUpdate(ctx, "rgconfigurationmanager", "abcde", "abcde", "abcde", "abcde", armworkloadorchestration.Execution{
		Properties: &armworkloadorchestration.ExecutionProperties{
			Specification:     map[string]any{},
			WorkflowVersionID: to.Ptr("xjsxzbfltzvbuvn"),
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
	// res = armworkloadorchestration.ExecutionsClientUpdateResponse{
	// 	Execution: &armworkloadorchestration.Execution{
	// 		Properties: &armworkloadorchestration.ExecutionProperties{
	// 			Specification: map[string]any{
	// 			},
	// 			Status: &armworkloadorchestration.ExecutionStatus{
	// 				Status: to.Ptr[int32](999),
	// 				StageHistory: []*armworkloadorchestration.StageStatus{
	// 					{
	// 						Status: to.Ptr[int32](999),
	// 						StatusMessage: to.Ptr("wfymzartwvvqrgrmdwyhfaftszoc"),
	// 						Stage: to.Ptr("gsostfpgjcsoeky"),
	// 						Nextstage: to.Ptr("wjxvqbrocjxzhzfgmgbzt"),
	// 						ErrorMessage: to.Ptr("xsvwgovyatvlacmp"),
	// 						IsActive: to.Ptr(armworkloadorchestration.ActiveStateActive),
	// 						Inputs: map[string]any{
	// 						},
	// 						Outputs: map[string]any{
	// 						},
	// 					},
	// 				},
	// 				UpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-17T13:55:20.922Z"); return t}()),
	// 				StatusMessage: to.Ptr("ebhukpnhnbu"),
	// 			},
	// 			WorkflowVersionID: to.Ptr("souenlqwltljsojdcbpc"),
	// 			ProvisioningState: to.Ptr(armworkloadorchestration.ProvisioningStateSucceeded),
	// 		},
	// 		ExtendedLocation: &armworkloadorchestration.ExtendedLocation{
	// 			Name: to.Ptr("ugf"),
	// 			Type: to.Ptr(armworkloadorchestration.ExtendedLocationTypeEdgeZone),
	// 		},
	// 		ETag: to.Ptr("rybzaakessjpkwsprpclpdp"),
	// 		ID: to.Ptr("/subscriptions/329B3ADC-9319-4E19-9CF0-58102804FD70/resourceGroups/myresourceGroupName/providers/Microsoft.Edge/workflows/wfName/versions/1.0.0/executions/executionName"),
	// 		Name: to.Ptr("pljfijtnrx"),
	// 		Type: to.Ptr("ngvrxwzxtpqawwvhf"),
	// 		SystemData: &armworkloadorchestration.SystemData{
	// 			CreatedBy: to.Ptr("favedmahrbemfqzeuggazxzrvwugxw"),
	// 			CreatedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-24T11:04:49.597Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lywqfnyqrutroctdfbxzytel"),
	// 			LastModifiedByType: to.Ptr(armworkloadorchestration.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-24T11:04:49.597Z"); return t}()),
	// 		},
	// 	},
	// }
}
