Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorder%2Farmedgeorder%2Fv0.1.0/sdk/resourcemanager/edgeorder/armedgeorder/README.md) on how to add the SDK to your project and authenticate.

```go
package armedgeorder_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateOrderItem.json
func ExampleEdgeOrderManagementClient_BeginUpdateOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armedgeorder.NewEdgeOrderManagementClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateOrderItem(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		armedgeorder.OrderItemUpdateParameter{
			Properties: &armedgeorder.OrderItemUpdateProperties{
				ForwardAddress: &armedgeorder.AddressProperties{
					ContactDetails: &armedgeorder.ContactDetails{
						ContactName: to.StringPtr("<contact-name>"),
						EmailList: []*string{
							to.StringPtr("testemail@microsoft.com")},
						Phone: to.StringPtr("<phone>"),
					},
				},
				Preferences: &armedgeorder.Preferences{
					TransportPreferences: &armedgeorder.TransportPreferences{
						PreferredShipmentType: armedgeorder.TransportShipmentTypesCustomerManaged.ToPtr(),
					},
				},
			},
			Tags: map[string]*string{
				"ant":    to.StringPtr("insect"),
				"pigeon": to.StringPtr("bird"),
				"tiger":  to.StringPtr("animal"),
			},
		},
		&armedgeorder.EdgeOrderManagementClientBeginUpdateOrderItemOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OrderItemResource.ID: %s\n", *res.ID)
}
```
