package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/GetStaticSiteBuildDatabaseConnectionWithDetails.json
func ExampleStaticSitesClient_GetBuildDatabaseConnectionWithDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().GetBuildDatabaseConnectionWithDetails(ctx, "rg", "testStaticSite0", "default", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseConnection = armappservice.DatabaseConnection{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Web/staticSites/builds/databaseConnections"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0/builds/default/databaseConnections/default"),
	// 	Properties: &armappservice.DatabaseConnectionProperties{
	// 		ConfigurationFiles: []*armappservice.StaticSiteDatabaseConnectionConfigurationFileOverview{
	// 			{
	// 				Type: to.Ptr("configuration"),
	// 				Contents: to.Ptr("base64encodeddatabaseconfiguration"),
	// 				FileName: to.Ptr("staticwebapp.database.config.json"),
	// 			},
	// 			{
	// 				Type: to.Ptr("graphqlschema"),
	// 				Contents: to.Ptr("base64encodeddatabasegraphqlschema"),
	// 				FileName: to.Ptr("staticwebapp.database.schema.gql"),
	// 		}},
	// 		ConnectionIdentity: to.Ptr("SystemAssigned"),
	// 		ConnectionString: to.Ptr("AccountEndpoint=https://exampleDatabaseName.documents.azure.com:443/;Database=mydb;"),
	// 		Region: to.Ptr("West US 2"),
	// 		ResourceID: to.Ptr("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/databaseRG/providers/Microsoft.DocumentDB/databaseAccounts/exampleDatabaseName"),
	// 	},
	// }
}
