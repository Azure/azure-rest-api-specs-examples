
/**
 * Samples for PrivateLinkServices GetPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * PrivateLinkServiceGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get private end point connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateEndPointConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateLinkServices().getPrivateEndpointConnectionWithResponse(
            "rg1", "testPls", "testPlePeConnection", null, com.azure.core.util.Context.NONE);
    }
}
