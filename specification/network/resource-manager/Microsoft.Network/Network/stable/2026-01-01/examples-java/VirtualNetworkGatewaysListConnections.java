
/**
 * Samples for VirtualNetworkGateways ListConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/VirtualNetworkGatewaysListConnections.json
     */
    /**
     * Sample code: VirtualNetworkGatewaysListConnections.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualNetworkGatewaysListConnections(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().listConnections("testrg", "test-vpn-gateway-1",
            com.azure.core.util.Context.NONE);
    }
}
