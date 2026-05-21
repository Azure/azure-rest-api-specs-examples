package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/SqlMiGetDatabaseMigration.json
func ExampleDatabaseMigrationsSQLMiClient_Get_getSqlMiDatabaseMigrationWithoutTheExpandParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseMigrationsSQLMiClient().Get(ctx, "testrg", "managedInstance1", "db1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatamigration.DatabaseMigrationsSQLMiClientGetResponse{
	// 	DatabaseMigrationSQLMi: armdatamigration.DatabaseMigrationSQLMi{
	// 		Name: to.Ptr("db1"),
	// 		Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname/providers/Microsoft.DataMigration/databaseMigrations/db1"),
	// 		Properties: &armdatamigration.DatabaseMigrationPropertiesSQLMi{
	// 			BackupConfiguration: &armdatamigration.BackupConfiguration{
	// 				SourceLocation: &armdatamigration.SourceLocation{
	// 					FileStorageType: to.Ptr("FileShare"),
	// 				},
	// 			},
	// 			EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
	// 			Kind: to.Ptr(armdatamigration.ResourceTypeSQLMi),
	// 			MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
	// 			MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
	// 			MigrationStatus: to.Ptr("InProgress"),
	// 			ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 			Scope: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/managedInstance1"),
	// 			SourceDatabaseName: to.Ptr("sourcename"),
	// 			SourceServerName: to.Ptr("sourceserver"),
	// 			SQLServerInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/instanceName"),
	// 			StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
	// 		},
	// 	},
	// }
}
