
/**
 * Samples for ConfigurationPolicyGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ConfigurationPolicyGroupGet.
     * json
     */
    /**
     * Sample code: ConfigurationPolicyGroupGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void configurationPolicyGroupGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getConfigurationPolicyGroups().getWithResponse("rg1",
            "vpnServerConfiguration1", "policyGroup1", com.azure.core.util.Context.NONE);
    }
}
