
/**
 * Samples for ConfigurationPolicyGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ConfigurationPolicyGroupDelete.json
     */
    /**
     * Sample code: ConfigurationPolicyGroupDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void configurationPolicyGroupDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getConfigurationPolicyGroups().delete("rg1", "vpnServerConfiguration1", "policyGroup1",
            com.azure.core.util.Context.NONE);
    }
}
