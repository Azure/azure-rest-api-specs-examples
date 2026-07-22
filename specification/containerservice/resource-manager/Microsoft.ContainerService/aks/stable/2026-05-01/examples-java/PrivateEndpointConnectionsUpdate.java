
import com.azure.resourcemanager.containerservice.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.containerservice.models.ConnectionStatus;
import com.azure.resourcemanager.containerservice.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/PrivateEndpointConnectionsUpdate.json
     */
    /**
     * Sample code: Update Private Endpoint Connection.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        updatePrivateEndpointConnection(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getPrivateEndpointConnections().updateWithResponse("rg1", "clustername1",
            "privateendpointconnection1",
            new PrivateEndpointConnectionInner().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(ConnectionStatus.APPROVED)),
            com.azure.core.util.Context.NONE);
    }
}
