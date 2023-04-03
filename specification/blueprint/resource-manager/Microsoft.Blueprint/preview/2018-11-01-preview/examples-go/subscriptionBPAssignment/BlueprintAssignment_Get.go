package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_Get.json
func ExampleAssignmentsClient_Get_assignmentAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().Get(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assignment = armblueprint.Assignment{
	// 	Name: to.Ptr("assignSimpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/Assignment"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint"),
	// 	Location: to.Ptr("eastus"),
	// 	Identity: &armblueprint.ManagedServiceIdentity{
	// 		Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	},
	// 	Properties: &armblueprint.AssignmentProperties{
	// 		Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
	// 		BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 		Parameters: map[string]*armblueprint.ParameterValue{
	// 			"costCenter": &armblueprint.ParameterValue{
	// 				Value: "Contoso/Online/Shopping/Production",
	// 			},
	// 			"owners": &armblueprint.ParameterValue{
	// 				Value: []any{
	// 					"johnDoe@contoso.com",
	// 					"johnsteam@contoso.com",
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterValue{
	// 				Value: "Standard_LRS",
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armblueprint.AssignmentProvisioningState("Succeeded")),
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
	// 			"storageRG": &armblueprint.ResourceGroupValue{
	// 				Name: to.Ptr("defaultRG"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 		},
	// 	},
	// }
}
