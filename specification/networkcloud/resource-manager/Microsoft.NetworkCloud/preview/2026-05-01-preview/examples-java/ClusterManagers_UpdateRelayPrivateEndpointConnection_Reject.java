
import com.azure.resourcemanager.networkcloud.models.ClusterManagerUpdateRelayPrivateEndpointConnectionParameters;
import com.azure.resourcemanager.networkcloud.models.RelayPrivateEndpointConnectionState;

/**
 * Samples for ClusterManagers UpdateRelayPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ClusterManagers_UpdateRelayPrivateEndpointConnection_Reject.json
     */
    /**
     * Sample code: Reject private endpoint connection.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        rejectPrivateEndpointConnection(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusterManagers().updateRelayPrivateEndpointConnection("resourceGroupName", "clusterManagerName",
            new ClusterManagerUpdateRelayPrivateEndpointConnectionParameters()
                .withConnectionState(RelayPrivateEndpointConnectionState.REJECTED)
                .withDescription("Rejecting private endpoint connection").withPrivateEndpointResourceId(
                    "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
            com.azure.core.util.Context.NONE);
    }
}
