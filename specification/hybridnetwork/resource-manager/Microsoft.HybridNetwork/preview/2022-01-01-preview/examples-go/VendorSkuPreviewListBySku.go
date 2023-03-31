package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/792db17291c758b2bfdbbc0d35d0e2f5b5a1bd05/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorSkuPreviewListBySku.json
func ExampleVendorSKUPreviewClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVendorSKUPreviewClient().NewListPager("TestVendor", "TestSku", nil)
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
		// page.PreviewSubscriptionsList = armhybridnetwork.PreviewSubscriptionsList{
		// 	Value: []*armhybridnetwork.PreviewSubscription{
		// 		{
		// 			Name: to.Ptr("previewSub1"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/vendors/vendorskus/previewsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestSku/previewsubscriptions/previewSub1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("previewSub2"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/vendors/vendorskus/previewsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestSku/previewsubscriptions/previewSub2"),
		// 	}},
		// }
	}
}
