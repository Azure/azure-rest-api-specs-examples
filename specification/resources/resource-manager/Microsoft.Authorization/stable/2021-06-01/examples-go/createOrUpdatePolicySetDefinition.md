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
	}
	ctx := context.Background()
	client, err := armpolicy.NewSetDefinitionsClient("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"CostManagement",
		armpolicy.SetDefinition{
			Properties: &armpolicy.SetDefinitionProperties{
				Description: to.Ptr("Policies to enforce low cost storage SKUs"),
				DisplayName: to.Ptr("Cost Management"),
				Metadata: map[string]interface{}{
					"category": "Cost Management",
				},
				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
					"namePrefix": {
						Type:         to.Ptr(armpolicy.ParameterTypeString),
						DefaultValue: "myPrefix",
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							DisplayName: to.Ptr("Prefix to enforce on resource names"),
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
						PolicyDefinitionID:          to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
						PolicyDefinitionReferenceID: to.Ptr("Limit_Skus"),
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
						PolicyDefinitionID:          to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
						PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.6.0/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.
