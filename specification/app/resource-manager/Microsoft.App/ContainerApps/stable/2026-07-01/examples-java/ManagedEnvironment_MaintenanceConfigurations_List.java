
/**
 * Samples for MaintenanceConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironment_MaintenanceConfigurations_List.json
     */
    /**
     * Sample code: ManagedEnvironmentMaintenanceConfigurationsList.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void managedEnvironmentMaintenanceConfigurationsList(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.maintenanceConfigurations().list("rg1", "managedEnv", com.azure.core.util.Context.NONE);
    }
}
