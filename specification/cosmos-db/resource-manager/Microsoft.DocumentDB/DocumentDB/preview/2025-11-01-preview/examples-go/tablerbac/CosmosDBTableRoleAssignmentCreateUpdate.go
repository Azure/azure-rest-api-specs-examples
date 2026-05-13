package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/tablerbac/CosmosDBTableRoleAssignmentCreateUpdate.json
func ExampleTableResourcesClient_BeginCreateUpdateTableRoleAssignment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTableResourcesClient().BeginCreateUpdateTableRoleAssignment(ctx, "myResourceGroupName", "myAccountName", "myRoleAssignmentId", armcosmos.TableRoleAssignmentResource{
		Properties: &armcosmos.TableRoleAssignmentResourceProperties{
			PrincipalID:      to.Ptr("myPrincipalId"),
			RoleDefinitionID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/tableRoleDefinitions/myRoleDefinitionId"),
			Scope:            to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
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
	// res = armcosmos.TableResourcesClientCreateUpdateTableRoleAssignmentResponse{
	// 	TableRoleAssignmentResource: &armcosmos.TableRoleAssignmentResource{
	// 		Name: to.Ptr("myRoleAssignmentId"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/tableRoleAssignments"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/tableRoleAssignments/myRoleAssignmentId"),
	// 		Properties: &armcosmos.TableRoleAssignmentResourceProperties{
	// 			PrincipalID: to.Ptr("myPrincipalId"),
	// 			RoleDefinitionID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/tableRoleDefinitions/myRoleDefinitionId"),
	// 			Scope: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
	// 		},
	// 	},
	// }
}
