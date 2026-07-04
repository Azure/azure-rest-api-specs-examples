
/**
 * Samples for ApplicationGatewayPrivateEndpointConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayPrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Lists all private endpoint connections on application gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listsAllPrivateEndpointConnectionsOnApplicationGateway(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGatewayPrivateEndpointConnections().list("rg1", "appgw",
            com.azure.core.util.Context.NONE);
    }
}
