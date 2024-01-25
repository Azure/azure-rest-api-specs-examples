package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/examples/CosmosDBMongoDBUserDefinitionGet.json
func ExampleMongoDBResourcesClient_GetMongoUserDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMongoDBResourcesClient().GetMongoUserDefinition(ctx, "myMongoUserDefinitionId", "myResourceGroupName", "myAccountName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MongoUserDefinitionGetResults = armcosmos.MongoUserDefinitionGetResults{
	// 	Name: to.Ptr("myUserId"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbUserDefinitions"),
	// 	ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongodbUserDefinitions/myUserId"),
	// 	Properties: &armcosmos.MongoUserDefinitionResource{
	// 		CustomData: to.Ptr("My custom data"),
	// 		DatabaseName: to.Ptr("sales"),
	// 		Mechanisms: to.Ptr("SCRAM-SHA-256"),
	// 		Roles: []*armcosmos.Role{
	// 			{
	// 				Db: to.Ptr("sales"),
	// 				Role: to.Ptr("myReadRole"),
	// 		}},
	// 		UserName: to.Ptr("myUserName"),
	// 	},
	// }
}
