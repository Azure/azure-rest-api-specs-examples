package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetZone.json
func ExampleZonesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewZonesClient().Get(ctx, "rg1", "zone1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Zone = armdns.Zone{
	// 	Name: to.Ptr("zone1"),
	// 	Type: to.Ptr("Microsoft.Network/dnsZones"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/zone1"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdns.ZoneProperties{
	// 		MaxNumberOfRecordSets: to.Ptr[int64](5000),
	// 		NameServers: []*string{
	// 			to.Ptr("ns1-01.azure-dns.com"),
	// 			to.Ptr("ns2-01.azure-dns.net"),
	// 			to.Ptr("ns3-01.azure-dns.org"),
	// 			to.Ptr("ns4-01.azure-dns.info")},
	// 			NumberOfRecordSets: to.Ptr[int64](2),
	// 		},
	// 	}
}
