
/**
 * Samples for MaintenanceConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * ManagedEnvironment_MaintenanceConfigurations_List.json
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
