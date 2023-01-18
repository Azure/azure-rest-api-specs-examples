/** Samples for PublicMaintenanceConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/PublicMaintenanceConfigurations_GetForResource.json
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
