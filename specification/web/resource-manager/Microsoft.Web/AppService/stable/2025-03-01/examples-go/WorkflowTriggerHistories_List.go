package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowTriggerHistories_List.json
func ExampleWorkflowTriggerHistoriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowTriggerHistoriesClient().NewListPager("testResourceGroup", "test-name", "testWorkflowName", "testTriggerName", &armappservice.WorkflowTriggerHistoriesClientListOptions{Top: nil,
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
		// page.WorkflowTriggerHistoryListResult = armappservice.WorkflowTriggerHistoryListResult{
		// 	Value: []*armappservice.WorkflowTriggerHistory{
		// 		{
		// 			ID: to.Ptr("/workflows/testWorkflowName/triggers/testTriggerName/histories/08586676746934337772206998657CU22"),
		// 			Name: to.Ptr("08586676746934337772206998657CU22"),
		// 			Type: to.Ptr("/workflows/triggers/histories"),
		// 			Properties: &armappservice.WorkflowTriggerHistoryProperties{
		// 				Code: to.Ptr("OK"),
		// 				Correlation: &armappservice.Correlation{
		// 					ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.298Z"); return t}()),
		// 				Fired: to.Ptr(true),
		// 				Run: &armappservice.ResourceReference{
		// 					Name: to.Ptr("08586676746934337772206998657CU22"),
		// 					Type: to.Ptr("/workflows/runs"),
		// 					ID: to.Ptr("/workflows/testWorkflowName/runs/08586676746934337772206998657CU22"),
		// 				},
		// 				ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}
