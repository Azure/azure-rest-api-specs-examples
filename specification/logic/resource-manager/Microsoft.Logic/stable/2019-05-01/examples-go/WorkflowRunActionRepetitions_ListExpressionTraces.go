package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionRepetitions_ListExpressionTraces.json
func ExampleWorkflowRunActionRepetitionsClient_NewListExpressionTracesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunActionRepetitionsClient().NewListExpressionTracesPager("testResourceGroup", "testFlow", "08586776228332053161046300351", "testAction", "000001", nil)
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
		// page.ExpressionTraces = armlogic.ExpressionTraces{
		// 	Inputs: []*armlogic.ExpressionRoot{
		// 		{
		// 			Text: to.Ptr("items('For_each')?['OccuringLocation']?['Environment']"),
		// 			Value: "PROD",
		// 			Path: to.Ptr(""),
		// 	}},
		// }
	}
}
