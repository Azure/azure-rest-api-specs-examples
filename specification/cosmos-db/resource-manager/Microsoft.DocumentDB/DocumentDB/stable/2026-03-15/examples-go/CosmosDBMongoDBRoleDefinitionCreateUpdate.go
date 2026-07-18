package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBMongoDBRoleDefinitionCreateUpdate.json
func ExampleMongoDBResourcesClient_BeginCreateUpdateMongoRoleDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMongoDBResourcesClient().BeginCreateUpdateMongoRoleDefinition(ctx, "myMongoRoleDefinitionId", "myResourceGroupName", "myAccountName", armcosmos.MongoRoleDefinitionCreateUpdateParameters{
		Properties: &armcosmos.MongoRoleDefinitionResource{
			RoleName:     to.Ptr("myRoleName"),
			DatabaseName: to.Ptr("sales"),
			Privileges: []*armcosmos.Privilege{
				{
					Resource: &armcosmos.PrivilegeResource{
						Db:         to.Ptr("sales"),
						Collection: to.Ptr("sales"),
					},
					Actions: []*string{
						to.Ptr("insert"),
						to.Ptr("find"),
					},
				},
			},
			Roles: []*armcosmos.Role{
				{
					Role: to.Ptr("myInheritedRole"),
					Db:   to.Ptr("sales"),
				},
			},
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
	// res = armcosmos.MongoDBResourcesClientCreateUpdateMongoRoleDefinitionResponse{
	// 	MongoRoleDefinitionGetResults: armcosmos.MongoRoleDefinitionGetResults{
	// 		ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongodbRoleDefinitions/myMongoDbRoleDefinitionId"),
	// 		Name: to.Ptr("myMongoDbRoleDefinitionId"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongoDbRoleDefinitionId"),
	// 		Properties: &armcosmos.MongoRoleDefinitionResource{
	// 			RoleName: to.Ptr("myRoleName"),
	// 			Type: to.Ptr(armcosmos.MongoRoleDefinitionTypeCustomRole),
	// 			DatabaseName: to.Ptr("sales"),
	// 			Privileges: []*armcosmos.Privilege{
	// 				{
	// 					Resource: &armcosmos.PrivilegeResource{
	// 						Db: to.Ptr("sales"),
	// 						Collection: to.Ptr("coll"),
	// 					},
	// 					Actions: []*string{
	// 						to.Ptr("find"),
	// 						to.Ptr("insert"),
	// 					},
	// 				},
	// 			},
	// 			Roles: []*armcosmos.Role{
	// 				{
	// 					Db: to.Ptr("sales"),
	// 					Role: to.Ptr("myReadRole"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
