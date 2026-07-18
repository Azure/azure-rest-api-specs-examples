package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBSqlRoleAssignmentCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateSQLRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginCreateUpdateSQLRoleAssignment(ctx, "myRoleAssignmentId", "myResourceGroupName", "myAccountName", armcosmos.SQLRoleAssignmentCreateUpdateParameters{
		Properties: &armcosmos.SQLRoleAssignmentResource{
			RoleDefinitionID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId"),
			Scope:            to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
			PrincipalID:      to.Ptr("myPrincipalId"),
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
	// res = armcosmos.SQLResourcesClientCreateUpdateSQLRoleAssignmentResponse{
	// 	SQLRoleAssignmentGetResults: armcosmos.SQLRoleAssignmentGetResults{
	// 		ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleAssignments/myRoleAssignmentId"),
	// 		Name: to.Ptr("myRoleAssignmentId"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"),
	// 		Properties: &armcosmos.SQLRoleAssignmentResource{
	// 			RoleDefinitionID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId"),
	// 			Scope: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
	// 			PrincipalID: to.Ptr("myPrincipalId"),
	// 		},
	// 	},
	// }
}
