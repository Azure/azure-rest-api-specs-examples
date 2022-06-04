Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv1.0.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseCreateRestoreExternalBackup.json
func ExampleManagedDatabasesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabasesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"Default-SQL-SouthEastAsia",
		"managedInstance",
		"managedDatabase",
		armsql.ManagedDatabase{
			Location: to.Ptr("southeastasia"),
			Properties: &armsql.ManagedDatabaseProperties{
				AutoCompleteRestore:      to.Ptr(true),
				Collation:                to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
				CreateMode:               to.Ptr(armsql.ManagedDatabaseCreateModeRestoreExternalBackup),
				LastBackupName:           to.Ptr("last_backup_name"),
				StorageContainerSasToken: to.Ptr("sv=2015-12-11&sr=c&sp=rl&sig=1234"),
				StorageContainerURI:      to.Ptr("https://myaccountname.blob.core.windows.net/backups"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
