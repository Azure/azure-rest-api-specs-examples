package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-02-15-preview/examples/CosmosDBSqlRoleDefinitionGet.json
func ExampleSQLResourcesClient_GetSQLRoleDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLResourcesClient().GetSQLRoleDefinition(ctx, "myRoleDefinitionId", "myResourceGroupName", "myAccountName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLRoleDefinitionGetResults = armcosmos.SQLRoleDefinitionGetResults{
	// 	Name: to.Ptr("myRoleDefinitionId"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlRoleDefinitions"),
	// 	ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId"),
	// 	Properties: &armcosmos.SQLRoleDefinitionResource{
	// 		Type: to.Ptr(armcosmos.RoleDefinitionTypeCustomRole),
	// 		AssignableScopes: []*string{
	// 			to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/sales"),
	// 			to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases")},
	// 			Permissions: []*armcosmos.Permission{
	// 				{
	// 					DataActions: []*string{
	// 						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/create"),
	// 						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read")},
	// 						NotDataActions: []*string{
	// 						},
	// 				}},
	// 				RoleName: to.Ptr("myRoleName"),
	// 			},
	// 		}
}
