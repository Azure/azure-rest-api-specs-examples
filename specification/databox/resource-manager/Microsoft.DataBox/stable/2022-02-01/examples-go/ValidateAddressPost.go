package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/ValidateAddressPost.json
func ExampleServiceClient_ValidateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().ValidateAddress(ctx, "westus", armdatabox.ValidateAddress{
		ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateAddress),
		DeviceType:     to.Ptr(armdatabox.SKUNameDataBox),
		ShippingAddress: &armdatabox.ShippingAddress{
			AddressType:     to.Ptr(armdatabox.AddressTypeCommercial),
			City:            to.Ptr("San Francisco"),
			CompanyName:     to.Ptr("Microsoft"),
			Country:         to.Ptr("US"),
			PostalCode:      to.Ptr("94107"),
			StateOrProvince: to.Ptr("CA"),
			StreetAddress1:  to.Ptr("16 TOWNSEND ST"),
			StreetAddress2:  to.Ptr("Unit 1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AddressValidationOutput = armdatabox.AddressValidationOutput{
	// 	Properties: &armdatabox.AddressValidationProperties{
	// 		ValidationType: to.Ptr(armdatabox.ValidationInputDiscriminatorValidateAddress),
	// 		AlternateAddresses: []*armdatabox.ShippingAddress{
	// 			{
	// 				AddressType: to.Ptr(armdatabox.AddressTypeNone),
	// 				City: to.Ptr("SAN FRANCISCO"),
	// 				Country: to.Ptr("US"),
	// 				PostalCode: to.Ptr("94107"),
	// 				StateOrProvince: to.Ptr("CA"),
	// 				StreetAddress1: to.Ptr("16 TOWNSEND ST"),
	// 				StreetAddress2: to.Ptr("Unit 1"),
	// 				StreetAddress3: to.Ptr(""),
	// 		}},
	// 		ValidationStatus: to.Ptr(armdatabox.AddressValidationStatusValid),
	// 	},
	// }
}
