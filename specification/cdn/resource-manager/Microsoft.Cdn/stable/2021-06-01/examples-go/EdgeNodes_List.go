package armcdn_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/EdgeNodes_List.json
func ExampleEdgeNodesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcdn.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeNodesClient().NewListPager(nil)
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
		// page.EdgenodeResult = armcdn.EdgenodeResult{
		// 	Value: []*armcdn.EdgeNode{
		// 		{
		// 			Name: to.Ptr("Standard_Verizon"),
		// 			Type: to.Ptr("Microsoft.Cdn/edgenodes"),
		// 			ID: to.Ptr("/providers/Microsoft.Cdn/edgenodes/Standard_Verizon"),
		// 			Properties: &armcdn.EdgeNodeProperties{
		// 				IPAddressGroups: []*armcdn.IPAddressGroup{
		// 					{
		// 						DeliveryRegion: to.Ptr("All"),
		// 						IPv4Addresses: []*armcdn.CidrIPAddress{
		// 							{
		// 								BaseIPAddress: to.Ptr("192.229.176.0"),
		// 								PrefixLength: to.Ptr[int32](24),
		// 							},
		// 							{
		// 								BaseIPAddress: to.Ptr("180.240.184.128"),
		// 								PrefixLength: to.Ptr[int32](25),
		// 							},
		// 							{
		// 								BaseIPAddress: to.Ptr("152.195.27.0"),
		// 								PrefixLength: to.Ptr[int32](24),
		// 						}},
		// 						IPv6Addresses: []*armcdn.CidrIPAddress{
		// 							{
		// 								BaseIPAddress: to.Ptr("2606:2800:60f2::"),
		// 								PrefixLength: to.Ptr[int32](48),
		// 							},
		// 							{
		// 								BaseIPAddress: to.Ptr("2606:2800:700c::"),
		// 								PrefixLength: to.Ptr[int32](48),
		// 						}},
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Premium_Verizon"),
		// 			Type: to.Ptr("Microsoft.Cdn/edgenodes"),
		// 			ID: to.Ptr("/providers/Microsoft.Cdn/edgenodes/Premium_Verizon"),
		// 			Properties: &armcdn.EdgeNodeProperties{
		// 				IPAddressGroups: []*armcdn.IPAddressGroup{
		// 					{
		// 						DeliveryRegion: to.Ptr("All"),
		// 						IPv4Addresses: []*armcdn.CidrIPAddress{
		// 							{
		// 								BaseIPAddress: to.Ptr("192.229.176.0"),
		// 								PrefixLength: to.Ptr[int32](24),
		// 							},
		// 							{
		// 								BaseIPAddress: to.Ptr("152.195.27.0"),
		// 								PrefixLength: to.Ptr[int32](24),
		// 						}},
		// 						IPv6Addresses: []*armcdn.CidrIPAddress{
		// 							{
		// 								BaseIPAddress: to.Ptr("2606:2800:60f2::"),
		// 								PrefixLength: to.Ptr[int32](48),
		// 							},
		// 							{
		// 								BaseIPAddress: to.Ptr("2606:2800:700c::"),
		// 								PrefixLength: to.Ptr[int32](48),
		// 						}},
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Custom_Verizon"),
		// 			Type: to.Ptr("Microsoft.Cdn/edgenodes"),
		// 			ID: to.Ptr("/providers/Microsoft.Cdn/edgenodes/Custom_Verizon"),
		// 			Properties: &armcdn.EdgeNodeProperties{
		// 				IPAddressGroups: []*armcdn.IPAddressGroup{
		// 					{
		// 						DeliveryRegion: to.Ptr("All"),
		// 						IPv4Addresses: []*armcdn.CidrIPAddress{
		// 							{
		// 								BaseIPAddress: to.Ptr("192.229.176.0"),
		// 								PrefixLength: to.Ptr[int32](24),
		// 							},
		// 							{
		// 								BaseIPAddress: to.Ptr("2606:2800:420b::"),
		// 								PrefixLength: to.Ptr[int32](48),
		// 							},
		// 							{
		// 								BaseIPAddress: to.Ptr("2606:2800:700c::"),
		// 								PrefixLength: to.Ptr[int32](48),
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
