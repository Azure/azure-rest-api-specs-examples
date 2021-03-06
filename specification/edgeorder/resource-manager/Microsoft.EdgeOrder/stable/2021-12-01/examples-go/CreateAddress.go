package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/CreateAddress.json
func ExampleManagementClient_BeginCreateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("YourSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateAddress(ctx,
		"TestAddressName2",
		"YourResourceGroupName",
		armedgeorder.AddressResource{
			Location: to.Ptr("eastus"),
			Properties: &armedgeorder.AddressProperties{
				ContactDetails: &armedgeorder.ContactDetails{
					ContactName: to.Ptr("XXXX XXXX"),
					EmailList: []*string{
						to.Ptr("xxxx@xxxx.xxx")},
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
