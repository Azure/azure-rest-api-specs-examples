package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/tablerbac/CosmosDBTableRoleAssignmentList.json
func ExampleTableResourcesClient_NewListTableRoleAssignmentsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTableResourcesClient().NewListTableRoleAssignmentsPager("myResourceGroupName", "myAccountName", nil)
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
		// page = armcosmos.TableResourcesClientListTableRoleAssignmentsResponse{
		// 	TableRoleAssignmentListResult: armcosmos.TableRoleAssignmentListResult{
		// 		Value: []*armcosmos.TableRoleAssignmentResource{
		// 			{
		// 				Name: to.Ptr("myRoleAssignmentId"),
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/tableRoleAssignments"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/tableRoleAssignments/myRoleAssignmentId"),
		// 				Properties: &armcosmos.TableRoleAssignmentResourceProperties{
		// 					PrincipalID: to.Ptr("myPrincipalId"),
		// 					RoleDefinitionID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/tableRoleDefinitions/myRoleDefinitionId"),
		// 					Scope: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
