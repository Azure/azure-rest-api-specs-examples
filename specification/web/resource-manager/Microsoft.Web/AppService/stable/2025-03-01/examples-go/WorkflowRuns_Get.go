package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowRuns_Get.json
func ExampleWorkflowRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunsClient().Get(ctx, "test-resource-group", "test-name", "test-workflow", "08586676746934337772206998657CU22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRun = armappservice.WorkflowRun{
	// 	ID: to.Ptr("/workflows/test-workflow/runs/08586676746934337772206998657CU22"),
	// 	Name: to.Ptr("08586676746934337772206998657CU22"),
	// 	Type: to.Ptr("workflows/runs"),
	// 	Properties: &armappservice.WorkflowRunProperties{
	// 		Correlation: &armappservice.Correlation{
	// 			ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.577Z"); return t}()),
	// 		Outputs: map[string]*armappservice.WorkflowOutputParameter{
	// 		},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
	// 		Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 		Trigger: &armappservice.WorkflowRunTrigger{
	// 			Name: to.Ptr("Recurrence"),
	// 			Code: to.Ptr("OK"),
	// 			Correlation: &armappservice.Correlation{
	// 				ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 			},
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
	// 			Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 		},
	// 		WaitEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.044Z"); return t}()),
	// 		Workflow: &armappservice.ResourceReference{
	// 			Name: to.Ptr("08586676754160363885"),
	// 			Type: to.Ptr("workflows/versions"),
	// 			ID: to.Ptr("/workflows/test-workflow/versions/08586676754160363885"),
	// 		},
	// 	},
	// }
}
