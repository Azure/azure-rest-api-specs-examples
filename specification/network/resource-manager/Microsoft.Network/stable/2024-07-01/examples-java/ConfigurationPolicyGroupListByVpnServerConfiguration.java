
/**
 * Samples for ConfigurationPolicyGroups ListByVpnServerConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * ConfigurationPolicyGroupListByVpnServerConfiguration.json
     */
    /**
     * Sample code: ConfigurationPolicyGroupListByVpnServerConfiguration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        configurationPolicyGroupListByVpnServerConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConfigurationPolicyGroups().listByVpnServerConfiguration("rg1",
            "vpnServerConfiguration1", com.azure.core.util.Context.NONE);
    }
}
