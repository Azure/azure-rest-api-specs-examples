package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBSqlRoleDefinitionList.json
func ExampleSQLResourcesClient_NewListSQLRoleDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewSQLResourcesClient("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListSQLRoleDefinitionsPager("myResourceGroupName", "myAccountName", nil)
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
		// page.SQLRoleDefinitionListResult = armcosmos.SQLRoleDefinitionListResult{
		// 	Value: []*armcosmos.SQLRoleDefinitionGetResults{
		// 		{
		// 			Name: to.Ptr("myRoleDefinitionId"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlRoleDefinitions"),
		// 			ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId"),
		// 			Properties: &armcosmos.SQLRoleDefinitionResource{
		// 				Type: to.Ptr(armcosmos.RoleDefinitionTypeCustomRole),
		// 				AssignableScopes: []*string{
		// 					to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/sales"),
		// 					to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases")},
		// 					Permissions: []*armcosmos.Permission{
		// 						{
		// 							DataActions: []*string{
		// 								to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/create"),
		// 								to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read")},
		// 								NotDataActions: []*string{
		// 								},
		// 						}},
		// 						RoleName: to.Ptr("myRoleName"),
		// 					},
		// 			}},
		// 		}
	}
}
