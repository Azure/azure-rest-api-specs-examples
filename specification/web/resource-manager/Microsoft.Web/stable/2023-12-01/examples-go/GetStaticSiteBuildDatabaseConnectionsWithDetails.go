package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetStaticSiteBuildDatabaseConnectionsWithDetails.json
func ExampleStaticSitesClient_NewGetBuildDatabaseConnectionsWithDetailsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStaticSitesClient().NewGetBuildDatabaseConnectionsWithDetailsPager("rg", "testStaticSite0", "default", nil)
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
		// page.DatabaseConnectionCollection = armappservice.DatabaseConnectionCollection{
		// 	Value: []*armappservice.DatabaseConnection{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Web/staticSites/builds/databaseConnections"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0/builds/default/databaseConnections/default"),
		// 			Properties: &armappservice.DatabaseConnectionProperties{
		// 				ConfigurationFiles: []*armappservice.StaticSiteDatabaseConnectionConfigurationFileOverview{
		// 					{
		// 						Type: to.Ptr("configuration"),
		// 						Contents: to.Ptr("base64encodeddatabaseconfiguration"),
		// 						FileName: to.Ptr("staticwebapp.database.config.json"),
		// 					},
		// 					{
		// 						Type: to.Ptr("graphqlschema"),
		// 						Contents: to.Ptr("base64encodeddatabasegraphqlschema"),
		// 						FileName: to.Ptr("staticwebapp.database.schema.gql"),
		// 				}},
		// 				ConnectionIdentity: to.Ptr("SystemAssigned"),
		// 				ConnectionString: to.Ptr("AccountEndpoint=https://exampleDatabaseName.documents.azure.com:443/;Database=mydb;"),
		// 				Region: to.Ptr("West US 2"),
		// 				ResourceID: to.Ptr("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/databaseRG/providers/Microsoft.DocumentDB/databaseAccounts/exampleDatabaseName"),
		// 			},
		// 	}},
		// }
	}
}
