Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.3.1/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinition.json
func ExampleDefinitionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	_, err = client.CreateOrUpdate(ctx,
		"<policy-definition-name>",
		armpolicy.Definition{
			Properties: &armpolicy.DefinitionProperties{
				Description: to.StringPtr("<description>"),
				DisplayName: to.StringPtr("<display-name>"),
				Metadata: map[string]interface{}{
					"category": "Naming",
				},
				Mode: to.StringPtr("<mode>"),
				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
					"prefix": {
						Type: armpolicy.ParameterType("String").ToPtr(),
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							Description: to.StringPtr("<description>"),
							DisplayName: to.StringPtr("<display-name>"),
						},
					},
					"suffix": {
						Type: armpolicy.ParameterType("String").ToPtr(),
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							Description: to.StringPtr("<description>"),
							DisplayName: to.StringPtr("<display-name>"),
						},
					},
				},
				PolicyRule: map[string]interface{}{
					"if": map[string]interface{}{
						"not": map[string]interface{}{
							"field": "name",
							"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
						},
					},
					"then": map[string]interface{}{
						"effect": "deny",
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
