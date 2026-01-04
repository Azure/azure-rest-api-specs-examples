
/**
 * Samples for P2SVpnGateways Reset.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/P2SVpnGatewayReset.json
     */
    /**
     * Sample code: ResetP2SVpnGateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetP2SVpnGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getP2SVpnGateways().reset("rg1", "p2sVpnGateway1",
            com.azure.core.util.Context.NONE);
    }
}
