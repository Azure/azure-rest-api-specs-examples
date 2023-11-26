package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ListCallbackUrl.json
func ExampleWorkflowsClient_ListCallbackURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowsClient().ListCallbackURL(ctx, "testResourceGroup", "testWorkflow", armlogic.GetCallbackURLParameters{
		KeyType:  to.Ptr(armlogic.KeyTypePrimary),
		NotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-19T16:00:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerCallbackURL = armlogic.WorkflowTriggerCallbackURL{
	// 	Method: to.Ptr("POST"),
	// 	BasePath: to.Ptr("https://prod-03.westus.logic.azure.com/workflows/d4690301f3b340de9330256bb2e83974"),
	// 	Queries: &armlogic.WorkflowTriggerListCallbackURLQueries{
	// 		APIVersion: to.Ptr("2018-07-01-preview"),
	// 		Se: to.Ptr("2018-04-19T16:00:00.0000000Z"),
	// 		Sig: to.Ptr("JXbHjs3qzLPDyk-IM_VzsN-WL_mNql3v_uWbBbKgtVk"),
	// 		Sp: to.Ptr("//*"),
	// 		Sv: to.Ptr("1.0"),
	// 	},
	// 	Value: to.Ptr("https://prod-03.westus.logic.azure.com/workflows/d4690301f3b340de9330256bb2e83974/triggers/requestTrigger/paths/invoke?api-version=2018-07-01-preview&se=2018-04-19T16%3A00%3A00.0000000Z&sp=%2Ftriggers%2FrequestTrigger%2Frun&sv=1.0&sig=JXbHjs3qzLPDyk-IM_VzsN-WL_mNql3v_uWbBbKgtVk"),
	// }
}
