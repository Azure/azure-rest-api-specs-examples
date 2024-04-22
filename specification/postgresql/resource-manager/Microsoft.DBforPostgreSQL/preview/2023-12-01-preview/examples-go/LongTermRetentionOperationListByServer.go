package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/LongTermRetentionOperationListByServer.json
func ExampleLtrBackupOperationsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLtrBackupOperationsClient().NewListByServerPager("rgLongTermRetention", "pgsqlltrtestserver", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.LtrServerBackupOperationList = armpostgresqlflexibleservers.LtrServerBackupOperationList{
		// 	Value: []*armpostgresqlflexibleservers.LtrServerBackupOperation{
		// 		{
		// 			Name: to.Ptr("backup1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/ltrbackupOperations"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/rgLongTermRetention/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgsqlltrtestserver"),
		// 			Properties: &armpostgresqlflexibleservers.LtrBackupOperationResponseProperties{
		// 				BackupMetadata: to.Ptr("backupMetadata"),
		// 				BackupName: to.Ptr("backup1"),
		// 				DataTransferredInBytes: to.Ptr[int64](9),
		// 				DatasourceSizeInBytes: to.Ptr[int64](21),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-13T01:32:30.128Z"); return t}()),
		// 				PercentComplete: to.Ptr[float64](4),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-13T01:32:30.128Z"); return t}()),
		// 				Status: to.Ptr(armpostgresqlflexibleservers.ExecutionStatusRunning),
		// 			},
		// 	}},
		// }
	}
}
