package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowTriggers_Get.json
func ExampleWorkflowTriggersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowTriggersClient().Get(ctx, "test-resource-group", "test-name", "test-workflow", "manual", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTrigger = armappservice.WorkflowTrigger{
	// 	ID: to.Ptr("/workflows/test-workflow/triggers/manual"),
	// 	Name: to.Ptr("manual"),
	// 	Type: to.Ptr("/workflows/triggers"),
	// 	Properties: &armappservice.WorkflowTriggerProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:47:49.528Z"); return t}()),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T17:32:30.249Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armappservice.WorkflowTriggerProvisioningStateSucceeded),
	// 		State: to.Ptr(armappservice.WorkflowStateEnabled),
	// 		Workflow: &armappservice.ResourceReference{
	// 			Name: to.Ptr("08586676800160476478"),
	// 			Type: to.Ptr("/workflows/versions"),
	// 			ID: to.Ptr("/workflows/test-workflow/versions/08586676800160476478"),
	// 		},
	// 	},
	// }
}
