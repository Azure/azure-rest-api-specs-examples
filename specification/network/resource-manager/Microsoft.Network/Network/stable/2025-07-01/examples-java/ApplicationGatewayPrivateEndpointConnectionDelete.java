
/**
 * Samples for ApplicationGatewayPrivateEndpointConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayPrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Delete Application Gateway Private Endpoint Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteApplicationGatewayPrivateEndpointConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGatewayPrivateEndpointConnections().delete("rg1", "appgw", "connection1",
            com.azure.core.util.Context.NONE);
    }
}
