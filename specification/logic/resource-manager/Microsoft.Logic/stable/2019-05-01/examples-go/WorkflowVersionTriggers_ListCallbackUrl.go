package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowVersionTriggers_ListCallbackUrl.json
func ExampleWorkflowVersionTriggersClient_ListCallbackURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowVersionTriggersClient().ListCallbackURL(ctx, "testResourceGroup", "testWorkflowName", "testWorkflowVersionId", "testTriggerName", &armlogic.WorkflowVersionTriggersClientListCallbackURLOptions{Parameters: &armlogic.GetCallbackURLParameters{
		KeyType:  to.Ptr(armlogic.KeyTypePrimary),
		NotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-05T08:00:00.000Z"); return t }()),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerCallbackURL = armlogic.WorkflowTriggerCallbackURL{
	// 	Method: to.Ptr("POST"),
	// 	BasePath: to.Ptr("http://test-host:100/workflows/fb9c8d79b15f41ce9b12861862f43546/versions/08587100027316071865/triggers/manualTrigger/paths/invoke"),
	// 	Queries: &armlogic.WorkflowTriggerListCallbackURLQueries{
	// 		APIVersion: to.Ptr("2015-08-01-preview"),
	// 		Sig: to.Ptr("IxEQ_ygZf6WNEQCbjV0Vs6p6Y4DyNEJVAa86U5B4xhk"),
	// 		Sp: to.Ptr("/versions/08587100027316071865/triggers/manualTrigger/run"),
	// 		Sv: to.Ptr("1.0"),
	// 	},
	// 	Value: to.Ptr("http://test-host:100/workflows/fb9c8d79b15f41ce9b12861862f43546/versions/08587100027316071865/triggers/manualTrigger/paths/invoke?api-version=2015-08-01-preview&sp=%2Fversions%2F08587100027316071865%2Ftriggers%2FmanualTrigger%2Frun&sv=1.0&sig=IxEQ_ygZf6WNEQCbjV0Vs6p6Y4DyNEJVAa86U5B4xhk"),
	// }
}
