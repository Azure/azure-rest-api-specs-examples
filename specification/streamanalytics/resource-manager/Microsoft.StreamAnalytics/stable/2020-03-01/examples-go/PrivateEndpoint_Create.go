package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/PrivateEndpoint_Create.json
func ExamplePrivateEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewPrivateEndpointsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"sjrg",
		"testcluster",
		"testpe",
		armstreamanalytics.PrivateEndpoint{
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
		},
		&armstreamanalytics.PrivateEndpointsClientCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
