package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/ARMTemplateArtifact_Create.json
func ExampleArtifactsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewArtifactsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"providers/Microsoft.Management/managementGroups/ContosoOnlineGroup",
		"simpleBlueprint",
		"storageTemplate",
		&armblueprint.TemplateArtifact{
			Kind: to.Ptr(armblueprint.ArtifactKindTemplate),
			Properties: &armblueprint.TemplateArtifactProperties{
				Parameters: map[string]*armblueprint.ParameterValue{
					"storageAccountType": {
						Value: "[parameters('storageAccountType')]",
					},
				},
				ResourceGroup: to.Ptr("storageRG"),
				Template: map[string]interface{}{
					"contentVersion": "1.0.0.0",
					"outputs": map[string]interface{}{
						"storageAccountName": map[string]interface{}{
							"type":  "string",
							"value": "[variables('storageAccountName')]",
						},
					},
					"parameters": map[string]interface{}{
						"storageAccountType": map[string]interface{}{
							"type": "string",
							"allowedValues": []interface{}{
								"Standard_LRS",
								"Standard_GRS",
								"Standard_ZRS",
								"Premium_LRS",
							},
							"defaultValue": "Standard_LRS",
							"metadata": map[string]interface{}{
								"description": "Storage Account type",
							},
						},
					},
					"resources": []interface{}{
						map[string]interface{}{
							"name":       "[variables('storageAccountName')]",
							"type":       "Microsoft.Storage/storageAccounts",
							"apiVersion": "2016-01-01",
							"kind":       "Storage",
							"location":   "[resourceGroup().location]",
							"properties": map[string]interface{}{},
							"sku": map[string]interface{}{
								"name": "[parameters('storageAccountType')]",
							},
						},
					},
					"variables": map[string]interface{}{
						"storageAccountName": "[concat(uniquestring(resourceGroup().id), 'standardsa')]",
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
