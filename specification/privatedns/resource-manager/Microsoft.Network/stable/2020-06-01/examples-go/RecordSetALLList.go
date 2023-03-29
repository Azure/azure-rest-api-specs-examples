package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetALLList.json
func ExampleRecordSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecordSetsClient().NewListPager("resourceGroup1", "privatezone1.com", &armprivatedns.RecordSetsClientListOptions{Top: nil,
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
		// page.RecordSetListResult = armprivatedns.RecordSetListResult{
		// 	Value: []*armprivatedns.RecordSet{
		// 		{
		// 			Name: to.Ptr("recorda"),
		// 			Type: to.Ptr("Microsoft.Network/privateDnsZones/A"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/A/recorda"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armprivatedns.RecordSetProperties{
		// 				ARecords: []*armprivatedns.ARecord{
		// 					{
		// 						IPv4Address: to.Ptr("1.2.3.4"),
		// 				}},
		// 				Fqdn: to.Ptr("recorda.privatezone1.com."),
		// 				IsAutoRegistered: to.Ptr(false),
		// 				Metadata: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 				TTL: to.Ptr[int64](3600),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("recordcname"),
		// 			Type: to.Ptr("Microsoft.Network/privateDnsZones/CNAME"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/CNAME/recordcname"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armprivatedns.RecordSetProperties{
		// 				CnameRecord: &armprivatedns.CnameRecord{
		// 					Cname: to.Ptr("contoso.com"),
		// 				},
		// 				Fqdn: to.Ptr("recordcname.privatezone1.com."),
		// 				IsAutoRegistered: to.Ptr(false),
		// 				Metadata: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 				TTL: to.Ptr[int64](3600),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("recordmx"),
		// 			Type: to.Ptr("Microsoft.Network/privateDnsZones/MX"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/MX/recordmx"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armprivatedns.RecordSetProperties{
		// 				Fqdn: to.Ptr("recordmx.privatezone1.com."),
		// 				IsAutoRegistered: to.Ptr(false),
		// 				Metadata: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 				MxRecords: []*armprivatedns.MxRecord{
		// 					{
		// 						Exchange: to.Ptr("mail.contoso1.com"),
		// 						Preference: to.Ptr[int32](0),
		// 				}},
		// 				TTL: to.Ptr[int64](3600),
		// 			},
		// 	}},
		// }
	}
}
