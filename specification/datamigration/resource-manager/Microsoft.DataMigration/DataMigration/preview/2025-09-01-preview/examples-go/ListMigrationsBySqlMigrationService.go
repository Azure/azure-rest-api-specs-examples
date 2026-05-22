package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/ListMigrationsBySqlMigrationService.json
func ExampleSQLMigrationServicesClient_NewListMigrationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLMigrationServicesClient().NewListMigrationsPager("testrg", "service1", nil)
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
		// page = armdatamigration.SQLMigrationServicesClientListMigrationsResponse{
		// 	DatabaseMigrationListResult: armdatamigration.DatabaseMigrationListResult{
		// 		Value: []*armdatamigration.DatabaseMigration{
		// 			{
		// 				Name: to.Ptr("targetdb"),
		// 				Type: to.Ptr("Microsoft.DataMigration/dataMigrations"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname/providers/Microsoft.DataMigration/databaseMigrations/targetdb"),
		// 				Properties: &armdatamigration.DatabaseMigrationPropertiesSQLMi{
		// 					EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
		// 					Kind: to.Ptr(armdatamigration.ResourceTypeSQLMi),
		// 					MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
		// 					MigrationStatus: to.Ptr("InProgress"),
		// 					MigrationStatusDetails: &armdatamigration.MigrationStatusDetails{
		// 						ActiveBackupSets: []*armdatamigration.SQLBackupSetInfo{
		// 						},
		// 						IsFullBackupRestored: to.Ptr(false),
		// 						MigrationState: to.Ptr("WaitForFullBackupUploadOperation"),
		// 					},
		// 					Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname"),
		// 					SourceDatabaseName: to.Ptr("sourcename"),
		// 					SQLServerInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/instanceName"),
		// 					StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("targetdb"),
		// 				Type: to.Ptr("Microsoft.DataMigration/dataMigrations"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname/providers/Microsoft.DataMigration/databaseMigrations/targetdb"),
		// 				Properties: &armdatamigration.DatabaseMigrationPropertiesSQLMi{
		// 					EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
		// 					Kind: to.Ptr(armdatamigration.ResourceTypeSQLMi),
		// 					MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
		// 					MigrationStatus: to.Ptr("InProgress"),
		// 					MigrationStatusDetails: &armdatamigration.MigrationStatusDetails{
		// 						ActiveBackupSets: []*armdatamigration.SQLBackupSetInfo{
		// 						},
		// 						IsFullBackupRestored: to.Ptr(false),
		// 						MigrationState: to.Ptr("WaitForFullBackupUploadOperation"),
		// 					},
		// 					Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname"),
		// 					SourceDatabaseName: to.Ptr("sourcename"),
		// 					SQLServerInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/instanceName"),
		// 					StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
