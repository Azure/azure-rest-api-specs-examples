
/**
 * Samples for PrivateLinkServices ListPrivateEndpointConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * PrivateLinkServiceListPrivateEndpointConnection.json
     */
    /**
     * Sample code: List private link service in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPrivateLinkServiceInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateLinkServices().listPrivateEndpointConnections("rg1",
            "testPls", com.azure.core.util.Context.NONE);
    }
}
