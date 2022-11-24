package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_CreateLinkedIntegrationRuntime.json
func ExampleIntegrationRuntimesClient_CreateLinkedIntegrationRuntime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewIntegrationRuntimesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateLinkedIntegrationRuntime(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", armdatafactory.CreateLinkedIntegrationRuntimeRequest{
		Name:                to.Ptr("bfa92911-9fb6-4fbe-8f23-beae87bc1c83"),
		DataFactoryLocation: to.Ptr("West US"),
		DataFactoryName:     to.Ptr("e9955d6d-56ea-4be3-841c-52a12c1a9981"),
		SubscriptionID:      to.Ptr("061774c7-4b5a-4159-a55b-365581830283"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
