
/**
 * Samples for VirtualNetworkGateways InvokeCommitMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayCommitMigration.json
     */
    /**
     * Sample code: VirtualNetworkGatewayCommitMigration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualNetworkGatewayCommitMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().invokeCommitMigration("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
