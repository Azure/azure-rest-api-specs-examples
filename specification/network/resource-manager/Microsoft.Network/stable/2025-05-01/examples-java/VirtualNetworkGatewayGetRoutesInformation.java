
/**
 * Samples for VirtualNetworkGateways GetRoutesInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayGetRoutesInformation.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayRoutesInformation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayRoutesInformation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().getRoutesInformation("rg1", "vpngw",
            false, com.azure.core.util.Context.NONE);
    }
}
