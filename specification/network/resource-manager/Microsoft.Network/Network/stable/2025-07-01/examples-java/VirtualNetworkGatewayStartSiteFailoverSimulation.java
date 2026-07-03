
/**
 * Samples for VirtualNetworkGateways StartExpressRouteSiteFailoverSimulation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayStartSiteFailoverSimulation.json
     */
    /**
     * Sample code: VirtualNetworkGatewayStartSiteFailoverSimulation.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        virtualNetworkGatewayStartSiteFailoverSimulation(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().startExpressRouteSiteFailoverSimulation("rg1", "ergw",
            "Vancouver", com.azure.core.util.Context.NONE);
    }
}
