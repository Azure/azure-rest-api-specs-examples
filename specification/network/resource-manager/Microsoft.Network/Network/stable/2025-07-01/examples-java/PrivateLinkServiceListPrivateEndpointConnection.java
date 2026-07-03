
/**
 * Samples for PrivateLinkServices ListPrivateEndpointConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateLinkServiceListPrivateEndpointConnection.json
     */
    /**
     * Sample code: List private link service in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listPrivateLinkServiceInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().listPrivateEndpointConnections("rg1", "testPls",
            com.azure.core.util.Context.NONE);
    }
}
