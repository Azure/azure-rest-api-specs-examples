package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_GetMigrationWithSuccessfulValidationOnly.json
func ExampleMigrationsClient_Get_migrationsGetMigrationWithSuccessfulValidationOnly() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationsClient().Get(ctx, "ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget", "testmigrationwithsuccessfulvalidationonly", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MigrationResource = armpostgresqlflexibleservers.MigrationResource{
	// 	Name: to.Ptr("testmigrationwithsuccessfulvalidationonly"),
	// 	Type: to.Ptr("Microsoft.DBForPostgreSql/flexibleServers/migrations"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget/migrations/testmigrationwithsuccessfulvalidationonly"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armpostgresqlflexibleservers.MigrationResourceProperties{
	// 		CurrentStatus: &armpostgresqlflexibleservers.MigrationStatus{
	// 			CurrentSubStateDetails: &armpostgresqlflexibleservers.MigrationSubStateDetails{
	// 				CurrentSubState: to.Ptr(armpostgresqlflexibleservers.MigrationSubStateCompleted),
	// 				DbDetails: map[string]*armpostgresqlflexibleservers.DbMigrationStatus{
	// 				},
	// 				ValidationDetails: &armpostgresqlflexibleservers.ValidationDetails{
	// 					DbLevelValidationDetails: []*armpostgresqlflexibleservers.DbLevelValidationStatus{
	// 						{
	// 							DatabaseName: to.Ptr("UnknownCollationTest"),
	// 							EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-16T05:29:39.248Z"); return t}()),
	// 							StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-16T05:29:38.904Z"); return t}()),
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
	// 					ValidationEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-16T05:29:39.666Z"); return t}()),
	// 					ValidationStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-16T05:29:35.634Z"); return t}()),
	// 				},
	// 			},
	// 			State: to.Ptr(armpostgresqlflexibleservers.MigrationStateSucceeded),
	// 		},
	// 		DbsToMigrate: []*string{
	// 			to.Ptr("UnknownCollationTest")},
	// 			MigrationID: to.Ptr("77840327-7be8-44b8-adc0-af0ccccfeb36"),
	// 			MigrationMode: to.Ptr(armpostgresqlflexibleservers.MigrationModeOffline),
	// 			MigrationOption: to.Ptr(armpostgresqlflexibleservers.MigrationOptionValidate),
	// 			MigrationWindowEndTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-16T05:30:05.885Z"); return t}()),
	// 			MigrationWindowStartTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-16T05:28:58.775Z"); return t}()),
	// 			OverwriteDbsInTarget: to.Ptr(armpostgresqlflexibleservers.OverwriteDbsInTargetEnumTrue),
	// 			SetupLogicalReplicationOnSourceDbIfNeeded: to.Ptr(armpostgresqlflexibleservers.LogicalReplicationOnSourceDbEnumTrue),
	// 			SourceDbServerResourceID: to.Ptr("20.228.214.65:5432@postgres"),
	// 			TargetDbServerResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/flexibleServers/testtarget"),
	// 			TriggerCutover: to.Ptr(armpostgresqlflexibleservers.TriggerCutoverEnumTrue),
	// 		},
	// 	}
}
