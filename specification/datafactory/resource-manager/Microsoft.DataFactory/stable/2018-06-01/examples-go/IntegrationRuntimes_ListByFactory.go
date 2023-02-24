package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4afa6837cfb404d8e5ffa8a604a5e09996d6f79e/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_ListByFactory.json
func ExampleIntegrationRuntimesClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewIntegrationRuntimesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IntegrationRuntimeListResponse = armdatafactory.IntegrationRuntimeListResponse{
		// 	Value: []*armdatafactory.IntegrationRuntimeResource{
		// 		{
		// 			Name: to.Ptr("exampleIntegrationRuntime"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/integrationruntimes"),
		// 			Etag: to.Ptr("0400f1a1-0000-0000-0000-5b2188640000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/integrationruntimes/exampleIntegrationRuntime"),
		// 			Properties: &armdatafactory.SelfHostedIntegrationRuntime{
		// 				Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeSelfHosted),
		// 				Description: to.Ptr("A selfhosted integration runtime"),
		// 			},
		// 	}},
		// }
	}
}
