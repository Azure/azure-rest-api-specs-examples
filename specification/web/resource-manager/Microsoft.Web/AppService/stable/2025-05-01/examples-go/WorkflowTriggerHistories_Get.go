package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/WorkflowTriggerHistories_Get.json
func ExampleWorkflowTriggerHistoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
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
	// res = armappservice.WorkflowTriggerHistoriesClientGetResponse{
	// 	WorkflowTriggerHistory: &armappservice.WorkflowTriggerHistory{
	// 		Name: to.Ptr("08586676746934337772206998657CU22"),
	// 		Type: to.Ptr("Microsoft.Web/sites/workflows/triggers/histories"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Web/sites/test-name/workflows/testWorkflowName/triggers/testTriggerName/histories/08586676746934337772206998657CU22"),
	// 		Properties: &armappservice.WorkflowTriggerHistoryProperties{
	// 			Code: to.Ptr("OK"),
	// 			Correlation: &armappservice.Correlation{
	// 				ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 			},
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.2987996Z"); return t}()),
	// 			Fired: to.Ptr(true),
	// 			Run: &armappservice.ResourceReference{
	// 				Name: to.Ptr("08586676746934337772206998657CU22"),
	// 				Type: to.Ptr("Microsoft.Web/sites/workflows/runs"),
	// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Web/sites/test-name/workflows/testWorkflowName/runs/08586676746934337772206998657CU22"),
	// 			},
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:31.6344174Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.0387927Z"); return t}()),
	// 			Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 		},
	// 	},
	// }
}
