import com.azure.core.util.Context;

/** Samples for VpnConnections ListByVpnGateway. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VpnConnectionList.json
     */
    /**
     * Sample code: VpnConnectionList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnConnections()
            .listByVpnGateway("rg1", "gateway1", Context.NONE);
    }
}
