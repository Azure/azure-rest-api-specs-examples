
/**
 * Samples for VpnServerConfigurations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnServerConfigurationListByResourceGroup.json
     */
    /**
     * Sample code: VpnServerConfigurationListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        vpnServerConfigurationListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnServerConfigurations().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
