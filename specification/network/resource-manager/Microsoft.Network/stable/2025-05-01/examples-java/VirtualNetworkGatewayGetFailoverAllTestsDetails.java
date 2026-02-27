
/**
 * Samples for VirtualNetworkGateways GetFailoverAllTestDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayGetFailoverAllTestsDetails.json
     */
    /**
     * Sample code: VirtualNetworkGatewayGetFailoverAllTestsDetails.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualNetworkGatewayGetFailoverAllTestsDetails(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().getFailoverAllTestDetails("rg1", "ergw",
            "SingleSiteFailover", true, com.azure.core.util.Context.NONE);
    }
}
