
/**
 * Samples for AzureFirewalls ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * AzureFirewallListByResourceGroup.json
     */
    /**
     * Sample code: List all Azure Firewalls for a given resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllAzureFirewallsForAGivenResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
