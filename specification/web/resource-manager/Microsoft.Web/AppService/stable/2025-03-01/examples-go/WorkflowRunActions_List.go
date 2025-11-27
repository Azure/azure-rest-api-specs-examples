package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowRunActions_List.json
func ExampleWorkflowRunActionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunActionsClient().NewListPager("test-resource-group", "test-name", "test-workflow", "08586676746934337772206998657CU22", &armappservice.WorkflowRunActionsClientListOptions{Top: nil,
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
		// page.WorkflowRunActionListResult = armappservice.WorkflowRunActionListResult{
		// 	Value: []*armappservice.WorkflowRunAction{
		// 		{
		// 			ID: to.Ptr("/workflows/test-workflow/runs/08586676746934337772206998657CU22/actions/HTTP"),
		// 			Name: to.Ptr("HTTP"),
		// 			Type: to.Ptr("/workflows/runs/actions"),
		// 			Properties: &armappservice.WorkflowRunActionProperties{
		// 				Code: to.Ptr("OK"),
		// 				Correlation: &armappservice.RunActionCorrelation{
		// 					ClientTrackingID: to.Ptr("08586676746934337772206998657CU22"),
		// 					ActionTrackingID: to.Ptr("56063357-45dd-4278-9be5-8220ce0cc9ca"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.545Z"); return t}()),
		// 				InputsLink: &armappservice.ContentLink{
		// 					ContentSize: to.Ptr[int64](138),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/000012345678900000000056a41a/runs/00000aaaabbbbccccddddeeeeffff/actions/HTTP/contents/ActionInputs?api-version=2018-11-01&code=examplecode=2022-04-13T03%3A00%3A00.0000000Z&sp=%2Fruns%2F00000aaaabbbbccccddddeeeeffff%2Factions%2FHTTP%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=dJ6F00000aaaabbbbccccddddeeeeffff78co"),
		// 				},
		// 				OutputsLink: &armappservice.ContentLink{
		// 					ContentSize: to.Ptr[int64](15660),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/000012345678900000000056a41a/runs/00000aaaabbbbccccddddeeeeffff/actions/HTTP/contents/ActionOutputs?api-version=2018-11-01&code=examplecode=2022-04-13T03%3A00%3A00.0000000Z&sp=%2Fruns%2F00000aaaabbbbccccddddeeeeffff%2Factions%2FHTTP%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=fIIgF00000aaaabbbbccccddddeeeeffffWRU0"),
		// 				},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T20:16:32.305Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 			},
		// 	}},
		// }
	}
}
