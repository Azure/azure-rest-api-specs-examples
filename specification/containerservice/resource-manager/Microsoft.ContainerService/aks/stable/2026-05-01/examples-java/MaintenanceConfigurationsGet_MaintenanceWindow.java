
/**
 * Samples for MaintenanceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/MaintenanceConfigurationsGet_MaintenanceWindow.json
     */
    /**
     * Sample code: Get Maintenance Configuration Configured With Maintenance Window.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getMaintenanceConfigurationConfiguredWithMaintenanceWindow(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceConfigurations().getWithResponse("rg1", "clustername1",
            "aksManagedNodeOSUpgradeSchedule", com.azure.core.util.Context.NONE);
    }
}
