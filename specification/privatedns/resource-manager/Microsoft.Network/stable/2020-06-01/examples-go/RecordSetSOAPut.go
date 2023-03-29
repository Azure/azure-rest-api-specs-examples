package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/RecordSetSOAPut.json
func ExampleRecordSetsClient_CreateOrUpdate_putPrivateDnsZoneSoaRecordSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecordSetsClient().CreateOrUpdate(ctx, "resourceGroup1", "privatezone1.com", armprivatedns.RecordTypeSOA, "@", armprivatedns.RecordSet{
		Properties: &armprivatedns.RecordSetProperties{
			Metadata: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			SoaRecord: &armprivatedns.SoaRecord{
				Email:        to.Ptr("azureprivatedns-hostmaster.microsoft.com"),
				ExpireTime:   to.Ptr[int64](2419200),
				Host:         to.Ptr("azureprivatedns.net"),
				RefreshTime:  to.Ptr[int64](3600),
				RetryTime:    to.Ptr[int64](300),
				SerialNumber: to.Ptr[int64](1),
			},
			TTL: to.Ptr[int64](3600),
		},
	}, &armprivatedns.RecordSetsClientCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecordSet = armprivatedns.RecordSet{
	// 	Name: to.Ptr("@"),
	// 	Type: to.Ptr("Microsoft.Network/privateDnsZones/SOA"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/SOA/@"),
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armprivatedns.RecordSetProperties{
	// 		Fqdn: to.Ptr("privatezone1.com."),
	// 		IsAutoRegistered: to.Ptr(false),
	// 		Metadata: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 		SoaRecord: &armprivatedns.SoaRecord{
	// 			Email: to.Ptr("azureprivatedns-hostmaster.microsoft.com"),
	// 			ExpireTime: to.Ptr[int64](2419200),
	// 			Host: to.Ptr("azureprivatedns.net"),
	// 			RefreshTime: to.Ptr[int64](3600),
	// 			RetryTime: to.Ptr[int64](300),
	// 			SerialNumber: to.Ptr[int64](1),
	// 		},
	// 		TTL: to.Ptr[int64](3600),
	// 	},
	// }
}
