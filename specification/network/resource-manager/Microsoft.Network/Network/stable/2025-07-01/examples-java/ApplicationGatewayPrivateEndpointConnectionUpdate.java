
import com.azure.resourcemanager.network.fluent.models.ApplicationGatewayPrivateEndpointConnectionInner;
import com.azure.resourcemanager.network.models.PrivateLinkServiceConnectionState;

/**
 * Samples for ApplicationGatewayPrivateEndpointConnections Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayPrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: Update Application Gateway Private Endpoint Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        updateApplicationGatewayPrivateEndpointConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGatewayPrivateEndpointConnections()
            .update("rg1", "appgw", "connection1",
                new ApplicationGatewayPrivateEndpointConnectionInner().withName("connection1")
                    .withPrivateLinkServiceConnectionState(new PrivateLinkServiceConnectionState()
                        .withStatus("Approved").withDescription("approved it for some reason.")),
                com.azure.core.util.Context.NONE);
    }
}
