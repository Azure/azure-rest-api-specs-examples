
/**
 * Samples for AzureFirewallFqdnTags List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * AzureFirewallFqdnTagsListBySubscription.json
     */
    /**
     * Sample code: List all Azure Firewall FQDN Tags for a given subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllAzureFirewallFQDNTagsForAGivenSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewallFqdnTags().list(com.azure.core.util.Context.NONE);
    }
}
