package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListAddressesAtResourceGroupLevel.json
func ExampleManagementClient_NewListAddressesAtResourceGroupLevelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewListAddressesAtResourceGroupLevelPager("YourResourceGroupName", &armedgeorder.ManagementClientListAddressesAtResourceGroupLevelOptions{Filter: nil,
		SkipToken: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AddressResourceList = armedgeorder.AddressResourceList{
		// 	Value: []*armedgeorder.AddressResource{
		// 		{
		// 			Name: to.Ptr("TestAddressName1"),
		// 			Type: to.Ptr("Microsoft.EdgeOrder/addresses"),
		// 			ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/addresses/TestAddressName1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armedgeorder.AddressProperties{
		// 				AddressValidationStatus: to.Ptr(armedgeorder.AddressValidationStatusValid),
		// 				ContactDetails: &armedgeorder.ContactDetails{
		// 					ContactName: to.Ptr("XXXX XXXX"),
		// 					EmailList: []*string{
		// 						to.Ptr("xxxx@xxxx.xxx")},
		// 						Phone: to.Ptr("0000000000"),
		// 						PhoneExtension: to.Ptr(""),
		// 					},
		// 					ShippingAddress: &armedgeorder.ShippingAddress{
		// 						AddressType: to.Ptr(armedgeorder.AddressTypeNone),
		// 						City: to.Ptr("San Francisco"),
		// 						CompanyName: to.Ptr("Microsoft"),
		// 						Country: to.Ptr("US"),
		// 						PostalCode: to.Ptr("94107"),
		// 						StateOrProvince: to.Ptr("CA"),
		// 						StreetAddress1: to.Ptr("16 TOWNSEND ST"),
		// 						StreetAddress2: to.Ptr("UNIT 1"),
		// 					},
		// 				},
		// 				SystemData: &armedgeorder.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestAddressName2"),
		// 				Type: to.Ptr("Microsoft.EdgeOrder/addresses"),
		// 				ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/addresses/TestAddressName2"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 				Properties: &armedgeorder.AddressProperties{
		// 					AddressValidationStatus: to.Ptr(armedgeorder.AddressValidationStatusValid),
		// 					ContactDetails: &armedgeorder.ContactDetails{
		// 						ContactName: to.Ptr("YYYY YYYY"),
		// 						EmailList: []*string{
		// 							to.Ptr("xxxx@xxxx.xxx")},
		// 							Phone: to.Ptr("0000000000"),
		// 							PhoneExtension: to.Ptr(""),
		// 						},
		// 						ShippingAddress: &armedgeorder.ShippingAddress{
		// 							AddressType: to.Ptr(armedgeorder.AddressTypeNone),
		// 							City: to.Ptr("San Francisco"),
		// 							CompanyName: to.Ptr("Microsoft"),
		// 							Country: to.Ptr("US"),
		// 							PostalCode: to.Ptr("94107"),
		// 							StateOrProvince: to.Ptr("CA"),
		// 							StreetAddress1: to.Ptr("16 TOWNSEND ST"),
		// 							StreetAddress2: to.Ptr("UNIT 1"),
		// 						},
		// 					},
		// 					SystemData: &armedgeorder.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}
