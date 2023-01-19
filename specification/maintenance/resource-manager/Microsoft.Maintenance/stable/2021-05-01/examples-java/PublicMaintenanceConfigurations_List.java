/** Samples for PublicMaintenanceConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/PublicMaintenanceConfigurations_List.json
     */
    /**
     * Sample code: PublicMaintenanceConfigurations_List.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void publicMaintenanceConfigurationsList(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.publicMaintenanceConfigurations().list(com.azure.core.util.Context.NONE);
    }
}
