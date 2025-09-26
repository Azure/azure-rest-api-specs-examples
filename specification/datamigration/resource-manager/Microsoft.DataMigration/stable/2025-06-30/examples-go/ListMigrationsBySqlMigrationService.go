package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/930e8030f5058d947fea4e2640725baab8a4561a/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/ListMigrationsBySqlMigrationService.json
func ExampleSQLMigrationServicesClient_NewListMigrationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.DatabaseMigrationListResult = armdatamigration.DatabaseMigrationListResult{
		// 	Value: []*armdatamigration.DatabaseMigration{
		// 		{
		// 			Name: to.Ptr("targetdb"),
		// 			Type: to.Ptr("Microsoft.DataMigration/dataMigrations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname/providers/Microsoft.DataMigration/databaseMigrations/targetdb"),
		// 			Properties: &armdatamigration.DatabaseMigrationPropertiesSQLMi{
		// 				EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				Kind: to.Ptr(armdatamigration.ResourceTypeSQLMi),
		// 				MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
		// 				MigrationStatus: to.Ptr("InProgress"),
		// 				Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname"),
		// 				StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				SourceDatabaseName: to.Ptr("sourcename"),
		// 				MigrationStatusDetails: &armdatamigration.MigrationStatusDetails{
		// 					ActiveBackupSets: []*armdatamigration.SQLBackupSetInfo{
		// 					},
		// 					IsFullBackupRestored: to.Ptr(false),
		// 					MigrationState: to.Ptr("WaitForFullBackupUploadOperation"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("targetdb"),
		// 			Type: to.Ptr("Microsoft.DataMigration/dataMigrations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname/providers/Microsoft.DataMigration/databaseMigrations/targetdb"),
		// 			Properties: &armdatamigration.DatabaseMigrationPropertiesSQLMi{
		// 				EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				Kind: to.Ptr(armdatamigration.ResourceTypeSQLMi),
		// 				MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
		// 				MigrationStatus: to.Ptr("InProgress"),
		// 				Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/miname"),
		// 				StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				SourceDatabaseName: to.Ptr("sourcename"),
		// 				MigrationStatusDetails: &armdatamigration.MigrationStatusDetails{
		// 					ActiveBackupSets: []*armdatamigration.SQLBackupSetInfo{
		// 					},
		// 					IsFullBackupRestored: to.Ptr(false),
		// 					MigrationState: to.Ptr("WaitForFullBackupUploadOperation"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
