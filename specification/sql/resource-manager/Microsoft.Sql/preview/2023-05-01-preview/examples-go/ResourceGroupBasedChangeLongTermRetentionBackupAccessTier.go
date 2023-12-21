package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ResourceGroupBasedChangeLongTermRetentionBackupAccessTier.json
func ExampleLongTermRetentionBackupsClient_BeginChangeAccessTierByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionBackupsClient().BeginChangeAccessTierByResourceGroup(ctx, "resourceGroupName", "japaneast", "serverName", "databaseName", "55555555-6666-7777-8888-999999999999;131637960820000000;Archive", armsql.ChangeLongTermRetentionBackupAccessTierParameters{
		BackupStorageAccessTier: to.Ptr("Hot"),
		OperationMode:           to.Ptr("Copy"),
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
	// res.LongTermRetentionBackup = armsql.LongTermRetentionBackup{
	// 	Name: to.Ptr("55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionServers/longTermRetentionDatabases/longTermRetentionBackups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.Sql/locations/japaneast/longTermRetentionServers/serverName/longTermRetentionDatabases/databaseName/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000;Archive"),
	// 	Properties: &armsql.LongTermRetentionBackupProperties{
	// 		BackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		BackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00.000Z"); return t}()),
	// 		DatabaseName: to.Ptr("databaseName"),
	// 		ServerCreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-06T08:00:00.000Z"); return t}()),
	// 		ServerName: to.Ptr("serverName"),
	// 	},
	// }
}
