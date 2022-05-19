Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblueprint%2Farmblueprint%2Fv0.5.0/sdk/resourcemanager/blueprint/armblueprint/README.md) on how to add the SDK to your project and authenticate.

```go
package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/Blueprint_Create.json
func ExampleBlueprintsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewBlueprintsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"providers/Microsoft.Management/managementGroups/ContosoOnlineGroup",
		"simpleBlueprint",
		armblueprint.Blueprint{
			Properties: &armblueprint.Properties{
				Description: to.Ptr("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
				Parameters: map[string]*armblueprint.ParameterDefinition{
					"costCenter": {
						Type: to.Ptr(armblueprint.TemplateParameterTypeString),
						Metadata: &armblueprint.ParameterDefinitionMetadata{
							DisplayName: to.Ptr("force cost center tag for all resources under given subscription."),
						},
					},
					"owners": {
						Type: to.Ptr(armblueprint.TemplateParameterTypeArray),
						Metadata: &armblueprint.ParameterDefinitionMetadata{
							DisplayName: to.Ptr("assign owners to subscription along with blueprint assignment."),
						},
					},
					"storageAccountType": {
						Type: to.Ptr(armblueprint.TemplateParameterTypeString),
						Metadata: &armblueprint.ParameterDefinitionMetadata{
							DisplayName: to.Ptr("storage account type."),
						},
					},
				},
				ResourceGroups: map[string]*armblueprint.ResourceGroupDefinition{
					"storageRG": {
						Metadata: &armblueprint.ParameterDefinitionMetadata{
							Description: to.Ptr("Contains storageAccounts that collect all shoebox logs."),
							DisplayName: to.Ptr("storage resource group"),
						},
					},
				},
				TargetScope: to.Ptr(armblueprint.BlueprintTargetScopeSubscription),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
