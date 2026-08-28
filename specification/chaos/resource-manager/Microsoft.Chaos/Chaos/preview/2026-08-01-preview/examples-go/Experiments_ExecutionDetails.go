package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v3"
)

// Generated from example definition: 2026-08-01-preview/Experiments_ExecutionDetails.json
func ExampleExperimentsClient_ExecutionDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
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
	// res = armchaos.ExperimentsClientExecutionDetailsResponse{
	// 	ExperimentExecutionDetails: armchaos.ExperimentExecutionDetails{
	// 		Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
	// 		Type: to.Ptr("Microsoft.Chaos/experiments/executions/getExecutionDetails"),
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executions/f24500ad-744e-4a26-864b-b76199eac333/getExecutionDetails"),
	// 		Properties: &armchaos.ExperimentExecutionDetailsProperties{
	// 			FailureReason: to.Ptr("Dependency failure"),
	// 			LastActionAt: to.Ptr(time.Date(2020, time.December, 14, 21, 52, 52, 255257400, time.UTC)),
	// 			RunInformation: &armchaos.ExperimentExecutionDetailsPropertiesRunInformation{
	// 				Steps: []*armchaos.StepStatus{
	// 					{
	// 						Branches: []*armchaos.BranchStatus{
	// 							{
	// 								Actions: []*armchaos.ActionStatus{
	// 									{
	// 										ActionID: to.Ptr("59499d33-6751-4b6e-a1f6-58f4d56a040a"),
	// 										ActionName: to.Ptr("urn:provider:agent-v2:Microsoft.Azure.Chaos.Fault.CPUPressureAllProcessors"),
	// 										EndTime: to.Ptr(time.Date(2020, time.December, 14, 21, 56, 13, 627015300, time.UTC)),
	// 										StartTime: to.Ptr(time.Date(2020, time.December, 14, 21, 56, 13, 627015300, time.UTC)),
	// 										Status: to.Ptr("failed"),
	// 										Targets: []*armchaos.ExperimentExecutionActionTargetDetailsProperties{
	// 											{
	// 												Status: to.Ptr("succeeded"),
	// 												Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 												TargetCompletedTime: to.Ptr(time.Date(2021, time.April, 2, 17, 30, 55, 0, time.UTC)),
	// 												TargetFailedTime: to.Ptr(time.Date(2021, time.April, 2, 16, 30, 55, 0, time.UTC)),
	// 											},
	// 											{
	// 												Status: to.Ptr("failed"),
	// 												Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
	// 												TargetCompletedTime: to.Ptr(time.Date(2021, time.April, 2, 17, 30, 55, 0, time.UTC)),
	// 												TargetFailedTime: to.Ptr(time.Date(2021, time.April, 2, 16, 30, 55, 0, time.UTC)),
	// 											},
	// 										},
	// 									},
	// 								},
	// 								BranchID: to.Ptr("FirstBranch"),
	// 								BranchName: to.Ptr("FirstBranch"),
	// 								Status: to.Ptr("failed"),
	// 							},
	// 						},
	// 						Status: to.Ptr("failed"),
	// 						StepID: to.Ptr("FirstStep"),
	// 						StepName: to.Ptr("FirstStep"),
	// 					},
	// 				},
	// 			},
	// 			StartedAt: to.Ptr(time.Date(2020, time.December, 14, 21, 52, 52, 255257400, time.UTC)),
	// 			Status: to.Ptr("failed"),
	// 			StoppedAt: to.Ptr(time.Date(2020, time.December, 14, 21, 56, 18, 928195600, time.UTC)),
	// 		},
	// 	},
	// }
}
