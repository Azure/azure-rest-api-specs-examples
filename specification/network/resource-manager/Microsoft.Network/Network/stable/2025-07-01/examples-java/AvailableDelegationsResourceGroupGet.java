
/**
 * Samples for AvailableResourceGroupDelegations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AvailableDelegationsResourceGroupGet.json
     */
    /**
     * Sample code: Get available delegations in the resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getAvailableDelegationsInTheResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAvailableResourceGroupDelegations().list("westcentralus", "rg1",
            com.azure.core.util.Context.NONE);
    }
}
