
/**
 * Samples for VirtualNetworkGateways InvokeCommitMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayCommitMigration.json
     */
    /**
     * Sample code: VirtualNetworkGatewayCommitMigration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualNetworkGatewayCommitMigration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().invokeCommitMigration("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
