package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: 2026-05-01-preview/ClusterManagers_UpdateRelayPrivateEndpointConnection_Approve.json
func ExampleClusterManagersClient_BeginUpdateRelayPrivateEndpointConnection_approvePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("123e4567-e89b-12d3-a456-426655440000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterManagersClient().BeginUpdateRelayPrivateEndpointConnection(ctx, "resourceGroupName", "clusterManagerName", &armnetworkcloud.ClusterManagersClientBeginUpdateRelayPrivateEndpointConnectionOptions{
		ClusterManagerUpdateRelayPrivateEndpointConnectionParameters: &armnetworkcloud.ClusterManagerUpdateRelayPrivateEndpointConnectionParameters{
			ConnectionState:           to.Ptr(armnetworkcloud.RelayPrivateEndpointConnectionStateApproved),
			Description:               to.Ptr("Approving private endpoint connection"),
			PrivateEndpointResourceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
