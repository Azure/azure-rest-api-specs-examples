package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowRunActionScopeRepetitions_List.json
func ExampleWorkflowRunActionScopeRepetitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunActionScopeRepetitionsClient().NewListPager("testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "for_each", nil)
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
		// page.WorkflowRunActionRepetitionDefinitionCollection = armappservice.WorkflowRunActionRepetitionDefinitionCollection{
		// 	Value: []*armappservice.WorkflowRunActionRepetitionDefinition{
		// 		{
		// 			Name: to.Ptr("000000"),
		// 			Type: to.Ptr("/workflows/runs/actions/scopeRepetitions"),
		// 			ID: to.Ptr("/workflows/testFlow/runs/08586776228332053161046300351/actions/for_each/scopeRepetitions/000000"),
		// 			Properties: &armappservice.WorkflowRunActionRepetitionProperties{
		// 				Code: to.Ptr("NotSpecified"),
		// 				Correlation: &armappservice.RunActionCorrelation{
		// 					ClientTrackingID: to.Ptr("08586775357427610445444523191"),
		// 					ActionTrackingID: to.Ptr("5c0e7c24-4891-44e8-b631-8084c5531dd5"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.624Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.209Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				RepetitionIndexes: []*armappservice.RepetitionIndex{
		// 					{
		// 						ItemIndex: to.Ptr[int32](0),
		// 						ScopeName: to.Ptr("For_each"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("000001"),
		// 			Type: to.Ptr("/workflows/runs/actions/scopeRepetitions"),
		// 			ID: to.Ptr("/workflows/testFlow/runs/08586776228332053161046300351/actions/for_each/scopeRepetitions/000001"),
		// 			Properties: &armappservice.WorkflowRunActionRepetitionProperties{
		// 				Code: to.Ptr("NotSpecified"),
		// 				Correlation: &armappservice.RunActionCorrelation{
		// 					ClientTrackingID: to.Ptr("08586775357427610445444523191"),
		// 					ActionTrackingID: to.Ptr("1682d0bd-7189-4478-b333-fe364a80bd69"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.227Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.068Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				RepetitionIndexes: []*armappservice.RepetitionIndex{
		// 					{
		// 						ItemIndex: to.Ptr[int32](1),
		// 						ScopeName: to.Ptr("For_each"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
