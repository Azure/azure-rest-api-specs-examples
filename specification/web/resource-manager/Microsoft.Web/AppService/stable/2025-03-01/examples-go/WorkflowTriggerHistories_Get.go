package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowTriggerHistories_Get.json
func ExampleWorkflowTriggerHistoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowTriggerHistoriesClient().Get(ctx, "testResourceGroup", "test-name", "testWorkflowName", "testTriggerName", "08586676746934337772206998657CU22", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerHistory = armappservice.WorkflowTriggerHistory{
	// 	ID: to.Ptr("/workflows/testWorkflowName/triggers/testTriggerName/histories/08586676746934337772206998657CU22"),
	// 	Name: to.Ptr("08586676746934337772206998657CU22"),
	// 	Type: to.Ptr("/workflows/triggers/histories"),
	// 	Properties: &armappservice.WorkflowTriggerHistoryProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armappservice.Correlation{
	// 			ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.298Z"); return t}()),
	// 		Fired: to.Ptr(true),
	// 		Run: &armappservice.ResourceReference{
	// 			Name: to.Ptr("08586676746934337772206998657CU22"),
	// 			Type: to.Ptr("/workflows/runs"),
	// 			ID: to.Ptr("/workflows/testWorkflowName/runs/08586676746934337772206998657CU22"),
	// 		},
	// 		ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.634Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.038Z"); return t}()),
	// 		Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 	},
	// }
}
