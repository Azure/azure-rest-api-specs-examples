
/**
 * Samples for ConfigurationPolicyGroups ListByVpnServerConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ConfigurationPolicyGroupListByVpnServerConfiguration.json
     */
    /**
     * Sample code: ConfigurationPolicyGroupListByVpnServerConfiguration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        configurationPolicyGroupListByVpnServerConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConfigurationPolicyGroups().listByVpnServerConfiguration("rg1",
            "vpnServerConfiguration1", com.azure.core.util.Context.NONE);
    }
}
