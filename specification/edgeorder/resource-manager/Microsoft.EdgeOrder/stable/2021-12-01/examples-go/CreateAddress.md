Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorder%2Farmedgeorder%2Fv0.4.0/sdk/resourcemanager/edgeorder/armedgeorder/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateAddress.json
func ExampleManagementClient_BeginCreateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateAddress(ctx,
		"<address-name>",
		"<resource-group-name>",
		armedgeorder.AddressResource{
			Location: to.Ptr("<location>"),
			Properties: &armedgeorder.AddressProperties{
				ContactDetails: &armedgeorder.ContactDetails{
					ContactName: to.Ptr("<contact-name>"),
					EmailList: []*string{
						to.Ptr("xxxx@xxxx.xxx")},
					Phone:          to.Ptr("<phone>"),
					PhoneExtension: to.Ptr("<phone-extension>"),
				},
				ShippingAddress: &armedgeorder.ShippingAddress{
					AddressType:     to.Ptr(armedgeorder.AddressTypeNone),
					City:            to.Ptr("<city>"),
					CompanyName:     to.Ptr("<company-name>"),
					Country:         to.Ptr("<country>"),
					PostalCode:      to.Ptr("<postal-code>"),
					StateOrProvince: to.Ptr("<state-or-province>"),
					StreetAddress1:  to.Ptr("<street-address1>"),
					StreetAddress2:  to.Ptr("<street-address2>"),
				},
			},
		},
		&armedgeorder.ManagementClientBeginCreateAddressOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
