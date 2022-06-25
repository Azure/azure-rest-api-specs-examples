import com.azure.core.util.Context;

/** Samples for VpnGateways Reset. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnGatewayReset.json
     */
    /**
     * Sample code: ResetVpnGateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetVpnGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnGateways().reset("rg1", "vpngw", Context.NONE);
    }
}
