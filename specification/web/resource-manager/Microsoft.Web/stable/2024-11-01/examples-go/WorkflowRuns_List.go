package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/WorkflowRuns_List.json
func ExampleWorkflowRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunsClient().NewListPager("test-resource-group", "test-name", "test-workflow", &armappservice.WorkflowRunsClientListOptions{Top: nil,
		Filter: nil,
	})
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
		// page.WorkflowRunListResult = armappservice.WorkflowRunListResult{
		// 	Value: []*armappservice.WorkflowRun{
		// 		{
		// 			ID: to.Ptr("workflows/test-workflow/runs/08586676746934337772206998657CU22"),
		// 			Name: to.Ptr("08586676746934337772206998657CU22"),
		// 			Type: to.Ptr("workflows/runs"),
		// 			Properties: &armappservice.WorkflowRunProperties{
		// 				Correlation: &armappservice.Correlation{
		// 					ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.577Z"); return t}()),
		// 				Outputs: map[string]*armappservice.WorkflowOutputParameter{
		// 				},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				Trigger: &armappservice.WorkflowRunTrigger{
		// 					Name: to.Ptr("Recurrence"),
		// 					Code: to.Ptr("OK"),
		// 					Correlation: &armappservice.Correlation{
		// 						ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 					},
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
		// 					ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
		// 					Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				},
		// 				WaitEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
		// 				Workflow: &armappservice.ResourceReference{
		// 					Name: to.Ptr("08586676754160363885"),
		// 					Type: to.Ptr("workflows/versions"),
		// 					ID: to.Ptr("workflows/test-workflow/versions/08586676754160363885"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
