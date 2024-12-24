
/**
 * Samples for VirtualNetworkGateways GetFailoverSingleTestDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/
     * VirtualNetworkGatewayGetFailoverSingleTestDetails.json
     */
    /**
     * Sample code: VirtualNetworkGatewayGetFailoverSingleTestDetails.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualNetworkGatewayGetFailoverSingleTestDetails(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().getFailoverSingleTestDetails("rg1",
            "ergw", "Vancouver", "fe458ae8-d2ae-4520-a104-44bc233bde7e", com.azure.core.util.Context.NONE);
    }
}
