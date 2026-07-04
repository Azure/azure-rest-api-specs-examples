
/**
 * Samples for VirtualNetworkGateways InvokeExecuteMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayExecuteMigration.json
     */
    /**
     * Sample code: VirtualNetworkGatewayExecuteMigration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualNetworkGatewayExecuteMigration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().invokeExecuteMigration("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
