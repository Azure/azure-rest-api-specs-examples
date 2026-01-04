
/**
 * Samples for VpnGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VpnGatewayGet.json
     */
    /**
     * Sample code: VpnGatewayGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnGatewayGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnGateways().getByResourceGroupWithResponse("rg1", "gateway1",
            com.azure.core.util.Context.NONE);
    }
}
