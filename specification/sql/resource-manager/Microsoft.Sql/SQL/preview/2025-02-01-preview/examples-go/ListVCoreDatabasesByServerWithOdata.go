package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ListVCoreDatabasesByServerWithOdata.json
func ExampleDatabasesClient_NewListByServerPager_getsAListOfDatabasesWithODataFiltering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByServerPager("Default-SQL-SouthEastAsia", "testsvr", nil)
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
		// page = armsql.DatabasesClientListByServerResponse{
		// 	DatabaseListResult: armsql.DatabaseListResult{
		// 		Value: []*armsql.Database{
		// 			{
		// 				SKU: &armsql.SKU{
		// 					Name: to.Ptr("BC_Gen5"),
		// 					Tier: to.Ptr("BusinessCritical"),
		// 					Capacity: to.Ptr[int32](2),
		// 				},
		// 				Kind: to.Ptr("v12.0,user,vcore"),
		// 				Properties: &armsql.DatabaseProperties{
		// 					Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 					MaxSizeBytes: to.Ptr[int64](268435456000),
		// 					Status: to.Ptr(armsql.DatabaseStatusOnline),
		// 					DatabaseID: to.Ptr("6c764297-577b-470f-9af4-96d3d41e2ba3"),
		// 					CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
		// 					CurrentServiceObjectiveName: to.Ptr("BC_Gen5_2"),
		// 					RequestedServiceObjectiveName: to.Ptr("BC_Gen5_2"),
		// 					DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 					CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
		// 					LicenseType: to.Ptr(armsql.DatabaseLicenseTypeLicenseIncluded),
		// 					MaxLogSizeBytes: to.Ptr[int64](104857600),
		// 					IsInfraEncryptionEnabled: to.Ptr(false),
		// 					ZoneRedundant: to.Ptr(false),
		// 					ReadScale: to.Ptr(armsql.DatabaseReadScaleEnabled),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:51:33.937Z"); return t}()),
		// 					MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_SouthEastAsia_1"),
		// 					CurrentSKU: &armsql.SKU{
		// 						Name: to.Ptr("BC_Gen5"),
		// 						Tier: to.Ptr("BusinessCritical"),
		// 						Capacity: to.Ptr[int32](2),
		// 					},
		// 					CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					IsLedgerOn: to.Ptr(false),
		// 				},
		// 				Location: to.Ptr("southeastasia"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/db2"),
		// 				Name: to.Ptr("db2"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/databases"),
		// 			},
		// 			{
		// 				SKU: &armsql.SKU{
		// 					Name: to.Ptr("BC_Gen5"),
		// 					Tier: to.Ptr("BusinessCritical"),
		// 					Capacity: to.Ptr[int32](2),
		// 				},
		// 				Kind: to.Ptr("v12.0,user,vcore"),
		// 				Properties: &armsql.DatabaseProperties{
		// 					Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 					MaxSizeBytes: to.Ptr[int64](268435456000),
		// 					Status: to.Ptr(armsql.DatabaseStatusOnline),
		// 					DatabaseID: to.Ptr("6c764297-577b-470f-9af4-96d3d41e2ba3"),
		// 					CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
		// 					CurrentServiceObjectiveName: to.Ptr("BC_Gen5_2"),
		// 					RequestedServiceObjectiveName: to.Ptr("BC_Gen5_2"),
		// 					DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 					CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
		// 					LicenseType: to.Ptr(armsql.DatabaseLicenseTypeLicenseIncluded),
		// 					MaxLogSizeBytes: to.Ptr[int64](104857600),
		// 					IsInfraEncryptionEnabled: to.Ptr(false),
		// 					ZoneRedundant: to.Ptr(false),
		// 					ReadScale: to.Ptr(armsql.DatabaseReadScaleEnabled),
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:51:33.937Z"); return t}()),
		// 					MaintenanceConfigurationID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_SouthEastAsia_1"),
		// 					CurrentSKU: &armsql.SKU{
		// 						Name: to.Ptr("BC_Gen5"),
		// 						Tier: to.Ptr("BusinessCritical"),
		// 						Capacity: to.Ptr[int32](2),
		// 					},
		// 					CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
		// 					IsLedgerOn: to.Ptr(false),
		// 				},
		// 				Location: to.Ptr("southeastasia"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/db1"),
		// 				Name: to.Ptr("db1"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/databases"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
