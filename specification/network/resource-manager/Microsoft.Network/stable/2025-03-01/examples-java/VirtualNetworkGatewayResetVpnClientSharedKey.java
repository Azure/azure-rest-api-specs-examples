
/**
 * Samples for VirtualNetworkGateways ResetVpnClientSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayResetVpnClientSharedKey.json
     */
    /**
     * Sample code: ResetVpnClientSharedKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetVpnClientSharedKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().resetVpnClientSharedKey("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
