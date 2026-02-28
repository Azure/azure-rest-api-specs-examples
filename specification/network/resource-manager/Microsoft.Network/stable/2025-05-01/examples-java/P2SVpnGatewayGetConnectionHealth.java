
/**
 * Samples for P2SVpnGateways GetP2SVpnConnectionHealth.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * P2SVpnGatewayGetConnectionHealth.json
     */
    /**
     * Sample code: P2SVpnGatewayGetConnectionHealth.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void p2SVpnGatewayGetConnectionHealth(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getP2SVpnGateways().getP2SVpnConnectionHealth("rg1",
            "p2sVpnGateway1", com.azure.core.util.Context.NONE);
    }
}
