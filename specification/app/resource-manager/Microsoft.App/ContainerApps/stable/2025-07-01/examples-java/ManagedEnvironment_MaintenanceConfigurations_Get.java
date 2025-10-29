
/**
 * Samples for MaintenanceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
     * ManagedEnvironment_MaintenanceConfigurations_Get.json
     */
    /**
     * Sample code: ManagedEnvironmentMaintenanceConfigurationsGet.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void managedEnvironmentMaintenanceConfigurationsGet(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.maintenanceConfigurations().getWithResponse("rg1", "managedEnv", "default",
            com.azure.core.util.Context.NONE);
    }
}
