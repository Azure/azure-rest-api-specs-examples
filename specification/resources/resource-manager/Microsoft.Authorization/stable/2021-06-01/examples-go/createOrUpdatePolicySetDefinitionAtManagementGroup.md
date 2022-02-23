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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicySetDefinitionAtManagementGroup.json
func ExampleSetDefinitionsClient_CreateOrUpdateAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewSetDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateAtManagementGroup(ctx,
		"<policy-set-definition-name>",
		"<management-group-id>",
		armpolicy.SetDefinition{
			Properties: &armpolicy.SetDefinitionProperties{
				Description: to.StringPtr("<description>"),
				DisplayName: to.StringPtr("<display-name>"),
				Metadata: map[string]interface{}{
					"category": "Cost Management",
				},
				PolicyDefinitions: []*armpolicy.DefinitionReference{
					{
						Parameters: map[string]*armpolicy.ParameterValuesValue{
							"listOfAllowedSKUs": {
								Value: map[string]interface{}{
									"0": "Standard_GRS",
									"1": "Standard_LRS",
								},
							},
						},
						PolicyDefinitionID:          to.StringPtr("<policy-definition-id>"),
						PolicyDefinitionReferenceID: to.StringPtr("<policy-definition-reference-id>"),
					},
					{
						Parameters: map[string]*armpolicy.ParameterValuesValue{
							"prefix": {
								Value: map[string]interface{}{
									"0": "D",
									"1": "e",
									"2": "p",
									"3": "t",
									"4": "A",
								},
							},
							"suffix": {
								Value: map[string]interface{}{
									"0": "-",
									"1": "L",
									"2": "C",
								},
							},
						},
						PolicyDefinitionID:          to.StringPtr("<policy-definition-id>"),
						PolicyDefinitionReferenceID: to.StringPtr("<policy-definition-reference-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SetDefinitionsClientCreateOrUpdateAtManagementGroupResult)
}
```
