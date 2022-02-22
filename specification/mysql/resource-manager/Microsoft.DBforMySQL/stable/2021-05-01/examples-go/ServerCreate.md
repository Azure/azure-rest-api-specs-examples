Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysqlflexibleservers%2Fv0.3.1/sdk/resourcemanager/mysql/armmysqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

```go
package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
)

// x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerCreate.json
func ExampleServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmysqlflexibleservers.NewServersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmysqlflexibleservers.Server{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"num": to.StringPtr("1"),
			},
			Properties: &armmysqlflexibleservers.ServerProperties{
				AdministratorLogin:         to.StringPtr("<administrator-login>"),
				AdministratorLoginPassword: to.StringPtr("<administrator-login-password>"),
				AvailabilityZone:           to.StringPtr("<availability-zone>"),
				Backup: &armmysqlflexibleservers.Backup{
					BackupRetentionDays: to.Int32Ptr(7),
					GeoRedundantBackup:  armmysqlflexibleservers.EnableStatusEnum("Disabled").ToPtr(),
				},
				CreateMode: armmysqlflexibleservers.CreateMode("Default").ToPtr(),
				HighAvailability: &armmysqlflexibleservers.HighAvailability{
					Mode:                    armmysqlflexibleservers.HighAvailabilityMode("ZoneRedundant").ToPtr(),
					StandbyAvailabilityZone: to.StringPtr("<standby-availability-zone>"),
				},
				Storage: &armmysqlflexibleservers.Storage{
					AutoGrow:      armmysqlflexibleservers.EnableStatusEnum("Disabled").ToPtr(),
					Iops:          to.Int32Ptr(600),
					StorageSizeGB: to.Int32Ptr(100),
				},
				Version: armmysqlflexibleservers.ServerVersion("5.7").ToPtr(),
			},
			SKU: &armmysqlflexibleservers.SKU{
				Name: to.StringPtr("<name>"),
				Tier: armmysqlflexibleservers.SKUTier("GeneralPurpose").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ServersClientCreateResult)
}
```
