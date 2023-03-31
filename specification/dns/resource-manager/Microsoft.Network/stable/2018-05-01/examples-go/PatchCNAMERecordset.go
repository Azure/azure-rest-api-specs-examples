package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/PatchCNAMERecordset.json
func ExampleRecordSetsClient_Update_patchCnameRecordset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecordSetsClient().Update(ctx, "rg1", "zone1", "record1", armdns.RecordTypeCNAME, armdns.RecordSet{
		Properties: &armdns.RecordSetProperties{
			Metadata: map[string]*string{
				"key2": to.Ptr("value2"),
			},
		},
	}, &armdns.RecordSetsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecordSet = armdns.RecordSet{
	// 	Name: to.Ptr("record1"),
	// 	Type: to.Ptr("Microsoft.Network/dnsZones/CNAME"),
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnsZones/zone1/CNAME/record1"),
	// 	Properties: &armdns.RecordSetProperties{
	// 		CnameRecord: &armdns.CnameRecord{
	// 			Cname: to.Ptr("contoso.com"),
	// 		},
	// 		TTL: to.Ptr[int64](3600),
	// 		Fqdn: to.Ptr("record1.zone1"),
	// 		Metadata: map[string]*string{
	// 			"key2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
