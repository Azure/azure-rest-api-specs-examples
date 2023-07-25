package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2020-10-01-preview/examples/GetRoleManagementPolicyAssignmentByName.json
func ExampleRoleManagementPolicyAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleManagementPolicyAssignmentsClient().Get(ctx, "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368", "b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoleManagementPolicyAssignment = armauthorization.RoleManagementPolicyAssignment{
	// 	Name: to.Ptr("b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 	Type: to.Ptr("Microsoft.Authorization/RoleManagementPolicyAssignment"),
	// 	ID: to.Ptr("/providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicyAssignment/b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 	Properties: &armauthorization.RoleManagementPolicyAssignmentProperties{
	// 		PolicyAssignmentProperties: &armauthorization.PolicyAssignmentProperties{
	// 			Policy: &armauthorization.PolicyAssignmentPropertiesPolicy{
	// 				ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
	// 				LastModifiedBy: &armauthorization.Principal{
	// 					DisplayName: to.Ptr("Admin"),
	// 				},
	// 			},
	// 			RoleDefinition: &armauthorization.PolicyAssignmentPropertiesRoleDefinition{
	// 				Type: to.Ptr("BuiltInRole"),
	// 				DisplayName: to.Ptr("FHIR Data Converter"),
	// 				ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 			},
	// 			Scope: &armauthorization.PolicyAssignmentPropertiesScope{
	// 				Type: to.Ptr("subscription"),
	// 				DisplayName: to.Ptr("Pay-As-You-Go"),
	// 				ID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
	// 			},
	// 		},
	// 		PolicyID: to.Ptr("/providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
	// 		RoleDefinitionID: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
	// 		Scope: to.Ptr("/subscriptions/129ff972-28f8-46b8-a726-e497be039368"),
	// 	},
	// }
}
