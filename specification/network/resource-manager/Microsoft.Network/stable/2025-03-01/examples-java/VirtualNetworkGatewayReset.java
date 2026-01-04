
/**
 * Samples for VirtualNetworkGateways Reset.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewayReset.
     * json
     */
    /**
     * Sample code: ResetVirtualNetworkGateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetVirtualNetworkGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().reset("rg1", "vpngw", null,
            com.azure.core.util.Context.NONE);
    }
}
