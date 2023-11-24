package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_GetMigrationWithSuccessfulValidationAndMigration.json
func ExampleMigrationsClient_Get_migrationsGetMigrationWithSuccessfulValidationAndMigration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Get(ctx, "ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget", "testmigrationwithsuccessfulvalidationandmigration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationResource = armpostgresqlflexibleservers.MigrationResource{
	// 	Name: to.Ptr("testmigrationwithsuccessfulvalidationandmigration"),
	// 	Type: to.Ptr("Microsoft.DBForPostgreSql/flexibleServers/migrations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget/migrations/testmigrationwithsuccessfulvalidationandmigration"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpostgresqlflexibleservers.MigrationResourceProperties{
	// 		CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 			CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubStateDetails{
	// 				CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubStateCompleted),
	// 				DbDetails: map[string]*armpostgresqlflexibleservers.DbMigrationStatus{
	// 					"testdb3": &armpostgresqlflexibleservers.DbMigrationStatus{
	// 						AppliedChanges: to.Ptr[int32](0),
	// 						CdcDeleteCounter: to.Ptr[int32](0),
	// 						CdcInsertCounter: to.Ptr[int32](0),
	// 						CdcUpdateCounter: to.Ptr[int32](0),
	// 						DatabaseName: to.Ptr("testdb3"),
	// 						EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:55:49.681Z"); return t}()),
	// 						FullLoadCompletedTables: to.Ptr[int32](10),
	// 						FullLoadErroredTables: to.Ptr[int32](0),
	// 						FullLoadLoadingTables: to.Ptr[int32](0),
	// 						FullLoadQueuedTables: to.Ptr[int32](0),
	// 						IncomingChanges: to.Ptr[int32](0),
	// 						Latency: to.Ptr[int32](0),
	// 						MigrationState: to.Ptr(armpostgresqlflexibleservers.MigrationDbStateSucceeded),
	// 						StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:55:28.995Z"); return t}()),
	// 					},
	// 				},
	// 				ValidationDetails: &armpostgresqlflexibleservers.ValidationDetails{
	// 					DbLevelValidationDetails: []*armpostgresqlflexibleservers.DbLevelValidationStatus{
	// 						{
	// 							DatabaseName: to.Ptr("testdb3"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:54:29.894Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:54:29.553Z"); return t}()),
	// 							Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 								{
	// 									Type: to.Ptr("SchemaValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 								},
	// 								{
	// 									Type: to.Ptr("ExtensionsValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 								},
	// 								{
	// 									Type: to.Ptr("CollationsValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							}},
	// 					}},
	// 					ServerLevelValidationDetails: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 						{
	// 							Type: to.Ptr("AuthenticationAndConnectivityValidation"),
	// 							State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 						},
	// 						{
	// 							Type: to.Ptr("SourceVersionValidation"),
	// 							State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 						},
	// 						{
	// 							Type: to.Ptr("ServerParametersValidation"),
	// 							State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 					}},
	// 					Status: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 					ValidationEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:55:00.499Z"); return t}()),
	// 					ValidationStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:54:28.077Z"); return t}()),
	// 				},
	// 			},
	// 			State: to.Ptr(armpostgresqlflexibleservers.MigrationStateSucceeded),
	// 		},
	// 		DbsToMigrate: []*string{
	// 			to.Ptr("testdb3")},
	// 			MigrationID: to.Ptr("f2354e72-2828-4a19-ad20-b4cd9e2673c1"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationOption: to.Ptr(armpostgresqlflexibleservers.MigrationOptionValidateAndMigrate),
	// 			MigrationWindowEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:57:30.736Z"); return t}()),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-12T05:54:15.999Z"); return t}()),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDbsInTargetEnumTrue),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceDbEnumTrue),
	// 			SourceDbServerResourceID: to.Ptr("20.228.214.65:5432@postgres"),
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget"),
	// 			TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverEnumTrue),
	// 		},
	// 	}
}
