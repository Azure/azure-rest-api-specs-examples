
/**
 * Samples for AzureFirewalls ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureFirewallListByResourceGroup.json
     */
    /**
     * Sample code: List all Azure Firewalls for a given resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllAzureFirewallsForAGivenResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAzureFirewalls().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
