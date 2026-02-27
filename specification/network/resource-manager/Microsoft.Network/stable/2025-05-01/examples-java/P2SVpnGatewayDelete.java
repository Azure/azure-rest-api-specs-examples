
/**
 * Samples for P2SVpnGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/P2SVpnGatewayDelete.json
     */
    /**
     * Sample code: P2SVpnGatewayDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void p2SVpnGatewayDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getP2SVpnGateways().delete("rg1", "p2sVpnGateway1",
            com.azure.core.util.Context.NONE);
    }
}
