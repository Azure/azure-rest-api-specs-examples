package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PublicIpPrefixList.json
func ExamplePublicIPPrefixesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicIPPrefixesClient().NewListPager("rg1", nil)
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
		// page.PublicIPPrefixListResult = armnetwork.PublicIPPrefixListResult{
		// 	Value: []*armnetwork.PublicIPPrefix{
		// 		{
		// 			Name: to.Ptr("test-ipprefix"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
		// 				IPPrefix: to.Ptr("40.85.154.2/30"),
		// 				IPTags: []*armnetwork.IPTag{
		// 					{
		// 						IPTagType: to.Ptr("FirstPartyUsage"),
		// 						Tag: to.Ptr("SQL"),
		// 				}},
		// 				PrefixLength: to.Ptr[int32](30),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.PublicIPPrefixSKU{
		// 				Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ipprefix03"),
		// 			Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/ipprefix03"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
		// 				IPPrefix: to.Ptr("40.85.153.2/31"),
		// 				IPTags: []*armnetwork.IPTag{
		// 				},
		// 				PrefixLength: to.Ptr[int32](31),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.PublicIPPrefixSKU{
		// 				Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
		// 			},
		// 	}},
		// }
	}
}
