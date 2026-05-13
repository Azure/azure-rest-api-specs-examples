package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBSqlRoleDefinitionCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateSQLRoleDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginCreateUpdateSQLRoleDefinition(ctx, "myRoleDefinitionId", "myResourceGroupName", "myAccountName", armcosmos.SQLRoleDefinitionCreateUpdateParameters{
		Properties: &armcosmos.SQLRoleDefinitionResource{
			Type: to.Ptr(armcosmos.RoleDefinitionTypeCustomRole),
			AssignableScopes: []*string{
				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/sales"),
				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases"),
			},
			Permissions: []*armcosmos.Permission{
				{
					DataActions: []*string{
						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/create"),
						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read"),
					},
					NotDataActions: []*string{},
				},
			},
			RoleName: to.Ptr("myRoleName"),
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
	// res = armcosmos.SQLResourcesClientCreateUpdateSQLRoleDefinitionResponse{
	// 	SQLRoleDefinitionGetResults: &armcosmos.SQLRoleDefinitionGetResults{
	// 		Name: to.Ptr("myRoleDefinitionId"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlRoleDefinitions"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId"),
	// 		Properties: &armcosmos.SQLRoleDefinitionResource{
	// 			Type: to.Ptr(armcosmos.RoleDefinitionTypeCustomRole),
	// 			AssignableScopes: []*string{
	// 				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/sales"),
	// 				to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases"),
	// 			},
	// 			Permissions: []*armcosmos.Permission{
	// 				{
	// 					DataActions: []*string{
	// 						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/create"),
	// 						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read"),
	// 					},
	// 				},
	// 			},
	// 			RoleName: to.Ptr("myRoleName"),
	// 		},
	// 	},
	// }
}
