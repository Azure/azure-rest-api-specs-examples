
/**
 * Samples for VirtualNetworkGateways GetRoutesInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGetRoutesInformation.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayRoutesInformation.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayRoutesInformation(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getRoutesInformation("rg1", "vpngw", false,
            com.azure.core.util.Context.NONE);
    }
}
