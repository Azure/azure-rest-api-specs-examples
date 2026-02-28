
/**
 * Samples for P2SVpnGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/P2SVpnGatewayGet.json
     */
    /**
     * Sample code: P2SVpnGatewayGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void p2SVpnGatewayGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getP2SVpnGateways().getByResourceGroupWithResponse("rg1",
            "p2sVpnGateway1", com.azure.core.util.Context.NONE);
    }
}
