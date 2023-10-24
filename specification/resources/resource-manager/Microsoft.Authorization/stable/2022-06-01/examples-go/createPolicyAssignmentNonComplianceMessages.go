package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignmentNonComplianceMessages.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithMultipleNonComplianceMessages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Create(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "securityInitAssignment", armpolicy.Assignment{
		Properties: &armpolicy.AssignmentProperties{
			DisplayName: to.Ptr("Enforce security policies"),
			NonComplianceMessages: []*armpolicy.NonComplianceMessage{
				{
					Message: to.Ptr("Resources must comply with all internal security policies. See <internal site URL> for more info."),
				},
				{
					Message:                     to.Ptr("Resource names must start with 'DeptA' and end with '-LC'."),
					PolicyDefinitionReferenceID: to.Ptr("10420126870854049575"),
				},
				{
					Message:                     to.Ptr("Storage accounts must have firewall rules configured."),
					PolicyDefinitionReferenceID: to.Ptr("8572513655450389710"),
				}},
			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/securityInitiative"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
