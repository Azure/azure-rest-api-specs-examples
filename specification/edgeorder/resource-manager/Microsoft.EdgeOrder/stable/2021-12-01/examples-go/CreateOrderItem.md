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

// x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateOrderItem.json
func ExampleEdgeOrderManagementClient_BeginCreateOrderItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armedgeorder.NewEdgeOrderManagementClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrderItem(ctx,
		"<order-item-name>",
		"<resource-group-name>",
		armedgeorder.OrderItemResource{
			TrackedResource: armedgeorder.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"carrot": to.StringPtr("vegetable"),
					"mango":  to.StringPtr("fruit"),
				},
			},
			Properties: &armedgeorder.OrderItemProperties{
				AddressDetails: &armedgeorder.AddressDetails{
					ForwardAddress: &armedgeorder.AddressProperties{
						ContactDetails: &armedgeorder.ContactDetails{
							ContactName: to.StringPtr("<contact-name>"),
							EmailList: []*string{
								to.StringPtr("ssemmail@microsoft.com"),
								to.StringPtr("vishwamdir@microsoft.com")},
							Phone: to.StringPtr("<phone>"),
						},
						ShippingAddress: &armedgeorder.ShippingAddress{
							AddressType:     armedgeorder.AddressTypeResidential.ToPtr(),
							City:            to.StringPtr("<city>"),
							CompanyName:     to.StringPtr("<company-name>"),
							Country:         to.StringPtr("<country>"),
							PostalCode:      to.StringPtr("<postal-code>"),
							StateOrProvince: to.StringPtr("<state-or-province>"),
							StreetAddress1:  to.StringPtr("<street-address1>"),
							StreetAddress2:  to.StringPtr("<street-address2>"),
							ZipExtendedCode: to.StringPtr("<zip-extended-code>"),
						},
					},
				},
				OrderID: to.StringPtr("<order-id>"),
				OrderItemDetails: &armedgeorder.OrderItemDetails{
					OrderItemType: armedgeorder.OrderItemTypePurchase.ToPtr(),
					Preferences: &armedgeorder.Preferences{
						TransportPreferences: &armedgeorder.TransportPreferences{
							PreferredShipmentType: armedgeorder.TransportShipmentTypesMicrosoftManaged.ToPtr(),
						},
					},
					ProductDetails: &armedgeorder.ProductDetails{
						HierarchyInformation: &armedgeorder.HierarchyInformation{
							ConfigurationName: to.StringPtr("<configuration-name>"),
							ProductFamilyName: to.StringPtr("<product-family-name>"),
							ProductLineName:   to.StringPtr("<product-line-name>"),
							ProductName:       to.StringPtr("<product-name>"),
						},
					},
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
	log.Printf("OrderItemResource.ID: %s\n", *res.ID)
}
```
