package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/updatePolicyAssignmentWithIdentityById.json
func ExamplePolicyAssignmentsClient_UpdateByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewPolicyAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateByID(ctx,
		"<policy-assignment-id>",
		armpolicy.PolicyAssignmentUpdate{
			Identity: &armpolicy.Identity{
				Type: armpolicy.ResourceIdentityTypeSystemAssigned.ToPtr(),
			},
			Location: to.StringPtr("<location>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PolicyAssignment.ID: %s\n", *res.ID)
}
