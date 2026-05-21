package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: 2023-07-01-preview/ListZonesByResourceGroup.json
func ExampleZonesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewZonesClient().NewListByResourceGroupPager("rg1", nil)
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
		// page = armdns.ZonesClientListByResourceGroupResponse{
		// 	ZoneListResult: armdns.ZoneListResult{
		// 		NextLink: to.Ptr("https://servicehost/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones?api-version=2023-07-01-preview&$skipToken=skipToken"),
		// 		Value: []*armdns.Zone{
		// 			{
		// 				Name: to.Ptr("zone1"),
		// 				Type: to.Ptr("Microsoft.Network/dnsZones"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/zone1"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armdns.ZoneProperties{
		// 					MaxNumberOfRecordSets: to.Ptr[int64](5000),
		// 					NameServers: []*string{
		// 						to.Ptr("ns1-01.azure-dns.com"),
		// 						to.Ptr("ns2-01.azure-dns.net"),
		// 						to.Ptr("ns3-01.azure-dns.org"),
		// 						to.Ptr("ns4-01.azure-dns.info"),
		// 					},
		// 					NumberOfRecordSets: to.Ptr[int64](2),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("zone2"),
		// 				Type: to.Ptr("Microsoft.Network/dnsZones"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/zone2"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armdns.ZoneProperties{
		// 					MaxNumberOfRecordSets: to.Ptr[int64](5000),
		// 					NameServers: []*string{
		// 						to.Ptr("ns1-02.azure-dns.com"),
		// 						to.Ptr("ns2-02.azure-dns.net"),
		// 						to.Ptr("ns3-02.azure-dns.org"),
		// 						to.Ptr("ns4-02.azure-dns.info"),
		// 					},
		// 					NumberOfRecordSets: to.Ptr[int64](300),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
