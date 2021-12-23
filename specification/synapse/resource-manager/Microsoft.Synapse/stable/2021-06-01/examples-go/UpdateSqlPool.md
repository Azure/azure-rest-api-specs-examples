Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.1.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateSqlPool.json
func ExampleSQLPoolsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.SQLPoolPatchInfo{
			Location: to.StringPtr("<location>"),
			Properties: &armsynapse.SQLPoolResourceProperties{
				Collation:             to.StringPtr("<collation>"),
				CreateMode:            armsynapse.CreateModeDefault.ToPtr(),
				CreationDate:          to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t }()),
				MaxSizeBytes:          to.Int64Ptr(0),
				RecoverableDatabaseID: to.StringPtr("<recoverable-database-id>"),
				RestorePointInTime:    to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t }()),
				SourceDatabaseID:      to.StringPtr("<source-database-id>"),
			},
			SKU: &armsynapse.SKU{
				Name: to.StringPtr("<name>"),
				Tier: to.StringPtr("<tier>"),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SQLPool.ID: %s\n", *res.ID)
}
```
