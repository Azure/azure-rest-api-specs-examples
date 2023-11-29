package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/CreateDatabasePITRMode.json
func ExampleDatabasesClient_BeginCreateOrUpdate_createsADatabaseFromPointInTimeRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "dbpitr", armsql.Database{
		Location: to.Ptr("southeastasia"),
		Properties: &armsql.DatabaseProperties{
			CreateMode:         to.Ptr(armsql.CreateModePointInTimeRestore),
			RestorePointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-22T05:35:31.503Z"); return t }()),
			SourceDatabaseID:   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SoutheastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
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
	// 	Name: to.Ptr("dbpitr"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SoutheastAsia/providers/Microsoft.Sql/servers/testsvr/databases/dbpitr"),
	// 	Location: to.Ptr("SoutheastAsia"),
	// 	Kind: to.Ptr("v12.0,user"),
	// 	Properties: &armsql.DatabaseProperties{
	// 		CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
	// 		Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-23T20:31:35.807Z"); return t}()),
	// 		CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		CurrentServiceObjectiveName: to.Ptr("S0"),
	// 		CurrentSKU: &armsql.SKU{
	// 			Name: to.Ptr("Standard"),
	// 			Capacity: to.Ptr[int32](10),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		DatabaseID: to.Ptr("2463ef48-53e2-4c26-9781-2d7138080bfd"),
	// 		DefaultSecondaryLocation: to.Ptr("SoutheastAsia"),
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-10-23T20:47:37.000Z"); return t}()),
	// 		IsInfraEncryptionEnabled: to.Ptr(false),
	// 		IsLedgerOn: to.Ptr(false),
	// 		MaxSizeBytes: to.Ptr[int64](268435456000),
	// 		ReadScale: to.Ptr(armsql.DatabaseReadScaleDisabled),
	// 		RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		RequestedServiceObjectiveName: to.Ptr("S0"),
	// 		Status: to.Ptr(armsql.DatabaseStatusOnline),
	// 		ZoneRedundant: to.Ptr(false),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("Standard"),
	// 		Capacity: to.Ptr[int32](10),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
