
/**
 * Samples for ApplyUpdates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/ApplyUpdates_List.json
     */
    /**
     * Sample code: ApplyUpdates_List.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void applyUpdatesList(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.applyUpdates().list(com.azure.core.util.Context.NONE);
    }
}
