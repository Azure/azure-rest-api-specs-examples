Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdataboxedge%2Farmdataboxedge%2Fv0.1.0/sdk/resourcemanager/databoxedge/armdataboxedge/README.md) on how to add the SDK to your project and authenticate.

```go
package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/OrderPut.json
func ExampleOrdersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewOrdersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.Order{
			Properties: &armdataboxedge.OrderProperties{
				ContactInformation: &armdataboxedge.ContactDetails{
					CompanyName:   to.StringPtr("<company-name>"),
					ContactPerson: to.StringPtr("<contact-person>"),
					EmailList: []*string{
						to.StringPtr("john@microsoft.com")},
					Phone: to.StringPtr("<phone>"),
				},
				ShippingAddress: &armdataboxedge.Address{
					AddressLine1: to.StringPtr("<address-line1>"),
					AddressLine2: to.StringPtr("<address-line2>"),
					AddressLine3: to.StringPtr("<address-line3>"),
					City:         to.StringPtr("<city>"),
					Country:      to.StringPtr("<country>"),
					PostalCode:   to.StringPtr("<postal-code>"),
					State:        to.StringPtr("<state>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Order.ID: %s\n", *res.ID)
}
```
