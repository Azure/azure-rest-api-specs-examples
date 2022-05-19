Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagedservices%2Farmmanagedservices%2Fv0.5.0/sdk/resourcemanager/managedservices/armmanagedservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetRegistrationAssignment.json
func ExampleRegistrationAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagedservices.NewRegistrationAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"subscription/0afefe50-734e-4610-8a82-a144ahf49dea",
		"26c128c2-fefa-4340-9bb1-6e081c90ada2",
		&armmanagedservices.RegistrationAssignmentsClientGetOptions{ExpandRegistrationDefinition: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
