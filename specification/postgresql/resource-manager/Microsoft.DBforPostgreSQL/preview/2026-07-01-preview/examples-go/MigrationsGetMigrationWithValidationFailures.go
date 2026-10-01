package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-07-01-preview/MigrationsGetMigrationWithValidationFailures.json
func ExampleMigrationsClient_Get_getInformationAboutAMigrationWithValidationFailures() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Get(ctx, "exampleresourcegroup", "exampleserver", "examplemigration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpostgresqlflexibleservers.MigrationsClientGetResponse{
	// 	Migration: armpostgresqlflexibleservers.Migration{
	// 		Name: to.Ptr("examplemigration"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/migrations"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampletarget/migrations/examplemigration"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armpostgresqlflexibleservers.MigrationProperties{
	// 			CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 				CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubstateDetails{
	// 					CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubstateCompleted),
	// 					DbDetails: map[string]*armpostgresqlflexibleservers.DatabaseMigrationState{
	// 					},
	// 					ValidationDetails: &armpostgresqlflexibleservers.ValidationDetails{
	// 						DbLevelValidationDetails: []*armpostgresqlflexibleservers.DbLevelValidationStatus{
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase1"),
	// 								EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 								StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 								Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 									{
	// 										Type: to.Ptr("SchemaValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("ExtensionsValidation"),
	// 										Messages: []*armpostgresqlflexibleservers.ValidationMessage{
	// 											{
	// 												Message: to.Ptr("Unsupported Extension. Single to Flex migration tool does not support migration of databases having postgres_fdw extension. Consider performing the migration through other migration tools such as pg_dump/pg_restore (https://aka.ms/migrate-using-pgdump-restore)"),
	// 												State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 											},
	// 										},
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 									},
	// 									{
	// 										Type: to.Ptr("CollationsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase2"),
	// 								EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 								StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 								Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 									{
	// 										Type: to.Ptr("SchemaValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("ExtensionsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("CollationsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase3"),
	// 								EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 								StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 								Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 									{
	// 										Type: to.Ptr("SchemaValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("ExtensionsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("CollationsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase4"),
	// 								EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 								StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 								Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 									{
	// 										Type: to.Ptr("SchemaValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("ExtensionsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase5"),
	// 								EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 								StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 								Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 									{
	// 										Type: to.Ptr("SchemaValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("ExtensionsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase6"),
	// 								EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 								StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 								Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 									{
	// 										Type: to.Ptr("SchemaValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("ExtensionsValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 								},
	// 							},
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase7"),
	// 								EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 								StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 								Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 									{
	// 										Type: to.Ptr("SchemaValidation"),
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 									},
	// 									{
	// 										Type: to.Ptr("ExtensionsValidation"),
	// 										Messages: []*armpostgresqlflexibleservers.ValidationMessage{
	// 											{
	// 												Message: to.Ptr("Unsupported Extension. Single to Flex migration tool does not support migration of databases having postgres_fdw extension. Consider performing the migration through other migration tools such as pg_dump/pg_restore (https://aka.ms/migrate-using-pgdump-restore)"),
	// 												State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 											},
	// 										},
	// 										State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						ServerLevelValidationDetails: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 							{
	// 								Type: to.Ptr("AuthenticationAndConnectivityValidation"),
	// 								State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							},
	// 							{
	// 								Type: to.Ptr("SourceVersionValidation"),
	// 								State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							},
	// 							{
	// 								Type: to.Ptr("ServerParametersValidation"),
	// 								State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							},
	// 						},
	// 						Status: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 						ValidationEndTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 						ValidationStartTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 					},
	// 				},
	// 				State: to.Ptr(armpostgresqlflexibleservers.MigrationStateValidationFailed),
	// 			},
	// 			DbsToMigrate: []*string{
	// 				to.Ptr("exampledatabase1"),
	// 				to.Ptr("exampledatabase2"),
	// 				to.Ptr("exampledatabase3"),
	// 				to.Ptr("exampledatabase4"),
	// 				to.Ptr("exampledatabase5"),
	// 				to.Ptr("exampledatabase6"),
	// 				to.Ptr("exampledatabase7"),
	// 			},
	// 			MigrateRoles: to.Ptr(armpostgresqlflexibleservers.MigrateRolesAndPermissionsFalse),
	// 			MigrationID: to.Ptr("a3e2d3cc-b139-4201-9431-e4f3003140fd"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationOption: to.Ptr(armpostgresqlflexibleservers.MigrationOptionValidate),
	// 			MigrationWindowEndTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDatabasesOnTargetServerTrue),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceServerTrue),
	// 			SourceDbServerResourceID: to.Ptr("20.228.214.65:5432@postgres"),
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampletarget"),
	// 			TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverTrue),
	// 		},
	// 	},
	// }
}
