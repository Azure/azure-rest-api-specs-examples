package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorListBySubscription.json
func ExampleVendorsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVendorsClient().NewListBySubscriptionPager(nil)
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
		// page.VendorListResult = armhybridnetwork.VendorListResult{
		// 	Value: []*armhybridnetwork.Vendor{
		// 		{
		// 			Name: to.Ptr("TestVendor"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/vendors"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor"),
		// 			Properties: &armhybridnetwork.VendorPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
		// 				SKUs: []*armhybridnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestVendorSku1"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("TestVendor2"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/vendors"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor2"),
		// 			Properties: &armhybridnetwork.VendorPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
		// 				SKUs: []*armhybridnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestVendorSku2"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
