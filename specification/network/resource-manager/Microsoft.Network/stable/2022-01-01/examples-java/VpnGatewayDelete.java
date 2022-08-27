import com.azure.core.util.Context;

/** Samples for VpnGateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnGatewayDelete.json
     */
    /**
     * Sample code: VpnGatewayDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnGatewayDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnGateways().delete("rg1", "gateway1", Context.NONE);
    }
}
