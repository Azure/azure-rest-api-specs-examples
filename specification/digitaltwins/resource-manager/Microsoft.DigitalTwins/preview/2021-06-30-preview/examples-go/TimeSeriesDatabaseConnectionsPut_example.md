Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdigitaltwins%2Farmdigitaltwins%2Fv0.4.0/sdk/resourcemanager/digitaltwins/armdigitaltwins/README.md) on how to add the SDK to your project and authenticate.

```go
package armdigitaltwins_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/TimeSeriesDatabaseConnectionsPut_example.json
func ExampleTimeSeriesDatabaseConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<time-series-database-connection-name>",
		armdigitaltwins.TimeSeriesDatabaseConnection{
			Properties: &armdigitaltwins.AzureDataExplorerConnectionProperties{
				ConnectionType:              to.Ptr(armdigitaltwins.ConnectionTypeAzureDataExplorer),
				AdxDatabaseName:             to.Ptr("<adx-database-name>"),
				AdxEndpointURI:              to.Ptr("<adx-endpoint-uri>"),
				AdxResourceID:               to.Ptr("<adx-resource-id>"),
				AdxTableName:                to.Ptr("<adx-table-name>"),
				EventHubEndpointURI:         to.Ptr("<event-hub-endpoint-uri>"),
				EventHubEntityPath:          to.Ptr("<event-hub-entity-path>"),
				EventHubNamespaceResourceID: to.Ptr("<event-hub-namespace-resource-id>"),
			},
		},
		&armdigitaltwins.TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
