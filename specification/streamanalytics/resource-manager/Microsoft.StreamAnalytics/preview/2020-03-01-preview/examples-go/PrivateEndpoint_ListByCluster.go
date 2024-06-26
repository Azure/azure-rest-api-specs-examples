package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/PrivateEndpoint_ListByCluster.json
func ExamplePrivateEndpointsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointsClient().NewListByClusterPager("sjrg", "testcluster", nil)
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
		// page.PrivateEndpointListResult = armstreamanalytics.PrivateEndpointListResult{
		// 	Value: []*armstreamanalytics.PrivateEndpoint{
		// 		{
		// 			Name: to.Ptr("An Example Private Endpoint"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/clusters/privateEndpoints"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/clusters/testcluster/privateEndpoints/AnExamplePrivateEndpoint"),
		// 			Etag: to.Ptr("F86B9B70-D5B1-451D-AFC8-0B42D4729B8C"),
		// 			Properties: &armstreamanalytics.PrivateEndpointProperties{
		// 				CreatedDate: to.Ptr("2020-03-01T01:00Z"),
		// 				ManualPrivateLinkServiceConnections: []*armstreamanalytics.PrivateLinkServiceConnection{
		// 					{
		// 						Properties: &armstreamanalytics.PrivateLinkServiceConnectionProperties{
		// 							GroupIDs: []*string{
		// 								to.Ptr("groupIdFromResource")},
		// 								PrivateLinkServiceConnectionState: &armstreamanalytics.PrivateLinkConnectionState{
		// 									Description: to.Ptr("Awaiting approval"),
		// 									ActionsRequired: to.Ptr("None"),
		// 									Status: to.Ptr("Pending"),
		// 								},
		// 								PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
		// 								RequestMessage: to.Ptr("Please approve my connection."),
		// 							},
		// 					}},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("A Different Private Endpoint"),
		// 				Type: to.Ptr("Microsoft.StreamAnalytics/clusters/privateEndpoints"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/clusters/testcluster/privateEndpoints/ADifferentPrivateEndpoint"),
		// 				Etag: to.Ptr("G97C0C81-D5B1-451D-AFC8-0B42D4729B8C"),
		// 				Properties: &armstreamanalytics.PrivateEndpointProperties{
		// 					CreatedDate: to.Ptr("2020-03-01T01:00Z"),
		// 					ManualPrivateLinkServiceConnections: []*armstreamanalytics.PrivateLinkServiceConnection{
		// 						{
		// 							Properties: &armstreamanalytics.PrivateLinkServiceConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("groupIdFromResource")},
		// 									PrivateLinkServiceConnectionState: &armstreamanalytics.PrivateLinkConnectionState{
		// 										Description: to.Ptr("Awaiting approval"),
		// 										ActionsRequired: to.Ptr("None"),
		// 										Status: to.Ptr("Pending"),
		// 									},
		// 									PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
		// 									RequestMessage: to.Ptr("Please approve my connection."),
		// 								},
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
