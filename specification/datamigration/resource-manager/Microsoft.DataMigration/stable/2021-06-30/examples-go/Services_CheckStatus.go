package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Services_CheckStatus.json
func ExampleServicesClient_CheckStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CheckStatus(ctx, "DmsSdkRg", "DmsSdkService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceStatusResponse = armdatamigration.ServiceStatusResponse{
	// 	AgentVersion: to.Ptr("3.3.3871.6"),
	// 	Status: to.Ptr("Online"),
	// 	SupportedTaskTypes: []*string{
	// 		to.Ptr("diagnostic.count"),
	// 		to.Ptr("diagnostic.echo"),
	// 		to.Ptr("diagnostic.error"),
	// 		to.Ptr("diagnostic.fastCount"),
	// 		to.Ptr("ConnectToSource.SqlServer.Sync"),
	// 		to.Ptr("SyncMigrationOperationalTelemetry.LogCollector"),
	// 		to.Ptr("GetUserTables.AzureSqlDb.Sync"),
	// 		to.Ptr("ConnectToTarget.SqlDb.Sync"),
	// 		to.Ptr("Migrate.MySql.AzureDbForMySql.Sync"),
	// 		to.Ptr("Migrate.SqlServer.AzureSqlDb.Sync"),
	// 		to.Ptr("ValidateMigrationInput.SqlServer.SqlDb.Sync"),
	// 		to.Ptr("DataMigration.AzureSqlDbPostMigrationValidationScenarioId"),
	// 		to.Ptr("GetTDECertificates.Sql"),
	// 		to.Ptr("Migrate.SqlServer.AzureSqlDbMI"),
	// 		to.Ptr("ValidateMigrationInput.SqlServer.AzureSqlDbMI"),
	// 		to.Ptr("ConnectToTarget.AzureDbForMySql"),
	// 		to.Ptr("ConnectToSource.SqlServer"),
	// 		to.Ptr("GetUserTables.Sql"),
	// 		to.Ptr("ConnectToTarget.AzureSqlDbMI"),
	// 		to.Ptr("ConnectToTarget.SqlDb"),
	// 		to.Ptr("Migrate.SqlServer.SqlDb"),
	// 		to.Ptr("ConnectToSource.MySql")},
	// 		VMSize: to.Ptr("Standard_A4_v2"),
	// 	}
}
