Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv0.5.0/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.

```go
package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBManagedCassandraDataCenterCreate.json
func ExampleCassandraDataCentersClient_BeginCreateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcosmos.NewCassandraDataCentersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<data-center-name>",
		armcosmos.DataCenterResource{
			Properties: &armcosmos.DataCenterResourceProperties{
				Base64EncodedCassandraYamlFragment: to.Ptr("<base64encoded-cassandra-yaml-fragment>"),
				DataCenterLocation:                 to.Ptr("<data-center-location>"),
				DelegatedSubnetID:                  to.Ptr("<delegated-subnet-id>"),
				NodeCount:                          to.Ptr[int32](9),
			},
		},
		&armcosmos.CassandraDataCentersClientBeginCreateUpdateOptions{ResumeToken: ""})
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
