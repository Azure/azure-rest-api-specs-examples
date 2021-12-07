Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbaremetalinfrastructure%2Farmbaremetalinfrastructure%2Fv0.1.0/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure/README.md) on how to add the SDK to your project and authenticate.

```go
package armbaremetalinfrastructure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/baremetalinfrastructure/armbaremetalinfrastructure"
)

// x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_ListBySubscription.json
func ExampleAzureBareMetalInstancesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbaremetalinfrastructure.NewAzureBareMetalInstancesClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("AzureBareMetalInstance.ID: %s\n", *v.ID)
		}
	}
}
```
