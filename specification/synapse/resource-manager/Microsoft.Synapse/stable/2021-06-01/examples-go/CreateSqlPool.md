Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.4.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPool.json
func ExampleSQLPoolsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.SQLPool{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armsynapse.SQLPoolResourceProperties{
				Collation:             to.Ptr("<collation>"),
				CreateMode:            to.Ptr(armsynapse.CreateMode("")),
				MaxSizeBytes:          to.Ptr[int64](0),
				RecoverableDatabaseID: to.Ptr("<recoverable-database-id>"),
				SourceDatabaseID:      to.Ptr("<source-database-id>"),
				StorageAccountType:    to.Ptr(armsynapse.StorageAccountTypeLRS),
			},
			SKU: &armsynapse.SKU{
				Name: to.Ptr("<name>"),
				Tier: to.Ptr("<tier>"),
			},
		},
		&armsynapse.SQLPoolsClientBeginCreateOptions{ResumeToken: ""})
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
