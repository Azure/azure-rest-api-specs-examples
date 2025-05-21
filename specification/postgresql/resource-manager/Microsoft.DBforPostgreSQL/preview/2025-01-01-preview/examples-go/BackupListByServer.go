package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/BackupListByServer.json
func ExampleBackupsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackupsClient().NewListByServerPager("TestGroup", "postgresqltestserver", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ServerBackupListResult = armpostgresqlflexibleservers.ServerBackupListResult{
		// 	Value: []*armpostgresqlflexibleservers.ServerBackup{
		// 		{
		// 			Name: to.Ptr("backup_638766209959406043"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/backup_638766209959406043"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T17:49:56.940Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ondemandbackup-03032025-1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/ondemandbackup-03032025-1"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginCustomerOnDemand),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T20:53:22.446Z"); return t}()),
		// 				Source: to.Ptr("Customer Initiated"),
		// 			},
		// 	}},
		// }
	}
}
