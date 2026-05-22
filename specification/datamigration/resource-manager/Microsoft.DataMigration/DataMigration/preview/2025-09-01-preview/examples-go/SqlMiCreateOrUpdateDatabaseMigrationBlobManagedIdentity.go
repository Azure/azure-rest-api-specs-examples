package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/SqlMiCreateOrUpdateDatabaseMigrationBlobManagedIdentity.json
func ExampleDatabaseMigrationsSQLMiClient_BeginCreateOrUpdate_createOrUpdateDatabaseMigrationResourceFromAzureBlobUsingManagedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseMigrationsSQLMiClient().BeginCreateOrUpdate(ctx, "testrg", "managedInstance1", "db1", armdatamigration.DatabaseMigrationSQLMi{
		Properties: &armdatamigration.DatabaseMigrationPropertiesSQLMi{
			BackupConfiguration: &armdatamigration.BackupConfiguration{
				SourceLocation: &armdatamigration.SourceLocation{
					AzureBlob: &armdatamigration.AzureBlob{
						AuthType:          to.Ptr(armdatamigration.AuthTypeManagedIdentity),
						BlobContainerName: to.Ptr("test"),
						Identity: &armdatamigration.ManagedServiceIdentity{
							Type: to.Ptr(armdatamigration.ManagedServiceIdentityTypeUserAssigned),
							UserAssignedIdentities: map[string]*armdatamigration.UserAssignedIdentity{
								"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity": {},
							},
						},
						StorageAccountResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount"),
					},
				},
			},
			Kind:             to.Ptr(armdatamigration.ResourceTypeSQLMi),
			MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
			OfflineConfiguration: &armdatamigration.OfflineConfiguration{
				LastBackupName: to.Ptr("last_backup_file_name"),
				Offline:        to.Ptr(true),
			},
			Scope:               to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance"),
			SourceDatabaseName:  to.Ptr("aaa"),
			SQLServerInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/instanceName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatamigration.DatabaseMigrationsSQLMiClientCreateOrUpdateResponse{
	// 	DatabaseMigrationSQLMi: armdatamigration.DatabaseMigrationSQLMi{
	// 		Name: to.Ptr("db1"),
	// 		Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance/providers/Microsoft.DataMigration/databaseMigrations/db1"),
	// 		Properties: &armdatamigration.DatabaseMigrationPropertiesSQLMi{
	// 			BackupConfiguration: &armdatamigration.BackupConfiguration{
	// 				SourceLocation: &armdatamigration.SourceLocation{
	// 					FileStorageType: to.Ptr("AzureBlob"),
	// 				},
	// 			},
	// 			EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
	// 			Kind: to.Ptr(armdatamigration.ResourceTypeSQLMi),
	// 			MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
	// 			MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
	// 			MigrationStatus: to.Ptr("InProgress"),
	// 			OfflineConfiguration: &armdatamigration.OfflineConfiguration{
	// 				LastBackupName: to.Ptr("last_backup_file_name"),
	// 				Offline: to.Ptr(true),
	// 			},
	// 			ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 			Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/instance"),
	// 			SourceDatabaseName: to.Ptr("sourcename"),
	// 			SourceServerName: to.Ptr("sourceserver"),
	// 			SQLServerInstanceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlServerInstances/instanceName"),
	// 			StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T08:00:00Z"); return t}()),
	// 		},
	// 	},
	// }
}
