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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createPolicyAssignment.json
func ExampleAssignmentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armpolicy.NewAssignmentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Create(ctx,
		"<scope>",
		"<policy-assignment-name>",
		armpolicy.Assignment{
			Properties: &armpolicy.AssignmentProperties{
				Description: to.Ptr("<description>"),
				DisplayName: to.Ptr("<display-name>"),
				Metadata: map[string]interface{}{
					"assignedBy": "Special Someone",
				},
				NonComplianceMessages: []*armpolicy.NonComplianceMessage{
					{
						Message: to.Ptr("<message>"),
					}},
				Parameters: map[string]*armpolicy.ParameterValuesValue{
					"prefix": {
						Value: "DeptA",
					},
					"suffix": {
						Value: "-LC",
					},
				},
				PolicyDefinitionID: to.Ptr("<policy-definition-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
