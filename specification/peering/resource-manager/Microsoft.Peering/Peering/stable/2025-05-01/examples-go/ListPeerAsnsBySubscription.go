package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering/v2"
)

// Generated from example definition: 2025-05-01/ListPeerAsnsBySubscription.json
func ExamplePeerAsnsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPeerAsnsClient().NewListBySubscriptionPager(nil)
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
		// page = armpeering.PeerAsnsClientListBySubscriptionResponse{
		// 	PeerAsnListResult: armpeering.PeerAsnListResult{
		// 		Value: []*armpeering.PeerAsn{
		// 			{
		// 				Name: to.Ptr("peerAsnName"),
		// 				Type: to.Ptr("Microsoft.Peering/peerAsns"),
		// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/peerAsnName"),
		// 				Properties: &armpeering.PeerAsnProperties{
		// 					PeerAsn: to.Ptr[int32](65000),
		// 					PeerContactDetail: []*armpeering.ContactDetail{
		// 						{
		// 							Email: to.Ptr("noc@contoso.com"),
		// 							Phone: to.Ptr("+1 (234) 567-8999"),
		// 							Role: to.Ptr(armpeering.RoleNoc),
		// 						},
		// 						{
		// 							Email: to.Ptr("abc@contoso.com"),
		// 							Phone: to.Ptr("+1 (234) 567-8900"),
		// 							Role: to.Ptr(armpeering.RolePolicy),
		// 						},
		// 						{
		// 							Email: to.Ptr("xyz@contoso.com"),
		// 							Phone: to.Ptr("+1 (234) 567-8900"),
		// 							Role: to.Ptr(armpeering.RoleTechnical),
		// 						},
		// 					},
		// 					PeerName: to.Ptr("Contoso"),
		// 					ValidationState: to.Ptr(armpeering.ValidationStateApproved),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
