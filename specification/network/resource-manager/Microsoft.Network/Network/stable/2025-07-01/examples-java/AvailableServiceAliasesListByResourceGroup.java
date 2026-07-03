
/**
 * Samples for AvailableServiceAliases ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AvailableServiceAliasesListByResourceGroup.json
     */
    /**
     * Sample code: Get available service aliases in the resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getAvailableServiceAliasesInTheResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAvailableServiceAliases().listByResourceGroup("rg1", "westcentralus",
            com.azure.core.util.Context.NONE);
    }
}
