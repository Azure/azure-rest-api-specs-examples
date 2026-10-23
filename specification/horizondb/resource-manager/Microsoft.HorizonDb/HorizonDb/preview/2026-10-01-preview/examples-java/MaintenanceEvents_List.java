
/**
 * Samples for MaintenanceEvents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MaintenanceEvents_List.json
     */
    /**
     * Sample code: List maintenance events for a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listMaintenanceEventsForAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.maintenanceEvents().list("example-resource-group", "example-cluster", "example-pool-01", null,
            com.azure.core.util.Context.NONE);
    }
}
