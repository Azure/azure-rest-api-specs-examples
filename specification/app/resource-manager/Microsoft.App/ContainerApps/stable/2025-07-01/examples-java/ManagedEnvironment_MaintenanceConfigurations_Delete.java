
/**
 * Samples for MaintenanceConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
     * ManagedEnvironment_MaintenanceConfigurations_Delete.json
     */
    /**
     * Sample code: ManagedEnvironmentMaintenanceConfigurationsDelete.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void managedEnvironmentMaintenanceConfigurationsDelete(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.maintenanceConfigurations().deleteWithResponse("rg1", "managedEnv", "default",
            com.azure.core.util.Context.NONE);
    }
}
