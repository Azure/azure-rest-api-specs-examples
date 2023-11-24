package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/DetailsExperiment.json
func ExampleExperimentsClient_ExecutionDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExperimentsClient().ExecutionDetails(ctx, "exampleRG", "exampleExperiment", "f24500ad-744e-4a26-864b-b76199eac333", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExperimentExecutionDetails = armchaos.ExperimentExecutionDetails{
	// 	Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
	// 	Type: to.Ptr("Microsoft.Chaos/experiments/executions/getExecutionDetails"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executions/f24500ad-744e-4a26-864b-b76199eac333/getExecutionDetails"),
	// 	Properties: &armchaos.ExperimentExecutionDetailsProperties{
	// 		StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.255Z"); return t}()),
	// 		Status: to.Ptr("failed"),
	// 		StoppedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.928Z"); return t}()),
	// 		FailureReason: to.Ptr("Dependency failure"),
	// 		LastActionAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.255Z"); return t}()),
	// 		RunInformation: &armchaos.ExperimentExecutionDetailsPropertiesRunInformation{
	// 			Steps: []*armchaos.StepStatus{
	// 				{
	// 					Branches: []*armchaos.BranchStatus{
	// 						{
	// 							Actions: []*armchaos.ActionStatus{
	// 								{
	// 									ActionID: to.Ptr("59499d33-6751-4b6e-a1f6-58f4d56a040a"),
	// 									ActionName: to.Ptr("urn:provider:agent-v2:Microsoft.Azure.Chaos.Fault.CPUPressureAllProcessors"),
	// 									EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:13.627Z"); return t}()),
	// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:13.627Z"); return t}()),
	// 									Status: to.Ptr("failed"),
	// 									Targets: []*armchaos.ExperimentExecutionActionTargetDetailsProperties{
	// 										{
	// 											Status: to.Ptr("succeeded"),
	// 											Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 											TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55.000Z"); return t}()),
	// 											TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55.000Z"); return t}()),
	// 										},
	// 										{
	// 											Status: to.Ptr("failed"),
	// 											Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 											TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55.000Z"); return t}()),
	// 											TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55.000Z"); return t}()),
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
	// 	},
	// }
}
