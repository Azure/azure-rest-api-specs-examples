package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/PatchStaticSiteDatabaseConnection.json
func ExampleStaticSitesClient_UpdateDatabaseConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStaticSitesClient().UpdateDatabaseConnection(ctx, "rg", "testStaticSite0", "default", armappservice.DatabaseConnectionPatchRequest{
		Properties: &armappservice.DatabaseConnectionPatchRequestProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseConnection = armappservice.DatabaseConnection{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Web/staticSites/databaseConnections"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0/databaseConnections/default"),
	// 	Properties: &armappservice.DatabaseConnectionProperties{
	// 		ConnectionIdentity: to.Ptr("SystemAssigned"),
	// 		Region: to.Ptr("West US 2"),
	// 		ResourceID: to.Ptr("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/databaseRG/providers/Microsoft.DocumentDB/databaseAccounts/exampleDatabaseName"),
	// 	},
	// }
}
