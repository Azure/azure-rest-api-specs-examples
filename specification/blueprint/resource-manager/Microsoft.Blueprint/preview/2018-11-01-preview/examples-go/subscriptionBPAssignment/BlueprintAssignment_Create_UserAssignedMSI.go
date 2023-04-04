package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignment_Create_UserAssignedMSI.json
func ExampleAssignmentsClient_CreateOrUpdate_assignmentWithUserAssignedManagedIdentityAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().CreateOrUpdate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", armblueprint.Assignment{
		Location: to.Ptr("eastus"),
		Identity: &armblueprint.ManagedServiceIdentity{
			Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armblueprint.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity": {},
			},
		},
		Properties: &armblueprint.AssignmentProperties{
			Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
			BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
			Parameters: map[string]*armblueprint.ParameterValue{
				"costCenter": {
					Value: "Contoso/Online/Shopping/Production",
				},
				"owners": {
					Value: []any{
						"johnDoe@contoso.com",
						"johnsteam@contoso.com",
					},
				},
				"storageAccountType": {
					Value: "Standard_LRS",
				},
			},
			ResourceGroups: map[string]*armblueprint.ResourceGroupValue{
				"storageRG": {
					Name:     to.Ptr("defaultRG"),
					Location: to.Ptr("eastus"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
