package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/IntegrationRuntimes_Get.json
func ExampleIntegrationRuntimesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatafactory.IntegrationRuntimesClientGetResponse{
	// 	IntegrationRuntimeResource: &armdatafactory.IntegrationRuntimeResource{
	// 		Name: to.Ptr("exampleIntegrationRuntime"),
	// 		Type: to.Ptr("Microsoft.DataFactory/factories/integrationruntimes"),
	// 		Etag: to.Ptr("15003c4f-0000-0200-0000-5cbe090b0000"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/integrationruntimes/exampleIntegrationRuntime"),
	// 		Properties: &armdatafactory.SelfHostedIntegrationRuntime{
	// 			Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeSelfHosted),
	// 			Description: to.Ptr("A selfhosted integration runtime"),
	// 		},
	// 	},
	// }
}
