Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupCopy.json
func ExampleLongTermRetentionBackupsClient_BeginCopyByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewLongTermRetentionBackupsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCopyByResourceGroup(ctx,
		"<resource-group-name>",
		"<location-name>",
		"<long-term-retention-server-name>",
		"<long-term-retention-database-name>",
		"<backup-name>",
		armsql.CopyLongTermRetentionBackupParameters{
			Properties: &armsql.CopyLongTermRetentionBackupParametersProperties{
				TargetBackupStorageRedundancy: armsql.BackupStorageRedundancy("Geo").ToPtr(),
				TargetDatabaseName:            to.StringPtr("<target-database-name>"),
				TargetServerResourceID:        to.StringPtr("<target-server-resource-id>"),
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
	log.Printf("Response result: %#v\n", res.LongTermRetentionBackupsClientCopyByResourceGroupResult)
}
```
