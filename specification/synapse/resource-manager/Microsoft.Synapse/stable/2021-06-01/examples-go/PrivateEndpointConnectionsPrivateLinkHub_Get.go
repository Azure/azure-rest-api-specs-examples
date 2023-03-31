package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/PrivateEndpointConnectionsPrivateLinkHub_Get.json
func ExamplePrivateEndpointConnectionsPrivateLinkHubClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsPrivateLinkHubClient().Get(ctx, "gh-res-grp", "pe0", "pe0-f3ed30f5-338c-4855-a542-24a403694ad2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionForPrivateLinkHub = armsynapse.PrivateEndpointConnectionForPrivateLinkHub{
	// 	ID: to.Ptr("/subscriptions/48b08652-d7a1-4d52-b13f-5a2471dce57b/resourceGroups/gh-res-grp/providers/Microsoft.Synapse/privateLinkHubs/plh900/privateEndpointConnections/pe0-f3ed30f5-338c-4855-a542-24a403694ad2"),
	// 	Properties: &armsynapse.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armsynapse.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/48b08652-d7a1-4d52-b13f-5a2471dce57b/resourceGroups/gh-res-grp/providers/Microsoft.Network/privateEndpoints/pe0"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armsynapse.PrivateLinkServiceConnectionState{
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// 	Name: to.Ptr("pe0-f3ed30f5-338c-4855-a542-24a403694ad2"),
	// 	Type: to.Ptr("Microsoft.Synapse/privateLinkHubs/privateEndpointConnections"),
	// }
}
