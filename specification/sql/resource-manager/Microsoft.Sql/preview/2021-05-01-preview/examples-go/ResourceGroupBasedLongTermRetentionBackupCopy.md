Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.5.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupCopy.json
func ExampleLongTermRetentionBackupsClient_BeginCopyByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCopyByResourceGroup(ctx,
		"<resource-group-name>",
		"<location-name>",
		"<long-term-retention-server-name>",
		"<long-term-retention-database-name>",
		"<backup-name>",
		armsql.CopyLongTermRetentionBackupParameters{
			Properties: &armsql.CopyLongTermRetentionBackupParametersProperties{
				TargetBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
				TargetDatabaseName:            to.Ptr("<target-database-name>"),
				TargetServerResourceID:        to.Ptr("<target-server-resource-id>"),
			},
		},
		&armsql.LongTermRetentionBackupsClientBeginCopyByResourceGroupOptions{ResumeToken: ""})
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
