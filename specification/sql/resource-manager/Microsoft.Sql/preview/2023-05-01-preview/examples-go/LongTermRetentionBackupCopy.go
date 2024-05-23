package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/LongTermRetentionBackupCopy.json
func ExampleLongTermRetentionBackupsClient_BeginCopy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionBackupsClient().BeginCopy(ctx, "japaneast", "testserver", "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000", armsql.CopyLongTermRetentionBackupParameters{
		Properties: &armsql.CopyLongTermRetentionBackupParametersProperties{
			TargetBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
			TargetDatabaseName:            to.Ptr("testDatabase2"),
			TargetServerResourceID:        to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver2"),
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
	// res.LongTermRetentionBackupOperationResult = armsql.LongTermRetentionBackupOperationResult{
	// 	Name: to.Ptr("a1aa7c77-961b-4fbb-bcd6-aa9acfcd1706"),
	// 	Type: to.Ptr("Microsoft.Sql/locations/longTermRetentionBackupOperationResults"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japanEast/longTermRetentionBackupOperationResults/a1aa7c77-961b-4fbb-bcd6-aa9acfcd1706"),
	// 	Properties: &armsql.LongTermRetentionOperationResultProperties{
	// 		FromBackupResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japanEast/longTermRetentionServers/testserver/longterRetentionDatabases/testDatabase/longTermRetentionBackups/55555555-6666-7777-8888-999999999999;131637960820000000"),
	// 		OperationType: to.Ptr("CopyBackup"),
	// 		RequestID: to.Ptr("a1aa7c77-961b-4fbb-bcd6-aa9acfcd1706"),
	// 		Status: to.Ptr("Succeeded"),
	// 		TargetBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
	// 		ToBackupResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup/providers/Microsoft.Sql/locations/japanEast/longTermRetentionServers/testserver2/longterRetentionDatabases/testDatabase2/longTermRetentionBackups/55555555-6666-7777-8888-111111111111;131637960820000000"),
	// 	},
	// }
}
