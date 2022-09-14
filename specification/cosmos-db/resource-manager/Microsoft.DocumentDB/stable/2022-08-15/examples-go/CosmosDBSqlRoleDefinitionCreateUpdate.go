package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBSqlRoleDefinitionCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateSQLRoleDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewSQLResourcesClient("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateSQLRoleDefinition(ctx, "myRoleDefinitionId", "myResourceGroupName", "myAccountName", armcosmos.SQLRoleDefinitionCreateUpdateParameters{
		Properties: &armcosmos.SQLRoleDefinitionResource{
			Type: to.Ptr(armcosmos.RoleDefinitionTypeCustomRole),
			AssignableScopes: []*string{
				to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/sales"),
				to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases")},
			Permissions: []*armcosmos.Permission{
				{
					DataActions: []*string{
						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/create"),
						to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read")},
					NotDataActions: []*string{},
				}},
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
	// TODO: use response item
	_ = res
}
