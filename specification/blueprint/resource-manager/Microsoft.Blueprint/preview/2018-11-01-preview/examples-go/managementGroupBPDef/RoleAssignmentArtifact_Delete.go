package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/RoleAssignmentArtifact_Delete.json
func ExampleArtifactsClient_Delete_mgRoleAssignmentArtifact() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactsClient().Delete(ctx, "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup", "simpleBlueprint", "ownerAssignment", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armblueprint.ArtifactsClientDeleteResponse{
	// 	                            ArtifactClassification: &armblueprint.RoleAssignmentArtifact{
	// 		Name: to.Ptr("ownerAssignment"),
	// 		Type: to.Ptr("Microsoft.Blueprint/blueprints/artifacts"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint/artifacts/ownerAssignment"),
	// 		Kind: to.Ptr(armblueprint.ArtifactKindRoleAssignment),
	// 		Properties: &armblueprint.RoleAssignmentArtifactProperties{
	// 			DisplayName: to.Ptr("enforce owners of given subscription"),
	// 			PrincipalIDs: "[parameters('owners')]",
	// 			RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7"),
	// 		},
	// 	},
	// 	                        }
}
