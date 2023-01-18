/** Samples for MaintenanceConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/MaintenanceConfigurations_List.json
     */
    /**
     * Sample code: MaintenanceConfigurations_List.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void maintenanceConfigurationsList(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.maintenanceConfigurations().list(com.azure.core.util.Context.NONE);
    }
}
