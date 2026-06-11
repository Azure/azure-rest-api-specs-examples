package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/PublicIpPrefixGet.json
func ExamplePublicIPPrefixesClient_Get_getPublicIPPrefix() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicIPPrefixesClient().Get(ctx, "rg1", "test-ipprefix", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.PublicIPPrefixesClientGetResponse{
	// 	PublicIPPrefix: armnetwork.PublicIPPrefix{
	// 		Name: to.Ptr("test-ipprefix"),
	// 		Type: to.Ptr("Microsoft.Network/publicIPPrefixes"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/test-ipprefix"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.PublicIPPrefixPropertiesFormat{
	// 			IPPrefix: to.Ptr("192.168.254.2/30"),
	// 			IPTags: []*armnetwork.IPTag{
	// 			},
	// 			PrefixLength: to.Ptr[int32](30),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 			PublicIPAddresses: []*armnetwork.ReferencedPublicIPAddress{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
	// 				},
	// 			},
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 		},
	// 		SKU: &armnetwork.PublicIPPrefixSKU{
	// 			Name: to.Ptr(armnetwork.PublicIPPrefixSKUNameStandard),
	// 		},
	// 	},
	// }
}
