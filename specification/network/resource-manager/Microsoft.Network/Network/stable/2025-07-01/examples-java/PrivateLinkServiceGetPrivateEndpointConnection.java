
/**
 * Samples for PrivateLinkServices GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateLinkServiceGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get private end point connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getPrivateEndPointConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().getPrivateEndpointConnectionWithResponse("rg1", "testPls",
            "testPlePeConnection", null, com.azure.core.util.Context.NONE);
    }
}
