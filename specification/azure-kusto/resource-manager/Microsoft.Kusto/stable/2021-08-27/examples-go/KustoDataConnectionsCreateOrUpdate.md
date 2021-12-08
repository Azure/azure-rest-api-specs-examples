Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.1.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

```go
package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDataConnectionsCreateOrUpdate.json
func ExampleDataConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewDataConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		"<data-connection-name>",
		&armkusto.EventHubDataConnection{
			DataConnection: armkusto.DataConnection{
				Kind:     armkusto.DataConnectionKindEventHub.ToPtr(),
				Location: to.StringPtr("<location>"),
			},
			Properties: &armkusto.EventHubConnectionProperties{
				ConsumerGroup:             to.StringPtr("<consumer-group>"),
				EventHubResourceID:        to.StringPtr("<event-hub-resource-id>"),
				ManagedIdentityResourceID: to.StringPtr("<managed-identity-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DataConnectionClassification.GetDataConnection().ID: %s\n", *res.GetDataConnection().ID)
}
```
