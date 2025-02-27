
/**
 * Samples for Inventory ListBySolutionConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/Inventory_ListBySolutionConfiguration.json
     */
    /**
     * Sample code: Inventory_ListBySolutionConfiguration.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void inventoryListBySolutionConfiguration(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.inventories().listBySolutionConfiguration("ymuj", "wsxt", com.azure.core.util.Context.NONE);
    }
}
