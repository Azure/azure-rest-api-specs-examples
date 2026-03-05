
/**
 * Samples for MaintenanceConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/MaintenanceConfigurations_DeleteForResource.json
     */
    /**
     * Sample code: MaintenanceConfigurations_DeleteForResource.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        maintenanceConfigurationsDeleteForResource(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.maintenanceConfigurations().deleteByResourceGroupWithResponse("examplerg", "example1",
            com.azure.core.util.Context.NONE);
    }
}
