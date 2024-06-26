package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/GetAlertDefinitions.json
func ExampleAlertDefinitionsClient_NewListForScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertDefinitionsClient().NewListForScopePager("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", nil)
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
		// page.AlertDefinitionListResult = armauthorization.AlertDefinitionListResult{
		// 	Value: []*armauthorization.AlertDefinition{
		// 		{
		// 			Name: to.Ptr("TooManyPermanentOwnersAssignedToResource"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleManagementAlertDefinitions"),
		// 			ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertDefinitions/TooManyPermanentOwnersAssignedToResource"),
		// 			Properties: &armauthorization.AlertDefinitionProperties{
		// 				Description: to.Ptr("The number of users set to never expire is too high. To enhance the security of your resources, we recommend requiring activation for role use. Take a moment to review the list of users, and suggested changes here."),
		// 				DisplayName: to.Ptr("Too many permanent owners assigned to a resource"),
		// 				HowToPrevent: to.Ptr("Enable “Activation Required” in the role settings menu. This will ensure newly added users must activate their role."),
		// 				IsConfigurable: to.Ptr(true),
		// 				IsRemediatable: to.Ptr(true),
		// 				MitigationSteps: to.Ptr("To mitigate this issue, require the user to activate the role before use."),
		// 				Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 				SecurityImpact: to.Ptr("Providing users permanent access in a role may leave resources vulnerable to accidental or malicious activity."),
		// 				SeverityLevel: to.Ptr(armauthorization.SeverityLevelMedium),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TooManyOwnersAssignedToResource"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleManagementAlertDefinitions"),
		// 			ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertDefinitions/TooManyOwnersAssignedToResource"),
		// 			Properties: &armauthorization.AlertDefinitionProperties{
		// 				Description: to.Ptr("The number of users with the Owner role is too high. We recommend assigning these individuals to less privileged roles or roles more suitable to their daily needs. Take a moment to review the current assignments, and suggested changes here."),
		// 				DisplayName: to.Ptr("Too many owners assigned to a resource"),
		// 				HowToPrevent: to.Ptr("Choose a role that provides the fewest privileges necessary for a user or group to complete their tasks."),
		// 				IsConfigurable: to.Ptr(true),
		// 				IsRemediatable: to.Ptr(true),
		// 				MitigationSteps: to.Ptr("To mitigate this issue, reduce the number of users in the Owner role. Review the list of users in the list, and reassign them to a less privileged role such as Contributor."),
		// 				Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 				SecurityImpact: to.Ptr("As the number of users with the owner role increases, so does the potential for malicious or mistaken actions affecting your resource."),
		// 				SeverityLevel: to.Ptr(armauthorization.SeverityLevelMedium),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DuplicateRoleCreated"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleManagementAlertDefinitions"),
		// 			ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertDefinitions/DuplicateRoleCreated"),
		// 			Properties: &armauthorization.AlertDefinitionProperties{
		// 				Description: to.Ptr("One or more custom roles have the same display name and/or permissions as a built-in or preexisting custom role. Please review the newly created role membership assignments and determine the appropriate action."),
		// 				DisplayName: to.Ptr("Duplicate role created"),
		// 				HowToPrevent: to.Ptr("Prior to creating a custom role, determine if a built-in or preexisting custom role aligns to your security requirements."),
		// 				IsConfigurable: to.Ptr(false),
		// 				IsRemediatable: to.Ptr(true),
		// 				MitigationSteps: to.Ptr("To mitigate this issue, review the newly created role, and determine if a built-in role is suitable."),
		// 				Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 				SecurityImpact: to.Ptr("Duplicate roles add confusion and increases the complexity of administration."),
		// 				SeverityLevel: to.Ptr(armauthorization.SeverityLevelMedium),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureRolesAssignedOutsidePimAlert"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleManagementAlertDefinitions"),
		// 			ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertDefinitions/AzureRolesAssignedOutsidePimAlert"),
		// 			Properties: &armauthorization.AlertDefinitionProperties{
		// 				Description: to.Ptr("2 privileged assignment(s) were made outisde of Azure AD PIM"),
		// 				DisplayName: to.Ptr("Roles are being assigned outside of Privileged Identity Management"),
		// 				HowToPrevent: to.Ptr("Investigate where users are being assigned privileged roles outside of Privileged Identity Management and prohibit future assignments from there."),
		// 				IsConfigurable: to.Ptr(false),
		// 				IsRemediatable: to.Ptr(true),
		// 				MitigationSteps: to.Ptr("Review the users in the list and remove them from privileged roles assigned outside of Privileged Identity Management."),
		// 				Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
		// 				SecurityImpact: to.Ptr("Privileged role assignments made outside of Privileged Identity Management are not properly monitored and may indicate an active attack."),
		// 				SeverityLevel: to.Ptr(armauthorization.SeverityLevelHigh),
		// 			},
		// 	}},
		// }
	}
}
