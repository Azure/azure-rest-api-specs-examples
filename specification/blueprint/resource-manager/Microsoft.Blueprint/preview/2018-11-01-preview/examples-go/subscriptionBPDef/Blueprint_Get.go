package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/Blueprint_Get.json
func ExampleBlueprintsClient_Get_subscriptionBlueprint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlueprintsClient().Get(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", "simpleBlueprint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Blueprint = armblueprint.Blueprint{
	// 	Name: to.Ptr("simpleBlueprint"),
	// 	Type: to.Ptr("Microsoft.Blueprint/blueprints"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
	// 	Properties: &armblueprint.Properties{
	// 		Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
	// 		Parameters: map[string]*armblueprint.ParameterDefinition{
	// 			"costCenter": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
	// 				},
	// 			},
	// 			"owners": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
	// 				},
	// 			},
	// 			"storageAccountType": &armblueprint.ParameterDefinition{
	// 				Type: to.Ptr(armblueprint.TemplateParameterTypeString),
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					DisplayName: to.Ptr("storage account type."),
	// 				},
	// 			},
	// 		},
	// 		ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
	// 			"storageRG": &armblueprint.ResourceGroupDefinition{
	// 				Metadata: &armblueprint.ParameterDefinitionMetadata{
	// 					Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
	// 					DisplayName: to.Ptr("storage resource group"),
	// 				},
	// 			},
	// 		},
	// 		TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
	// 	},
	// }
}
