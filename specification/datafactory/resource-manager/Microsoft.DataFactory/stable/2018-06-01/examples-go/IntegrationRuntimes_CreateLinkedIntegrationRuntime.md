Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_CreateLinkedIntegrationRuntime.json
func ExampleIntegrationRuntimesClient_CreateLinkedIntegrationRuntime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewIntegrationRuntimesClient("<subscription-id>", cred, nil)
	res, err := client.CreateLinkedIntegrationRuntime(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<integration-runtime-name>",
		armdatafactory.CreateLinkedIntegrationRuntimeRequest{
			Name:                to.StringPtr("<name>"),
			DataFactoryLocation: to.StringPtr("<data-factory-location>"),
			DataFactoryName:     to.StringPtr("<data-factory-name>"),
			SubscriptionID:      to.StringPtr("<subscription-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IntegrationRuntimesClientCreateLinkedIntegrationRuntimeResult)
}
```
