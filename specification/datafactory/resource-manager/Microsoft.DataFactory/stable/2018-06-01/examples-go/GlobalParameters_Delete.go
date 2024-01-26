package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be0d1c383d683a8eb22ab5fd5b75e380ac3c2eea/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Delete.json
func ExampleGlobalParametersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGlobalParametersClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
