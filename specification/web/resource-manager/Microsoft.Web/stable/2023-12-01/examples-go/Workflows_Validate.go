package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/Workflows_Validate.json
func ExampleWorkflowsClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWorkflowsClient().Validate(ctx, "test-resource-group", "test-name", "test-workflow", armappservice.Workflow{
		Properties: &armappservice.WorkflowProperties{
			Definition: map[string]any{
				"$schema":        "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
				"actions":        map[string]any{},
				"contentVersion": "1.0.0.0",
				"outputs":        map[string]any{},
				"parameters":     map[string]any{},
				"triggers":       map[string]any{},
			},
			Kind: to.Ptr(armappservice.KindStateful),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
