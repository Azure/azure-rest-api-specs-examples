package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Create.json
func ExampleIntegrationRuntimesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", armdatafactory.IntegrationRuntimeResource{
		Properties: &armdatafactory.SelfHostedIntegrationRuntime{
			Type:        to.Ptr(armdatafactory.IntegrationRuntimeTypeSelfHosted),
			Description: to.Ptr("A selfhosted integration runtime"),
		},
	}, &armdatafactory.IntegrationRuntimesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeResource = armdatafactory.IntegrationRuntimeResource{
	// 	Name: to.Ptr("exampleIntegrationRuntime"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/integrationruntimes"),
	// 	Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/integrationruntimes/exampleIntegrationRuntime"),
	// 	Properties: &armdatafactory.SelfHostedIntegrationRuntime{
	// 		Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeSelfHosted),
	// 		Description: to.Ptr("A selfhosted integration runtime"),
	// 	},
	// }
}
