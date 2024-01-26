
/**
 * Samples for P2SVpnGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/
     * P2SVpnGatewayListByResourceGroup.json
     */
    /**
     * Sample code: P2SVpnGatewayListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void p2SVpnGatewayListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getP2SVpnGateways().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
