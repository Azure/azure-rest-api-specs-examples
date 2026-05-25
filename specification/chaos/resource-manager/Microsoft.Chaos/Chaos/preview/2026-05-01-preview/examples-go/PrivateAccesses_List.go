package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/PrivateAccesses_List.json
func ExamplePrivateAccessesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateAccessesClient().NewListPager("myResourceGroup", nil)
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
		// page = armchaos.PrivateAccessesClientListResponse{
		// 	PrivateAccessListResult: armchaos.PrivateAccessListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Chaos/privateAccesses?continuationToken=&api-version=2026-05-01-preview"),
		// 		Value: []*armchaos.PrivateAccess{
		// 			{
		// 				Name: to.Ptr("myPrivateAccess2"),
		// 				Type: to.Ptr("Microsoft.Chaos/privateAccesses"),
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Chaos/privateAccesses/myPrivateAccess2"),
		// 				Location: to.Ptr("centraluseuap"),
		// 				Properties: &armchaos.PrivateAccessProperties{
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myPrivateAccess"),
		// 				Type: to.Ptr("Microsoft.Chaos/privateAccesses"),
		// 				ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Chaos/privateAccesses/myPrivateAccess"),
		// 				Location: to.Ptr("centraluseuap"),
		// 				Properties: &armchaos.PrivateAccessProperties{
		// 					PrivateEndpointConnections: []*armchaos.PrivateEndpointConnection{
		// 						{
		// 							Name: to.Ptr("myPrivateAccess.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
		// 							Type: to.Ptr("Microsoft.Chaos/privateAccesses/PrivateEndpointConnections"),
		// 							ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Chaos/privateAccesses/myPrivateAccess/privateEndpoinConnections/myPrivateAccess.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
		// 							Properties: &armchaos.PrivateEndpointConnectionProperties{
		// 								PrivateEndpoint: &armchaos.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armchaos.PrivateLinkServiceConnectionState{
		// 									Description: to.Ptr("Auto-Approved"),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Status: to.Ptr(armchaos.PrivateEndpointServiceConnectionStatusApproved),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					PublicNetworkAccess: to.Ptr(armchaos.PublicNetworkAccessOptionEnabled),
		// 				},
		// 				SystemData: &armchaos.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T00:00:00.0Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T00:00:00.0Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
