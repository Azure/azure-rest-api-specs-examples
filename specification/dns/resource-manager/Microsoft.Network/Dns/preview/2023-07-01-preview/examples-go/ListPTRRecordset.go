package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: 2023-07-01-preview/ListPTRRecordset.json
func ExampleRecordSetsClient_NewListByTypePager_listPtrRecordsets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecordSetsClient().NewListByTypePager("rg1", "0.0.127.in-addr.arpa", armdns.RecordTypePTR, nil)
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
		// page = armdns.RecordSetsClientListByTypeResponse{
		// 	RecordSetListResult: armdns.RecordSetListResult{
		// 		NextLink: to.Ptr("https://servicehost/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/0.0.127.in-addr.arpa/PTR?api-version=2023-07-01-preview&$skipToken=skipToken"),
		// 		Value: []*armdns.RecordSet{
		// 			{
		// 				Name: to.Ptr("1"),
		// 				Type: to.Ptr("Microsoft.Network/dnsZones/PTR"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/0.0.127.in-addr.arpa/PTR/1"),
		// 				Properties: &armdns.RecordSetProperties{
		// 					PtrRecords: []*armdns.PtrRecord{
		// 						{
		// 							Ptrdname: to.Ptr("localhost"),
		// 						},
		// 					},
		// 					TTL: to.Ptr[int64](3600),
		// 					Fqdn: to.Ptr("1.0.0.127.in-addr.arpa"),
		// 					Metadata: map[string]*string{
		// 						"key1": to.Ptr("value1"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
