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

// x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/UpdateAddress.json
func ExampleEdgeOrderManagementClient_BeginUpdateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armedgeorder.NewEdgeOrderManagementClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateAddress(ctx,
		"<address-name>",
		"<resource-group-name>",
		armedgeorder.AddressUpdateParameter{
			Properties: &armedgeorder.AddressUpdateProperties{
				ContactDetails: &armedgeorder.ContactDetails{
					ContactName: to.StringPtr("<contact-name>"),
					EmailList: []*string{
						to.StringPtr("ssemcr@microsoft.com")},
					Phone:          to.StringPtr("<phone>"),
					PhoneExtension: to.StringPtr("<phone-extension>"),
				},
				ShippingAddress: &armedgeorder.ShippingAddress{
					AddressType:     armedgeorder.AddressTypeNone.ToPtr(),
					City:            to.StringPtr("<city>"),
					CompanyName:     to.StringPtr("<company-name>"),
					Country:         to.StringPtr("<country>"),
					PostalCode:      to.StringPtr("<postal-code>"),
					StateOrProvince: to.StringPtr("<state-or-province>"),
					StreetAddress1:  to.StringPtr("<street-address1>"),
					StreetAddress2:  to.StringPtr("<street-address2>"),
				},
			},
			Tags: map[string]*string{
				"Hobby":    to.StringPtr("Web Series Added"),
				"Name":     to.StringPtr("Smile-Updated"),
				"WhatElse": to.StringPtr("Web Series Added"),
				"Work":     to.StringPtr("Engineering"),
			},
		},
		&armedgeorder.EdgeOrderManagementClientBeginUpdateAddressOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AddressResource.ID: %s\n", *res.ID)
}
```
