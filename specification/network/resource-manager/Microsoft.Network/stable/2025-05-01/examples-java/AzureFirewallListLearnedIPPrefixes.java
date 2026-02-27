
/**
 * Samples for AzureFirewalls ListLearnedPrefixes.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * AzureFirewallListLearnedIPPrefixes.json
     */
    /**
     * Sample code: AzureFirewallListLearnedPrefixes.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void azureFirewallListLearnedPrefixes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().listLearnedPrefixes("rg1", "azureFirewall1",
            com.azure.core.util.Context.NONE);
    }
}
