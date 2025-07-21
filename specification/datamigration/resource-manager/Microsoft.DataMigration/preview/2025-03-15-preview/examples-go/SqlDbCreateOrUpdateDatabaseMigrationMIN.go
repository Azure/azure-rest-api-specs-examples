package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/SqlDbCreateOrUpdateDatabaseMigrationMIN.json
func ExampleDatabaseMigrationsSQLDbClient_BeginCreateOrUpdate_createOrUpdateDatabaseMigrationResourceWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseMigrationsSQLDbClient().BeginCreateOrUpdate(ctx, "testrg", "sqldbinstance", "db1", armdatamigration.DatabaseMigrationSQLDb{
		Properties: &armdatamigration.DatabaseMigrationPropertiesSQLDb{
			Kind:               to.Ptr(armdatamigration.ResourceTypeSQLDb),
			MigrationService:   to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
			Scope:              to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance"),
			SourceDatabaseName: to.Ptr("aaa"),
			SourceSQLConnection: &armdatamigration.SQLConnectionInformation{
				Authentication:         to.Ptr("WindowsAuthentication"),
				DataSource:             to.Ptr("aaa"),
				EncryptConnection:      to.Ptr(true),
				Password:               to.Ptr("placeholder"),
				TrustServerCertificate: to.Ptr(true),
				UserName:               to.Ptr("bbb"),
			},
			TargetSQLConnection: &armdatamigration.SQLConnectionInformation{
				Authentication:         to.Ptr("SqlAuthentication"),
				DataSource:             to.Ptr("sqldbinstance"),
				EncryptConnection:      to.Ptr(true),
				Password:               to.Ptr("placeholder"),
				TrustServerCertificate: to.Ptr(true),
				UserName:               to.Ptr("bbb"),
			},
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
	// res.DatabaseMigrationSQLDb = armdatamigration.DatabaseMigrationSQLDb{
	// 	Name: to.Ptr("db1"),
	// 	Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance/providers/Microsoft.DataMigration/databaseMigrations/db1"),
	// 	Properties: &armdatamigration.DatabaseMigrationPropertiesSQLDb{
	// 		EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		Kind: to.Ptr(armdatamigration.ResourceTypeSQLDb),
	// 		MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
	// 		MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
	// 		MigrationStatus: to.Ptr("InProgress"),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 		Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance"),
	// 		StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		SourceDatabaseName: to.Ptr("sourcename"),
	// 		SourceServerName: to.Ptr("sourceserver"),
	// 		OfflineConfiguration: &armdatamigration.SQLDbOfflineConfiguration{
	// 			Offline: to.Ptr(true),
	// 		},
	// 	},
	// }
}
