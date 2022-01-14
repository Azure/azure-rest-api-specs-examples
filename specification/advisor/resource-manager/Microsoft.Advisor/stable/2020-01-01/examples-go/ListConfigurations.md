Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fadvisor%2Farmadvisor%2Fv0.2.0/sdk/resourcemanager/advisor/armadvisor/README.md) on how to add the SDK to your project and authenticate.

```go
package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListConfigurations.json
func ExampleConfigurationsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armadvisor.NewConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.ListByResourceGroup(ctx,
		"<resource-group>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConfigurationsClientListByResourceGroupResult)
}
```
