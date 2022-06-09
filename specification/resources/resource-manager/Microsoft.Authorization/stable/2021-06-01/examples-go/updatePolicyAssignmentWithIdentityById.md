```go
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmpolicy%2Fv0.1.1/sdk/resourcemanager/resources/armpolicy/README.md) on how to add the SDK to your project and authenticate.
