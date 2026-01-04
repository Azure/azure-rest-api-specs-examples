
/**
 * Samples for AzureFirewalls Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/AzureFirewallDelete.json
     */
    /**
     * Sample code: Delete Azure Firewall.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAzureFirewall(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().delete("rg1", "azurefirewall",
            com.azure.core.util.Context.NONE);
    }
}
