Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresqlflexibleservers%2Fv0.3.1/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers"
)

// x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ServerUpdate.json
func ExampleServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpostgresqlflexibleservers.NewServersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armpostgresqlflexibleservers.ServerForUpdate{
			Location: to.StringPtr("<location>"),
			Properties: &armpostgresqlflexibleservers.ServerPropertiesForUpdate{
				AdministratorLoginPassword: to.StringPtr("<administrator-login-password>"),
				Backup: &armpostgresqlflexibleservers.Backup{
					BackupRetentionDays: to.Int32Ptr(20),
				},
				CreateMode: armpostgresqlflexibleservers.CreateModeForUpdate("Update").ToPtr(),
				Storage: &armpostgresqlflexibleservers.Storage{
					StorageSizeGB: to.Int32Ptr(1024),
				},
			},
			SKU: &armpostgresqlflexibleservers.SKU{
				Name: to.StringPtr("<name>"),
				Tier: armpostgresqlflexibleservers.SKUTier("GeneralPurpose").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ServersClientUpdateResult)
}
```
