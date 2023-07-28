package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/PatchVCoreDatabase.json
func ExampleDatabasesClient_BeginUpdate_updatesADatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", armsql.DatabaseUpdate{
		Properties: &armsql.DatabaseUpdateProperties{
			LicenseType:  to.Ptr(armsql.DatabaseLicenseTypeLicenseIncluded),
			MaxSizeBytes: to.Ptr[int64](1073741824),
		},
		SKU: &armsql.SKU{
			Name: to.Ptr("BC_Gen4_4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Database = armsql.Database{
	// 	Name: to.Ptr("testdb"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Kind: to.Ptr("v12.0,user,vcore"),
	// 	Properties: &armsql.DatabaseProperties{
	// 		CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
	// 		CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		CurrentServiceObjectiveName: to.Ptr("BC_Gen4_2"),
	// 		CurrentSKU: &armsql.SKU{
	// 			Name: to.Ptr("BC_Gen4"),
	// 			Capacity: to.Ptr[int32](4),
	// 			Tier: to.Ptr("BusinessCritical"),
	// 		},
	// 		DatabaseID: to.Ptr("6c764297-577b-470f-9af4-96d3d41e2ba3"),
	// 		DefaultSecondaryLocation: to.Ptr("North Europe"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:51:33.937Z"); return t}()),
	// 		IsInfraEncryptionEnabled: to.Ptr(false),
	// 		IsLedgerOn: to.Ptr(false),
	// 		LicenseType: to.Ptr(armsql.DatabaseLicenseTypeLicenseIncluded),
	// 		MaxLogSizeBytes: to.Ptr[int64](104857600),
	// 		MaxSizeBytes: to.Ptr[int64](1073741824),
	// 		ReadScale: to.Ptr(armsql.DatabaseReadScaleEnabled),
	// 		RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyLocal),
	// 		RequestedServiceObjectiveName: to.Ptr("BC_Gen4_2"),
	// 		Status: to.Ptr(armsql.DatabaseStatusOnline),
	// 		ZoneRedundant: to.Ptr(false),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("BC_Gen4"),
	// 		Capacity: to.Ptr[int32](4),
	// 		Tier: to.Ptr("BusinessCritical"),
	// 	},
	// }
}
