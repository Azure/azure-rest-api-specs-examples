
/**
 * Samples for MaintenanceEvents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MaintenanceEvents_ListWithFilter.json
     */
    /**
     * Sample code: List upcoming maintenance events for a HorizonDB pool.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listUpcomingMaintenanceEventsForAHorizonDBPool(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.maintenanceEvents().list("example-resource-group", "example-cluster", "example-pool-01",
            "properties/status eq 'Planned' or properties/status eq 'InProgress'", com.azure.core.util.Context.NONE);
    }
}
