
/**
 * Samples for ApplicationGatewayPrivateEndpointConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayPrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: Get Application Gateway Private Endpoint Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getApplicationGatewayPrivateEndpointConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGatewayPrivateEndpointConnections().getWithResponse("rg1", "appgw",
            "connection1", com.azure.core.util.Context.NONE);
    }
}
