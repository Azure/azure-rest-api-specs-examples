package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPDef/Artifact_List.json
func ExampleArtifactsClient_NewListPager_subArtifactList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArtifactsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000", "simpleBlueprint", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ArtifactList = armblueprint.ArtifactList{
		// 	Value: []armblueprint.ArtifactClassification{
		// 		&armblueprint.RoleAssignmentArtifact{
		// 			Name: to.Ptr("ownerAssignment"),
		// 			Type: to.Ptr("Microsoft.Blueprint/blueprints/artifacts"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprints/simpleBlueprint/artifacts/ownerAssignment"),
		// 			Kind: to.Ptr(armblueprint.ArtifactKindRoleAssignment),
		// 			Properties: &armblueprint.RoleAssignmentArtifactProperties{
		// 				DisplayName: to.Ptr("enforce owners of given subscription"),
		// 				PrincipalIDs: "[parameters('owners')]",
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Authorization/roleDefinitions/acdd72a7-3385-48ef-bd42-f606fba81ae7"),
		// 			},
		// 		},
		// 		&armblueprint.PolicyAssignmentArtifact{
		// 			Name: to.Ptr("costCenterPolicy"),
		// 			Type: to.Ptr("Microsoft.Blueprint/blueprints/artifacts"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprints/simpleBlueprint/artifacts/costCenterPolicy"),
		// 			Kind: to.Ptr(armblueprint.ArtifactKindPolicyAssignment),
		// 			Properties: &armblueprint.PolicyAssignmentArtifactProperties{
		// 				DisplayName: to.Ptr("force costCenter tag on all resources"),
		// 				Parameters: map[string]*armblueprint.ParameterValue{
		// 					"tagName": &armblueprint.ParameterValue{
		// 						Value: "costCenter",
		// 					},
		// 					"tagValue": &armblueprint.ParameterValue{
		// 						Value: "[parameter('costCenter')]",
		// 					},
		// 				},
		// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/1e30110a-5ceb-460c-a204-c1c3969c6d62"),
		// 			},
		// 		},
		// 		&armblueprint.TemplateArtifact{
		// 			Name: to.Ptr("storageTemplate"),
		// 			Type: to.Ptr("Microsoft.Blueprint/blueprints/artifacts"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprints/simpleBlueprint/artifacts/storageTemplate"),
		// 			Kind: to.Ptr(armblueprint.ArtifactKindTemplate),
		// 			Properties: &armblueprint.TemplateArtifactProperties{
		// 				Parameters: map[string]*armblueprint.ParameterValue{
		// 					"storageAccountType": &armblueprint.ParameterValue{
		// 						Value: "[parameters('storageAccountType')]",
		// 					},
		// 				},
		// 				ResourceGroup: to.Ptr("storageRG"),
		// 				Template: map[string]any{
		// 					"contentVersion": "1.0.0.0",
		// 					"outputs":map[string]any{
		// 						"storageAccountName":map[string]any{
		// 							"type": "string",
		// 							"value": "[variables('storageAccountName')]",
		// 						},
		// 					},
		// 					"parameters":map[string]any{
		// 						"storageAccountType":map[string]any{
		// 							"type": "string",
		// 							"allowedValues":[]any{
		// 								"Standard_LRS",
		// 								"Standard_GRS",
		// 								"Standard_ZRS",
		// 								"Premium_LRS",
		// 							},
		// 							"defaultValue": "Standard_LRS",
		// 							"metadata":map[string]any{
		// 								"description": "Storage Account type",
		// 							},
		// 						},
		// 					},
		// 					"resources":[]any{
		// 						map[string]any{
		// 							"name": "[variables('storageAccountName')]",
		// 							"type": "Microsoft.Storage/storageAccounts",
		// 							"apiVersion": "2016-01-01",
		// 							"kind": "Storage",
		// 							"location": "[resourceGroup().location]",
		// 							"properties":map[string]any{
		// 							},
		// 							"sku":map[string]any{
		// 								"name": "[parameters('storageAccountType')]",
		// 							},
		// 						},
		// 					},
		// 					"variables":map[string]any{
		// 						"storageAccountName": "[concat(uniquestring(resourceGroup().id), 'standardsa')]",
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
