package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListSRVRecordset.json
func ExampleRecordSetsClient_NewListByTypePager_listSrvRecordsets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecordSetsClient().NewListByTypePager("rg1", "zone1", armdns.RecordTypeSRV, &armdns.RecordSetsClientListByTypeOptions{Top: nil,
		Recordsetnamesuffix: nil,
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
		// page.RecordSetListResult = armdns.RecordSetListResult{
		// 	Value: []*armdns.RecordSet{
		// 		{
		// 			Name: to.Ptr("record1"),
		// 			Type: to.Ptr("Microsoft.Network/dnsZones/SRV"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/zone1/SRV/record1"),
		// 			Properties: &armdns.RecordSetProperties{
		// 				SrvRecords: []*armdns.SrvRecord{
		// 					{
		// 						Port: to.Ptr[int32](80),
		// 						Priority: to.Ptr[int32](0),
		// 						Target: to.Ptr("contoso.com"),
		// 						Weight: to.Ptr[int32](10),
		// 				}},
		// 				TTL: to.Ptr[int64](3600),
		// 				Fqdn: to.Ptr("record1.zone1"),
		// 				Metadata: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
