
/**
 * Samples for MaintenanceConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
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
