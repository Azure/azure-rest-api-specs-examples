
/**
 * Samples for MaintenanceWindows GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MaintenanceWindowsGet.json
     */
    /**
     * Sample code: Get Maintenance Window.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getMaintenanceWindow(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceWindows().getByResourceGroupWithResponse("rg-maintenance",
            "production-weekends", com.azure.core.util.Context.NONE);
    }
}
