Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysqlflexibleservers%2Fv1.0.0/sdk/resourcemanager/mysql/armmysqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

```go
package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerCreate.json
func ExampleServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmysqlflexibleservers.NewServersClient("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"testrg",
		"mysqltestserver",
		armmysqlflexibleservers.Server{
			Location: to.Ptr("southeastasia"),
			Tags: map[string]*string{
				"num": to.Ptr("1"),
			},
			Properties: &armmysqlflexibleservers.ServerProperties{
				AdministratorLogin:         to.Ptr("cloudsa"),
				AdministratorLoginPassword: to.Ptr("your_password"),
				AvailabilityZone:           to.Ptr("1"),
				Backup: &armmysqlflexibleservers.Backup{
					BackupRetentionDays: to.Ptr[int32](7),
					GeoRedundantBackup:  to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
				},
				CreateMode: to.Ptr(armmysqlflexibleservers.CreateModeDefault),
				HighAvailability: &armmysqlflexibleservers.HighAvailability{
					Mode:                    to.Ptr(armmysqlflexibleservers.HighAvailabilityModeZoneRedundant),
					StandbyAvailabilityZone: to.Ptr("3"),
				},
				Storage: &armmysqlflexibleservers.Storage{
					AutoGrow:      to.Ptr(armmysqlflexibleservers.EnableStatusEnumDisabled),
					Iops:          to.Ptr[int32](600),
					StorageSizeGB: to.Ptr[int32](100),
				},
				Version: to.Ptr(armmysqlflexibleservers.ServerVersionFive7),
			},
			SKU: &armmysqlflexibleservers.SKU{
				Name: to.Ptr("Standard_D2ds_v4"),
				Tier: to.Ptr(armmysqlflexibleservers.SKUTierGeneralPurpose),
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
