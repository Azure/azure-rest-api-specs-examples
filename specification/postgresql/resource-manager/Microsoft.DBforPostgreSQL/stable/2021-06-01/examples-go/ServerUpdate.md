Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresqlflexibleservers%2Fv0.5.0/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ServerUpdate.json
func ExampleServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpostgresqlflexibleservers.NewServersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armpostgresqlflexibleservers.ServerForUpdate{
			Location: to.Ptr("<location>"),
			Properties: &armpostgresqlflexibleservers.ServerPropertiesForUpdate{
				AdministratorLoginPassword: to.Ptr("<administrator-login-password>"),
				Backup: &armpostgresqlflexibleservers.Backup{
					BackupRetentionDays: to.Ptr[int32](20),
				},
				CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeForUpdateUpdate),
				Storage: &armpostgresqlflexibleservers.Storage{
					StorageSizeGB: to.Ptr[int32](1024),
				},
			},
			SKU: &armpostgresqlflexibleservers.SKU{
				Name: to.Ptr("<name>"),
				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
			},
		},
		&armpostgresqlflexibleservers.ServersClientBeginUpdateOptions{ResumeToken: ""})
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
