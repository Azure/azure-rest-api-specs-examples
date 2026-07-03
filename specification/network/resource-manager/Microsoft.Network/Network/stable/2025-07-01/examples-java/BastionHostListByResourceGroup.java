
/**
 * Samples for BastionHosts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostListByResourceGroup.json
     */
    /**
     * Sample code: List all Bastion Hosts for a given resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllBastionHostsForAGivenResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
