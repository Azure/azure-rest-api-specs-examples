Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv0.5.0/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.

```go
package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/data-transfer-service/CosmosDBDataTransferJobCreate.json
func ExampleDataTransferJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcosmos.NewDataTransferJobsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<job-name>",
		armcosmos.CreateJobRequest{
			Properties: &armcosmos.DataTransferJobProperties{
				Destination: &armcosmos.AzureBlobDataTransferDataSourceSink{
					Component:     to.Ptr(armcosmos.DataTransferComponentAzureBlobStorage),
					ContainerName: to.Ptr("<container-name>"),
					EndpointURL:   to.Ptr("<endpoint-url>"),
				},
				Source: &armcosmos.CassandraDataTransferDataSourceSink{
					Component:    to.Ptr(armcosmos.DataTransferComponentCosmosDBCassandra),
					KeyspaceName: to.Ptr("<keyspace-name>"),
					TableName:    to.Ptr("<table-name>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
