package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActions_Get.json
func ExampleWorkflowRunActionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunActionsClient().Get(ctx, "test-resource-group", "test-workflow", "08586676746934337772206998657CU22", "HTTP", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRunAction = armlogic.WorkflowRunAction{
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/workflows/test-workflow/runs/08586676746934337772206998657CU22/actions/HTTP"),
	// 	Name: to.Ptr("HTTP"),
	// 	Type: to.Ptr("Microsoft.Logic/workflows/runs/actions"),
	// 	Properties: &armlogic.WorkflowRunActionProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armlogic.RunActionCorrelation{
	// 			ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
	// 			ActionTrackingID: to.Ptr("56063357-45dd-4278-9be5-8220ce0cc9ca"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.545Z"); return t}()),
	// 		InputsLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("XXUACojv0aBfuP56os3CWw=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](46),
	// 			ContentVersion: to.Ptr("5XUACojv0aBfuP56os3CWw=="),
	// 			URI: to.Ptr("https://tempuri.org"),
	// 		},
	// 		OutputsLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("XX6KRrevI6AAxEvo9FeBYQ=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](11873),
	// 			ContentVersion: to.Ptr("6X6KRrevI6AAxEvo9FeBYQ=="),
	// 			URI: to.Ptr("https://tempuri.org"),
	// 		},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.305Z"); return t}()),
	// 		Status: to.Ptr(armlogic.WorkflowStatusSucceeded),
	// 	},
	// }
}
