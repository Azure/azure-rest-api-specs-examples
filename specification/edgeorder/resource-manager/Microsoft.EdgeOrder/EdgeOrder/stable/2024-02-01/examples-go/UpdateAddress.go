package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder/v2"
)

// Generated from example definition: 2024-02-01/UpdateAddress.json
func ExampleAddressesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("eb5dc900-6186-49d8-b7d7-febd866fdc1d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAddressesClient().BeginUpdate(ctx, "YourResourceGroupName", "TestAddressName2", armedgeorder.AddressUpdateParameter{
		Properties: &armedgeorder.AddressUpdateProperties{
			ContactDetails: &armedgeorder.ContactDetails{
				ContactName: to.Ptr("YYYY YYYY"),
				EmailList: []*string{
					to.Ptr("xxxx@xxxx.xxx"),
				},
				Phone:          to.Ptr("0000000000"),
				PhoneExtension: to.Ptr(""),
			},
			ShippingAddress: &armedgeorder.ShippingAddress{
				AddressType:     to.Ptr(armedgeorder.AddressTypeNone),
				City:            to.Ptr("San Francisco"),
				CompanyName:     to.Ptr("Microsoft"),
				Country:         to.Ptr("US"),
				PostalCode:      to.Ptr("94107"),
				StateOrProvince: to.Ptr("CA"),
				StreetAddress1:  to.Ptr("16 TOWNSEND ST"),
				StreetAddress2:  to.Ptr("UNIT 1"),
			},
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armedgeorder.AddressesClientUpdateResponse{
	// 	AddressResource: &armedgeorder.AddressResource{
	// 		Name: to.Ptr("TestAddressName2"),
	// 		Type: to.Ptr("Microsoft.EdgeOrder/addresses"),
	// 		ID: to.Ptr("/subscriptions/eb5dc900-6186-49d8-b7d7-febd866fdc1d/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/addresses/TestAddressName2"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armedgeorder.AddressProperties{
	// 			AddressValidationStatus: to.Ptr(armedgeorder.AddressValidationStatusValid),
	// 			ContactDetails: &armedgeorder.ContactDetails{
	// 				ContactName: to.Ptr("YYYY YYYY"),
	// 				EmailList: []*string{
	// 					to.Ptr("xxxx@xxxx.xxx"),
	// 				},
	// 				Phone: to.Ptr("0000000000"),
	// 				PhoneExtension: to.Ptr(""),
	// 			},
	// 			ShippingAddress: &armedgeorder.ShippingAddress{
	// 				AddressType: to.Ptr(armedgeorder.AddressTypeNone),
	// 				City: to.Ptr("San Francisco"),
	// 				CompanyName: to.Ptr("Microsoft"),
	// 				Country: to.Ptr("US"),
	// 				PostalCode: to.Ptr("94107"),
	// 				StateOrProvince: to.Ptr("CA"),
	// 				StreetAddress1: to.Ptr("16 TOWNSEND ST"),
	// 				StreetAddress2: to.Ptr("UNIT 1"),
	// 			},
	// 		},
	// 		SystemData: &armedgeorder.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
