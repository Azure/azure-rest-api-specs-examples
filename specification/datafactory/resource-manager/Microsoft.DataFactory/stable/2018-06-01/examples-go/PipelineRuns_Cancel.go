package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Cancel.json
func ExamplePipelineRunsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPipelineRunsClient().Cancel(ctx, "exampleResourceGroup", "exampleFactoryName", "16ac5348-ff82-4f95-a80d-638c1d47b721", &armdatafactory.PipelineRunsClientCancelOptions{IsRecursive: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
