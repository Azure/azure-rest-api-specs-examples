
/**
 * Samples for MaintenanceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/MaintenanceConfigurationsGet_LinkedMaintenanceWindow.json
     */
    /**
     * Sample code: Get a Linked Maintenance Configuration.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getALinkedMaintenanceConfiguration(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceConfigurations().getWithResponse("rg1", "clustername1",
            "aksManagedAutoUpgradeSchedule", com.azure.core.util.Context.NONE);
    }
}
