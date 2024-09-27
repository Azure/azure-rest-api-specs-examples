package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4991356eec55985c1af47096c9c2091126a7d0f/specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/RecordSetSOAList.json
func ExampleRecordSetsClient_NewListByTypePager_getPrivateDnsZoneSoaRecordSets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecordSetsClient().NewListByTypePager("resourceGroup1", "privatezone1.com", armprivatedns.RecordTypeSOA, &armprivatedns.RecordSetsClientListByTypeOptions{Top: nil,
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
		// 			Name: to.Ptr("@"),
		// 			Type: to.Ptr("Microsoft.Network/privateDnsZones/SOA"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/SOA/@"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armprivatedns.RecordSetProperties{
		// 				Fqdn: to.Ptr("privatezone1.com."),
		// 				IsAutoRegistered: to.Ptr(false),
		// 				Metadata: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 				},
		// 				SoaRecord: &armprivatedns.SoaRecord{
		// 					Email: to.Ptr("azureprivatedns-hostmaster.microsoft.com"),
		// 					ExpireTime: to.Ptr[int64](2419200),
		// 					Host: to.Ptr("azureprivatedns.net"),
		// 					MinimumTTL: to.Ptr[int64](300),
		// 					RefreshTime: to.Ptr[int64](3600),
		// 					RetryTime: to.Ptr[int64](300),
		// 					SerialNumber: to.Ptr[int64](1),
		// 				},
		// 				TTL: to.Ptr[int64](3600),
		// 			},
		// 	}},
		// }
	}
}
