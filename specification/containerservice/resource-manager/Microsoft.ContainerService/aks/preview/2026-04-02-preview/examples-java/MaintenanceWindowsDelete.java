
/**
 * Samples for MaintenanceWindows Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MaintenanceWindowsDelete.json
     */
    /**
     * Sample code: Delete Maintenance Window.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteMaintenanceWindow(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceWindows().delete("rg-maintenance", "production-weekends",
            com.azure.core.util.Context.NONE);
    }
}
