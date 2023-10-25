package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/BackupGet.json
func ExampleBackupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupsClient().Get(ctx, "TestGroup", "postgresqltestserver", "daily_20210615T160516", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerBackup = armpostgresqlflexibleservers.ServerBackup{
	// 	Name: to.Ptr("daily_20210615T160516"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210615T160516"),
	// 	Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
	// 		BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
	// 		CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T16:05:19.902522+00:00"); return t}()),
	// 		Source: to.Ptr("Automatic"),
	// 	},
	// }
}
