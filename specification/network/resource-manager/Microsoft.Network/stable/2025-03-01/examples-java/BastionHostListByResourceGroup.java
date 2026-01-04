
/**
 * Samples for BastionHosts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * BastionHostListByResourceGroup.json
     */
    /**
     * Sample code: List all Bastion Hosts for a given resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllBastionHostsForAGivenResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getBastionHosts().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
