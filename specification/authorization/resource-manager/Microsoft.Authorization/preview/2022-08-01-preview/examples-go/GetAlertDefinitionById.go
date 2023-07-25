package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/GetAlertDefinitionById.json
func ExampleAlertDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertDefinitionsClient().Get(ctx, "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", "TooManyPermanentOwnersAssignedToResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertDefinition = armauthorization.AlertDefinition{
	// 	Name: to.Ptr("TooManyPermanentOwnersAssignedToResource"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleManagementAlertDefinitions"),
	// 	ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlertDefinitions/TooManyPermanentOwnersAssignedToResource"),
	// 	Properties: &armauthorization.AlertDefinitionProperties{
	// 		Description: to.Ptr("The number of users set to never expire is too high. To enhance the security of your resources, we recommend requiring activation for role use. Take a moment to review the list of users, and suggested changes here."),
	// 		DisplayName: to.Ptr("Too many permanent owners assigned to a resource"),
	// 		HowToPrevent: to.Ptr("Enable “Activation Required” in the role settings menu. This will ensure newly added users must activate their role."),
	// 		IsConfigurable: to.Ptr(true),
	// 		IsRemediatable: to.Ptr(true),
	// 		MitigationSteps: to.Ptr("To mitigate this issue, require the user to activate the role before use."),
	// 		Scope: to.Ptr("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f"),
	// 		SecurityImpact: to.Ptr("Providing users permanent access in a role may leave resources vulnerable to accidental or malicious activity."),
	// 		SeverityLevel: to.Ptr(armauthorization.SeverityLevelMedium),
	// 	},
	// }
}
