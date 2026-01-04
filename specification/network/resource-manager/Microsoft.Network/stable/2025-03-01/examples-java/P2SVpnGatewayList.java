
/**
 * Samples for P2SVpnGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/P2SVpnGatewayList.json
     */
    /**
     * Sample code: P2SVpnGatewayListBySubscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void p2SVpnGatewayListBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getP2SVpnGateways().list(com.azure.core.util.Context.NONE);
    }
}
