Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.4.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

```go
package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionRoleAssignment.json
func ExampleRoleAssignmentsClient_GetByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armbilling.NewRoleAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetByInvoiceSection(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		"<invoice-section-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
