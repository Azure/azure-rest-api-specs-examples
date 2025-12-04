package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/BackupsLongTermRetentionStart.json
func ExampleBackupsLongTermRetentionClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupsLongTermRetentionClient().BeginStart(ctx, "exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.BackupsLongTermRetentionRequest{
		BackupSettings: &armpostgresqlflexibleservers.BackupSettings{
			BackupName: to.Ptr("exampleltrbackup"),
		},
		TargetDetails: &armpostgresqlflexibleservers.BackupStoreDetails{
			SasURIList: []*string{
				to.Ptr("sasuri")},
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
	// res.BackupsLongTermRetentionResponse = armpostgresqlflexibleservers.BackupsLongTermRetentionResponse{
	// 	Properties: &armpostgresqlflexibleservers.LtrBackupOperationResponseProperties{
	// 		BackupMetadata: to.Ptr("backupmetadata"),
	// 		DataTransferredInBytes: to.Ptr[int64](23),
	// 		DatasourceSizeInBytes: to.Ptr[int64](23),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:35:22.123Z"); return t}()),
	// 		PercentComplete: to.Ptr[float64](100),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T18:30:22.123Z"); return t}()),
	// 		Status: to.Ptr(armpostgresqlflexibleservers.ExecutionStatusRunning),
	// 	},
	// }
}
