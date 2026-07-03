
/**
 * Samples for VirtualNetworkGateways GetFailoverAllTestDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGetFailoverAllTestsDetails.json
     */
    /**
     * Sample code: VirtualNetworkGatewayGetFailoverAllTestsDetails.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        virtualNetworkGatewayGetFailoverAllTestsDetails(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getFailoverAllTestDetails("rg1", "ergw",
            "SingleSiteFailover", true, com.azure.core.util.Context.NONE);
    }
}
