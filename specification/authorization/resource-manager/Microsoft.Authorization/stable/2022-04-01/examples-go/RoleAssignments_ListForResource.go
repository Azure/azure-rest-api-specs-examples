package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_ListForResource.json
func ExampleRoleAssignmentsClient_NewListForResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleAssignmentsClient().NewListForResourcePager("testrg", "Microsoft.DocumentDb", "databaseAccounts", "test-db-account", &armauthorization.RoleAssignmentsClientListForResourceOptions{Filter: nil,
		TenantID: nil,
	})
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
		// page.RoleAssignmentListResult = armauthorization.RoleAssignmentListResult{
		// 	Value: []*armauthorization.RoleAssignment{
		// 		{
		// 			Name: to.Ptr("b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/b0f43c54-e787-4862-89b1-a653fa9cf747"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("96786e4b-dede-4c2e-8736-8ab911987f08"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.Authorization/roleAssignments/96786e4b-dede-4c2e-8736-8ab911987f08"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("05c5a614-a7d6-4502-b150-c2fb455033ff"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
		// 			ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account/providers/Microsoft.Authorization/roleAssignments/05c5a614-a7d6-4502-b150-c2fb455033ff"),
		// 			Properties: &armauthorization.RoleAssignmentProperties{
		// 				PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
		// 				PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		// 				Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg/providers/Microsoft.DocumentDb/databaseAccounts/test-db-account"),
		// 			},
		// 	}},
		// }
	}
}
