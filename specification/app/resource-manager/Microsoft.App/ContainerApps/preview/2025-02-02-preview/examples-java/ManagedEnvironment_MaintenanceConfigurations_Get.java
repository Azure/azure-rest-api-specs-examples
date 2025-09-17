
/**
 * Samples for MaintenanceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
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
