package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/PolicyAssignmentArtifact_Create.json
func ExampleArtifactsClient_CreateOrUpdate_subPolicyAssignmentArtifact() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewArtifactsClient().CreateOrUpdate(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "simpleBlueprint", "costCenterPolicy", &armblueprint.PolicyAssignmentArtifact{
		Kind: to.Ptr(armblueprint.ArtifactKindPolicyAssignment),
		Properties: &armblueprint.PolicyAssignmentArtifactProperties{
			DisplayName: to.Ptr("force costCenter tag on all resources"),
			Parameters: map[string]*armblueprint.ParameterValue{
				"tagName": {
					Value: "costCenter",
				},
				"tagValue": {
					Value: "[parameter('costCenter')]",
				},
			},
			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/1e30110a-5ceb-460c-a204-c1c3969c6d62"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
