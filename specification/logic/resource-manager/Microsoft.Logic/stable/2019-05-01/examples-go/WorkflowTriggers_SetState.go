package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_SetState.json
func ExampleWorkflowTriggersClient_SetState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogic.NewWorkflowTriggersClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.SetState(ctx,
		"testResourceGroup",
		"testWorkflow",
		"testTrigger",
		armlogic.SetTriggerStateActionDefinition{
			Source: &armlogic.WorkflowTriggerReference{
				ID: to.Ptr("subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sourceResGroup/providers/Microsoft.Logic/workflows/sourceWorkflow/triggers/sourceTrigger"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
