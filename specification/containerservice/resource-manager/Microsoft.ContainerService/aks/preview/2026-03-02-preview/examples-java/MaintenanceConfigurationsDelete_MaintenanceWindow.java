
/**
 * Samples for MaintenanceConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/MaintenanceConfigurationsDelete_MaintenanceWindow.json
     */
    /**
     * Sample code: Delete Maintenance Configuration For Node OS Upgrade.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void deleteMaintenanceConfigurationForNodeOSUpgrade(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceConfigurations().deleteWithResponse("rg1", "clustername1",
            "aksManagedNodeOSUpgradeSchedule", com.azure.core.util.Context.NONE);
    }
}
