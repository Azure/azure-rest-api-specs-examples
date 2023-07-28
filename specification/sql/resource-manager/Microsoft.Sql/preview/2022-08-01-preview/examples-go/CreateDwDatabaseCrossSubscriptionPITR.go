package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/CreateDwDatabaseCrossSubscriptionPITR.json
func ExampleDatabasesClient_BeginCreateOrUpdate_createsADataWarehouseDatabaseAsACrossSubscriptionRestoreFromARestorePointOfAnExistingDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdw", armsql.Database{
		Location: to.Ptr("southeastasia"),
		Properties: &armsql.DatabaseProperties{
			CreateMode:         to.Ptr(armsql.CreateModePointInTimeRestore),
			RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-22T05:35:31.503Z"); return t }()),
			SourceResourceID:   to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/srcsvr/databases/srcdw"),
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
	// 	Name: to.Ptr("testdw"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdw"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Kind: to.Ptr("v12.0,user,datawarehouse,gen2"),
	// 	Properties: &armsql.DatabaseProperties{
	// 		CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-24T06:46:14.99Z"); return t}()),
	// 		CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		CurrentServiceObjectiveName: to.Ptr("DW1000c"),
	// 		CurrentSKU: &armsql.SKU{
	// 			Name: to.Ptr("DataWarehouse"),
	// 			Capacity: to.Ptr[int32](9000),
	// 			Tier: to.Ptr("DataWarehouse"),
	// 		},
	// 		DatabaseID: to.Ptr("188784c9-d602-4684-86cf-e67b6f03551a"),
	// 		DefaultSecondaryLocation: to.Ptr("eastus"),
	// 		MaxSizeBytes: to.Ptr[int64](263882790666240),
	// 		ReadScale: to.Ptr(armsql.DatabaseReadScaleDisabled),
	// 		RequestedServiceObjectiveName: to.Ptr("DW1000c"),
	// 		Status: to.Ptr(armsql.DatabaseStatusOnline),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("DataWarehouse"),
	// 		Capacity: to.Ptr[int32](9000),
	// 		Tier: to.Ptr("DataWarehouse"),
	// 	},
	// }
}
