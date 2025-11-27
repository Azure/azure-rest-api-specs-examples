package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowTriggers_ListCallbackUrl.json
func ExampleWorkflowTriggersClient_ListCallbackURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowTriggersClient().ListCallbackURL(ctx, "test-resource-group", "test-name", "test-workflow", "manual", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerCallbackURL = armappservice.WorkflowTriggerCallbackURL{
	// 	Method: to.Ptr("POST"),
	// 	BasePath: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/fb9c8d79b15f41ce9b12861862f43546/versions/08587100027316071865/triggers/manualTrigger/paths/invoke"),
	// 	Queries: &armappservice.WorkflowTriggerListCallbackURLQueries{
	// 		APIVersion: to.Ptr("2018-07-01-preview"),
	// 		Sig: to.Ptr("IxEQ_ygZf6WNEQCbjV0Vs6p6Y4DyNEJVAa86U5B4xhk"),
	// 		Sp: to.Ptr("/versions/08587100027316071865/triggers/manualTrigger/run"),
	// 		Sv: to.Ptr("1.0"),
	// 	},
	// 	Value: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/fb9c8d79b15f41ce9b12861862f43546/versions/08587100027316071865/triggers/manualTrigger/paths/invoke?api-version=2015-08-01-preview&sp=%2Fversions%2F08587100027316071865%2Ftriggers%2FmanualTrigger%2Frun&sv=1.0&sig=IxEQ_ygZf6WNEQCbjV0Vs6p6Y4DyNEJVAa86U5B4xhk"),
	// }
}
