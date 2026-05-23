package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns/v2"
)

// Generated from example definition: 2024-06-01/RecordSetAList.json
func ExampleRecordSetsClient_NewListByTypePager_getPrivateDnsZoneARecordSets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecordSetsClient().NewListByTypePager("resourceGroup1", "privatezone1.com", armprivatedns.RecordTypeA, nil)
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
		// page = armprivatedns.RecordSetsClientListByTypeResponse{
		// 	RecordSetListResult: armprivatedns.RecordSetListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/A?api-version=2024-06-01&$skipToken=skipToken"),
		// 		Value: []*armprivatedns.RecordSet{
		// 			{
		// 				Name: to.Ptr("recorda1"),
		// 				Type: to.Ptr("Microsoft.Network/privateDnsZones/A"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/A/recorda1"),
		// 				Properties: &armprivatedns.RecordSetProperties{
		// 					ARecords: []*armprivatedns.ARecord{
		// 						{
		// 							IPv4Address: to.Ptr("1.2.3.4"),
		// 						},
		// 					},
		// 					Fqdn: to.Ptr("recorda1.privatezone1.com."),
		// 					IsAutoRegistered: to.Ptr(false),
		// 					Metadata: map[string]*string{
		// 						"key1": to.Ptr("value1"),
		// 					},
		// 					TTL: to.Ptr[int64](3600),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("recorda2"),
		// 				Type: to.Ptr("Microsoft.Network/privateDnsZones/A"),
		// 				Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/A/recorda2"),
		// 				Properties: &armprivatedns.RecordSetProperties{
		// 					ARecords: []*armprivatedns.ARecord{
		// 						{
		// 							IPv4Address: to.Ptr("5.6.7.8"),
		// 						},
		// 					},
		// 					Fqdn: to.Ptr("recorda2.privatezone1.com."),
		// 					IsAutoRegistered: to.Ptr(false),
		// 					TTL: to.Ptr[int64](3600),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
