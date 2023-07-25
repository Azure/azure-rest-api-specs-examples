package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2020-10-01-preview/examples/PutRoleManagementPolicyAssignment.json
func ExampleRoleManagementPolicyAssignmentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRoleManagementPolicyAssignmentsClient().Create(ctx, "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368", "b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24", armauthorization.RoleManagementPolicyAssignment{
		Properties: &armauthorization.RoleManagementPolicyAssignmentProperties{
			PolicyID:         to.Ptr("/providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
			RoleDefinitionID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
			Scope:            to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
