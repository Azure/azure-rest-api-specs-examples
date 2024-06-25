package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ec7ee8842bf615c2f0354bf8b5b8725fdac9454a/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/CosmosDBMongoDBRoleDefinitionList.json
func ExampleMongoDBResourcesClient_NewListMongoRoleDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMongoDBResourcesClient().NewListMongoRoleDefinitionsPager("myResourceGroupName", "myAccountName", nil)
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
		// page.MongoRoleDefinitionListResult = armcosmos.MongoRoleDefinitionListResult{
		// 	Value: []*armcosmos.MongoRoleDefinitionGetResults{
		// 		{
		// 			Name: to.Ptr("myRoleDefinitionId"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbRoleDefinitions"),
		// 			ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongodbRoleDefinitions/myRoleDefinitionId"),
		// 			Properties: &armcosmos.MongoRoleDefinitionResource{
		// 				Type: to.Ptr(armcosmos.MongoRoleDefinitionTypeCustomRole),
		// 				DatabaseName: to.Ptr("sales"),
		// 				Privileges: []*armcosmos.Privilege{
		// 					{
		// 						Actions: []*string{
		// 							to.Ptr("find"),
		// 							to.Ptr("insert")},
		// 							Resource: &armcosmos.PrivilegeResource{
		// 								Collection: to.Ptr("coll"),
		// 								Db: to.Ptr("sales"),
		// 							},
		// 					}},
		// 					RoleName: to.Ptr("myRoleName"),
		// 					Roles: []*armcosmos.Role{
		// 						{
		// 							Db: to.Ptr("sales"),
		// 							Role: to.Ptr("myReadRole"),
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
