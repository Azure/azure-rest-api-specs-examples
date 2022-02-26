Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.2.0/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_DeleteResourceGroupScope.json
func ExampleRemediationsClient_DeleteAtResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicyinsights.NewRemediationsClient("<subscription-id>", cred, nil)
	res, err := client.DeleteAtResourceGroup(ctx,
		"<resource-group-name>",
		"<remediation-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RemediationsClientDeleteAtResourceGroupResult)
}
```
