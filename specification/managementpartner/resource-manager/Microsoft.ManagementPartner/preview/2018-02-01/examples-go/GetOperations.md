Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagementpartner%2Farmmanagementpartner%2Fv0.2.1/sdk/resourcemanager/managementpartner/armmanagementpartner/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagementpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner"
)

// x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/GetOperations.json
func ExampleOperationClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagementpartner.NewOperationClient(cred, nil)
	pager := client.List(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
