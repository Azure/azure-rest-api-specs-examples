package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionScopeRepetitions_Get.json
func ExampleWorkflowRunActionScopeRepetitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunActionScopeRepetitionsClient().Get(ctx, "testResourceGroup", "testFlow", "08586776228332053161046300351", "for_each", "000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRunActionRepetitionDefinition = armlogic.WorkflowRunActionRepetitionDefinition{
	// 	Name: to.Ptr("000000"),
	// 	Type: to.Ptr("Microsoft.Logic/workflows/runs/actions/scopeRepetitions"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testFlow/runs/08586776228332053161046300351/actions/for_each/scopeRepetitions/000000"),
	// 	Properties: &armlogic.WorkflowRunActionRepetitionProperties{
	// 		Code: to.Ptr("NotSpecified"),
	// 		Correlation: &armlogic.RunActionCorrelation{
	// 			ClientTrackingID: to.Ptr("08586775357427610445444523191"),
	// 			ActionTrackingID: to.Ptr("5c0e7c24-4891-44e8-b631-8084c5531dd5"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.624Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.209Z"); return t}()),
	// 		Status: to.Ptr(armlogic.WorkflowStatusSucceeded),
	// 		RepetitionIndexes: []*armlogic.RepetitionIndex{
	// 			{
	// 				ItemIndex: to.Ptr[int32](0),
	// 				ScopeName: to.Ptr("For_each"),
	// 		}},
	// 	},
	// }
}
