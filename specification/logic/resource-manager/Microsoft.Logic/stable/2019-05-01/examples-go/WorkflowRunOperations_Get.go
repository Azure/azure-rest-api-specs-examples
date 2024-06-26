package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunOperations_Get.json
func ExampleWorkflowRunOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunOperationsClient().Get(ctx, "testResourceGroup", "testFlow", "08586774142730039209110422528", "ebdcbbde-c4db-43ec-987c-fd0f7726f43b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRun = armlogic.WorkflowRun{
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testFlow/runs/08586774142730039209110422528"),
	// 	Name: to.Ptr("08586774142730039209110422528"),
	// 	Type: to.Ptr("Microsoft.Logic/workflows/runs"),
	// 	Properties: &armlogic.WorkflowRunProperties{
	// 		Correlation: &armlogic.Correlation{
	// 			ClientTrackingID: to.Ptr("08586774142730039209110422528"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-20T02:50:13.717Z"); return t}()),
	// 		Outputs: map[string]*armlogic.WorkflowOutputParameter{
	// 		},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-20T02:50:12.474Z"); return t}()),
	// 		Status: to.Ptr(armlogic.WorkflowStatusSucceeded),
	// 		Trigger: &armlogic.WorkflowRunTrigger{
	// 			Name: to.Ptr("Recurrence"),
	// 			Code: to.Ptr("OK"),
	// 			Correlation: &armlogic.Correlation{
	// 				ClientTrackingID: to.Ptr("08586774142730039209110422528"),
	// 			},
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-20T02:50:12.460Z"); return t}()),
	// 			ScheduledTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-20T02:50:12.141Z"); return t}()),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-20T02:50:12.460Z"); return t}()),
	// 			Status: to.Ptr(armlogic.WorkflowStatusSucceeded),
	// 		},
	// 		WaitEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-20T02:50:12.474Z"); return t}()),
	// 		Workflow: &armlogic.ResourceReference{
	// 			Name: to.Ptr("08586993867806980512"),
	// 			Type: to.Ptr("Microsoft.Logic/workflows/versions"),
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/workflows/testFlow/versions/08586993867806980512"),
	// 		},
	// 	},
	// }
}
