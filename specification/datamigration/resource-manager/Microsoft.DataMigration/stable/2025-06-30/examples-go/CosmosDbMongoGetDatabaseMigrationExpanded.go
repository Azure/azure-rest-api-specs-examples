package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/930e8030f5058d947fea4e2640725baab8a4561a/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/CosmosDbMongoGetDatabaseMigrationExpanded.json
func ExampleDatabaseMigrationsMongoToCosmosDbvCoreMongoClient_Get_getMongoToCosmosDbMongoVCoreDatabaseMigrationWithTheExpandParameter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseMigrationsMongoToCosmosDbvCoreMongoClient().Get(ctx, "testrg", "targetCosmosDbClusterName", "migrationRequest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseMigrationCosmosDbMongo = armdatamigration.DatabaseMigrationCosmosDbMongo{
	// 	Name: to.Ptr("migrationRequest"),
	// 	Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName/providers/Microsoft.DataMigration/databaseMigrations/migrationRequest"),
	// 	Properties: &armdatamigration.DatabaseMigrationPropertiesCosmosDbMongo{
	// 		EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T09:00:00.000Z"); return t}()),
	// 		Kind: to.Ptr(armdatamigration.ResourceTypeMongoToCosmosDbMongo),
	// 		MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
	// 		MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/testMigrationService"),
	// 		MigrationStatus: to.Ptr("Succeeded"),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 		Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
	// 		StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
	// 		CollectionList: []*armdatamigration.MongoMigrationCollection{
	// 			{
	// 				MigrationProgressDetails: &armdatamigration.MongoMigrationProgressDetails{
	// 					DurationInSeconds: to.Ptr[int32](400),
	// 					MigrationStatus: to.Ptr(armdatamigration.MongoMigrationStatusCompleted),
	// 					ProcessedDocumentCount: to.Ptr[int64](25544),
	// 					SourceDocumentCount: to.Ptr[int64](25544),
	// 				},
	// 				SourceCollection: to.Ptr("sourceCol1"),
	// 				SourceDatabase: to.Ptr("sourceDb1"),
	// 				TargetCollection: to.Ptr("targetCol1"),
	// 				TargetDatabase: to.Ptr("targetDb1"),
	// 			},
	// 			{
	// 				MigrationProgressDetails: &armdatamigration.MongoMigrationProgressDetails{
	// 					DurationInSeconds: to.Ptr[int32](100),
	// 					MigrationStatus: to.Ptr(armdatamigration.MongoMigrationStatusCompleted),
	// 					ProcessedDocumentCount: to.Ptr[int64](255),
	// 					SourceDocumentCount: to.Ptr[int64](255),
	// 				},
	// 				SourceCollection: to.Ptr("sourceCol2"),
	// 				SourceDatabase: to.Ptr("sourceDb2"),
	// 				TargetCollection: to.Ptr("sourceCol2"),
	// 				TargetDatabase: to.Ptr("sourceDb2"),
	// 		}},
	// 		SourceMongoConnection: &armdatamigration.MongoConnectionInformation{
	// 			Host: to.Ptr("abc.mongodb.com"),
	// 			Port: to.Ptr[int32](88),
	// 		},
	// 		TargetMongoConnection: &armdatamigration.MongoConnectionInformation{
	// 			Host: to.Ptr("xyz.mongocluster.cosmos.azure.com"),
	// 			Port: to.Ptr[int32](10255),
	// 		},
	// 	},
	// }
}
