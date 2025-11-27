package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowRunActionRepetitions_Get.json
func ExampleWorkflowRunActionRepetitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunActionRepetitionsClient().Get(ctx, "testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "testAction", "000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRunActionRepetitionDefinition = armappservice.WorkflowRunActionRepetitionDefinition{
	// 	Name: to.Ptr("000001"),
	// 	Type: to.Ptr("/workflows/runs/actions/repetitions"),
	// 	ID: to.Ptr("/workflows/testFlow/runs/08586776228332053161046300351/actions/testAction/repetitions/000001"),
	// 	Properties: &armappservice.WorkflowRunActionRepetitionProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armappservice.RunActionCorrelation{
	// 			ClientTrackingID: to.Ptr("08586775357427610445444523191"),
	// 			ActionTrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.101Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.092Z"); return t}()),
	// 		Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 		InputsLink: &armappservice.ContentLink{
	// 			ContentHash: &armappservice.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](6),
	// 			ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionInputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=dLmnt50joimEMK4k9rR6njHQh94iSFJ9rrDxFbkEg5M"),
	// 		},
	// 		OutputsLink: &armappservice.ContentLink{
	// 			ContentHash: &armappservice.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](6),
	// 			ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionOutputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=B3-X5sqIAv1Lb31GOD34ZgIRUXGuiM2QllWiNwXFYAw"),
	// 		},
	// 		TrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
	// 		RepetitionIndexes: []*armappservice.RepetitionIndex{
	// 			{
	// 				ItemIndex: to.Ptr[int32](1),
	// 				ScopeName: to.Ptr("For_each"),
	// 		}},
	// 	},
	// }
}
