
/**
 * Samples for VirtualNetworkGateways InvokeAbortMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayAbortMigration.json
     */
    /**
     * Sample code: VirtualNetworkGatewayAbortMigration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualNetworkGatewayAbortMigration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().invokeAbortMigration("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
