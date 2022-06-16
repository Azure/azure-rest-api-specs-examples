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
	}
	ctx := context.Background()
	client, err := armpolicy.NewAssignmentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx,
		"subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2",
		"EnforceNaming",
		armpolicy.Assignment{
			Properties: &armpolicy.AssignmentProperties{
				Description: to.Ptr("Force resource names to begin with given DeptA and end with -LC"),
				DisplayName: to.Ptr("Enforce resource naming rules"),
				Metadata: map[string]interface{}{
					"assignedBy": "Special Someone",
				},
				NonComplianceMessages: []*armpolicy.NonComplianceMessage{
					{
						Message: to.Ptr("Resource names must start with 'DeptA' and end with '-LC'."),
					}},
				Parameters: map[string]*armpolicy.ParameterValuesValue{
					"prefix": {
						Value: "DeptA",
					},
					"suffix": {
						Value: "-LC",
					},
				},
				PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
