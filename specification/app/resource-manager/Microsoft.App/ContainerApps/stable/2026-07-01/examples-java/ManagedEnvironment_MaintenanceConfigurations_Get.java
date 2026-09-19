
/**
 * Samples for MaintenanceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironment_MaintenanceConfigurations_Get.json
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
