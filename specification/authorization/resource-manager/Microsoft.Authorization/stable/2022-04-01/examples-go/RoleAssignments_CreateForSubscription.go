package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_CreateForSubscription.json
func ExampleRoleAssignmentsClient_Create_createRoleAssignmentForSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleAssignmentsClient().Create(ctx, "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2", "05c5a614-a7d6-4502-b150-c2fb455033ff", armauthorization.RoleAssignmentCreateParameters{
		Properties: &armauthorization.RoleAssignmentProperties{
			PrincipalID:      to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
			PrincipalType:    to.Ptr(armauthorization.PrincipalTypeUser),
			RoleDefinitionID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleAssignment = armauthorization.RoleAssignment{
	// 	Name: to.Ptr("05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleAssignments"),
	// 	ID: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleAssignments/05c5a614-a7d6-4502-b150-c2fb455033ff"),
	// 	Properties: &armauthorization.RoleAssignmentProperties{
	// 		PrincipalID: to.Ptr("ce2ce14e-85d7-4629-bdbc-454d0519d987"),
	// 		PrincipalType: to.Ptr(armauthorization.PrincipalTypeUser),
	// 		RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"),
	// 		Scope: to.Ptr("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2"),
	// 	},
	// }
}
