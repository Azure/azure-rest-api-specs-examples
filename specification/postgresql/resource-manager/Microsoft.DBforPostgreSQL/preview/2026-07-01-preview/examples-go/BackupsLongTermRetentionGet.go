package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-07-01-preview/BackupsLongTermRetentionGet.json
func ExampleBackupsLongTermRetentionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackupsLongTermRetentionClient().Get(ctx, "exampleresourcegroup", "exampleserver", "exampleltrbackup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpostgresqlflexibleservers.BackupsLongTermRetentionClientGetResponse{
	// 	BackupsLongTermRetentionOperation: armpostgresqlflexibleservers.BackupsLongTermRetentionOperation{
	// 		Name: to.Ptr("exampleltrbackup"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/ltrbackupOperations"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver"),
	// 		Properties: &armpostgresqlflexibleservers.LtrBackupOperationResponseProperties{
	// 			BackupMetadata: to.Ptr("backupMetadata"),
	// 			BackupName: to.Ptr("exampleltrbackup"),
	// 			DataTransferredInBytes: to.Ptr[int64](9),
	// 			DatasourceSizeInBytes: to.Ptr[int64](21),
	// 			EndTime: to.Ptr(time.Date(2025, time.June, 1, 18, 35, 22, 123000000, time.UTC)),
	// 			PercentComplete: to.Ptr[float64](4),
	// 			StartTime: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123000000, time.UTC)),
	// 			Status: to.Ptr(armpostgresqlflexibleservers.ExecutionStatusRunning),
	// 		},
	// 	},
	// }
}
