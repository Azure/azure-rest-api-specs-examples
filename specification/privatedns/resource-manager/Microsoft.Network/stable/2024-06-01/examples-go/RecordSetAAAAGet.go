package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4991356eec55985c1af47096c9c2091126a7d0f/specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/RecordSetAAAAGet.json
func ExampleRecordSetsClient_Get_getPrivateDnsZoneAaaaRecordSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecordSetsClient().Get(ctx, "resourceGroup1", "privatezone1.com", armprivatedns.RecordTypeAAAA, "recordAAAA", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecordSet = armprivatedns.RecordSet{
	// 	Name: to.Ptr("recordaaaa"),
	// 	Type: to.Ptr("Microsoft.Network/privateDnsZones/AAAA"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/AAAA/recordaaaa"),
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armprivatedns.RecordSetProperties{
	// 		AaaaRecords: []*armprivatedns.AaaaRecord{
	// 			{
	// 				IPv6Address: to.Ptr("::1"),
	// 		}},
	// 		Fqdn: to.Ptr("recordaaaa.privatezone1.com."),
	// 		IsAutoRegistered: to.Ptr(false),
	// 		Metadata: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 		TTL: to.Ptr[int64](3600),
	// 	},
	// }
}
