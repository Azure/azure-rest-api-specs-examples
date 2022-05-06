Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.4.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDataConnectionEventGridValidationAsync.json
func ExampleDataConnectionsClient_BeginDataConnectionValidation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armkusto.NewDataConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDataConnectionValidation(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		armkusto.DataConnectionValidation{
			DataConnectionName: to.Ptr("<data-connection-name>"),
			Properties: &armkusto.EventGridDataConnection{
				Kind: to.Ptr(armkusto.DataConnectionKindEventGrid),
				Properties: &armkusto.EventGridConnectionProperties{
					BlobStorageEventType:      to.Ptr(armkusto.BlobStorageEventTypeMicrosoftStorageBlobCreated),
					ConsumerGroup:             to.Ptr("<consumer-group>"),
					DataFormat:                to.Ptr(armkusto.EventGridDataFormatJSON),
					DatabaseRouting:           to.Ptr(armkusto.DatabaseRoutingSingle),
					EventGridResourceID:       to.Ptr("<event-grid-resource-id>"),
					EventHubResourceID:        to.Ptr("<event-hub-resource-id>"),
					IgnoreFirstRecord:         to.Ptr(false),
					ManagedIdentityResourceID: to.Ptr("<managed-identity-resource-id>"),
					MappingRuleName:           to.Ptr("<mapping-rule-name>"),
					StorageAccountResourceID:  to.Ptr("<storage-account-resource-id>"),
					TableName:                 to.Ptr("<table-name>"),
				},
			},
		},
		&armkusto.DataConnectionsClientBeginDataConnectionValidationOptions{ResumeToken: ""})
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
