import com.azure.core.util.Context;

/** Samples for VpnConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VpnConnectionGet.json
     */
    /**
     * Sample code: VpnConnectionGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnConnections()
            .getWithResponse("rg1", "gateway1", "vpnConnection1", Context.NONE);
    }
}
