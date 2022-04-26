Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.5.0/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicySetDefinition.json
func ExampleSetDefinitionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpolicy.NewSetDefinitionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<policy-set-definition-name>",
		armpolicy.SetDefinition{
			Properties: &armpolicy.SetDefinitionProperties{
				Description: to.Ptr("<description>"),
				DisplayName: to.Ptr("<display-name>"),
				Metadata: map[string]interface{}{
					"category": "Cost Management",
				},
				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
					"namePrefix": {
						Type:         to.Ptr(armpolicy.ParameterTypeString),
						DefaultValue: "myPrefix",
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							DisplayName: to.Ptr("<display-name>"),
						},
					},
				},
				PolicyDefinitions: []*armpolicy.DefinitionReference{
					{
						Parameters: map[string]*armpolicy.ParameterValuesValue{
							"listOfAllowedSKUs": {
								Value: []interface{}{
									"Standard_GRS",
									"Standard_LRS",
								},
							},
						},
						PolicyDefinitionID:          to.Ptr("<policy-definition-id>"),
						PolicyDefinitionReferenceID: to.Ptr("<policy-definition-reference-id>"),
					},
					{
						Parameters: map[string]*armpolicy.ParameterValuesValue{
							"prefix": {
								Value: "[parameters('namePrefix')]",
							},
							"suffix": {
								Value: "-LC",
							},
						},
						PolicyDefinitionID:          to.Ptr("<policy-definition-id>"),
						PolicyDefinitionReferenceID: to.Ptr("<policy-definition-reference-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
