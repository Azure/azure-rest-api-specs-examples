package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/RoleAssignmentArtifact_Create.json
func ExampleArtifactsClient_CreateOrUpdate_mgRoleAssignmentArtifact() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewArtifactsClient().CreateOrUpdate(ctx, "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup", "simpleBlueprint", "ownerAssignment", &armblueprint.RoleAssignmentArtifact{
		Kind: to.Ptr(armblueprint.ArtifactKindRoleAssignment),
		Properties: &armblueprint.RoleAssignmentArtifactProperties{
			DisplayName:      to.Ptr("enforce owners of given subscription"),
			PrincipalIDs:     "[parameters('owners')]",
			RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
