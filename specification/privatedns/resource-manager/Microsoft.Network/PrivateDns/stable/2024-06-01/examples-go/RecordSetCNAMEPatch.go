package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns/v2"
)

// Generated from example definition: 2024-06-01/RecordSetCNAMEPatch.json
func ExampleRecordSetsClient_Update_patchPrivateDnsZoneCnameRecordSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatedns.NewClientFactory("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecordSetsClient().Update(ctx, "resourceGroup1", "privatezone1.com", armprivatedns.RecordTypeCNAME, "recordCNAME", armprivatedns.RecordSet{
		Properties: &armprivatedns.RecordSetProperties{
			Metadata: map[string]*string{
				"key2": to.Ptr("value2"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprivatedns.RecordSetsClientUpdateResponse{
	// 	RecordSet: &armprivatedns.RecordSet{
	// 		Name: to.Ptr("recordcname"),
	// 		Type: to.Ptr("Microsoft.Network/privateDnsZones/CNAME"),
	// 		Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateDnsZones/privatezone1.com/CNAME/recordcname"),
	// 		Properties: &armprivatedns.RecordSetProperties{
	// 			CnameRecord: &armprivatedns.CnameRecord{
	// 				Cname: to.Ptr("contoso.com"),
	// 			},
	// 			Fqdn: to.Ptr("recordcname.privatezone1.com."),
	// 			IsAutoRegistered: to.Ptr(false),
	// 			Metadata: map[string]*string{
	// 				"key2": to.Ptr("value2"),
	// 			},
	// 			TTL: to.Ptr[int64](3600),
	// 		},
	// 	},
	// }
}
