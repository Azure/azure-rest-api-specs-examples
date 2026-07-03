
/**
 * Samples for VpnServerConfigurations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnServerConfigurationGet.json
     */
    /**
     * Sample code: VpnServerConfigurationGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnServerConfigurationGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnServerConfigurations().getByResourceGroupWithResponse("rg1",
            "vpnServerConfiguration1", com.azure.core.util.Context.NONE);
    }
}
