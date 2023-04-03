package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/ResourceGroupWithTags.json
func ExampleBlueprintsClient_CreateOrUpdate_resourceGroupWithTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBlueprintsClient().CreateOrUpdate(ctx, "providers/Microsoft.Management/managementGroups/{ManagementGroupId}", "simpleBlueprint", armblueprint.Blueprint{
		Properties: &armblueprint.Properties{
			Description: to.Ptr("An example blueprint containing an RG with two tags."),
			ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
				"myRGName": {
					Name:     to.Ptr("myRGName"),
					Location: to.Ptr("westus"),
					Metadata: &armblueprint.ParameterDefinitionMetadata{
						DisplayName: to.Ptr("My Resource Group"),
					},
					Tags: map[string]*string{
						"costcenter":  to.Ptr("123456"),
						"nameOnlyTag": to.Ptr(""),
					},
				},
			},
			TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
