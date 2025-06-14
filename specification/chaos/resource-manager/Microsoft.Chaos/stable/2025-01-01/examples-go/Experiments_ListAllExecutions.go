package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2025-01-01/Experiments_ListAllExecutions.json
func ExampleExperimentsClient_NewListAllExecutionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExperimentsClient().NewListAllExecutionsPager("exampleRG", "exampleExperiment", nil)
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
		// page = armchaos.ExperimentsClientListAllExecutionsResponse{
		// 	ExperimentExecutionListResult: armchaos.ExperimentExecutionListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executionDetails?continuationToken=&api-version=2025-01-01"),
		// 		Value: []*armchaos.ExperimentExecution{
		// 			{
		// 				Name: to.Ptr("f24500ad-744e-4a26-864b-b76199eac333"),
		// 				Type: to.Ptr("Microsoft.Chaos/experiments/executions"),
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executions/f24500ad-744e-4a26-864b-b76199eac333"),
		// 				Properties: &armchaos.ExperimentExecutionProperties{
		// 					StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
		// 					Status: to.Ptr("failed"),
		// 					StoppedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.9281956Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("14d98367-52ef-4596-be4f-53fc81bbfc33"),
		// 				Type: to.Ptr("Microsoft.Chaos/experiments/executions"),
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/executions/14d98367-52ef-4596-be4f-53fc81bbfc33"),
		// 				Properties: &armchaos.ExperimentExecutionProperties{
		// 					StartedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:52:52.2552574Z"); return t}()),
		// 					Status: to.Ptr("success"),
		// 					StoppedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-14T21:56:18.9281956Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
