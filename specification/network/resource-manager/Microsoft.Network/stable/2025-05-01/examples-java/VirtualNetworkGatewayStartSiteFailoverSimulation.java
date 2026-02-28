
/**
 * Samples for VirtualNetworkGateways StartExpressRouteSiteFailoverSimulation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayStartSiteFailoverSimulation.json
     */
    /**
     * Sample code: VirtualNetworkGatewayStartSiteFailoverSimulation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualNetworkGatewayStartSiteFailoverSimulation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways()
            .startExpressRouteSiteFailoverSimulation("rg1", "ergw", "Vancouver", com.azure.core.util.Context.NONE);
    }
}
