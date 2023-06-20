package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBMongoDBUserDefinitionList.json
func ExampleMongoDBResourcesClient_NewListMongoUserDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMongoDBResourcesClient().NewListMongoUserDefinitionsPager("myResourceGroupName", "myAccountName", nil)
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
		// page.MongoUserDefinitionListResult = armcosmos.MongoUserDefinitionListResult{
		// 	Value: []*armcosmos.MongoUserDefinitionGetResults{
		// 		{
		// 			Name: to.Ptr("myUserId"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbUserDefinitions"),
		// 			ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongodbUserDefinitions/myUserId"),
		// 			Properties: &armcosmos.MongoUserDefinitionResource{
		// 				CustomData: to.Ptr("My custom data"),
		// 				DatabaseName: to.Ptr("sales"),
		// 				Mechanisms: to.Ptr("SCRAM-SHA-256"),
		// 				Roles: []*armcosmos.Role{
		// 					{
		// 						Db: to.Ptr("sales"),
		// 						Role: to.Ptr("myReadRole"),
		// 				}},
		// 				UserName: to.Ptr("myUserName"),
		// 			},
		// 	}},
		// }
	}
}
