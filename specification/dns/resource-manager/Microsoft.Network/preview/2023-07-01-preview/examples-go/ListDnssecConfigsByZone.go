package armdns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/987a8f38ab2a8359d085e149be042267a9ecc66f/specification/dns/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/ListDnssecConfigsByZone.json
func ExampleDnssecConfigsClient_NewListByDNSZonePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdns.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDnssecConfigsClient().NewListByDNSZonePager("rg1", "zone1", nil)
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
		// page.DnssecConfigListResult = armdns.DnssecConfigListResult{
		// 	Value: []*armdns.DnssecConfig{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Network/dnszones/dnssecConfigs"),
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/dnszones/zone1/dnssecConfigs/default"),
		// 			Properties: &armdns.DnssecProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SigningKeys: []*armdns.SigningKey{
		// 					{
		// 						DelegationSignerInfo: []*armdns.DelegationSignerInfo{
		// 						},
		// 						Flags: to.Ptr[int32](256),
		// 						KeyTag: to.Ptr[int32](37721),
		// 						PublicKey: to.Ptr("publicKey1"),
		// 						SecurityAlgorithmType: to.Ptr[int32](13),
		// 						Protocol: to.Ptr[int32](3),
		// 					},
		// 					{
		// 						DelegationSignerInfo: []*armdns.DelegationSignerInfo{
		// 							{
		// 								DigestAlgorithmType: to.Ptr[int32](2),
		// 								DigestValue: to.Ptr("digestValue1"),
		// 								Record: to.Ptr("11920 13 2 digestValue1"),
		// 						}},
		// 						Flags: to.Ptr[int32](257),
		// 						KeyTag: to.Ptr[int32](11920),
		// 						PublicKey: to.Ptr("publicKey2"),
		// 						SecurityAlgorithmType: to.Ptr[int32](13),
		// 						Protocol: to.Ptr[int32](3),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
