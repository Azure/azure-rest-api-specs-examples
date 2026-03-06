
/**
 * Samples for PublicMaintenanceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/PublicMaintenanceConfigurations_GetForResource.json
     */
    /**
     * Sample code: PublicMaintenanceConfigurations_GetForResource.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void publicMaintenanceConfigurationsGetForResource(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.publicMaintenanceConfigurations().getWithResponse("configuration1", com.azure.core.util.Context.NONE);
    }
}
