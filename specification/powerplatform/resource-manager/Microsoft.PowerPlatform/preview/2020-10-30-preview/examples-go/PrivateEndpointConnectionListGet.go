package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/powerplatform/resource-manager/Microsoft.PowerPlatform/preview/2020-10-30-preview/examples/PrivateEndpointConnectionListGet.json
func ExamplePrivateEndpointConnectionsClient_NewListByEnterprisePolicyPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByEnterprisePolicyPager("rg1", "ddb1", nil)
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
		// page.PrivateEndpointConnectionListResult = armpowerplatform.PrivateEndpointConnectionListResult{
		// 	Value: []*armpowerplatform.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("privateEndpointConnectionName"),
		// 			Type: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.PowerPlatform/enterprisePolicies/ddb1/privateEndpointConnections/privateEndpointConnectionName"),
		// 			Properties: &armpowerplatform.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armpowerplatform.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armpowerplatform.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armpowerplatform.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("privateEndpointConnectionName"),
		// 			Type: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.PowerPlatform/enterprisePolicies/ddb1/privateEndpointConnections/privateEndpointConnectionName2"),
		// 			Properties: &armpowerplatform.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armpowerplatform.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1Network/providers/Microsoft.Network/privateEndpoints/privateEndpointName2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armpowerplatform.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armpowerplatform.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 			},
		// 			SystemData: &armpowerplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
