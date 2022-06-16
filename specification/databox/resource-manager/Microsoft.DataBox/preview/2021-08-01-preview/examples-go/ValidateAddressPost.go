package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/ValidateAddressPost.json
func ExampleServiceClient_ValidateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	res, err := client.ValidateAddress(ctx,
		"<location>",
		armdatabox.ValidateAddress{
			ValidationType: armdatabox.ValidationInputDiscriminatorValidateAddress.ToPtr(),
			DeviceType:     armdatabox.SKUNameDataBox.ToPtr(),
			ShippingAddress: &armdatabox.ShippingAddress{
				AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
				City:            to.StringPtr("<city>"),
				CompanyName:     to.StringPtr("<company-name>"),
				Country:         to.StringPtr("<country>"),
				PostalCode:      to.StringPtr("<postal-code>"),
				StateOrProvince: to.StringPtr("<state-or-province>"),
				StreetAddress1:  to.StringPtr("<street-address1>"),
				StreetAddress2:  to.StringPtr("<street-address2>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServiceClientValidateAddressResult)
}
