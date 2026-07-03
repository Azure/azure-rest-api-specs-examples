
/**
 * Samples for PrivateLinkServices DeletePrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateLinkServiceDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: delete private end point connection for a private link service.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletePrivateEndPointConnectionForAPrivateLinkService(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().deletePrivateEndpointConnection("rg1", "testPls",
            "testPlePeConnection", com.azure.core.util.Context.NONE);
    }
}
