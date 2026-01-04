
/**
 * Samples for VpnGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VpnGatewayList.json
     */
    /**
     * Sample code: VpnGatewayListBySubscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnGatewayListBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnGateways().list(com.azure.core.util.Context.NONE);
    }
}
