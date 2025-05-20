package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/BackupCreate.json
func ExampleBackupsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBackupsClient().BeginCreate(ctx, "TestGroup", "postgresqltestserver", "backup_20250303T160516", nil)
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
	// res.ServerBackup = armpostgresqlflexibleservers.ServerBackup{
	// 	Name: to.Ptr("backup_20250303T160516"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/postgresqltestserver/backups/backup_20250303T160516"),
	// 	Properties: &armpostgresqlflexibleservers.ServerBackupProperties{
	// 		BackupType: to.Ptr(armpostgresqlflexibleservers.OriginCustomerOnDemand),
	// 		CompletedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-03T16:05:19.902Z"); return t}()),
	// 		Source: to.Ptr("Customer Initiated"),
	// 	},
	// }
}
