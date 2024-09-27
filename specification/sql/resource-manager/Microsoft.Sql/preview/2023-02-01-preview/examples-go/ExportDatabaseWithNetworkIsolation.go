package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/ExportDatabaseWithNetworkIsolation.json
func ExampleDatabasesClient_BeginExport_exportsADatabaseUsingPrivateLinkToCommunicateWithSqlServerAndStorageAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabasesClient().BeginExport(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", armsql.ExportDatabaseDefinition{
		AdministratorLogin:         to.Ptr("login"),
		AdministratorLoginPassword: to.Ptr("password"),
		AuthenticationType:         to.Ptr("Sql"),
		NetworkIsolation: &armsql.NetworkIsolationSettings{
			SQLServerResourceID:      to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr"),
			StorageAccountResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Storage/storageAccounts/test-privatelink"),
		},
		StorageKey:     to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=="),
		StorageKeyType: to.Ptr(armsql.StorageKeyTypeStorageAccessKey),
		StorageURI:     to.Ptr("https://test.blob.core.windows.net/test.bacpac"),
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
	// res.ImportExportOperationResult = armsql.ImportExportOperationResult{
	// 	Name: to.Ptr("9d9a794a-5cec-4f23-af70-d29511b522a4"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/importExportOperationResults"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/importExportOperationResults/9d9a794a-5cec-4f23-af70-d29511b522a4"),
	// 	Properties: &armsql.ImportExportOperationResultProperties{
	// 		BlobURI: to.Ptr("https://test.blob.core.windows.net/test.bacpac"),
	// 		DatabaseName: to.Ptr("testdb"),
	// 		LastModifiedTime: to.Ptr("2/2/2020 8:34:47 PM"),
	// 		QueuedTime: to.Ptr("2/2/2020 8:33:27 PM"),
	// 		RequestID: to.Ptr("9d9a794a-5cec-4f23-af70-d29511b522a4"),
	// 		RequestType: to.Ptr("Export"),
	// 		ServerName: to.Ptr("testsvr.database.windows.net"),
	// 		Status: to.Ptr("Completed"),
	// 	},
	// }
}
