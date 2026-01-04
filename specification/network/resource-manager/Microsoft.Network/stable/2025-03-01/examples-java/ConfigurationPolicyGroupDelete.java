
/**
 * Samples for ConfigurationPolicyGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ConfigurationPolicyGroupDelete.json
     */
    /**
     * Sample code: ConfigurationPolicyGroupDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationPolicyGroupDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConfigurationPolicyGroups().delete("rg1",
            "vpnServerConfiguration1", "policyGroup1", com.azure.core.util.Context.NONE);
    }
}
