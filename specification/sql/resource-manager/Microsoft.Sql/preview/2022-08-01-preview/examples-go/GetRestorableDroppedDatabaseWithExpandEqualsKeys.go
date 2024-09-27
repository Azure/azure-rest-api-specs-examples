package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/GetRestorableDroppedDatabaseWithExpandEqualsKeys.json
func ExampleRestorableDroppedDatabasesClient_Get_getsARestorableDroppedDatabaseWithExpandEqualsKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorableDroppedDatabasesClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000", &armsql.RestorableDroppedDatabasesClientGetOptions{Expand: to.Ptr("keys"),
		Filter: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorableDroppedDatabase = armsql.RestorableDroppedDatabase{
	// 	Name: to.Ptr("testdb,131403269876900000"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/restorableDroppedDatabases"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/restorableDroppedDatabases/testdb"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Properties: &armsql.RestorableDroppedDatabaseProperties{
	// 		BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-07T04:41:33.937Z"); return t}()),
	// 		DatabaseName: to.Ptr("testdb"),
	// 		DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T02:49:47.690Z"); return t}()),
	// 		Keys: map[string]*armsql.DatabaseKey{
	// 			"https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion": &armsql.DatabaseKey{
	// 			},
	// 			"https://your-key-vault-name.vault.azure.net/yourKey2/yourKey2Version": &armsql.DatabaseKey{
	// 			},
	// 		},
	// 		MaxSizeBytes: to.Ptr[int64](268435456000),
	// 	},
	// 	SKU: &armsql.SKU{
	// 		Name: to.Ptr("BC_Gen4_2"),
	// 		Tier: to.Ptr("BusinessCritical"),
	// 	},
	// }
}
