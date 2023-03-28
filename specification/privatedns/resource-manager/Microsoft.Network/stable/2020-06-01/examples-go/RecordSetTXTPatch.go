package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetTXTPatch.json
func ExampleRecordSetsClient_Update_patchPrivateDnsZoneTxtRecordSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecordSetsClient().Update(ctx, "resourceGroup1", "privatezone1.com", armprivatedns.RecordTypeTXT, "recordTXT", armprivatedns.RecordSet{
		Properties: &armprivatedns.RecordSetProperties{
			Metadata: map[string]*string{
				"key2": to.Ptr("value2"),
			},
		},
	}, &armprivatedns.RecordSetsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecordSet = armprivatedns.RecordSet{
	// 	Name: to.Ptr("recordtxt"),
	// 	Type: to.Ptr("Microsoft.Network/privateDnsZones/TXT"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/TXT/recordtxt"),
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armprivatedns.RecordSetProperties{
	// 		Fqdn: to.Ptr("recordtxt.privatezone1.com."),
	// 		IsAutoRegistered: to.Ptr(false),
	// 		Metadata: map[string]*string{
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 		TTL: to.Ptr[int64](3600),
	// 		TxtRecords: []*armprivatedns.TxtRecord{
	// 			{
	// 				Value: []*string{
	// 					to.Ptr("string1"),
	// 					to.Ptr("string2")},
	// 			}},
	// 		},
	// 	}
}
