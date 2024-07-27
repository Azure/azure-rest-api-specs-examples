
/**
 * Samples for VpnServerConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/VpnServerConfigurationDelete.
     * json
     */
    /**
     * Sample code: VpnServerConfigurationDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnServerConfigurationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnServerConfigurations().delete("rg1", "vpnServerConfiguration1",
            com.azure.core.util.Context.NONE);
    }
}
