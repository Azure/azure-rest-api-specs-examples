package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-07-01-preview/MigrationsGetMigrationWithSuccessfulValidationAndMigration.json
func ExampleMigrationsClient_Get_getInformationAboutAMigrationWithSuccessfulValidationAndSuccessfulMigration() {
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
	// 		Type: to.Ptr("Microsoft.DBForPostgreSql/flexibleServers/migrations"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/flexibleServers/exampletarget/migrations/examplemigration"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armpostgresqlflexibleservers.MigrationProperties{
	// 			CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 				CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubstateDetails{
	// 					CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubstateCompleted),
	// 					DbDetails: map[string]*armpostgresqlflexibleservers.DatabaseMigrationState{
	// 						"exampledatabase": &armpostgresqlflexibleservers.DatabaseMigrationState{
	// 							AppliedChanges: to.Ptr[int32](0),
	// 							CdcDeleteCounter: to.Ptr[int32](0),
	// 							CdcInsertCounter: to.Ptr[int32](0),
	// 							CdcUpdateCounter: to.Ptr[int32](0),
	// 							DatabaseName: to.Ptr("exampledatabase"),
	// 							EndedOn: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 							FullLoadCompletedTables: to.Ptr[int32](10),
	// 							FullLoadErroredTables: to.Ptr[int32](0),
	// 							FullLoadLoadingTables: to.Ptr[int32](0),
	// 							FullLoadQueuedTables: to.Ptr[int32](0),
	// 							IncomingChanges: to.Ptr[int32](0),
	// 							Latency: to.Ptr[int32](0),
	// 							MigrationState: to.Ptr(armpostgresqlflexibleservers.MigrationDatabaseStateSucceeded),
	// 							StartedOn: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 						},
	// 					},
	// 					ValidationDetails: &armpostgresqlflexibleservers.ValidationDetails{
	// 						DbLevelValidationDetails: []*armpostgresqlflexibleservers.DbLevelValidationStatus{
	// 							{
	// 								DatabaseName: to.Ptr("exampledatabase"),
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
	// 						Status: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 						ValidationEndTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 						ValidationStartTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 					},
	// 				},
	// 				State: to.Ptr(armpostgresqlflexibleservers.MigrationStateSucceeded),
	// 			},
	// 			DbsToMigrate: []*string{
	// 				to.Ptr("exampledatabase"),
	// 			},
	// 			MigrateRoles: to.Ptr(armpostgresqlflexibleservers.MigrateRolesAndPermissionsFalse),
	// 			MigrationID: to.Ptr("f2354e72-2828-4a19-ad20-b4cd9e2673c1"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationOption: to.Ptr(armpostgresqlflexibleservers.MigrationOptionValidateAndMigrate),
	// 			MigrationWindowEndTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 20, 30, 22, 123456000, time.UTC)),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(time.Date(2025, time.June, 1, 18, 30, 22, 123456000, time.UTC)),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDatabasesOnTargetServerTrue),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceServerTrue),
	// 			SourceDbServerResourceID: to.Ptr("20.228.214.65:5432@postgres"),
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBForPostgreSql/flexibleServers/exampletarget"),
	// 			TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverTrue),
	// 		},
	// 	},
	// }
}
