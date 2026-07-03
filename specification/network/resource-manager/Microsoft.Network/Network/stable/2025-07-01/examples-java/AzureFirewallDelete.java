
/**
 * Samples for AzureFirewalls Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureFirewallDelete.json
     */
    /**
     * Sample code: Delete Azure Firewall.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteAzureFirewall(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAzureFirewalls().delete("rg1", "azurefirewall", com.azure.core.util.Context.NONE);
    }
}
