import com.azure.core.util.Context;

/** Samples for VpnGateways ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnGatewayListByResourceGroup.json
     */
    /**
     * Sample code: VpnGatewayListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnGatewayListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnGateways().listByResourceGroup("rg1", Context.NONE);
    }
}
