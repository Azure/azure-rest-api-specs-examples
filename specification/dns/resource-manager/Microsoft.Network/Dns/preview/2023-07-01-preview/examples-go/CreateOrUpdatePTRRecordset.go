package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: 2023-07-01-preview/CreateOrUpdatePTRRecordset.json
func ExampleRecordSetsClient_CreateOrUpdate_createPtrRecordset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecordSetsClient().CreateOrUpdate(ctx, "rg1", "0.0.127.in-addr.arpa", "1", armdns.RecordTypePTR, armdns.RecordSet{
		Properties: &armdns.RecordSetProperties{
			PtrRecords: []*armdns.PtrRecord{
				{
					Ptrdname: to.Ptr("localhost"),
				},
			},
			TTL: to.Ptr[int64](3600),
			Metadata: map[string]*string{
				"key1": to.Ptr("value1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdns.RecordSetsClientCreateOrUpdateResponse{
	// 	RecordSet: &armdns.RecordSet{
	// 		Name: to.Ptr("1"),
	// 		Type: to.Ptr("Microsoft.Network/dnsZones/PTR"),
	// 		Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/0.0.127.in-addr.arpa/PTR/1"),
	// 		Properties: &armdns.RecordSetProperties{
	// 			PtrRecords: []*armdns.PtrRecord{
	// 				{
	// 					Ptrdname: to.Ptr("localhost"),
	// 				},
	// 			},
	// 			TTL: to.Ptr[int64](3600),
	// 			Fqdn: to.Ptr("1.0.0.127.in-addr.arpa"),
	// 			Metadata: map[string]*string{
	// 				"key1": to.Ptr("value1"),
	// 			},
	// 		},
	// 	},
	// }
}
