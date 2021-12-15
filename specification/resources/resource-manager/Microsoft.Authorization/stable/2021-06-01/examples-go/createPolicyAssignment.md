Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.1.1/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createPolicyAssignment.json
func ExamplePolicyAssignmentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewPolicyAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<scope>",
		"<policy-assignment-name>",
		armpolicy.PolicyAssignment{
			Properties: &armpolicy.PolicyAssignmentProperties{
				Description: to.StringPtr("<description>"),
				DisplayName: to.StringPtr("<display-name>"),
				Metadata: map[string]interface{}{
					"assignedBy": "Special Someone",
				},
				NonComplianceMessages: []*armpolicy.NonComplianceMessage{
					{
						Message: to.StringPtr("<message>"),
					}},
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
				PolicyDefinitionID: to.StringPtr("<policy-definition-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PolicyAssignment.ID: %s\n", *res.ID)
}
```
