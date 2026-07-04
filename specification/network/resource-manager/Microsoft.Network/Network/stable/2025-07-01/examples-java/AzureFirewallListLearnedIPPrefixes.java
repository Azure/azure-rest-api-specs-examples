
/**
 * Samples for AzureFirewalls ListLearnedPrefixes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureFirewallListLearnedIPPrefixes.json
     */
    /**
     * Sample code: AzureFirewallListLearnedPrefixes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void azureFirewallListLearnedPrefixes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAzureFirewalls().listLearnedPrefixes("rg1", "azureFirewall1",
            com.azure.core.util.Context.NONE);
    }
}
