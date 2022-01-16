Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblueprint%2Farmblueprint%2Fv0.2.0/sdk/resourcemanager/blueprint/armblueprint/README.md) on how to add the SDK to your project and authenticate.

```go
package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// x-ms-original-file: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/ARMTemplateArtifact_Create.json
func ExampleArtifactsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblueprint.NewArtifactsClient(cred, nil)
	_, err = client.CreateOrUpdate(ctx,
		"<resource-scope>",
		"<blueprint-name>",
		"<artifact-name>",
		&armblueprint.TemplateArtifact{
			Kind: armblueprint.ArtifactKind("template").ToPtr(),
			Properties: &armblueprint.TemplateArtifactProperties{
				Parameters: map[string]*armblueprint.ParameterValue{
					"storageAccountType": {
						Value: map[string]interface{}{
							"0":  "[",
							"1":  "p",
							"2":  "a",
							"3":  "r",
							"4":  "a",
							"5":  "m",
							"6":  "e",
							"7":  "t",
							"8":  "e",
							"9":  "r",
							"10": "s",
							"11": "(",
							"12": "'",
							"13": "s",
							"14": "t",
							"15": "o",
							"16": "r",
							"17": "a",
							"18": "g",
							"19": "e",
							"20": "A",
							"21": "c",
							"22": "c",
							"23": "o",
							"24": "u",
							"25": "n",
							"26": "t",
							"27": "T",
							"28": "y",
							"29": "p",
							"30": "e",
							"31": "'",
							"32": ")",
							"33": "]",
						},
					},
				},
				ResourceGroup: to.StringPtr("<resource-group>"),
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
		log.Fatal(err)
	}
}
```
