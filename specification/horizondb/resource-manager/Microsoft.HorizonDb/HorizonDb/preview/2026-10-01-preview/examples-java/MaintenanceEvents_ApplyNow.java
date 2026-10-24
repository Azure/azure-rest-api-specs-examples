
/**
 * Samples for MaintenanceEvents ApplyNow.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MaintenanceEvents_ApplyNow.json
     */
    /**
     * Sample code: Apply a maintenance event immediately for a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void applyAMaintenanceEventImmediatelyForAHorizonDBPool(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.maintenanceEvents().applyNow("example-resource-group", "example-cluster", "example-pool-01", "ABCD-123",
            com.azure.core.util.Context.NONE);
    }
}
