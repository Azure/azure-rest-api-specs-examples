package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be0d1c383d683a8eb22ab5fd5b75e380ac3c2eea/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimeObjectMetadata_Refresh.json
func ExampleIntegrationRuntimeObjectMetadataClient_BeginRefresh() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimeObjectMetadataClient().BeginRefresh(ctx, "exampleResourceGroup", "exampleFactoryName", "testactivityv2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SsisObjectMetadataStatusResponse = armdatafactory.SsisObjectMetadataStatusResponse{
	// 	Name: to.Ptr("ca63c855b72d44959653ffcc6eb0b96c"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
