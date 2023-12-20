package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_GetMigrationWithValidationFailures.json
func ExampleMigrationsClient_Get_migrationsGetMigrationWithValidationFailures() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Get(ctx, "ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget", "testmigrationwithvalidationfailure", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationResource = armpostgresqlflexibleservers.MigrationResource{
	// 	Name: to.Ptr("testmigrationwithvalidationfailure"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/migrations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testtarget/migrations/testmigrationwithvalidationfailure"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armpostgresqlflexibleservers.MigrationResourceProperties{
	// 		CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 			CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubStateDetails{
	// 				CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubStateCompleted),
	// 				DbDetails: map[string]*armpostgresqlflexibleservers.DbMigrationStatus{
	// 				},
	// 				ValidationDetails: &armpostgresqlflexibleservers.ValidationDetails{
	// 					DbLevelValidationDetails: []*armpostgresqlflexibleservers.DbLevelValidationStatus{
	// 						{
	// 							DatabaseName: to.Ptr("c.utf-8"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:02.810Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:01.830Z"); return t}()),
	// 							Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 								{
	// 									Type: to.Ptr("SchemaValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 								},
	// 								{
	// 									Type: to.Ptr("ExtensionsValidation"),
	// 									Messages: []*armpostgresqlflexibleservers.ValidationMessage{
	// 										{
	// 											Message: to.Ptr("Unsupported Extension. Single to Flex migration tool does not support migration of databases having postgres_fdw extension. Consider performing the migration through other migration tools such as pg_dump/pg_restore (https://aka.ms/migrate-using-pgdump-restore)"),
	// 											State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 									}},
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 								},
	// 								{
	// 									Type: to.Ptr("CollationsValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							}},
	// 						},
	// 						{
	// 							DatabaseName: to.Ptr("postgres"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:02.810Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:01.830Z"); return t}()),
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
	// 						},
	// 						{
	// 							DatabaseName: to.Ptr("readme_to_recover"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:02.810Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:01.830Z"); return t}()),
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
	// 						},
	// 						{
	// 							DatabaseName: to.Ptr("testdb1"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:02.810Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:01.830Z"); return t}()),
	// 							Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 								{
	// 									Type: to.Ptr("SchemaValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 								},
	// 								{
	// 									Type: to.Ptr("ExtensionsValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							}},
	// 						},
	// 						{
	// 							DatabaseName: to.Ptr("testdb2"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:02.810Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:01.830Z"); return t}()),
	// 							Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 								{
	// 									Type: to.Ptr("SchemaValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 								},
	// 								{
	// 									Type: to.Ptr("ExtensionsValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							}},
	// 						},
	// 						{
	// 							DatabaseName: to.Ptr("testSchema"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:02.810Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:01.830Z"); return t}()),
	// 							Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 								{
	// 									Type: to.Ptr("SchemaValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 								},
	// 								{
	// 									Type: to.Ptr("ExtensionsValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 							}},
	// 						},
	// 						{
	// 							DatabaseName: to.Ptr("UnknownCollationTest"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:02.810Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:01.830Z"); return t}()),
	// 							Summary: []*armpostgresqlflexibleservers.ValidationSummaryItem{
	// 								{
	// 									Type: to.Ptr("SchemaValidation"),
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateSucceeded),
	// 								},
	// 								{
	// 									Type: to.Ptr("ExtensionsValidation"),
	// 									Messages: []*armpostgresqlflexibleservers.ValidationMessage{
	// 										{
	// 											Message: to.Ptr("Unsupported Extension. Single to Flex migration tool does not support migration of databases having postgres_fdw extension. Consider performing the migration through other migration tools such as pg_dump/pg_restore (https://aka.ms/migrate-using-pgdump-restore)"),
	// 											State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 									}},
	// 									State: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
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
	// 					Status: to.Ptr(armpostgresqlflexibleservers.ValidationStateFailed),
	// 					ValidationEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:20:33.136Z"); return t}()),
	// 					ValidationStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:19:59.917Z"); return t}()),
	// 				},
	// 			},
	// 			State: to.Ptr(armpostgresqlflexibleservers.MigrationStateValidationFailed),
	// 		},
	// 		DbsToMigrate: []*string{
	// 			to.Ptr("c.utf-8"),
	// 			to.Ptr("postgres"),
	// 			to.Ptr("readme_to_recover"),
	// 			to.Ptr("testdb1"),
	// 			to.Ptr("testdb2"),
	// 			to.Ptr("testSchema"),
	// 			to.Ptr("UnknownCollationTest")},
	// 			MigrationID: to.Ptr("a3e2d3cc-b139-4201-9431-e4f3003140fd"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationOption: to.Ptr(armpostgresqlflexibleservers.MigrationOptionValidate),
	// 			MigrationWindowEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:21:00.043Z"); return t}()),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-17T19:19:47.448Z"); return t}()),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDbsInTargetEnumTrue),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceDbEnumTrue),
	// 			SourceDbServerResourceID: to.Ptr("20.228.214.65:5432@postgres"),
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testtarget"),
	// 			TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverEnumTrue),
	// 		},
	// 	}
}
