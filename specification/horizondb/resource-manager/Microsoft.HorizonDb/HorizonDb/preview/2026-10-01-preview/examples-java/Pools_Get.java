
/**
 * Samples for Pools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Pools_Get.json
     */
    /**
     * Sample code: Get a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.pools().getWithResponse("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE);
    }
}
