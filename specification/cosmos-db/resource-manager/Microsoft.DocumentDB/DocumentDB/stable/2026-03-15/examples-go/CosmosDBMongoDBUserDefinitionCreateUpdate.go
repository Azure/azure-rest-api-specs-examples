package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBMongoDBUserDefinitionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoUserDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoDBResourcesClient().BeginCreateUpdateMongoUserDefinition(ctx, "myMongoUserDefinitionId", "myResourceGroupName", "myAccountName", armcosmos.MongoUserDefinitionCreateUpdateParameters{
		Properties: &armcosmos.MongoUserDefinitionResource{
			UserName:     to.Ptr("myUserName"),
			Password:     to.Ptr("myPassword"),
			DatabaseName: to.Ptr("sales"),
			CustomData:   to.Ptr("My custom data"),
			Roles: []*armcosmos.Role{
				{
					Role: to.Ptr("myReadRole"),
					Db:   to.Ptr("sales"),
				},
			},
			Mechanisms: to.Ptr("SCRAM-SHA-256"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.MongoDBResourcesClientCreateUpdateMongoUserDefinitionResponse{
	// 	MongoUserDefinitionGetResults: armcosmos.MongoUserDefinitionGetResults{
	// 		ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongodbUserDefinitions/myUserId"),
	// 		Name: to.Ptr("myUserName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbUserDefinitions"),
	// 		Properties: &armcosmos.MongoUserDefinitionResource{
	// 			UserName: to.Ptr("myUserName"),
	// 			DatabaseName: to.Ptr("sales"),
	// 			CustomData: to.Ptr("My custom data"),
	// 			Roles: []*armcosmos.Role{
	// 				{
	// 					Db: to.Ptr("sales"),
	// 					Role: to.Ptr("myReadRole"),
	// 				},
	// 			},
	// 			Mechanisms: to.Ptr("SCRAM-SHA-256"),
	// 		},
	// 	},
	// }
}
