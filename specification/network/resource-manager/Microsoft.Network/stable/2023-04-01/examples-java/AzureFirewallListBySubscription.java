/** Samples for AzureFirewalls List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/AzureFirewallListBySubscription.json
     */
    /**
     * Sample code: List all Azure Firewalls for a given subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllAzureFirewallsForAGivenSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().list(com.azure.core.util.Context.NONE);
    }
}
