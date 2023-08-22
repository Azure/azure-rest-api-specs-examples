/** Samples for ApplyUpdates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ApplyUpdates_List.json
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
