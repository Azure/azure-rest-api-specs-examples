package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/SealedPolicyAssignmentArtifact_Get.json
func ExamplePublishedArtifactsClient_Get_mgPolicyAssignmentArtifact() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublishedArtifactsClient().Get(ctx, "providers/Microsoft.Management/managementGroups/ContosoOnlineGroup", "simpleBlueprint", "V2", "costCenterPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armblueprint.PublishedArtifactsClientGetResponse{
	// 	                            ArtifactClassification: &armblueprint.PolicyAssignmentArtifact{
	// 		Name: to.Ptr("costCenterPolicy"),
	// 		Type: to.Ptr("Microsoft.Blueprint/blueprints/versions/artifacts"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint/versions/V2/artifacts/costCenterPolicy"),
	// 		Kind: to.Ptr(armblueprint.ArtifactKindPolicyAssignment),
	// 		Properties: &armblueprint.PolicyAssignmentArtifactProperties{
	// 			DisplayName: to.Ptr("force costCenter tag on all resources"),
	// 			Parameters: map[string]*armblueprint.ParameterValue{
	// 				"tagName": &armblueprint.ParameterValue{
	// 					Value: "costCenter",
	// 				},
	// 				"tagValue": &armblueprint.ParameterValue{
	// 					Value: "[parameter('costCenter')]",
	// 				},
	// 			},
	// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/1e30110a-5ceb-460c-a204-c1c3969c6d62"),
	// 		},
	// 	},
	// 	                        }
}
