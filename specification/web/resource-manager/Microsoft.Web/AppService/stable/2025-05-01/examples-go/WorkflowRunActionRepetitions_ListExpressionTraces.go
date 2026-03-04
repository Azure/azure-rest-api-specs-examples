package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/WorkflowRunActionRepetitions_ListExpressionTraces.json
func ExampleWorkflowRunActionRepetitionsClient_NewListExpressionTracesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunActionRepetitionsClient().NewListExpressionTracesPager("testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "testAction", "000001", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Inputs {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armappservice.WorkflowRunActionRepetitionsClientListExpressionTracesResponse{
		// 	ExpressionTraces: armappservice.ExpressionTraces{
		// 		Inputs: []*armappservice.ExpressionRoot{
		// 			{
		// 				Path: to.Ptr(""),
		// 				Text: to.Ptr("items('For_each')?['OccuringLocation']?['Environment']"),
		// 				Value: "PROD",
		// 			},
		// 		},
		// 	},
		// }
	}
}
