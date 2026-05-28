package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/CreateDatabaseWithEncryptionProtectorAutoRotation.json
func ExampleDatabasesClient_BeginCreateOrUpdate_createsADatabaseWithEncryptionProtectorAutoRotation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", armsql.Database{
		Identity: &armsql.DatabaseIdentity{
			Type: to.Ptr(armsql.DatabaseIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armsql.DatabaseUserIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi": {},
			},
		},
		Location: to.Ptr("southeastasia"),
		Properties: &armsql.DatabaseProperties{
			Collation:                       to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
			CreateMode:                      to.Ptr(armsql.CreateModeDefault),
			EncryptionProtector:             to.Ptr("https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion"),
			EncryptionProtectorAutoRotation: to.Ptr(true),
			MaxSizeBytes:                    to.Ptr[int64](1073741824),
		},
		SKU: &armsql.SKU{
			Name: to.Ptr("S0"),
			Tier: to.Ptr("Standard"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.DatabasesClientCreateOrUpdateResponse{
	// 	Database: armsql.Database{
	// 		Name: to.Ptr("testdb"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/databases"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
	// 		Kind: to.Ptr("v12.0,user"),
	// 		Location: to.Ptr("southeastasia"),
	// 		Properties: &armsql.DatabaseProperties{
	// 			CatalogCollation: to.Ptr(armsql.CatalogCollationTypeSQLLatin1GeneralCP1CIAS),
	// 			Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:41:33.937Z"); return t}()),
	// 			CurrentBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 			CurrentServiceObjectiveName: to.Ptr("S0"),
	// 			CurrentSKU: &armsql.SKU{
	// 				Name: to.Ptr("Standard"),
	// 				Capacity: to.Ptr[int32](10),
	// 				Tier: to.Ptr("Standard"),
	// 			},
	// 			DatabaseID: to.Ptr("6c764297-577b-470f-9af4-96d3d41e2ba3"),
	// 			DefaultSecondaryLocation: to.Ptr("North Europe"),
	// 			EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-07T04:51:33.937Z"); return t}()),
	// 			EncryptionProtector: to.Ptr("https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion"),
	// 			EncryptionProtectorAutoRotation: to.Ptr(true),
	// 			IsInfraEncryptionEnabled: to.Ptr(false),
	// 			IsLedgerOn: to.Ptr(false),
	// 			MaxSizeBytes: to.Ptr[int64](1073741824),
	// 			ReadScale: to.Ptr(armsql.DatabaseReadScaleDisabled),
	// 			RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 			RequestedServiceObjectiveName: to.Ptr("S0"),
	// 			Status: to.Ptr(armsql.DatabaseStatusOnline),
	// 			ZoneRedundant: to.Ptr(false),
	// 		},
	// 		SKU: &armsql.SKU{
	// 			Name: to.Ptr("S0"),
	// 			Capacity: to.Ptr[int32](10),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 	},
	// }
}
