
/**
 * Samples for VirtualNetworkGateways InvokeAbortMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayAbortMigration.json
     */
    /**
     * Sample code: VirtualNetworkGatewayAbortMigration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualNetworkGatewayAbortMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().invokeAbortMigration("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
