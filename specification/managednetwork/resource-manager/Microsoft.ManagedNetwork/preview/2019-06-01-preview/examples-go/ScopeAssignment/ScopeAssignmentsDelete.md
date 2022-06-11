```go
package armmanagednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ScopeAssignment/ScopeAssignmentsDelete.json
func ExampleScopeAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagednetwork.NewScopeAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"subscriptions/subscriptionC",
		"subscriptionCAssignment",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagednetwork%2Farmmanagednetwork%2Fv0.1.0/sdk/resourcemanager/managednetwork/armmanagednetwork/README.md) on how to add the SDK to your project and authenticate.
