package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/930e8030f5058d947fea4e2640725baab8a4561a/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/ListMigrationsByMigrationService.json
func ExampleMigrationServicesClient_NewListMigrationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMigrationServicesClient().NewListMigrationsPager("testrg", "testMigrationService", nil)
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
		// page.DatabaseMigrationBaseListResult = armdatamigration.DatabaseMigrationBaseListResult{
		// 	Value: []*armdatamigration.DatabaseMigrationBase{
		// 		{
		// 			Name: to.Ptr("migrationRequest1"),
		// 			Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName/providers/Microsoft.DataMigration/databaseMigrations/migrationRequest1"),
		// 			Properties: &armdatamigration.DatabaseMigrationPropertiesCosmosDbMongo{
		// 				Kind: to.Ptr(armdatamigration.ResourceTypeMongoToCosmosDbMongo),
		// 				MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/testMigrationService"),
		// 				MigrationStatus: to.Ptr("Creating"),
		// 				ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
		// 				Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
		// 				SourceMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 					Host: to.Ptr("abc.mongodb.com"),
		// 					Port: to.Ptr[int32](88),
		// 				},
		// 				TargetMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 					Host: to.Ptr("xyz.mongocluster.cosmos.azure.com"),
		// 					Port: to.Ptr[int32](10255),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("migrationRequest2"),
		// 			Type: to.Ptr("Microsoft.DataMigration/databaseMigrations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName/providers/Microsoft.DataMigration/databaseMigrations/migrationRequest2"),
		// 			Properties: &armdatamigration.DatabaseMigrationPropertiesCosmosDbMongo{
		// 				EndedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				Kind: to.Ptr(armdatamigration.ResourceTypeMongoToCosmosDbMongo),
		// 				MigrationFailureError: &armdatamigration.ErrorInfo{
		// 					Code: to.Ptr("400"),
		// 					Message: to.Ptr("Source or Target database connectivity could not be validated."),
		// 				},
		// 				MigrationOperationID: to.Ptr("858ba109-5ab7-4fa1-8aea-bea487cacdcd"),
		// 				MigrationService: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/testMigrationService"),
		// 				MigrationStatus: to.Ptr("Failed"),
		// 				ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
		// 				Scope: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DocumentDB/mongoClusters/targetCosmosDbClusterName"),
		// 				StartedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T08:00:00.000Z"); return t}()),
		// 				SourceMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 					Host: to.Ptr("abc.mongodb.com"),
		// 					Port: to.Ptr[int32](88),
		// 				},
		// 				TargetMongoConnection: &armdatamigration.MongoConnectionInformation{
		// 					Host: to.Ptr("xyz.mongocluster.cosmos.azure.com"),
		// 					Port: to.Ptr[int32](10255),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
