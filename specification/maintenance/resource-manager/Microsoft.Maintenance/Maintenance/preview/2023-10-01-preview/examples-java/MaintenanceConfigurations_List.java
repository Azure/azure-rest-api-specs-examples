
/**
 * Samples for MaintenanceConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/MaintenanceConfigurations_List.json
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
