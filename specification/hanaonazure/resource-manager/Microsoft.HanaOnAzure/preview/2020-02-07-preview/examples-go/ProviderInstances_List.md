Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhanaonazure%2Farmhanaonazure%2Fv0.1.0/sdk/resourcemanager/hanaonazure/armhanaonazure/README.md) on how to add the SDK to your project and authenticate.

```go
package armhanaonazure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/ProviderInstances_List.json
func ExampleProviderInstancesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhanaonazure.NewProviderInstancesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<sap-monitor-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("ProviderInstance.ID: %s\n", *v.ID)
		}
	}
}
```
