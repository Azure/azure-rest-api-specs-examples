
/**
 * Samples for VirtualNetworkGateways GetResiliencyInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayGetResiliencyInformation.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayResiliencyInformation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getVirtualNetworkGatewayResiliencyInformation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().getResiliencyInformation("rg1", "vpngw",
            true, com.azure.core.util.Context.NONE);
    }
}
