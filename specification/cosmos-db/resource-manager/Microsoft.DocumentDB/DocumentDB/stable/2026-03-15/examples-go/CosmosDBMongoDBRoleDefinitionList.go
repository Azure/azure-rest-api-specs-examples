package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBMongoDBRoleDefinitionList.json
func ExampleMongoDBResourcesClient_NewListMongoRoleDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
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
		// page = armcosmos.MongoDBResourcesClientListMongoRoleDefinitionsResponse{
		// 	MongoRoleDefinitionListResult: armcosmos.MongoRoleDefinitionListResult{
		// 		Value: []*armcosmos.MongoRoleDefinitionGetResults{
		// 			{
		// 				ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/mongodbRoleDefinitions/myRoleDefinitionId"),
		// 				Name: to.Ptr("myRoleDefinitionId"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/mongodbRoleDefinitions"),
		// 				Properties: &armcosmos.MongoRoleDefinitionResource{
		// 					RoleName: to.Ptr("myRoleName"),
		// 					Type: to.Ptr(armcosmos.MongoRoleDefinitionTypeCustomRole),
		// 					DatabaseName: to.Ptr("sales"),
		// 					Privileges: []*armcosmos.Privilege{
		// 						{
		// 							Resource: &armcosmos.PrivilegeResource{
		// 								Db: to.Ptr("sales"),
		// 								Collection: to.Ptr("coll"),
		// 							},
		// 							Actions: []*string{
		// 								to.Ptr("find"),
		// 								to.Ptr("insert"),
		// 							},
		// 						},
		// 					},
		// 					Roles: []*armcosmos.Role{
		// 						{
		// 							Db: to.Ptr("sales"),
		// 							Role: to.Ptr("myReadRole"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
