
/**
 * Samples for LocalNetworkGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LocalNetworkGatewayDelete.
     * json
     */
    /**
     * Sample code: DeleteLocalNetworkGateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteLocalNetworkGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLocalNetworkGateways().delete("rg1", "localgw",
            com.azure.core.util.Context.NONE);
    }
}
