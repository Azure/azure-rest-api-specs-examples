/** Samples for MaintenanceConfigurations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/MaintenanceConfigurations_GetForResource.json
     */
    /**
     * Sample code: MaintenanceConfigurations_GetForResource.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void maintenanceConfigurationsGetForResource(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .maintenanceConfigurations()
            .getByResourceGroupWithResponse("examplerg", "configuration1", com.azure.core.util.Context.NONE);
    }
}
