
/**
 * Samples for VirtualNetworkGateways InvokeExecuteMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VirtualNetworkGatewayExecuteMigration.json
     */
    /**
     * Sample code: VirtualNetworkGatewayExecuteMigration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualNetworkGatewayExecuteMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().invokeExecuteMigration("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
