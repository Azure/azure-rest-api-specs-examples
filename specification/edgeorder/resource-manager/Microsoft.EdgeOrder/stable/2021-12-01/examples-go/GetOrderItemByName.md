Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorder%2Farmedgeorder%2Fv0.1.0/sdk/resourcemanager/edgeorder/armedgeorder/README.md) on how to add the SDK to your project and authenticate.

```go
package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetOrderItemByName.json
func ExampleEdgeOrderManagementClient_GetOrderItemByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armedgeorder.NewEdgeOrderManagementClient("<subscription-id>", cred, nil)
	res, err := client.GetOrderItemByName(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		&armedgeorder.EdgeOrderManagementClientGetOrderItemByNameOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OrderItemResource.ID: %s\n", *res.ID)
}
```
