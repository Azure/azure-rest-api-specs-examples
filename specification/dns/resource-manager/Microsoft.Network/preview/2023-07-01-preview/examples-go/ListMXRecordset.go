package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/987a8f38ab2a8359d085e149be042267a9ecc66f/specification/dns/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/ListMXRecordset.json
func ExampleRecordSetsClient_NewListByTypePager_listMxRecordsets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecordSetsClient().NewListByTypePager("rg1", "zone1", armdns.RecordTypeMX, &armdns.RecordSetsClientListByTypeOptions{Top: nil,
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
		// 			Type: to.Ptr("Microsoft.Network/dnsZones/MX"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/zone1/MX/record1"),
		// 			Properties: &armdns.RecordSetProperties{
		// 				MxRecords: []*armdns.MxRecord{
		// 					{
		// 						Exchange: to.Ptr("mail.contoso.com"),
		// 						Preference: to.Ptr[int32](0),
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
