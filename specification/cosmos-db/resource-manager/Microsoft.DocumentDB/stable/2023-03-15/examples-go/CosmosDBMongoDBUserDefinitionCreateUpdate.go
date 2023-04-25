package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b065afacb99e7ec65787383550ee233e0e02a6e6/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBMongoDBUserDefinitionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoUserDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoDBResourcesClient().BeginCreateUpdateMongoUserDefinition(ctx, "myMongoUserDefinitionId", "myResourceGroupName", "myAccountName", armcosmos.MongoUserDefinitionCreateUpdateParameters{
		Properties: &armcosmos.MongoUserDefinitionResource{
			CustomData:   to.Ptr("My custom data"),
			DatabaseName: to.Ptr("sales"),
			Mechanisms:   to.Ptr("SCRAM-SHA-256"),
			Password:     to.Ptr("myPassword"),
			Roles: []*armcosmos.Role{
				{
					Db:   to.Ptr("sales"),
					Role: to.Ptr("myReadRole"),
				}},
			UserName: to.Ptr("myUserName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MongoUserDefinitionGetResults = armcosmos.MongoUserDefinitionGetResults{
	// 	Name: to.Ptr("myUserName"),
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
