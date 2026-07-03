
/**
 * Samples for VpnServerConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnServerConfigurationDelete.json
     */
    /**
     * Sample code: VpnServerConfigurationDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnServerConfigurationDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnServerConfigurations().delete("rg1", "vpnServerConfiguration1",
            com.azure.core.util.Context.NONE);
    }
}
