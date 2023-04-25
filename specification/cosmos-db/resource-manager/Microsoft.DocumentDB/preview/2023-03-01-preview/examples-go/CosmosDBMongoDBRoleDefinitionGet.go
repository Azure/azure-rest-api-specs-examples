package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-01-preview/examples/CosmosDBMongoDBRoleDefinitionGet.json
func ExampleMongoDBResourcesClient_GetMongoRoleDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMongoDBResourcesClient().GetMongoRoleDefinition(ctx, "myMongoRoleDefinitionId", "myResourceGroupName", "myAccountName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MongoRoleDefinitionGetResults = armcosmos.MongoRoleDefinitionGetResults{
	// 	Name: to.Ptr("myMongoRoleDefinitionId"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbRoleDefinitions"),
	// 	ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongodbRoleDefinitions/myMongoDbRoleDefinitionId"),
	// 	Properties: &armcosmos.MongoRoleDefinitionResource{
	// 		Type: to.Ptr(armcosmos.MongoRoleDefinitionTypeCustomRole),
	// 		DatabaseName: to.Ptr("sales"),
	// 		Privileges: []*armcosmos.Privilege{
	// 			{
	// 				Actions: []*string{
	// 					to.Ptr("find"),
	// 					to.Ptr("insert")},
	// 					Resource: &armcosmos.PrivilegeResource{
	// 						Collection: to.Ptr("coll"),
	// 						Db: to.Ptr("sales"),
	// 					},
	// 			}},
	// 			RoleName: to.Ptr("myRoleName"),
	// 			Roles: []*armcosmos.Role{
	// 				{
	// 					Db: to.Ptr("sales"),
	// 					Role: to.Ptr("myReadRole"),
	// 			}},
	// 		},
	// 	}
}
