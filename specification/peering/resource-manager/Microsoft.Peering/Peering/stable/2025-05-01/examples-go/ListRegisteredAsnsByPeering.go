package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/ListRegisteredAsnsByPeering.json
func ExampleRegisteredAsnsClient_NewListByPeeringPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegisteredAsnsClient().NewListByPeeringPager("rgName", "peeringName", nil)
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
		// page = armpeering.RegisteredAsnsClientListByPeeringResponse{
		// 	RegisteredAsnListResult: armpeering.RegisteredAsnListResult{
		// 		Value: []*armpeering.RegisteredAsn{
		// 			{
		// 				Name: to.Ptr("registeredAsnName0"),
		// 				Type: to.Ptr("Microsoft.Peering/registeredAsns"),
		// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredAsns/registeredAsnName0"),
		// 				Properties: &armpeering.RegisteredAsnProperties{
		// 					Asn: to.Ptr[int32](65000),
		// 					PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("registeredAsnName1"),
		// 				Type: to.Ptr("Microsoft.Peering/registeredAsns"),
		// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredAsns/registeredAsnName1"),
		// 				Properties: &armpeering.RegisteredAsnProperties{
		// 					Asn: to.Ptr[int32](65001),
		// 					ProvisioningState: to.Ptr(armpeering.ProvisioningStateFailed),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("registeredAsnName2"),
		// 				Type: to.Ptr("Microsoft.Peering/registeredAsns"),
		// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName/registeredAsns/registeredAsnName2"),
		// 				Properties: &armpeering.RegisteredAsnProperties{
		// 					Asn: to.Ptr[int32](65002),
		// 					PeeringServicePrefixKey: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
