Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysqlflexibleservers%2Fv0.6.0/sdk/resourcemanager/mysql/armmysqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerCreate.json
func ExampleServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmysqlflexibleservers.NewServersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<server-name>",
		armmysqlflexibleservers.Server{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"num": to.Ptr("1"),
			},
			Properties: &armmysqlflexibleservers.ServerProperties{
				AdministratorLogin:         to.Ptr("<administrator-login>"),
				AdministratorLoginPassword: to.Ptr("<administrator-login-password>"),
				AvailabilityZone:           to.Ptr("<availability-zone>"),
				Backup: &armmysqlflexibleservers.Backup{
					BackupRetentionDays: to.Ptr[int32](7),
					GeoRedundantBackup:  to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
				},
				CreateMode: to.Ptr(armmysqlflexibleservers.CreateModeDefault),
				HighAvailability: &armmysqlflexibleservers.HighAvailability{
					Mode:                    to.Ptr(armmysqlflexibleservers.HighAvailabilityModeZoneRedundant),
					StandbyAvailabilityZone: to.Ptr("<standby-availability-zone>"),
				},
				Storage: &armmysqlflexibleservers.Storage{
					AutoGrow:      to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
					Iops:          to.Ptr[int32](600),
					StorageSizeGB: to.Ptr[int32](100),
				},
				Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
			},
			SKU: &armmysqlflexibleservers.SKU{
				Name: to.Ptr("<name>"),
				Tier: to.Ptr(armmysqlflexibleservers.SKUTierGeneralPurpose),
			},
		},
		&armmysqlflexibleservers.ServersClientBeginCreateOptions{ResumeToken: ""})
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
