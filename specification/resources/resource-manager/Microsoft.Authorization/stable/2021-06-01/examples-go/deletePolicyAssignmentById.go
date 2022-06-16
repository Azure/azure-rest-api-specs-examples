package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/deletePolicyAssignmentById.json
func ExamplePolicyAssignmentsClient_DeleteByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicy.NewPolicyAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.DeleteByID(ctx,
		"<policy-assignment-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PolicyAssignment.ID: %s\n", *res.ID)
}
