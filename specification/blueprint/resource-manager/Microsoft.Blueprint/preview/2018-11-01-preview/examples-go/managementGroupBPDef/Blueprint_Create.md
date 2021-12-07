Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblueprint%2Farmblueprint%2Fv0.1.0/sdk/resourcemanager/blueprint/armblueprint/README.md) on how to add the SDK to your project and authenticate.

```go
package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// x-ms-original-file: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Blueprint_Create.json
func ExampleBlueprintsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblueprint.NewBlueprintsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-scope>",
		"<blueprint-name>",
		armblueprint.Blueprint{
			Properties: &armblueprint.BlueprintProperties{
				SharedBlueprintProperties: armblueprint.SharedBlueprintProperties{
					BlueprintResourcePropertiesBase: armblueprint.BlueprintResourcePropertiesBase{
						Description: to.StringPtr("<description>"),
					},
					Parameters: map[string]*armblueprint.ParameterDefinition{
						"costCenter": {
							Type: armblueprint.TemplateParameterTypeString.ToPtr(),
							Metadata: &armblueprint.ParameterDefinitionMetadata{
								DisplayName: to.StringPtr("<display-name>"),
							},
						},
						"owners": {
							Type: armblueprint.TemplateParameterTypeArray.ToPtr(),
							Metadata: &armblueprint.ParameterDefinitionMetadata{
								DisplayName: to.StringPtr("<display-name>"),
							},
						},
						"storageAccountType": {
							Type: armblueprint.TemplateParameterTypeString.ToPtr(),
							Metadata: &armblueprint.ParameterDefinitionMetadata{
								DisplayName: to.StringPtr("<display-name>"),
							},
						},
					},
					ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
						"storageRG": {
							Metadata: &armblueprint.ParameterDefinitionMetadata{
								Description: to.StringPtr("<description>"),
								DisplayName: to.StringPtr("<display-name>"),
							},
						},
					},
					TargetScope: armblueprint.BlueprintTargetScopeSubscription.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Blueprint.ID: %s\n", *res.ID)
}
```
