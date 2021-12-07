Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.1.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

```go
package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleAssignmentDelete.json
func ExampleBillingRoleAssignmentsClient_DeleteByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingRoleAssignmentsClient(cred, nil)
	res, err := client.DeleteByBillingAccount(ctx,
		"<billing-account-name>",
		"<billing-role-assignment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingRoleAssignment.ID: %s\n", *res.ID)
}
```
