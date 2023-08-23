package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/ListExperimentExecutionsDetails.json
func ExampleExperimentsClient_NewListExecutionDetailsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExperimentsClient().NewListExecutionDetailsPager("exampleRG", "exampleExperiment", nil)
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
		// page.ExperimentExecutionDetailsListResult = armchaos.ExperimentExecutionDetailsListResult{
		// 	Value: []*armchaos.ExperimentExecutionDetails{
		// 		{
		// 			Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
		// 			Type: to.Ptr("Microsoft.Chaos/experiments/executionDetails"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executionDetails/f24500ad-744e-4a26-864b-b76199eac333"),
		// 			Properties: &armchaos.ExperimentExecutionDetailsProperties{
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:42.917983Z"); return t}()),
		// 				ExperimentID: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
		// 				FailureReason: to.Ptr("Dependency failure"),
		// 				LastActionDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
		// 				RunInformation: &armchaos.ExperimentExecutionDetailsPropertiesRunInformation{
		// 					Steps: []*armchaos.StepStatus{
		// 						{
		// 							Branches: []*armchaos.BranchStatus{
		// 								{
		// 									Actions: []*armchaos.ActionStatus{
		// 										{
		// 											ActionID: to.Ptr("59499d33-6751-4b6e-a1f6-58f4d56a040a"),
		// 											ActionName: to.Ptr("urn:provider:agent-v2:Microsoft.Azure.Chaos.Fault.CPUPressureAllProcessors"),
		// 											EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T13:56:13.6270153-08:00"); return t}()),
		// 											StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T13:56:13.6270153-08:00"); return t}()),
		// 											Status: to.Ptr("failed"),
		// 											Targets: []*armchaos.ExperimentExecutionActionTargetDetailsProperties{
		// 												{
		// 													Status: to.Ptr("succeeded"),
		// 													Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
		// 													TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55+00:00"); return t}()),
		// 													TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55+00:00"); return t}()),
		// 											}},
		// 									}},
		// 									BranchID: to.Ptr("FirstBranch"),
		// 									BranchName: to.Ptr("FirstBranch"),
		// 									Status: to.Ptr("failed"),
		// 							}},
		// 							Status: to.Ptr("failed"),
		// 							StepID: to.Ptr("FirstStep"),
		// 							StepName: to.Ptr("FirstStep"),
		// 					}},
		// 				},
		// 				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
		// 				Status: to.Ptr("failed"),
		// 				StopDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.9281956Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("14d98367-52ef-4596-be4f-53fc81bbfc33"),
		// 			Type: to.Ptr("Microsoft.Chaos/experiments/executionDetails"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executionDetails/14d98367-52ef-4596-be4f-53fc81bbfc33"),
		// 			Properties: &armchaos.ExperimentExecutionDetailsProperties{
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:42.917983Z"); return t}()),
		// 				ExperimentID: to.Ptr("14d98367-52ef-4596-be4f-53fc81bbfc33"),
		// 				FailureReason: to.Ptr(""),
		// 				LastActionDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
		// 				RunInformation: &armchaos.ExperimentExecutionDetailsPropertiesRunInformation{
		// 					Steps: []*armchaos.StepStatus{
		// 						{
		// 							Branches: []*armchaos.BranchStatus{
		// 								{
		// 									Actions: []*armchaos.ActionStatus{
		// 										{
		// 											ActionID: to.Ptr("59499d33-6751-4b6e-a1f6-58f4d56a040a"),
		// 											ActionName: to.Ptr("urn:provider:agent-v2:Microsoft.Azure.Chaos.Fault.CPUPressureAllProcessors"),
		// 											EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T13:56:13.6270153-08:00"); return t}()),
		// 											StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T13:56:13.6270153-08:00"); return t}()),
		// 											Status: to.Ptr("success"),
		// 											Targets: []*armchaos.ExperimentExecutionActionTargetDetailsProperties{
		// 												{
		// 													Status: to.Ptr("succeeded"),
		// 													Target: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/VM1"),
		// 													TargetCompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T17:30:55+00:00"); return t}()),
		// 													TargetFailedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T16:30:55+00:00"); return t}()),
		// 											}},
		// 									}},
		// 									BranchID: to.Ptr("FirstBranch"),
		// 									BranchName: to.Ptr("FirstBranch"),
		// 									Status: to.Ptr("success"),
		// 							}},
		// 							Status: to.Ptr("success"),
		// 							StepID: to.Ptr("FirstStep"),
		// 							StepName: to.Ptr("FirstStep"),
		// 					}},
		// 				},
		// 				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
		// 				Status: to.Ptr("success"),
		// 				StopDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.9281956Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
