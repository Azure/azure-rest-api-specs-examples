package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Create_SystemAssignedMSI.json
func ExampleAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"managementGroups/ContosoOnlineGroup",
		"assignSimpleBlueprint",
		armblueprint.Assignment{
			Location: to.Ptr("eastus"),
			Identity: &armblueprint.ManagedServiceIdentity{
				Type: to.Ptr(armblueprint.ManagedServiceIdentityTypeSystemAssigned),
			},
			Properties: &armblueprint.AssignmentProperties{
				Description: to.Ptr("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
				BlueprintID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
				Parameters: map[string]*armblueprint.ParameterValue{
					"costCenter": {
						Value: "Contoso/Online/Shopping/Production",
					},
					"owners": {
						Value: []interface{}{
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
				Scope: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
