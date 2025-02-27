
/**
 * Samples for Inventory Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/Inventory_Get.json
     */
    /**
     * Sample code: Inventory_Get.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void inventoryGet(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.inventories().getWithResponse("ymuj", "zarfsraogroxlaqjjnwixtn", "xofprmcboosrbd",
            com.azure.core.util.Context.NONE);
    }
}
