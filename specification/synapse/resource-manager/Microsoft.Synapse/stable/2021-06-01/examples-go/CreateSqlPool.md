Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPool.json
func ExampleSQLPoolsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.SQLPool{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armsynapse.SQLPoolResourceProperties{
				Collation:             to.StringPtr("<collation>"),
				CreateMode:            armsynapse.CreateMode("").ToPtr(),
				MaxSizeBytes:          to.Int64Ptr(0),
				RecoverableDatabaseID: to.StringPtr("<recoverable-database-id>"),
				SourceDatabaseID:      to.StringPtr("<source-database-id>"),
				StorageAccountType:    armsynapse.StorageAccountType("LRS").ToPtr(),
			},
			SKU: &armsynapse.SKU{
				Name: to.StringPtr("<name>"),
				Tier: to.StringPtr("<tier>"),
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
	log.Printf("Response result: %#v\n", res.SQLPoolsClientCreateResult)
}
```
