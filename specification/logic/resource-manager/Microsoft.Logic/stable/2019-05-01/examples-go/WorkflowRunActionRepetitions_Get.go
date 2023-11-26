package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionRepetitions_Get.json
func ExampleWorkflowRunActionRepetitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunActionRepetitionsClient().Get(ctx, "testResourceGroup", "testFlow", "08586776228332053161046300351", "testAction", "000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRunActionRepetitionDefinition = armlogic.WorkflowRunActionRepetitionDefinition{
	// 	Name: to.Ptr("000001"),
	// 	Type: to.Ptr("Microsoft.Logic/workflows/runs/actions/repetitions"),
	// 	ID: to.Ptr("api/management/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testFlow/runs/08586776228332053161046300351/actions/testAction/repetitions/000001"),
	// 	Properties: &armlogic.WorkflowRunActionRepetitionProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armlogic.RunActionCorrelation{
	// 			ClientTrackingID: to.Ptr("08586775357427610445444523191"),
	// 			ActionTrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.101Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.092Z"); return t}()),
	// 		Status: to.Ptr(armlogic.WorkflowStatusSucceeded),
	// 		InputsLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](6),
	// 			ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			URI: to.Ptr("https://prod-08.northcentralus.logic.azure.com:443/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionInputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=dLmnt50joimEMK4k9rR6njHQh94iSFJ9rrDxFbkEg5M"),
	// 		},
	// 		OutputsLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](6),
	// 			ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			URI: to.Ptr("https://prod-08.northcentralus.logic.azure.com:443/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionOutputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=B3-X5sqIAv1Lb31GOD34ZgIRUXGuiM2QllWiNwXFYAw"),
	// 		},
	// 		TrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
	// 		RepetitionIndexes: []*armlogic.RepetitionIndex{
	// 			{
	// 				ItemIndex: to.Ptr[int32](1),
	// 				ScopeName: to.Ptr("For_each"),
	// 		}},
	// 	},
	// }
}
