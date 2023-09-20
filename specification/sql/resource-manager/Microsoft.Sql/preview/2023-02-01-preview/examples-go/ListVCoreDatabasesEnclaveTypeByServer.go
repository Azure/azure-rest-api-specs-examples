package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/ListVCoreDatabasesEnclaveTypeByServer.json
func ExampleDatabasesClient_NewListByServerPager_getsAListOfDatabasesConfiguredWithEnclaveType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByServerPager("Default-SQL-SouthEastAsia", "testsvr", &armsql.DatabasesClientListByServerOptions{SkipToken: nil})
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
		// page.DatabaseListResult = armsql.DatabaseListResult{
		// 	Value: []*armsql.Database{
		// 		{
		// 			Name: to.Ptr("testdb"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Kind: to.Ptr("v12.0,user,vcore"),
		// 			Properties: &armsql.DatabaseProperties{
		// 				CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
		// 				CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyZone),
		// 				CurrentServiceObjectiveName: to.Ptr("BC_Gen4_2"),
		// 				CurrentSKU: &armsql.SKU{
		// 					Name: to.Ptr("BC_Gen4"),
		// 					Capacity: to.Ptr[int32](2),
		// 					Tier: to.Ptr("BusinessCritical"),
		// 				},
		// 				DatabaseID: to.Ptr("6c764297-577b-470f-9af4-96d3d41e2ba3"),
		// 				DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:51:33.937Z"); return t}()),
		// 				IsInfraEncryptionEnabled: to.Ptr(false),
		// 				IsLedgerOn: to.Ptr(false),
		// 				LicenseType: to.Ptr(armsql.DatabaseLicenseTypeLicenseIncluded),
		// 				MaxLogSizeBytes: to.Ptr[int64](104857600),
		// 				MaxSizeBytes: to.Ptr[int64](268435456000),
		// 				PreferredEnclaveType: to.Ptr(armsql.AlwaysEncryptedEnclaveTypeDefault),
		// 				ReadScale: to.Ptr(armsql.DatabaseReadScaleEnabled),
		// 				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyZone),
		// 				RequestedServiceObjectiveName: to.Ptr("BC_Gen4_2"),
		// 				Status: to.Ptr(armsql.DatabaseStatusOnline),
		// 				ZoneRedundant: to.Ptr(false),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("BC_Gen4"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Tier: to.Ptr("BusinessCritical"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("master"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/master"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Kind: to.Ptr("v12.0,system"),
		// 			Properties: &armsql.DatabaseProperties{
		// 				CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
		// 				Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:23:42.537Z"); return t}()),
		// 				CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyLocal),
		// 				CurrentServiceObjectiveName: to.Ptr("System0"),
		// 				CurrentSKU: &armsql.SKU{
		// 					Name: to.Ptr("System0"),
		// 					Capacity: to.Ptr[int32](0),
		// 					Tier: to.Ptr("System"),
		// 				},
		// 				DatabaseID: to.Ptr("e6be351f-2cc9-4604-9e52-b0b28b2710b0"),
		// 				DefaultSecondaryLocation: to.Ptr("North Europe"),
		// 				IsInfraEncryptionEnabled: to.Ptr(false),
		// 				IsLedgerOn: to.Ptr(false),
		// 				MaxSizeBytes: to.Ptr[int64](32212254720),
		// 				PreferredEnclaveType: to.Ptr(armsql.AlwaysEncryptedEnclaveTypeVBS),
		// 				ReadScale: to.Ptr(armsql.DatabaseReadScaleDisabled),
		// 				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyLocal),
		// 				RequestedServiceObjectiveName: to.Ptr("System0"),
		// 				Status: to.Ptr(armsql.DatabaseStatusOnline),
		// 				ZoneRedundant: to.Ptr(false),
		// 			},
		// 			SKU: &armsql.SKU{
		// 				Name: to.Ptr("System0"),
		// 				Capacity: to.Ptr[int32](0),
		// 				Tier: to.Ptr("System"),
		// 			},
		// 	}},
		// }
	}
}
