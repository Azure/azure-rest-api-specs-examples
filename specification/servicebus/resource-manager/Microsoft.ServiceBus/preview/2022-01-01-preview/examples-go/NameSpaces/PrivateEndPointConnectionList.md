```go
package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/NameSpaces/PrivateEndPointConnectionList.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewPrivateEndpointConnectionsClient("subID", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("SDK-ServiceBus-4794",
		"sdk-Namespace-5828",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicebus%2Farmservicebus%2Fv2.0.0-beta.1/sdk/resourcemanager/servicebus/armservicebus/README.md) on how to add the SDK to your project and authenticate.
