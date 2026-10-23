
/**
 * Samples for MaintenanceEvents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MaintenanceEvents_Get.json
     */
    /**
     * Sample code: Get information about a maintenance event for a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getInformationAboutAMaintenanceEventForAHorizonDBPool(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.maintenanceEvents().getWithResponse("example-resource-group", "example-cluster", "example-pool-01",
            "ABCD-123", com.azure.core.util.Context.NONE);
    }
}
