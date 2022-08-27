import com.azure.core.util.Context;

/** Samples for VpnConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnConnectionDelete.json
     */
    /**
     * Sample code: VpnConnectionDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnConnectionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnConnections()
            .delete("rg1", "gateway1", "vpnConnection1", Context.NONE);
    }
}
