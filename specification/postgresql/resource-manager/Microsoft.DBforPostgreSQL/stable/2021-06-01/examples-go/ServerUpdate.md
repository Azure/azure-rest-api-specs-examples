Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresqlflexibleservers%2Fv1.0.0/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ServerUpdate.json
func ExampleServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpostgresqlflexibleservers.NewServersClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"TestGroup",
		"pgtestsvc4",
		armpostgresqlflexibleservers.ServerForUpdate{
			Location: to.Ptr("westus"),
			Properties: &armpostgresqlflexibleservers.ServerPropertiesForUpdate{
				AdministratorLoginPassword: to.Ptr("newpassword"),
				Backup: &armpostgresqlflexibleservers.Backup{
					BackupRetentionDays: to.Ptr[int32](20),
				},
				CreateMode: to.Ptr(armpostgresqlflexibleservers.CreateModeForUpdateUpdate),
				Storage: &armpostgresqlflexibleservers.Storage{
					StorageSizeGB: to.Ptr[int32](1024),
				},
			},
			SKU: &armpostgresqlflexibleservers.SKU{
				Name: to.Ptr("Standard_D8s_v3"),
				Tier: to.Ptr(armpostgresqlflexibleservers.SKUTierGeneralPurpose),
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
