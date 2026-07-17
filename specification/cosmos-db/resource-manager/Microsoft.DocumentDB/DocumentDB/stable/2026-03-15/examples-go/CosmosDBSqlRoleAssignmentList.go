package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBSqlRoleAssignmentList.json
func ExampleSQLResourcesClient_NewListSQLRoleAssignmentsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLResourcesClient().NewListSQLRoleAssignmentsPager("myResourceGroupName", "myAccountName", nil)
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
		// page = armcosmos.SQLResourcesClientListSQLRoleAssignmentsResponse{
		// 	SQLRoleAssignmentListResult: armcosmos.SQLRoleAssignmentListResult{
		// 		Value: []*armcosmos.SQLRoleAssignmentGetResults{
		// 			{
		// 				ID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleAssignments/myRoleAssignmentId"),
		// 				Name: to.Ptr("myRoleAssignmentId"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"),
		// 				Properties: &armcosmos.SQLRoleAssignmentResource{
		// 					RoleDefinitionID: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId"),
		// 					Scope: to.Ptr("/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
		// 					PrincipalID: to.Ptr("myPrincipalId"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
