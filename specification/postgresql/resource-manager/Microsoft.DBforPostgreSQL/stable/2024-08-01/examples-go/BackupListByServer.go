package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/BackupListByServer.json
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
		// 			Name: to.Ptr("daily_20210615T160516"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210615T160516"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-15T16:05:19.902Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210616T160520"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210616T160520"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T16:05:23.924Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210617T160525"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210617T160525"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-17T16:05:28.124Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210618T160529"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210618T160529"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-18T16:05:32.273Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210619T160533"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210619T160533"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-19T16:05:36.860Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210620T160538"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210620T160538"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-20T16:05:41.920Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210621T160543"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210621T160543"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T16:05:48.852Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210622T160803"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210622T160803"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-22T16:08:06.312Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210622T210807"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210622T210807"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-22T21:08:10.505Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210623T212413"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210623T212413"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-23T21:24:16.940Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("daily_20210624T061328"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/daily_20210624T061328"),
		// 			Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
		// 				BackupType: to.Ptr(armpostgresqlflexibleservers.OriginFull),
		// 				CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-24T06:13:31.496Z"); return t}()),
		// 				Source: to.Ptr("Automatic"),
		// 			},
		// 	}},
		// }
	}
}
