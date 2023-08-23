package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/GetAExperimentExecutionDetails.json
func ExampleExperimentsClient_GetExecutionDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExperimentsClient().GetExecutionDetails(ctx, "exampleRG", "exampleExperiment", "f24500ad-744e-4a26-864b-b76199eac333", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExperimentExecutionDetails = armchaos.ExperimentExecutionDetails{
	// 	Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
	// 	Type: to.Ptr("Microsoft.Chaos/experiments/executionDetails"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executionDetails/f24500ad-744e-4a26-864b-b76199eac333"),
	// 	Properties: &armchaos.ExperimentExecutionDetailsProperties{
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:42.917983Z"); return t}()),
	// 		ExperimentID: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
	// 		FailureReason: to.Ptr("Dependency failure"),
	// 		LastActionDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
	// 		RunInformation: &armchaos.ExperimentExecutionDetailsPropertiesRunInformation{
	// 			Steps: []*armchaos.StepStatus{
	// 				{
	// 					Branches: []*armchaos.BranchStatus{
	// 						{
	// 							Actions: []*armchaos.ActionStatus{
	// 								{
	// 									ActionID: to.Ptr("59499d33-6751-4b6e-a1f6-58f4d56a040a"),
	// 									ActionName: to.Ptr("urn:provider:agent-v2:Microsoft.Azure.Chaos.Fault.CPUPressureAllProcessors"),
	// 									EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T13:56:13.6270153-08:00"); return t}()),
	// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T13:56:13.6270153-08:00"); return t}()),
	// 									Status: to.Ptr("failed"),
	// 									Targets: []*armchaos.ExperimentExecutionActionTargetDetailsProperties{
	// 										{
	// 											Status: to.Ptr("succeeded"),
	// 											Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 											TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55+00:00"); return t}()),
	// 											TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55+00:00"); return t}()),
	// 										},
	// 										{
	// 											Status: to.Ptr("failed"),
	// 											Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 											TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55+00:00"); return t}()),
	// 											TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55+00:00"); return t}()),
	// 									}},
	// 							}},
	// 							BranchID: to.Ptr("FirstBranch"),
	// 							BranchName: to.Ptr("FirstBranch"),
	// 							Status: to.Ptr("failed"),
	// 					}},
	// 					Status: to.Ptr("failed"),
	// 					StepID: to.Ptr("FirstStep"),
	// 					StepName: to.Ptr("FirstStep"),
	// 			}},
	// 		},
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
	// 		Status: to.Ptr("failed"),
	// 		StopDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.9281956Z"); return t}()),
	// 	},
	// }
}
