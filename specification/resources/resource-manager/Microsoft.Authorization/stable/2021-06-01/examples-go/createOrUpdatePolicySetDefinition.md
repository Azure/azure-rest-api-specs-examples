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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicySetDefinition.json
func ExampleSetDefinitionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewSetDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<policy-set-definition-name>",
		armpolicy.SetDefinition{
			Properties: &armpolicy.SetDefinitionProperties{
				Description: to.StringPtr("<description>"),
				DisplayName: to.StringPtr("<display-name>"),
				Metadata: map[string]interface{}{
					"category": "Cost Management",
				},
				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
					"namePrefix": {
						Type: armpolicy.ParameterType("String").ToPtr(),
						DefaultValue: map[string]interface{}{
							"0": "m",
							"1": "y",
							"2": "P",
							"3": "r",
							"4": "e",
							"5": "f",
							"6": "i",
							"7": "x",
						},
						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
							DisplayName: to.StringPtr("<display-name>"),
						},
					},
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
									"13": "n",
									"14": "a",
									"15": "m",
									"16": "e",
									"17": "P",
									"18": "r",
									"19": "e",
									"20": "f",
									"21": "i",
									"22": "x",
									"23": "'",
									"24": ")",
									"25": "]",
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
	log.Printf("Response result: %#v\n", res.SetDefinitionsClientCreateOrUpdateResult)
}
```
