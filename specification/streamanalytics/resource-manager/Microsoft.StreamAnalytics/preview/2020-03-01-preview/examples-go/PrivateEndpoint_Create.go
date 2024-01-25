package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/PrivateEndpoint_Create.json
func ExamplePrivateEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointsClient().CreateOrUpdate(ctx, "sjrg", "testcluster", "testpe", armstreamanalytics.PrivateEndpoint{
		Properties: &armstreamanalytics.PrivateEndpointProperties{
			ManualPrivateLinkServiceConnections: []*armstreamanalytics.PrivateLinkServiceConnection{
				{
					Properties: &armstreamanalytics.PrivateLinkServiceConnectionProperties{
						GroupIDs: []*string{
							to.Ptr("groupIdFromResource")},
						PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
					},
				}},
		},
	}, &armstreamanalytics.PrivateEndpointsClientCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpoint = armstreamanalytics.PrivateEndpoint{
	// 	Name: to.Ptr("An Example Private Endpoint"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/clusters/privateEndpoints"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/clusters/testcluster/privateEndpoints/AnExamplePrivateEndpoint"),
	// 	Etag: to.Ptr("F86B9B70-D5B1-451D-AFC8-0B42D4729B8C"),
	// 	Properties: &armstreamanalytics.PrivateEndpointProperties{
	// 		CreatedDate: to.Ptr("2020-03-01T01:00Z"),
	// 		ManualPrivateLinkServiceConnections: []*armstreamanalytics.PrivateLinkServiceConnection{
	// 			{
	// 				Properties: &armstreamanalytics.PrivateLinkServiceConnectionProperties{
	// 					GroupIDs: []*string{
	// 						to.Ptr("groupIdFromResource")},
	// 						PrivateLinkServiceConnectionState: &armstreamanalytics.PrivateLinkConnectionState{
	// 							Description: to.Ptr("Awaiting approval"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr("Pending"),
	// 						},
	// 						PrivateLinkServiceID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
	// 						RequestMessage: to.Ptr("Please approve my connection."),
	// 					},
	// 			}},
	// 		},
	// 	}
}
